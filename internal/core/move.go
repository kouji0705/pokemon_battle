package core

import (
	"math/rand"
	"time"
)

// Move represents a Pokemon move with its properties.
type Move struct {
	Name     string
	Power    int    // 技の威力
	Accuracy int    // 命中率（0-100%）
	Type     string // 技の属性（例: 電気、炎など）
}

// CanHit checks if the move hits based on its accuracy.
func (m *Move) CanHit() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) < m.Accuracy
}

// CalculateDamage calculates the damage this move does to the target.
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
