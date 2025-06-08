package defense

import (
	"MultiTetris/blockShape"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Coordinate 입력받을 좌표
type Coordinate struct {
	X int
	Y int
}

func InitDefense() {
	// 주석 좀 달아줘 코드에

	fmt.Println("GoGoGo")
	PrintGroundWithoutFallingBlock() // FallingBlock 숨긴 Ground 출력
	pos := GetDefenseCoordinate()    // 입력받기
	if CheckFallingBlock(pos.X, pos.Y) {
		fmt.Println("수비 성공! : 어케 찾음?👍👍👍😁")
		blockShape.DeleteFallingBlock()
	}
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

	blockShape.PrintArray(copyGround)
}
func GetDefenseCoordinate() Coordinate {
	reader := bufio.NewReader(os.Stdin) //입력 이게 맞나
	for {
		fmt.Print("수비할 좌표를 입력하세요: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("흠? 뭐지 이건 :", err)
			continue
		}

		input = strings.TrimSpace(input) // 입력하기, 근데 문자열
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("2개만 입력해라 : ")
			continue
		}

		//정수로 변환
		x, errX := strconv.Atoi(parts[0])
		y, errY := strconv.Atoi(parts[1])
		if errX != nil || errY != nil {
			fmt.Println("버그 발견! : ¥¥¥ 이상한 형식의 입력 발생 ¥¥¥")
			continue
		}
		return Coordinate{X: x, Y: y}
	}
}
func CheckFallingBlock(x, y int) bool {
	return blockShape.Ground[x][y].Id == blockShape.FallingBlock.Id
}
