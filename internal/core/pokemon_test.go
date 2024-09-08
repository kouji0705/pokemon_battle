package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ポケモンのHPが正しく減少するかをテスト
func TestPokemonAttack_FixedValues(t *testing.T) {
	pikachu := &Pokemon{
		Name:    "Pikachu",
		HP:      100,
		Attack:  55,
		Defense: 40,
		Speed:   90,
		Moves: []Move{
			*NewMove("Thunderbolt", 90, 95, "Electric"),
		},
	}

	charmander := &Pokemon{
		Name:    "Charmander",
		HP:      100,
		Attack:  52,
		Defense: 43,
		Speed:   65,
		Moves: []Move{
			*NewMove("Flamethrower", 90, 95, "Fire"),
		},
	}

	// PikachuがCharmanderを攻撃（乱数を固定）
	move := pikachu.Moves[0]
	damage, err := pikachu.AttackPokemon(charmander, move)
	fmt.Println(pikachu.Name, "技：", move.Name, ", ", damage, "ダメージ！")
	// エラーがないことを確認
	assert.NoError(t, err)

	// ダメージを固定された値でテスト
	expectedDamage := 104 // 例: 期待する固定値
	assert.Equal(t, expectedDamage, damage, "ダメージが期待値と異なります")

	// 期待結果と実際のHPを比較
	expectedHP := 0
	assert.Equal(t, expectedHP, charmander.HP, "CharmanderのHPが期待値と異なります")
}

// 技の命中判定をテスト
func TestMoveCanHit(t *testing.T) {
	thunderbolt := NewMove("Thunderbolt", 90, 95, "Electric")

	hitCount := 0
	trials := 1000

	for i := 0; i < trials; i++ {
		if thunderbolt.CanHit() {
			hitCount++
		}
	}

	hitRate := float64(hitCount) / float64(trials)
	expectedHitRate := 0.95

	// 実際の命中率が期待範囲内か確認
	if hitRate < expectedHitRate-0.05 || hitRate > expectedHitRate+0.05 {
		t.Errorf("技の命中率が期待値と異なります。期待: %.2f, 実際: %.2f", expectedHitRate, hitRate)
	}
}
