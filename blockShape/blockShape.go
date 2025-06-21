package blockShape

import (
	"MultiTetris/user"
	"fmt"
	"github.com/fatih/color"
	"github.com/inancgumus/screen"
	"math/rand"
	"time"
)

type Block struct {
	Id        int  `json:"id"`        // 블럭 연결을 인식하기 위한 아이디
	TypeId    int  `json:"typeId"`    // 블럭 모양을 인식하기 위한 아이디
	IsFalling bool `json:"isFalling"` // 떨어지는 중인지 확인
	IsNotNone bool `json:"isNotNone"` // nil이 아닌지 == 블럭이 차 있는지
	Direction rune `json:"direction"` // 상t하b좌l우r
}

type BlockInfo struct {
	Id        int  `json:"id"`
	I         int  `json:"i"` // 해당 블럭의 가장 위쪽(i)
	J         int  `json:"j"` // 가장 왼쪽(j)
	BlockType int  `json:"blockType"`
	Direction rune `json:"direction"`
}

// 지금 떨어지고 있는 블럭의 정보

// 이건 걍 비어있는거 표시하는거
var NoneBlock = Block{0, 0, false, false, 't'}
var CursorBlock = Block{-1, -1, false, false, 't'}

// 블럭에 아이디 부여하기 위한거
var GlobalBlockId = 1

// 지금 떨어지고 있는 블럭의 정보 객체
var FallingBlock BlockInfo

// 게임판
var Ground [10][10]Block

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// 블럭 하나 만드는건데 귀찮아서 매크로한거임
func CreateBlock(typeId int) Block {
	b := Block{0, typeId, true, true, 't'}
	return b
}

// 블럭 타입 아이디 별로 컬러 지정한거임
var blockColorMap = map[int]*color.Color{
	-1: color.New(color.BgRed),
	0:  color.New(color.BgHiBlack),
	1:  color.New(color.BgCyan),
	2:  color.New(color.BgHiRed),
	3:  color.New(color.BgMagenta),
	4:  color.New(color.BgYellow),
	5:  color.New(color.BgBlue),
	6:  color.New(color.BgGreen),
	7:  color.New(color.BgRed),
}

// [타입아이디-1][direction]으로 넣으면 블럭모양 [4][4] 로 가져올 수 있음
var blockShapeList [7]map[rune][4][4]Block

// 게임판 출력해주는거
func PrintArray(a [10][10]Block) {
	screen.Clear()
	fmt.Println()
	for _, i := range a {
		for _, j := range i {

			_, _ = blockColorMap[j.TypeId].Print("  ")

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

	CreateBlockGroup(4, rand.Intn(6)+1, 't')
}

// 해당 블럭 모양을 만듦
func CreateBlockGroup(j int, blockType int, d rune) {
	count := 0
	iiP := 0
	for iP := range 4 {
		for jP := range 4 {
			if blockShapeList[blockType-1][d][iP][jP].TypeId != 0 {
				Ground[0+iP+iiP][j+jP] = blockShapeList[blockType-1][d][iP][jP]
				Ground[0+iP+iiP][j+jP].Id = GlobalBlockId
				count++
			}
		}
		if count == 0 {
			iiP--
		}
	}
	FallingBlock.Id = GlobalBlockId
	GlobalBlockId++
	FallingBlock.I = 0
	FallingBlock.J = j
	FallingBlock.BlockType = blockType
	FallingBlock.Direction = d
}

func DrawFallingBlock() {
	block := FallingBlock
	blockShape := blockShapeList[block.BlockType-1][block.Direction]
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if blockShape[i][j].TypeId == 0 {
				continue
			}
			ii := block.I + i
			jj := block.J + j
			if ii >= 0 && ii < len(Ground) && jj >= 0 && jj < len(Ground[0]) {
				// 이미 고정된 블럭이 있는 경우 건너뜀 (덮지 않음)
				if Ground[ii][jj].TypeId != 0 && !Ground[ii][jj].IsFalling {
					continue
				}
				Ground[ii][jj] = Block{
					TypeId:    block.BlockType,
					IsFalling: true,
					Id:        block.Id,
				}
			}
		}
	}
}

func CanFall() bool {
	block := FallingBlock
	blockShape := blockShapeList[block.BlockType-1][block.Direction]
	for i := 3; i >= 0; i-- {
		for j := 0; j < 4; j++ {
			if blockShape[i][j].TypeId == 0 {
				continue
			}
			curI := block.I + i
			curJ := block.J + j
			if curI+1 >= len(Ground) {
				return false
			}
			if Ground[curI+1][curJ].TypeId != 0 && Ground[curI+1][curJ].IsFalling == false {
				return false
			}
		}
	}
	return true
}

