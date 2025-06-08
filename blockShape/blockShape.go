package blockShape

import (
	"fmt"
	"github.com/fatih/color"
)

type Block struct {
	Id        int  // 블럭 연결을 인식하기 위한 아이디
	typeId    int  // 블럭 모양을 인식하기 위한 아이디
	isFalling bool // 떨어지는 중인지 확인
	isNotNone bool // nil이 아닌지 == 블럭이 차 있는지
	direction rune // 상t하b좌l우r
}

type blockInfo struct {
	Id        int
	i         int // 해당 블럭의 가장 왼쪽(j), 가장 위쪽(i)
	j         int
	blockType int
	direction rune
}

// 지금 떨어지고 있는 블럭의 정보

// 이건 걍 비어있는거 표시하는거
var NoneBlock = Block{0, 0, false, false, 't'}

// 블럭에 아이디 부여하기 위한거
var GlobalBlockId = 1

// 지금 떨어지고 있는 블럭의 정보 객체
var FallingBlock blockInfo

// 게임판
var Ground [22][12]Block

// 블럭 하나 만드는건데 귀찮아서 매크로한거임
func CreateBlock(typeId int) Block {
	b := Block{0, typeId, true, true, 't'}
	return b
}

// 블럭 타입 아이디 별로 컬러 지정한거임
var blockColorMap = map[int]*color.Color{
	0: color.New(color.BgBlack),
	1: color.New(color.BgCyan),
	2: color.New(color.BgHiRed),
	3: color.New(color.BgMagenta),
	4: color.New(color.BgYellow),
	5: color.New(color.BgBlue),
	6: color.New(color.BgGreen),
	7: color.New(color.BgRed),
}

// [타입아이디-1][direction]으로 넣으면 블럭모양 [4][4] 로 가져올 수 있음
var blockShapeList [7]map[rune][4][4]Block

// 게임판 출력해주는거
func PrintArray(a [22][12]Block) {
	// screen.Clear()
	fmt.Println()
	for _, i := range a {
		for _, j := range i {
			_, _ = blockColorMap[j.typeId].Print("  ")

		}
		fmt.Println()
	}

}

