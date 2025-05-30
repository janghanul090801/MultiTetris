package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var ground [4][4][2]int

func main() {

	// init
	for range 2 {
		setGroundWithRND()
	}

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
		printArray(ground)
		if gameOver(ground) {
			break
		}
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

func printArray(a [4][4][2]int) {
	screen.Clear()
	fmt.Println()
	var colorFunc [4]*color.Color
	for _, i := range a {
		for idx, j := range i {
			colorFunc[idx] = color.New(color.FgWhite)

			switch j[0] {
			case 2:
				colorFunc[idx] = color.BgRGB(57, 42, 26)
			case 4:
				colorFunc[idx] = color.BgRGB(71, 42, 22)
			case 8:
				colorFunc[idx] = color.BgRGB(127, 65, 11)
			case 16:
				colorFunc[idx] = color.BgRGB(141, 54, 8)
			case 32:
				colorFunc[idx] = color.BgRGB(145, 33, 7)
			case 64:
				colorFunc[idx] = color.BgRGB(167, 37, 7)
			case 128:
				colorFunc[idx] = color.BgRGB(97, 77, 12)
			case 256:
				colorFunc[idx] = color.BgRGB(75, 83, 12)
			case 512:
				colorFunc[idx] = color.BgRGB(113, 90, 12)
			case 1024:
				colorFunc[idx] = color.BgRGB(131, 96, 11)
			case 2048:
				colorFunc[idx] = color.BgRGB(129, 103, 11)
			}
		}
		for k := range 4 {
			fmt.Print(colorFunc[k].SprintFunc()("       "))
		}
		fmt.Println()
		for k := range 4 {
			fmt.Print(colorFunc[k].SprintFunc()(centerAlignNumber(i[k][0], 7)))
		}
		fmt.Println()
		for k := range 4 {
			fmt.Print(colorFunc[k].SprintFunc()("       "))
		}
		fmt.Print("\n")
	}
	for i := range 4 {
		for j := range 4 {
			ground[i][j][1] = 0
		}
	}

}

func setGround(i int, j int, val int) {
	if ground[i][j][0] != 0 {
		setGround(seededRand.Intn(4), seededRand.Intn(4), val)
		return
	}
	ground[i][j][0] = val
}

func setGroundWithRND() {
	var count int
OUTFOR:
	for _, i := range ground {
		for _, j := range i {
			if j[0] == 0 {
				count++
				break OUTFOR
			}
		}
	}
	if count == 0 {
		os.Exit(0)
		return
	}
	randN := seededRand.Intn(11)
	if randN <= 9 {
		setGround(seededRand.Intn(4), seededRand.Intn(4), 2)
	} else {
		setGround(seededRand.Intn(4), seededRand.Intn(4), 4)
	}
}

func move(q rune) {
	var moveToI int
	var moveToJ int
	switch q {
	case 'w':
		moveToI = -1
	case 's':
		moveToI = 1
	case 'd':
		moveToJ = 1
	case 'a':
		moveToJ = -1
	}
	var count int
	var changed [4][4][2]int = ground

	for x := range 4 {
		for y := range 4 {
			switch {
			case moveToI == -1 && ground[y][x][0] != 0:
				changed = moveTo(y, x, moveToI, moveToJ, ground)
			case moveToI == 1 && ground[3-y][x][0] != 0:
				changed = moveTo(3-y, x, moveToI, moveToJ, ground)
			case moveToJ == -1 && ground[x][y][0] != 0:
				changed = moveTo(x, y, moveToI, moveToJ, ground)
			case moveToJ == 1 && ground[x][3-y][0] != 0:
				changed = moveTo(x, 3-y, moveToI, moveToJ, ground)
			}

			if ground != changed {
				ground = changed
				count++
			}

		}
	}
	if count > 0 {
		setGroundWithRND()
	}
}

func moveTo(i int, j int, moveToI int, moveToJ int, ground [4][4][2]int) [4][4][2]int {
	from := ground[i][j]
	ground[i][j][0] = 0
	for !((i == 0 && moveToI == -1) || (i == 3 && moveToI == 1) || (j == 0 && moveToJ == -1) || (j == 3 && moveToJ == 1)) {
		if ground[i+moveToI][j+moveToJ][0] == 0 {
			i += moveToI
			j += moveToJ
		} else if ground[i+moveToI][j+moveToJ] == from && from[1] != 1 {
			from[0] = from[0] + from[0]
			i += moveToI
			j += moveToJ
			ground[i][j][0] = 0
			from[1] = 1

		} else {
			break
		}
	}
	ground[i][j] = from
	return ground
}

func gameOver(ground [4][4][2]int) bool {
	clone := ground
	b := true
	for i := range 4 {
		for j := range 4 {
			if ground[i][j][0] == 0 {
				b = false
				goto RETURN
			}
			if clone == moveTo(i, j, 1, 0, clone) {
				if clone == moveTo(i, j, -1, 0, clone) {
					if clone == moveTo(i, j, 0, 1, clone) {
						if clone == moveTo(i, j, 0, -1, clone) {
							b = true
						} else {
							b = false
							goto RETURN
						}
					} else {
						b = false
						goto RETURN
					}
				} else {
					b = false
					goto RETURN
				}
			} else {
				b = false
				goto RETURN
			}
		}
	}
RETURN:
	return b
}
