package attack

import (
	"MultiTetris/blockShape"
	"MultiTetris/user"
	"fmt"
	"strconv"
	"time"
)

// Coordinate 입력받을 좌표
type Coordinate struct {
	X int
	Y int
}

var Wasd = [8]Coordinate{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
	{1, -1},
	{1, 1},
	{-1, -1},
	{-1, 1},
}
var cursorX, cursorY int = 0, 0
var InputChan = make(chan rune)
var ErrChan = make(chan error)

func Attack(Ground [10][10]blockShape.Block, FallingBlock blockShape.BlockInfo, timeoutChan <-chan bool) (string, bool) {
	var inputTimer *time.Timer
	// 단일 키 입력 고루틴 (계속 유지)

	for {
		PrintGroundWithoutFallingBlock(Ground, FallingBlock)
		inputTimer = time.NewTimer(5 * time.Second)
		select {
		case <-timeoutChan:
			fmt.Println("\n전체 게임 시간 종료!")
			return strconv.Itoa(cursorX) + "," + strconv.Itoa(cursorY), false

		case err := <-ErrChan:
			fmt.Println("키 입력 에러:", err)
			return strconv.Itoa(cursorX) + "," + strconv.Itoa(cursorY), false

		case ch := <-InputChan:
			inputTimer.Stop()
			switch ch {
			case 'w':
				if isCanMoveCursor(Ground, FallingBlock, -1, 0) {
					cursorX--
				}
			case 's':
				if isCanMoveCursor(Ground, FallingBlock, 1, 0) {
					cursorX++
				}
			case 'a':
				if isCanMoveCursor(Ground, FallingBlock, 0, -1) {
					cursorY--
				}
			case 'd':
				if isCanMoveCursor(Ground, FallingBlock, 0, 1) {
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
func isCanMoveCursor(Ground [10][10]blockShape.Block, FallingBlock blockShape.BlockInfo, dx, dy int) bool {
	newX := cursorX + dx
	newY := cursorY + dy

	if newX < 0 || newX >= len(Ground) || newY < 0 || newY >= len(Ground[0]) {
		return false
	}

	// 블록 존재 여부 체크

	for _, dir := range Wasd {
		checkX := newX + dir.X
		checkY := newY + dir.Y

		// 경계 내에 있는 경우에만 체크
		if checkX >= 0 && checkX < len(Ground) && checkY >= 0 && checkY < len(Ground[0]) {
			// 블록이 있고 떨어지는 블록이 아닌 경우
			if Ground[checkX][checkY].Id != 0 && Ground[checkX][checkY].Id != FallingBlock.Id {
				return false
			}
		}
	}
	return Ground[newX][newY].Id == 0 || Ground[newX][newY].Id == FallingBlock.Id
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
