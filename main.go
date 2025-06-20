package main

import (
	"MultiTetris/attack"
	"MultiTetris/blockShape"
	"MultiTetris/soket"
	"MultiTetris/user"
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
	"os"
	"time"
)

func main() {
	fmt.Println("1 -> 서버 테스트 실행")
	var ch string
	fmt.Scanln(&ch)
	if ch == "1" {
		if err := keyboard.Open(); err != nil {
			panic(err)
		}
		defer keyboard.Close()
		soket.GetUrl()
		soket.StartServerSide()

		blockShape.InitEnv()

		<-soket.WaitConnect // 서버로부터 connect 수신 대기
		go func() {
			time.Sleep(3 * time.Minute)

			<-soket.WaitEnd
			user.Me.PrintScore()
			user.Other.PrintScore()
			os.Exit(0)
		}()
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
			gameState := soket.GamingClientSide(url, ",", false)
			go func() {
				time.Sleep(3 * time.Minute)
				soket.SendUserData(url, user.Me)
				user.Me.PrintScore()
				user.Other.PrintScore()
				os.Exit(0)
			}()
			for {
				s, b := attack.Attack(gameState.Ground, gameState.FallingBlock)
				fmt.Println("서버턴 대기중....")
				gameState = soket.GamingClientSide(url, s, b)
				screen.Clear()

				<-soket.WaitClient
			}
		}

	}
}
