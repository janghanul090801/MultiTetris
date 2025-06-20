package user

import "fmt"

type User struct {
	Points      int `json:"points"`
	Lines       int `json:"lines"`
	Combos      int `json:"combos"`
	AttackCount int `json:"attackcount"`
}

// 현재 플레이어와 상대
var (
	Me    = User{}
	Other = User{}
)

// 줄 제거 시 점수 증가
func (u *User) LineClearSuccess(lines int) {
	u.Lines += lines
	u.Points += lines * 100 * 20
}

// 공격 성공 시 콤보에 맞춰  증가
func (u *User) AttackSuccess() {
	u.AttackCount++
	u.Points += int(1000 * float64(u.Combos) * 0.85)
	u.Combos++
}

// 공격 실패 시 콤보 초기화
func (u *User) AttackFailed() {
	u.Combos = 1
}

func (u *User) PrintScore() {
	fmt.Println("Points:", u.Points)
	fmt.Println("Lines:", u.Lines)

}
