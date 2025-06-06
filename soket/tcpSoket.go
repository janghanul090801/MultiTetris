package soket

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

//func main() {
//	//go StartServerSide()
//	//time.Sleep(1 * time.Second)
//	//ConnectServerSide()
//
//	StartServerSide()
//}

func StartServerSide() {
	//listener, err := net.Listen("tcp", "0.0.0.0:9000")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("서버 시작: 포트 9000")
	//
	//conn, err := listener.Accept()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("클라이언트 연결됨:", conn.RemoteAddr())
	//
	//defer conn.Close()
	//
	//buffer := make([]byte, 1024)
	//
	//for {
	//	count, err := conn.Read(buffer)
	//	if err != nil {
	//		if io.EOF == err {
	//			fmt.Println("연결 종료 : " + conn.RemoteAddr().String())
	//		} else {
	//			fmt.Printf("수신 실패 : %v\n", err)
	//		}
	//	}
	//	if 0 < count {
	//		data := buffer[:count]
	//		fmt.Println(string(data))
	//	}
	//
	//}
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, _ := ioutil.ReadAll(r.Body)
			fmt.Println("받은 메시지:", string(body))
			w.Write([]byte("메시지 수신 완료"))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("서버 시작 : 8080 포트")
	http.ListenAndServe(":8080", nil)
}
func ConnectServerSide() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("서버 아이피를 입력 : ")
	//ip, _ := reader.ReadString('\n')
	//ip = strings.TrimSpace(ip) // 개행 문자 제거
	//
	////go func() {
	////	time.Sleep(time.Second * 10)
	////	fmt.Println("연결 실패, 재실행하십시오")
	////	os.Exit(0)
	////}()
	//conn, err := net.Dial("tcp", ip+":9000")
	//if err != nil {
	//	panic(err)
	//} else {
	//	_, err = conn.Write([]byte("세계최고장한울"))
	//	if err == nil {
	//		fmt.Println("전송 성공")
	//	}
	//	time.Sleep(time.Millisecond * 100)
	//	conn.Close()
	//}
	message := []byte("세계최고장한울")
	resp, err := http.Post(" https://c1c4-221-151-14-28.ngrok-free.app/message", "text/plain", bytes.NewBuffer(message))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("서버 응답:", resp.Status)
}
