package attack

import (
	"MultiTetris/blockShape"
	"MultiTetris/user"
	"fmt"
	"github.com/eiannone/keyboard"
	"strconv"
	"time"
)

// Coordinate 입력받을 좌표
type Coordinate struct {
	X int
	Y int
}

var cursorX, cursorY int = 0, 0

func Attack(Ground [10][10]blockShape.Block, FallingBlock blockShape.BlockInfo, timeoutChan <-chan bool) (string, bool) {
	inputChan := make(chan rune)
	errChan := make(chan error)
	var inputTimer *time.Timer
	// 단일 키 입력 고루틴 (계속 유지)
	go func() {
		for {
			ch, _, err := keyboard.GetKey()
			if err != nil {
				errChan <- err
				return
			}
			inputChan <- ch
		}
	}()

	for {
		PrintGroundWithoutFallingBlock(Ground, FallingBlock)
		inputTimer = time.NewTimer(5 * time.Second)
		select {
		case <-timeoutChan:
			fmt.Println("\n전체 게임 시간 종료!")
			return strconv.Itoa(cursorX) + "," + strconv.Itoa(cursorY), false

		case err := <-errChan:
			fmt.Println("키 입력 에러:", err)
			return strconv.Itoa(cursorX) + "," + strconv.Itoa(cursorY), false

		case ch := <-inputChan:
			inputTimer.Stop()
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
				return isAttackSuccessful(Ground, FallingBlock)
			}

		case <-inputTimer.C:
			fmt.Println("\n입력 시간 초과 : 자동 턴 종료")
			return isAttackSuccessful(Ground, FallingBlock)
		}
	}
}
func isAttackSuccessful(Ground [10][10]blockShape.Block, FallingBlock blockShape.BlockInfo) (string, bool) {
	if CheckFallingBlock(cursorX, cursorY, Ground, FallingBlock) {
		fmt.Println("공격 성공! : 대단하시네요 ")
		PrintGroundWithFallingBlock(Ground)
		user.Me.AttackSuccess()
		return strconv.Itoa(cursorX) + "," + strconv.Itoa(cursorY), true
	} else {
		fmt.Println("공격에 실패하셨습니다\n 현재 블록의 위치:")
		PrintGroundWithFallingBlock(Ground)
		user.Me.AttackFailed()
		return strconv.Itoa(cursorX) + "," + strconv.Itoa(cursorY), false
	}
}
func PrintGroundWithFallingBlock(Ground [10][10]blockShape.Block) {
	copyGround := Ground
	copyGround[cursorX][cursorY] = blockShape.CursorBlock
	blockShape.PrintArray(copyGround)
}
func PrintGroundWithoutFallingBlock(Ground [10][10]blockShape.Block, FallingBlock blockShape.BlockInfo) {
	copyGround := Ground // 현재 테트리스 판 복제해서 사용

	// FallingBlock인 블록들을 빈 블록으로 대체해버리기
	for i := 0; i < len(copyGround); i++ {
		for j := 0; j < len(copyGround[0]); j++ {
			if copyGround[i][j].Id == FallingBlock.Id {
				copyGround[i][j] = blockShape.Block{}
			}
		}
	}
	//현재 선택중인 커서 위치 설정
	copyGround[cursorX][cursorY] = blockShape.CursorBlock
	blockShape.PrintArray(copyGround)

}

func CheckFallingBlock(x, y int, Ground [10][10]blockShape.Block, FallingBlock blockShape.BlockInfo) bool {
	return Ground[x][y].Id == FallingBlock.Id
}