// 게임 시작하기 전 환경설정
func InitEnv() {
	color.NoColor = false

	for i := range 7 {
		blockShapeList[i] = make(map[rune][4][4]Block)
	}

	blockShapeList[0]['t'] = [4][4]Block{
		{},
		{CreateBlock(1), CreateBlock(1), CreateBlock(1), CreateBlock(1)},
	}
	blockShapeList[0]['r'] = [4][4]Block{
		{NoneBlock, NoneBlock, CreateBlock(1)},
		{NoneBlock, NoneBlock, CreateBlock(1)},
		{NoneBlock, NoneBlock, CreateBlock(1)},
		{NoneBlock, NoneBlock, CreateBlock(1)},
	}
	blockShapeList[0]['b'] = [4][4]Block{
		{},
		{},
		{CreateBlock(1), CreateBlock(1), CreateBlock(1), CreateBlock(1)},
	}
	blockShapeList[0]['l'] = [4][4]Block{
		{NoneBlock, CreateBlock(1)},
		{NoneBlock, CreateBlock(1)},
		{NoneBlock, CreateBlock(1)},
		{NoneBlock, CreateBlock(1)},
	}

	blockShapeList[1]['t'] = [4][4]Block{
		{CreateBlock(2)},
		{CreateBlock(2), CreateBlock(2), CreateBlock(2)},
	}
	blockShapeList[1]['r'] = [4][4]Block{
		{NoneBlock, CreateBlock(2), CreateBlock(2)},
		{NoneBlock, CreateBlock(2)},
		{NoneBlock, CreateBlock(2)},
	}
	blockShapeList[1]['b'] = [4][4]Block{
		{},
		{CreateBlock(2), CreateBlock(2), CreateBlock(2)},
		{NoneBlock, NoneBlock, CreateBlock(2)},
	}
	blockShapeList[1]['l'] = [4][4]Block{
		{NoneBlock, CreateBlock(2)},
		{NoneBlock, CreateBlock(2)},
		{CreateBlock(2), CreateBlock(2)},
	}

	blockShapeList[2]['t'] = [4][4]Block{
		{NoneBlock, CreateBlock(3)},
		{CreateBlock(3), CreateBlock(3), CreateBlock(3)},
	}
	blockShapeList[2]['r'] = [4][4]Block{
		{NoneBlock, CreateBlock(3)},
		{NoneBlock, CreateBlock(3), CreateBlock(3)},
		{NoneBlock, CreateBlock(3)},
	}
	blockShapeList[2]['b'] = [4][4]Block{
		{},
		{CreateBlock(3), CreateBlock(3), CreateBlock(3)},
		{NoneBlock, CreateBlock(3)},
	}
	blockShapeList[2]['l'] = [4][4]Block{
		{NoneBlock, CreateBlock(3)},
		{CreateBlock(3), CreateBlock(3)},
		{NoneBlock, CreateBlock(3)},
	}

	blockShapeList[3]['t'] = [4][4]Block{
		{CreateBlock(4), CreateBlock(4)},
		{CreateBlock(4), CreateBlock(4)},
	}
	blockShapeList[3]['r'] = [4][4]Block{
		{CreateBlock(4), CreateBlock(4)},
		{CreateBlock(4), CreateBlock(4)},
	}
	blockShapeList[3]['b'] = [4][4]Block{
		{CreateBlock(4), CreateBlock(4)},
		{CreateBlock(4), CreateBlock(4)},
	}
	blockShapeList[3]['l'] = [4][4]Block{
		{CreateBlock(4), CreateBlock(4)},
		{CreateBlock(4), CreateBlock(4)},
	}

	blockShapeList[4]['t'] = [4][4]Block{
		{NoneBlock, NoneBlock, CreateBlock(5)},
		{CreateBlock(5), CreateBlock(5), CreateBlock(5)},
	}
	blockShapeList[4]['r'] = [4][4]Block{
		{NoneBlock, CreateBlock(5)},
		{NoneBlock, CreateBlock(5)},
		{NoneBlock, CreateBlock(5), CreateBlock(5)},
	}
	blockShapeList[4]['b'] = [4][4]Block{
		{},
		{CreateBlock(5), CreateBlock(5), CreateBlock(5)},
		{CreateBlock(5)},
	}
	blockShapeList[4]['l'] = [4][4]Block{
		{CreateBlock(5), CreateBlock(5)},
		{NoneBlock, CreateBlock(5)},
		{NoneBlock, CreateBlock(5)},
	}

	// S자 블록 (6)
	blockShapeList[5]['t'] = [4][4]Block{
		{NoneBlock, CreateBlock(6), CreateBlock(6)},
		{CreateBlock(6), CreateBlock(6)},
	}
	blockShapeList[5]['r'] = [4][4]Block{
		{NoneBlock, CreateBlock(6)},
		{NoneBlock, CreateBlock(6), CreateBlock(6)},
		{NoneBlock, NoneBlock, CreateBlock(6)},
	}
	blockShapeList[5]['b'] = [4][4]Block{
		{},
		{NoneBlock, CreateBlock(6), CreateBlock(6)},
		{CreateBlock(6), CreateBlock(6)},
	}
	blockShapeList[5]['l'] = [4][4]Block{
		{CreateBlock(6)},
		{CreateBlock(6), CreateBlock(6)},
		{NoneBlock, CreateBlock(6)},
	}

	// Z자 블록 (7)
	blockShapeList[6]['t'] = [4][4]Block{
		{CreateBlock(7), CreateBlock(7)},
		{NoneBlock, CreateBlock(7), CreateBlock(7)},
	}
	blockShapeList[6]['r'] = [4][4]Block{
		{NoneBlock, NoneBlock, CreateBlock(7)},
		{NoneBlock, CreateBlock(7), CreateBlock(7)},
		{NoneBlock, CreateBlock(7)},
	}
	blockShapeList[6]['b'] = [4][4]Block{
		{},
		{CreateBlock(7), CreateBlock(7)},
		{NoneBlock, CreateBlock(7), CreateBlock(7)},
	}
	blockShapeList[6]['l'] = [4][4]Block{
		{NoneBlock, CreateBlock(7)},
		{CreateBlock(7), CreateBlock(7)},
		{CreateBlock(7)},
	}
}

// 해당 블럭 모양의 가장 오른쪽을 리턴
func GetRightestByBlockType(blockType int, d rune) int {
	blockShape := blockShapeList[blockType-1][d]
	var rightest = 0
	for i := range 4 {
		for j := range 4 {
			if blockShape[i][j].typeId != 0 {
				if j > rightest {
					rightest = j
				}
			}
		}
	}
	return rightest
}

// 해당 블럭 모양의 가장 아래쪽을 리턴
func GetBottomestByBlockType(blockType int, d rune) int {
	if blockType == 0 {
		return 0
	}
	blockShape := blockShapeList[blockType-1][d]
	var bottomest = 0
	for i := range 4 {
		for j := range 4 {
			if blockShape[i][j].typeId != 0 {
				if i > bottomest {
					bottomest = i
				}
			}
		}
	}
	return bottomest
}

// 해당 블럭 모양을 만듦
func CreateBlockGroup(j int, blockType int, d rune) {
	count := 0
	iiP := 0
	for iP := range 4 {
		for jP := range 4 {
			Ground[0+iP+iiP][j+jP] = blockShapeList[blockType-1][d][iP][jP]
			Ground[0+iP+iiP][j+jP].Id = GlobalBlockId
			if blockShapeList[blockType-1][d][iP][jP].typeId != 0 {
				count++
			}
		}
		if count == 0 {
			iiP--
		}
	}
	FallingBlock.Id = GlobalBlockId
	GlobalBlockId++
	FallingBlock.i = 0
	FallingBlock.j = j
	FallingBlock.blockType = blockType
	FallingBlock.direction = d
}

