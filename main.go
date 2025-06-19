package main

import (
	"MultiTetris/attack"
	"MultiTetris/blockShape"
	"MultiTetris/soket"
	"fmt"
	"github.com/eiannone/keyboard"
)

func main() {
	blockShape.InitEnv()
	blockShape.PrintArray(blockShape.Ground)
	fmt.Println("1 -> 서버 테스트 실행")
	var ch string
	fmt.Scanln(&ch)
	if ch == "1" {
		if err := keyboard.Open(); err != nil {
			panic(err)
		}
		defer keyboard.Close()
		soket.GetUrl()
		soket.StartServerSide() // <- 여기서 gaming 핸들러도 같이 등록됨

		blockShape.InitEnv()

		<-soket.WaitConnect // 서버로부터 connect 수신 대기

		blockShape.PrintArray(blockShape.Ground)

		// 서버는 계속 대기
		select {}
	} else if ch == "2" {
		isConnected, url := soket.ConnectServerSide()

		if err := keyboard.Open(); err != nil {
			panic(err)
		}
		defer keyboard.Close()

		if isConnected {
			Ground := soket.GamingClientSide(url, "_", false)
			for {
				s, b := attack.Attack(Ground)
				soket.GamingClientSide(url, s, b)
				<-soket.WaitConnect
			}
		}

	}
}
