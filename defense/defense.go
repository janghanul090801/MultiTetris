package defense

import (
	"MultiTetris/blockShape"
	"fmt"
)

func InitDefense() {
	//깊은 복사 있었으면.... 왜 새벽 4시냐 한거 없는데
	// 주석 좀 달아줘 코드에
	// DefenseGround := blockShape.copyGround
	blockShape.PrintArray(blockShape.Ground)
	fmt.Println("GoGoGo")

}

//func DeployShield(x, y int) bool {
//if i < 0 || i >= len(DefenseGround) || j < 0 || j >= len(DefenseGround[0]) {
//	return false
//}
//
//targetBlock := DefenseGround[i][j]
//
////블록이 없거나 Falling X 이면 제거 불가
//if targetBlock.isNotNone == false || targetBlock.isFalling == false {
//	return false
//}
//
//blockId := targetBlock.id
//
//// 같은 ID 제거
//for x := 0; x < len(Ground); x++ {
//	for y := 0; y < len(Ground[0]); y++ {
//		if Ground[x][y].id == blockId {
//			Ground[x][y] = noneBlock
//		}
//	}
//}
//
////공격 성공 시 블록 초기화
//if fallingBlock.id == blockId {
//	fallingBlock = blockInfo{}
//}

//	return true
//}
