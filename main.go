package main

import (
	"MultiTetris/attack"
	"MultiTetris/blockShape"
	"MultiTetris/soket"
	"fmt"
	"github.com/eiannone/keyboard"
)

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	blockShape.InitEnv()
	blockShape.PrintArray(blockShape.Ground)
	fmt.Println("1 -> 서버 테스트 실행")
	ch, _, _ := keyboard.GetKey()
	if ch == '1' {
		soket.GetUrl()
		soket.StartServerSide() // <- 여기서 gaming 핸들러도 같이 등록됨

		blockShape.InitEnv()

		<-soket.WaitConnect // 서버로부터 connect 수신 대기

		blockShape.PrintArray(blockShape.Ground)

		// 서버는 계속 대기
		select {}
	} else if ch == '2' {
		isConnected, url := soket.ConnectServerSide()

		if isConnected {
			Ground := soket.GamingClientSide(url, "", false)
			for {
				b := attack.Attack(Ground)
				soket.GamingClientSide(url, "", b)
			}
		}

	}
}
