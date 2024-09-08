package core

import (
	"math/rand"
	"time"
)

type Move struct {
	Name             string
	Power            int            // 技の威力
	Accuracy         int            // 命中率（0-100%）
	Type             string         // 技の属性（例: 電気、炎など）
	randSource       *rand.Rand     // 乱数生成器を保持
	randomFactorFunc func() float64 // ランダムなダメージ倍率を計算する関数
	hitCheckFunc     func() bool    // 命中判定の関数
}

// NewMove creates a new move and initializes the random source.
func NewMove(name string, power, accuracy int, moveType string) *Move {
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	move := &Move{
		Name:       name,
		Power:      power,
		Accuracy:   accuracy,
		Type:       moveType,
		randSource: randSource,
		randomFactorFunc: func() float64 {
			// デフォルトのダメージ倍率：0.85～1.0
			return 0.85 + randSource.Float64()*(1.0-0.85)
		},
		hitCheckFunc: func() bool {
			// デフォルトの命中判定：Accuracyに基づくランダムな命中判定
			return randSource.Intn(100) < accuracy
		},
	}
	return move
}

// CanHit checks if the move hits based on its accuracy.
func (m *Move) CanHit() bool {
	return m.hitCheckFunc() // 注入された命中判定関数を使用
}

// CalculateDamage calculates the damage of the move.
func (m *Move) CalculateDamage(attacker *Pokemon, defender *Pokemon) int {
	// ダメージ計算式
	damage := (attacker.Attack * m.Power) / defender.Defense

	// ランダム要素を使用してダメージを計算
	randomFactor := m.randomFactorFunc() // 注入された関数を使用
	damage = int(float64(damage) * randomFactor)

	if damage < 0 {
		damage = 0
	}
	return damage
}
