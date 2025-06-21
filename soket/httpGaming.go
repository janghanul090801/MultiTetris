package soket

import (
	"MultiTetris/blockShape"
	"MultiTetris/user"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eiannone/keyboard"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"time"
)

type NgrokTunnel struct {
	PublicURL string `json:"public_url"`
}

type NgrokTunnels struct {
	Tunnels []NgrokTunnel `json:"tunnels"`
}

type GameState struct {
	Ground       [10][10]blockShape.Block `json:"ground"`
	FallingBlock blockShape.BlockInfo     `json:"falling_block"`
}

var cmd *exec.Cmd
var WaitConnect = make(chan struct{})

var WaitClient = make(chan struct{})
var WaitEnd = make(chan struct{})

func GetUrl() {
	cmd = exec.Command("soket/ngrok.exe", "http", "8080")
	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	// 2. ngrok API가 활성화될 때까지 잠시 대기
	time.Sleep(2 * time.Second)

	// 3. ngrok API 호출해서 public URL 가져오기
	resp, err := http.Get("http://127.0.0.1:4040/api/tunnels")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var tunnels NgrokTunnels
	err = json.NewDecoder(resp.Body).Decode(&tunnels)
	if err != nil {
		panic(err)
	}

	// 4. public URL 출력
	for _, tunnel := range tunnels.Tunnels {
		fmt.Println("공개 URL:", tunnel.PublicURL)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt) // Ctrl+C

	go func() {
		<-sig
		fmt.Println("Ctrl+C 감지, ngrok 종료")
		_ = cmd.Process.Kill()
		os.Exit(0)
	}()

}

func StartServerSide() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	inputChan := make(chan rune)
	errChan := make(chan error)
	go func() {
		for {
			ch, _, err := keyboard.GetKey()
			if err != nil {
				errChan <- err
				return
			}
			inputChan <- ch
		}
	}()

	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("받은 메시지:", string(body))
		_, _ = w.Write([]byte("O"))
		close(WaitConnect) // 기존 채널 닫기
	})

	http.HandleFunc("/gaming", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		body, _ := io.ReadAll(r.Body)
		attackSuccess := false
		if string(body) != ",,0" {
			parts := strings.Split(string(body), ",")
			if len(parts) >= 3 && parts[2] == "1" {
				attackSuccess = true
			}
		}

		blockShape.FallingDown()
		if attackSuccess {
			blockShape.DeleteFallingBlock()
		}

		var input rune
		var gotInput bool
		timeout := time.After(5 * time.Second)

	WAIT_INPUT:
		for {
			select {
			case ch := <-inputChan:
				input = ch
				gotInput = true
				break WAIT_INPUT
			case err := <-errChan:
				fmt.Println("키 입력 에러:", err)
				break WAIT_INPUT
			case <-timeout:
				fmt.Println("입력 시간 초과! 자동 턴 종료")
				break WAIT_INPUT
			}
		}

		if gotInput {
			blockShape.Move(input)
		}

		state := GameState{
			Ground:       blockShape.Ground,
			FallingBlock: blockShape.FallingBlock,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(state); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println("클라이언트턴 대기중...")
	})

	http.HandleFunc("/end", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var u user.User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("받은 user data: %+v\n", u)
		user.Other = u

		j, err := json.Marshal(user.Me)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(j)
		close(WaitEnd)
	})

	fmt.Println("서버 시작 : 8080 포트")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func ConnectServerSide() (bool, string) {
	fmt.Println("서버 url를 입력 : ")
	var url string
	fmt.Scanln(&url)
	fmt.Println(url)
	url = strings.TrimSpace(url)

	if url == "" {
		fmt.Println("⚠️ 입력된 URL이 비어 있습니다. ngrok URL을 정확히 입력하세요.")
		os.Exit(1)
	}
	message := []byte("0,0")
	resp, err := http.Post(url+"connect", "text/plain", bytes.NewBuffer(message))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	bodyS := string(body)
	fmt.Println("body:" + bodyS)
	if bodyS == "O" {
		return true, url
	} else {
		return false, url
	}
}

func GamingClientSide(url string, message string, attackSuccess bool) GameState {
	WaitClient = make(chan struct{})
	if attackSuccess {
		message += ",1"
	} else {
		message += ",0"
	}
	resp, err := http.Post(url+"gaming", "text/plain", bytes.NewBuffer([]byte(message)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var state GameState
	if err := json.Unmarshal(body, &state); err != nil {
		panic(err)
	}
	close(WaitClient)
	return state
}
func SendUserData(url string, u user.User) {
	data, err := json.Marshal(u) // 이것은 User 구조체 데이터에 작성한 `json:"머시기"`에 맞춰 json 데이터로 바꿔주는 멋있는 함수
	if err != nil {
		log.Fatal("json.Marshal에서 에러가 나요ㅠㅠ", err)
	}

	resp, err := http.Post(url+"end", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("http.Post에서 에러가 나요ㅠㅠ", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var other user.User
	if err := json.Unmarshal(body, &other); err != nil {
		panic(err)
	}
	user.Other = other
}
