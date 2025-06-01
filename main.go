package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"math/rand"
	"time"

	"MultiTetris/blockShape"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	blockShape.InitEnv()
	blockShape.CreateBlockGroup(5, 1, 't')

	blockShape.PrintArray(blockShape.Ground)

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

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
