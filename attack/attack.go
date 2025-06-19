package attack

import (
	"MultiTetris/blockShape"
	"MultiTetris/user"
	"fmt"
	"github.com/eiannone/keyboard"
)

// Coordinate 입력받을 좌표
type Coordinate struct {
	X int
	Y int
}

var cursorX, cursorY int = 0, 0

func Attack(Ground [22][12]blockShape.Block) bool {

	fmt.Println("GoGoGo")
	_ = keyboard.Open()

	for {
		//fmt.Print("\033[H")
		//fmt.Printf("\033[%d;0H", len(blockShape.Ground)+3)

		PrintGroundWithoutFallingBlock(Ground)
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
			if cursorX < len(Ground)-1 {
				cursorX++
			}
		case 'a':
			if cursorY > 0 {
				cursorY--
			}
		case 'd':
			if cursorY < len(Ground[0])-1 {
				cursorY++
			}
		case 'f':
			// 선택 완료 시 처리
			if CheckFallingBlock(cursorX, cursorY, Ground) {
				//fmt.Printf("\033[%d;0H", len(blockShape.Ground))
				fmt.Println("수비 성공! : 대단하시네요 ")
				PrintGroundWithFallingBlock(Ground)
				// blockShape.DeleteFallingBlock() 이거 서버로 옮기기
				user.Me.AttackSuccess()
				return true
			} else {
				//fmt.Printf("\033[%d;0H", len(blockShape.Ground))
				fmt.Println("공격에 실패하셨습니다\n 현재 블록의 위치:")
				PrintGroundWithFallingBlock(Ground)
				user.Me.AttackFailed()

				return false
			}
		}
	}
	return false
}
func PrintGroundWithFallingBlock(Ground [22][12]blockShape.Block) {
	copyGround := Ground
	copyGround[cursorX][cursorY] = blockShape.CursorBlock
	blockShape.PrintArray(copyGround)
}
func PrintGroundWithoutFallingBlock(Ground [22][12]blockShape.Block) {
	copyGround := Ground // 현재 테트리스 판 복제해서 사용

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

func CheckFallingBlock(x, y int, Ground [22][12]blockShape.Block) bool {
	return Ground[x][y].Id == blockShape.FallingBlock.Id
}
