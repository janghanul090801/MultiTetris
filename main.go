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
	attack.Attack()
	fmt.Println("1 -> 서버 테스트 실행")
	ch, _, _ := keyboard.GetKey()
	if ch == '1' {
		soket.GetUrl()
		soket.StartServerSide() // <- 여기서 gaming 핸들러도 같이 등록됨

		isConnect, url := soket.ConnectServerSide()

		<-soket.WaitConnect // 서버로부터 connect 수신 대기

		// client-side simulation
		go func(isConnect bool, url string) {
			if isConnect {
				Ground := soket.GamingClientSide(url, "test-move")
				blockShape.PrintArray(Ground)
			}
		}(isConnect, url)

		blockShape.InitEnv()
		blockShape.PrintArray(blockShape.Ground)

		// 서버는 계속 대기
		select {}
	}
}
