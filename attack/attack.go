package attack

import (
	"MultiTetris/blockShape"
	"fmt"
	"github.com/eiannone/keyboard"
)

// Coordinate 입력받을 좌표
type Coordinate struct {
	X int
	Y int
}

const CursorTypeId = 99

var cursorX, cursorY int = 0, 0

func Attack() {
	// 주석 좀 달아줘 코드에

	fmt.Println("GoGoGo")
	_ = keyboard.Open()

	for {
		//fmt.Print("\033[H")
		//fmt.Printf("\033[%d;0H", len(blockShape.Ground)+3)

		PrintGroundWithoutFallingBlock()
		ch, _, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("키 입력 에러:", err)
			break
		}

		switch ch {
		case 'w':
			if cursorX > 0 {
				cursorX--
			}
		case 's':
			if cursorX < len(blockShape.Ground)-1 {
				cursorX++
			}
		case 'a':
			if cursorY > 0 {
				cursorY--
			}
		case 'd':
			if cursorY < len(blockShape.Ground[0])-1 {
				cursorY++
			}
		case 'f':
			// 선택 완료! 커서 위치 반환하거나 처리
			if blockShape.Ground[cursorX][cursorY].Id == blockShape.FallingBlock.Id {
				//fmt.Printf("\033[%d;0H", len(blockShape.Ground))
				fmt.Println("수비 성공! : 대단하시네요 ")
				PrintGroundWithFallingBlock()
				blockShape.DeleteFallingBlock()
				return
			} else {
				//fmt.Printf("\033[%d;0H", len(blockShape.Ground))
				fmt.Println("공격에 실패하셨습니다\n 현재 블록의 위치:")
				PrintGroundWithFallingBlock()
				return
			}
		}
	}
}
func PrintGroundWithFallingBlock() {
	copyGround := blockShape.Ground
	copyGround[cursorX][cursorY] = blockShape.CursorBlock
	blockShape.PrintArray(copyGround)
}
func PrintGroundWithoutFallingBlock() {
	copyGround := blockShape.Ground // 현재 테트리스 판 복제해서 사용

	// FallingBlock인 블록들을 빈 블록으로 대체해버리기
	for i := 0; i < len(copyGround); i++ {
		for j := 0; j < len(copyGround[0]); j++ {
			if copyGround[i][j].Id == blockShape.FallingBlock.Id {
				copyGround[i][j] = blockShape.Block{}
			}
		}
	}
	//현재 선택중인 커서 위치 설정
	copyGround[cursorX][cursorY] = blockShape.CursorBlock
	blockShape.PrintArray(copyGround)
}

func CheckFallingBlock(x, y int) bool {
	return blockShape.Ground[x][y].Id == blockShape.FallingBlock.Id
}