// 지금 떨어지고 있는 블럭 한칸 아래로 떨어지는거
func FallingDown() {
	iStart := FallingBlock.i
	jStart := FallingBlock.j

	for i := 3; i >= 0; i-- {
		for j := 0; j < 4; j++ {
			curI := iStart + i
			curJ := jStart + j
			if curI >= len(Ground) || curJ >= len(Ground[0]) {
				continue
			}
			if Ground[curI][curJ].Id == FallingBlock.Id {
				belowI := curI + 1
				if belowI >= len(Ground) {
					continue
				}
				if Ground[belowI][curJ].typeId != 0 && GetBottomestByBlockType(FallingBlock.blockType, FallingBlock.direction)+belowI >= len(Ground) {
					SetBlocksIsFallingFalse(FallingBlock.Id)
					return
				}
			}
			if Ground[curI][curJ].Id == FallingBlock.Id {
				Ground[curI+1][curJ] = Ground[curI][curJ]
				Ground[curI][curJ] = NoneBlock
			}
		}
	}

	FallingBlock.i++
}

// 바닥에 떨어지면 fallingBlock에서 삭제하는거
// 이거 새 블럭 만들어서 fallingBlock으로 지정해주는 로직 필요함
func SetBlocksIsFallingFalse(blockId int) {
	count := 0
	startI := FallingBlock.i
	startJ := FallingBlock.j

	for i := startI; i < startI+4 && i < len(Ground); i++ {
		for j := startJ; j < startJ+4 && j < len(Ground[0]); j++ {
			if Ground[i][j].Id == blockId {
				Ground[i][j].isFalling = false
				count++
				if count == 4 {
					FallingBlock = blockInfo{}
					return
				}
			}
		}
	}
}

// wasd 받아서 해당 방향으로 이동
func Move(q rune) {
	if FallingBlock.Id == 0 {
		return
	}
	rightest := GetRightestByBlockType(FallingBlock.blockType, FallingBlock.direction)
	switch q {
	case 'd':

		if FallingBlock.j+rightest >= len(Ground[0]) {
			return
		}

		for i := 0; i < 4; i++ {
			for j := rightest; j >= 0; j-- {
				if FallingBlock.j+rightest+1 >= len(Ground[0]) {
					return
				}
				b := Ground[FallingBlock.i+i][FallingBlock.j+j]
				if b.Id == FallingBlock.Id {
					if Ground[FallingBlock.i+i][FallingBlock.j+j+1].isNotNone &&
						Ground[FallingBlock.i+i][FallingBlock.j+j+1].Id != FallingBlock.Id {
						return
					}
				}
			}
		}

		// 오른쪽으로 한 칸 이동 (오른쪽 끝부터 복사)
		for i := 0; i < 4; i++ {
			for j := rightest; j >= 0; j-- {
				b := Ground[FallingBlock.i+i][FallingBlock.j+j]
				if b.Id == FallingBlock.Id {
					Ground[FallingBlock.i+i][FallingBlock.j+j+1] = b
					Ground[FallingBlock.i+i][FallingBlock.j+j] = NoneBlock
				}
			}
		}
		FallingBlock.j++

	case 'a': // 왼쪽 이동
		// 왼쪽 경계 체크
		if FallingBlock.j <= 0 {
			return
		}

		// 왼쪽 칸이 비어있는지 확인 (충돌 방지)
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if FallingBlock.j+j >= len(Ground[0]) {
					break
				}
				b := Ground[FallingBlock.i+i][FallingBlock.j+j]

				if b.Id == FallingBlock.Id {
					if Ground[FallingBlock.i+i][FallingBlock.j+j-1].isNotNone &&
						Ground[FallingBlock.i+i][FallingBlock.j+j-1].Id != FallingBlock.Id {
						return
					}
				}
			}
		}

		// 왼쪽으로 한 칸 이동 (왼쪽 끝부터 복사)
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if FallingBlock.j+j >= len(Ground[0]) {
					break
				}
				b := Ground[FallingBlock.i+i][FallingBlock.j+j]
				if b.Id == FallingBlock.Id {
					Ground[FallingBlock.i+i][FallingBlock.j+j-1] = b
					Ground[FallingBlock.i+i][FallingBlock.j+j] = NoneBlock
				}
			}
		}
		FallingBlock.j--
	}
}
func DeleteFallingBlock() { //수비 성공시 FallingBlock 삭제
	for i := FallingBlock.i; i < FallingBlock.i+4 && i < len(Ground); i++ {
		for j := FallingBlock.j; j < FallingBlock.j+4 && j < len(Ground[0]); j++ {
			if Ground[i][j].Id == FallingBlock.Id {
				Ground[i][j] = NoneBlock
			}
		}
	}
	FallingBlock = blockInfo{} // FallingBlock 초기화
}
