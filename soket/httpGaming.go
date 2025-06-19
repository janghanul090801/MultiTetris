package soket

import (
	"MultiTetris/blockShape"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eiannone/keyboard"
	"io"
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

var cmd *exec.Cmd
var WaitConnect = make(chan struct{})

// var WaitClient = make(chan struct{})

//func main() {
//	getUrl()
//	go StartServerSide()
//
//	ConnectServerSide()
//
//	//StartServerSide()
//	fmt.Println("Kill ngrok")
//	_ = cmd.Process.Kill()
//	_ = cmd.Wait()
//}

func GetUrl() {
	cmd = exec.Command("C:/Users/JUYOUNG/Desktop/Coding/MultiTetris/soket/ngrok.exe", "http", "8080")
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
	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := io.ReadAll(r.Body)
			if err == nil {
				fmt.Println("받은 메시지:", string(body))
				_, _ = w.Write([]byte("O"))
				close(WaitConnect)
			} else {
				_, _ = w.Write([]byte(err.Error()))
			}
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/gaming", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		body, _ := io.ReadAll(r.Body)
		fmt.Println("클라 메시지:", string(body))
		var attackSuccess string
		if string(body) == "" {
			goto RETURN
		}
		attackSuccess = strings.Split(string(body), ",")[1]

		if attackSuccess == "1" {
			blockShape.DeleteFallingBlock()
		}
		for {
			fmt.Print("서버 입력: ")
			ch, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}
			if key == keyboard.KeyEsc {
				os.Exit(0)
			}

			blockShape.Move(ch)
			if ch == 'f' {
				blockShape.FallingDown()
				break
			}
			blockShape.PrintArray(blockShape.Ground)
		}
	RETURN:
		GroundJson, err := json.Marshal(blockShape.Ground)
		if err != nil {
			panic(err)
		}
		_, _ = w.Write(GroundJson)
	})

	fmt.Println("서버 시작 : 8080 포트")
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}
	}()
}

func ConnectServerSide() (bool, string) {
	fmt.Println("서버 url를 입력 : ")
	url, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	url = strings.TrimSpace(url)
	message := []byte("0,0")
	resp, err := http.Post(url+"/connect", "text/plain", bytes.NewBuffer(message))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	bodyS := string(body)

	if bodyS == "O" {
		return true, url
	} else {
		return false, url
	}
}

func GamingClientSide(url string, message string, attackSuccess bool) [22][12]blockShape.Block {
	if attackSuccess {
		message += ",1"
	} else {
		message += ",0"
	}
	resp, err := http.Post(url+"/gaming", "text/plain", bytes.NewBuffer([]byte(message)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var bodyS [22][12]blockShape.Block
	err = json.Unmarshal(body, &bodyS)
	if err != nil {
		panic(err)
	}

	return bodyS
}
