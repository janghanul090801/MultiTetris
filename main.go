package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

type block struct {
	id        int
	typeId    int
	isFalling bool
	isNotNone bool
	direction rune // 상t하b좌l우r
}

var noneBlock block = block{0, 0, false, false, 't'}
var globalBlockId = 1

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
	setGround(0, 0, globalBlockId, 1)
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
		fallingDown()
		printArray(ground)
	}
}

func centerAlignNumber(n int, width int) string {
	s := fmt.Sprintf("%d", n)
	padding := width - len(s)
	if padding <= 0 {
		return s
	}
	left := padding / 2
	right := padding - left
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
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

func setGround(i int, j int, id int, blockType int) {
	for iP := range 4 {
		for jP := range 4 {
			ground[i+iP][j+jP] = blockShapeList[blockType-1][iP][jP]
			ground[i+iP][j+jP].id = id
		}
	}
}

func fallingDown() {
	for i := range 22 {
		for j := range 12 {
			blockPtr := &ground[21-i][11-j]
			if blockPtr.typeId > 0 && blockPtr.isFalling {
				tmp := *blockPtr
				if 21-i == 0 {
					*blockPtr = noneBlock
				} else {
					*blockPtr = ground[21-i-1][11-j]
				}
				if 21-i+1 == 22 {
					setBlocksIsFallingFalse(blockPtr.id)
				} else {
					ground[21-i+1][11-j] = tmp
				}
			}
		}
	}
}

// 해당 블럭의 가장 왼쪽(j), 가장 위쪽(i)
func setBlocksIsFallingFalse(blockId int) {
	count := 0
	for i := range 22 {
		for j := range 12 {
			if ground[21-i][11-j].id == blockId {
				ground[21-i][11-j].isFalling = false
				count++
			}
			if count == 4 {
				return
			}
		}
	}
}

func move(q rune) {
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
