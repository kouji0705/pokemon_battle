package core

import (
	"math/rand"
	"time"
)

type Move struct {
	Name     string
	Power    int    // 技の威力
	Accuracy int    // 命中率（0-100%）
	Type     string // 技の属性（例: 電気、炎など）

	randSource *rand.Rand // 乱数生成器を保持
}

// NewMove creates a new move and initializes the random source.
func NewMove(name string, power, accuracy int, moveType string) *Move {
	return &Move{
		Name:       name,
		Power:      power,
		Accuracy:   accuracy,
		Type:       moveType,
		randSource: rand.New(rand.NewSource(time.Now().UnixNano())), // 新しい乱数生成器を初期化
	}
}

// CanHit checks if the move hits based on its accuracy.
func (m *Move) CanHit() bool {
	// move に保持している乱数生成器を使って乱数を生成
	return m.randSource.Intn(100) < m.Accuracy
}

// / 技（Move）のダメージを計算する
func (m *Move) CalculateDamage(attacker *Pokemon, defender *Pokemon) int {
	// ダメージ計算式
	damage := (attacker.Attack * m.Power) / defender.Defense

	// ダメージにランダム要素を追加（0.85倍～1.0倍）
	randomFactor := 0.85 + rand.Float64()*(1.0-0.85)
	damage = int(float64(damage) * randomFactor)

	if damage < 0 {
		damage = 0
	}
	return damage
}
