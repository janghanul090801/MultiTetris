package main

import (
	"MultiTetris/blockShape"
	"fmt"
	"github.com/eiannone/keyboard"
)

func main() {

	blockShape.InitEnv()
	blockShape.CreateBlockGroup(5, 3, 't')
	blockShape.PrintArray(blockShape.Ground)
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

		// defense.InitDefense()
		blockShape.FallingDown()
		blockShape.PrintArray(blockShape.Ground)
	}
}
