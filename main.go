package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

type block struct {
	id        int
	typeId    int
	isFalling bool
	isNotNone bool
	direction rune // 상t하b좌l우r
}

type blockInfo struct {
	id        int
	i         int
	j         int
	blockType int
}

// 해당 블럭의 가장 왼쪽(j), 가장 위쪽(i)

var noneBlock block = block{0, 0, false, false, 't'}
var globalBlockId = 1
var fallingBlock blockInfo

func createBlock(typeId int) block {
	b := block{0, typeId, true, true, 't'}
	return b
}

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

var blockShapeList [7][4][4]block

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var ground [22][12]block

func main() {
	initEnv()
	createBlockGroup(5, globalBlockId, 1)
	globalBlockId++

	printArray(ground)

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

		move(ch)

		if ch != 'f' {
			printArray(ground)
			continue
		}

		fallingDown()
		printArray(ground)
	}
}

func printArray(a [22][12]block) {
	screen.Clear()
	fmt.Println()
	for _, i := range a {
		for _, j := range i {
			_, _ = blockColorMap[j.typeId].Print("  ")

		}
		fmt.Println()
	}

}

func getRightestByBlockType(blockType int) int {
	blockShape := blockShapeList[blockType]
	var rightest int = 0
	for i := range 4 {
		for j := range 4 {
			if blockShape[i][j].id != 0 {
				rightest = j
			}
		}
	}
	return rightest
}

func createBlockGroup(j int, id int, blockType int) {
	for iP := range 4 {
		for jP := range 4 {
			ground[0+iP][j+jP] = blockShapeList[blockType-1][iP][jP]
			ground[0+iP][j+jP].id = id
		}
	}
	fallingBlock.id = id
	fallingBlock.i = 0
	fallingBlock.j = j
	fallingBlock.blockType = blockType
}
func fallingDown() {
	iStart := fallingBlock.i
	jStart := fallingBlock.j

	for i := 3; i >= 0; i-- {
		for j := 0; j < 4; j++ {
			curI := iStart + i
			curJ := jStart + j
			if curI >= len(ground) || curJ >= len(ground[0]) {
				continue
			}
			if ground[curI][curJ].id == fallingBlock.id {
				belowI := curI + 1
				if belowI >= len(ground) || (ground[belowI][curJ].typeId > 0 && ground[belowI][curJ].id != fallingBlock.id) {
					setBlocksIsFallingFalse(fallingBlock.id)
					return
				}
			}
			if ground[curI][curJ].id == fallingBlock.id {
				ground[curI+1][curJ] = ground[curI][curJ]
				ground[curI][curJ] = noneBlock
			}
		}
	}

	fallingBlock.i++
}

func setBlocksIsFallingFalse(blockId int) {
	count := 0
	startI := fallingBlock.i
	startJ := fallingBlock.j

	for i := startI; i < startI+4 && i < len(ground); i++ {
		for j := startJ; j < startJ+4 && j < len(ground[0]); j++ {
			if ground[i][j].id == blockId {
				ground[i][j].isFalling = false
				count++
				if count == 4 {
					fallingBlock = blockInfo{}
					return
				}
			}
		}
	}
}

func move(q rune) {
	if fallingBlock.id == 0 {
		return
	}
	rightest := getRightestByBlockType(fallingBlock.blockType)
	switch q {
	case 'd':

		if fallingBlock.j+rightest >= len(ground[0]) {
			return
		}

		for i := 0; i < 4; i++ {
			for j := rightest; j >= 0; j-- {
				if fallingBlock.j+rightest+1 >= len(ground[0]) {
					return
				}
				b := ground[fallingBlock.i+i][fallingBlock.j+j]
				if b.id == fallingBlock.id {
					if ground[fallingBlock.i+i][fallingBlock.j+j+1].isNotNone &&
						ground[fallingBlock.i+i][fallingBlock.j+j+1].id != fallingBlock.id {
						return
					}
				}
			}
		}

		// 오른쪽으로 한 칸 이동 (오른쪽 끝부터 복사)
		for i := 0; i < 4; i++ {
			for j := rightest; j >= 0; j-- {
				b := ground[fallingBlock.i+i][fallingBlock.j+j]
				if b.id == fallingBlock.id {
					ground[fallingBlock.i+i][fallingBlock.j+j+1] = b
					ground[fallingBlock.i+i][fallingBlock.j+j] = noneBlock
				}
			}
		}
		fallingBlock.j++

	case 'a': // 왼쪽 이동
		// 왼쪽 경계 체크
		if fallingBlock.j <= 0 {
			return
		}

		// 왼쪽 칸이 비어있는지 확인 (충돌 방지)
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if fallingBlock.j+j >= len(ground[0]) {
					break
				}
				b := ground[fallingBlock.i+i][fallingBlock.j+j]

				if b.id == fallingBlock.id {
					if ground[fallingBlock.i+i][fallingBlock.j+j-1].isNotNone &&
						ground[fallingBlock.i+i][fallingBlock.j+j-1].id != fallingBlock.id {
						return
					}
				}
			}
		}

		// 왼쪽으로 한 칸 이동 (왼쪽 끝부터 복사)
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if fallingBlock.j+j >= len(ground[0]) {
					break
				}
				b := ground[fallingBlock.i+i][fallingBlock.j+j]
				if b.id == fallingBlock.id {
					ground[fallingBlock.i+i][fallingBlock.j+j-1] = b
					ground[fallingBlock.i+i][fallingBlock.j+j] = noneBlock
				}
			}
		}
		fallingBlock.j--
	}
}

func initEnv() {
	color.NoColor = false
	// I자 블록 (1)
	blockShapeList[0] = [4][4]block{
		{createBlock(1)},
		{createBlock(1)},
		{createBlock(1)},
		{createBlock(1)},
	}

	// L자 블록 (2)
	blockShapeList[1] = [4][4]block{
		{createBlock(2)},
		{createBlock(2)},
		{createBlock(2), createBlock(2)},
	}

	// T자 블록 (3)
	blockShapeList[2] = [4][4]block{
		{noneBlock, createBlock(3)},
		{createBlock(3), createBlock(3)},
		{createBlock(3)},
	}

	// 정사각형 블록 (4)
	blockShapeList[3] = [4][4]block{
		{createBlock(4), createBlock(4)},
		{createBlock(4), createBlock(4)},
	}

	// J자 블록 (5)
	blockShapeList[4] = [4][4]block{
		{noneBlock, createBlock(5)},
		{noneBlock, createBlock(5)},
		{createBlock(5), createBlock(5)},
	}

	// S자 블록 (6)
	blockShapeList[5] = [4][4]block{
		{noneBlock, createBlock(6), createBlock(6)},
		{createBlock(6), createBlock(6)},
	}

	// Z자 블록 (7)
	blockShapeList[6] = [4][4]block{
		{createBlock(7), createBlock(7)},
		{noneBlock, createBlock(7), createBlock(7)},
	}
}
