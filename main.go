package main

import (
	"MultiTetris/attack"
	"MultiTetris/blockShape"
	"MultiTetris/soket"
	"MultiTetris/user"
	"fmt"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
)

func main() {
	fmt.Println("1 -> 서버, 2 -> 클라이언트")
	var ch string
	fmt.Scanln(&ch)
	user.Me.Status = ch
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
			for {
				ch, key, err := keyboard.GetKey()
				if err != nil {
					soket.ErrChan <- err
					continue
				}
				soket.InputChan <- ch
				soket.KeyChan <- key
			}
		}()
		go func() {
			time.Sleep(3 * time.Minute)

			<-soket.WaitEnd
			fmt.Println("전체 게임 시간 종료!")
			user.Me.PrintScore()
			user.Other.PrintScore()
			os.Exit(0)
		}()
		blockShape.PrintArray(blockShape.Ground)

		// 서버는 계속 대기
		select {}
	} else if ch == "2" {
		isConnected, url := soket.ConnectServerSide()

		if isConnected {
			if err := keyboard.Open(); err != nil {
				panic(err)
			}
			defer keyboard.Close()

			go func() {
				for {
					ch, _, err := keyboard.GetKey()
					if err != nil {
						attack.ErrChan <- err
						return
					}
					attack.InputChan <- ch
				}
			}()
			gameState := soket.GamingClientSide(url, ",", false)
			timeoutChan := make(chan bool)

			go func() {
				time.Sleep(3 * time.Minute)
				timeoutChan <- true
				soket.SendUserData(url, user.Me)
				user.Me.PrintScore()
				user.Other.PrintScore()
				os.Exit(0)
			}()
			for {
				s, b := attack.Attack(gameState.Ground, gameState.FallingBlock, timeoutChan)
				fmt.Println("서버턴 대기중....")
				gameState = soket.GamingClientSide(url, s, b)
				screen.Clear()

				<-soket.WaitClient
			}
		}

	}
}
