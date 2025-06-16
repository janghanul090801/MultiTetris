package blockShape

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/inancgumus/screen"
	"math"
	"math/rand"
	"time"
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

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

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

// 회전 관련
var RotationRune = [4]rune{'t', 'r', 'b', 'l'}
var globalRotateNum = 0

// 게임판 출력해주는거
func PrintArray(a [22][12]Block) {
	screen.Clear()
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

func GetLeftestByBlockType(blockType int, d rune) int {
	if blockType == 0 {
		return 0
	}
	blockShape := blockShapeList[blockType-1][d]
	var leftest = 3
	for i := range 4 {
		for j := range 4 {
			if blockShape[i][j].typeId != 0 {
				if j < leftest {
					leftest = j
				}
			}
		}
	}

	return leftest
}

func GetTopestByBlockType(blockType int, d rune) int {
	if blockType == 0 {
		return 0
	}
	blockShape := blockShapeList[blockType-1][d]
	var topest = 3
	for i := range 4 {
		for j := range 4 {
			if blockShape[i][j].typeId != 0 {
				if i < topest {
					topest = i
				}
			}
		}
	}

	return topest
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

func DrawFallingBlock() {
	block := FallingBlock
	blockShape := blockShapeList[block.blockType-1][block.direction]
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if blockShape[i][j].typeId != 0 {
				ii := block.i + i
				jj := block.j + j
				if ii >= 0 && ii < len(Ground) && jj >= 0 && jj < len(Ground[0]) {
					Ground[ii][jj] = Block{
						typeId:    block.blockType,
						isFalling: true,
						Id:        block.Id,
					}
				}
			}
		}
	}
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
			if curJ < 0 {
				curJ = 0
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

OUT:
	for i := startI; i < startI+4 && i < len(Ground); i++ {
		for j := startJ; j < startJ+4 && j < len(Ground[0]); j++ {
			if j < 0 {
				j = 0
			}
			if Ground[i][j].Id == blockId {
				Ground[i][j].isFalling = false
				count++
				if count == 4 {
					FallingBlock = blockInfo{}
					break OUT
				}
			}
		}
	}

	CreateBlockGroup(5, 3, 't')
	CheckLine()
}

func EraseBlock() {
	block := FallingBlock
	blockShape := blockShapeList[block.blockType-1][block.direction]
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if blockShape[i][j].typeId != 0 {
				ii := block.i + i
				jj := block.j + j
				if ii >= 0 && ii < len(Ground) && jj >= 0 && jj < len(Ground[0]) {
					if Ground[ii][jj].Id == block.Id {
						Ground[ii][jj] = NoneBlock
					}
				}
			}
		}
	}
}

func CanMove(q int) bool {
	block := FallingBlock
	shape := blockShapeList[block.blockType-1][block.direction]
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if shape[i][j].typeId != 0 {
				ii := block.i + i
				jj := block.j + j + q
				if jj < 0 || jj >= len(Ground[0]) {
					return false
				}
				if ii >= 0 && Ground[ii][jj].typeId != 0 && Ground[ii][jj].Id != block.Id {
					return false
				}
			}
		}
	}
	return true
}

// wasd 받아서 해당 방향으로 이동
func Move(q rune) {
	if FallingBlock.Id == 0 {
		return
	}
	switch q {
	case 'd':
		if CanMove(1) {
			EraseBlock()
			FallingBlock.j++
			DrawFallingBlock()
		}
	case 'a': // 왼쪽 이동
		if CanMove(-1) {
			EraseBlock()
			FallingBlock.j--
			DrawFallingBlock()
		}
	case 'e':
		RotateBlock(1)
	case 'q':
		RotateBlock(-1)
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

func CanRotate(q int) bool {
	block := FallingBlock
	direction := RotationRune[int(math.Abs(float64(globalRotateNum+q)))%4]
	shape := blockShapeList[block.blockType-1][direction]
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if shape[i][j].typeId != 0 {
				ii := block.i + i
				jj := block.j + j
				if jj < 0 || jj >= len(Ground[0]) || ii >= len(Ground) {
					return false
				}
				if ii >= 0 && Ground[ii][jj].typeId != 0 && Ground[ii][jj].Id != block.Id {
					return false
				}
			}
		}
	}
	return true
}

func RotateBlock(q int) {
	if CanRotate(q) {
		EraseBlock()
		FallingBlock.direction = RotationRune[int(math.Abs(float64(globalRotateNum+q)))%4]
		globalRotateNum++
		DrawFallingBlock()
	}
}

func IsLimited(i int) bool {
	for j := range len(Ground[0]) {
		if Ground[i][j].typeId == 0 {
			return false
		}
	}
	return true
}

func ClearLine(ii int) {
	for i := ii; i > 0; i-- {
		for j := 0; j < len(Ground[i]); j++ {
			Ground[i][j] = Ground[i-1][j]
		}
	}
	for j := 0; j < len(Ground[0]); j++ {
		Ground[0][j] = Block{typeId: 0}
	}
}

func CheckLine() {
	for i := len(Ground) - 1; i >= 0; i-- {
		if IsLimited(i) {
			ClearLine(i)
			i++
		}
	}
}
