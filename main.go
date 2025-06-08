package main

import (
	"MultiTetris/blockShape"
	"MultiTetris/defense"
	"bytes"
	"fmt"
	"github.com/eiannone/keyboard"
	"io"
	"math/rand"
	"net"
	"net/http"
	"time"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	blockShape.InitEnv()
	blockShape.CreateBlockGroup(5, 1, 't')
	defense.InitDefense()
	blockShape.PrintArray(blockShape.Ground)
	//StartServerSide()
	ConnectServerSide()
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		err := keyboard.Close()
		if err != nil {

		}
	}()

	for {
		fmt.Print(":")
		ch, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEsc {
			break
		}

		blockShape.Move(ch)

		if ch != 'f' {
			blockShape.PrintArray(blockShape.Ground)
			continue
		}

		blockShape.FallingDown()
		blockShape.PrintArray(blockShape.Ground)
	}
}
func StartServerSide() {
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
func ConnectServerSide() {
	message := []byte("세계최고장한울")
	resp, err := http.Post("https://c1c4-221-151-14-28.ngrok-free.app/message", "text/plain", bytes.NewBuffer(message))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("서버 응답:", resp.Status)
}
