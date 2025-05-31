package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	go startServerSide()
	time.Sleep(time.Second)
	connectServerSide()
}

func startServerSide() {
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		panic(err)
	}

	fmt.Println("서버 시작: 포트 9000")

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	fmt.Println("클라이언트 연결됨:", conn.RemoteAddr())

	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		count, err := conn.Read(buffer)
		if err != nil {
			if io.EOF == err {
				fmt.Println("연결 종료 : " + conn.RemoteAddr().String())
			} else {
				fmt.Printf("수신 실패 : %v\n", err)
			}
		}
		if 0 < count {
			data := buffer[:count]
			fmt.Println(string(data))
		}

	}
}
func connectServerSide() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("서버 아이피를 입력 : ")
	ip, _ := reader.ReadString('\n')
	ip = strings.TrimSpace(ip) // 개행 문자 제거

	conn, err := net.Dial("tcp", ":9000")
	if err != nil {
		panic(err)
	} else {
		_, err = conn.Write([]byte("세계최고장한울"))
		if err == nil {
			fmt.Println("전송 성공")
		}
		time.Sleep(time.Millisecond * 100)
		conn.Close()
	}
}