func ClearFallingBlock() {
	for i := range Ground {
		for j := range Ground[i] {
			if Ground[i][j].IsFalling {
				Ground[i][j] = NoneBlock
			}
		}
	}
}

// 지금 떨어지고 있는 블럭 한칸 아래로 떨어지는거
// 기존 방식은 연속된 입력에서 불안정함. 완전히 지우고 다시 그리는 식으로 만들음
func FallingDown() {
	if !CanFall() {
		SetBlocksIsFallingFalse(FallingBlock.Id)
		return
	}

	ClearFallingBlock()
	FallingBlock.I++
	DrawFallingBlock()
}

// 바닥에 떨어지면 fallingBlock에서 삭제하는거
// 이거 새 블럭 만들어서 fallingBlock으로 지정해주는 로직 필요함
func SetBlocksIsFallingFalse(blockId int) {
	count := 0
	startI := FallingBlock.I
	startJ := FallingBlock.J

OUT:
	for i := startI; i < startI+4 && i < len(Ground); i++ {
		for j := startJ; j < startJ+4 && j < len(Ground[0]); j++ {
			if j < 0 {
				j = 0
			}
			if Ground[i][j].Id == blockId {
				Ground[i][j].IsFalling = false
				count++
				if count == 4 {
					FallingBlock = BlockInfo{}
					break OUT
				}
			}
		}
	}

	CreateBlockGroup(5, rand.Intn(6)+1, 't')
	CheckLine()
}

func EraseBlock() {
	block := FallingBlock
	blockShape := blockShapeList[block.BlockType-1][block.Direction]
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if blockShape[i][j].TypeId != 0 {
				ii := block.I + i
				jj := block.J + j
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
	shape := blockShapeList[block.BlockType-1][block.Direction]
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if shape[i][j].TypeId != 0 {
				ii := block.I + i
				jj := block.J + j + q
				if jj < 0 || jj >= len(Ground[0]) {
					return false
				}
				if ii >= 0 && Ground[ii][jj].TypeId != 0 && Ground[ii][jj].Id != block.Id {
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
			FallingBlock.J++
			DrawFallingBlock()
		}
	case 'a': // 왼쪽 이동
		if CanMove(-1) {
			EraseBlock()
			FallingBlock.J--
			DrawFallingBlock()
		}
	case 'e':
		RotateBlock(1)
	case 'q':
		RotateBlock(-1)
	}
}
func DeleteFallingBlock() {
	for i := FallingBlock.I; i < FallingBlock.I+4 && i < len(Ground); i++ {
		for j := FallingBlock.J; j < FallingBlock.J+4 && j < len(Ground[0]); j++ {
			jj := j
			if j < 0 {
				jj = 0
			}
			if Ground[i][jj].Id == FallingBlock.Id {
				Ground[i][jj] = NoneBlock
			}
		}
	}
	CreateBlockGroup(5, rand.Intn(6)+1, 't') // FallingBlock 초기화
	fmt.Println("공격에 당했습니다!")
}

func getNextDirection(current rune, q int) rune {
	// 시계 방향: +1, 반시계 방향: -1
	idxMap := map[rune]int{'t': 0, 'r': 1, 'b': 2, 'l': 3}
	runes := []rune{'t', 'r', 'b', 'l'}

	currentIdx := idxMap[current]
	nextIdx := (currentIdx + q + 4) % 4
	return runes[nextIdx]
}

func CanRotateTo(dir rune) bool {
	block := FallingBlock
	shape := blockShapeList[block.BlockType-1][dir]
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if shape[i][j].TypeId == 0 {
				continue
			}
			ii := block.I + i
			jj := block.J + j

			if ii >= len(Ground) || jj < 0 || jj >= len(Ground[0]) {
				return false
			}
			if Ground[ii][jj].TypeId != 0 && Ground[ii][jj].Id != block.Id {
				return false
			}
		}
	}
	return true
}

func RotateBlock(q int) {
	if FallingBlock.Id == 0 {
		return
	}

	nextDir := getNextDirection(FallingBlock.Direction, q)

	if CanRotateTo(nextDir) {
		EraseBlock()
		FallingBlock.Direction = nextDir
		DrawFallingBlock()
	}
}

func IsLimited(i int) bool {
	for j := range len(Ground[0]) {
		if Ground[i][j].TypeId == 0 {
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
		Ground[0][j] = Block{TypeId: 0}
	}
	user.Me.LineClearSuccess(1)
}

func CheckLine() {
	for i := len(Ground) - 1; i >= 0; i-- {
		if IsLimited(i) {
			ClearLine(i)
			i++
		}
	}
}
