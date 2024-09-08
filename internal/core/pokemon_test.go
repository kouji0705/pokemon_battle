package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ポケモンのHPが正しく減少するかをテスト
func TestPokemonAttack_FixedValues(t *testing.T) {
	// Pikachu と Charmander のセットアップ
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

	// Pikachu の技を取得
	move := &pikachu.Moves[0]

	// 乱数要素と命中判定を固定化する
	move.randomFactorFunc = func() float64 {
		return 1.0 // ダメージのランダム倍率を固定（1.0倍）
	}
	move.hitCheckFunc = func() bool {
		return true // 命中判定を固定（常に命中）
	}

	// Pikachu が Charmander を攻撃
	damage, err := pikachu.AttackPokemon(charmander, *move)

	// 結果を表示
	fmt.Println(pikachu.Name, "技：", move.Name, ", ", damage, "ダメージ！")

	// エラーがないことを確認
	assert.NoError(t, err)

	// ダメージが期待通りか確認
	expectedDamage := 115 // 期待する固定ダメージ（例: 計算に基づく値）
	assert.Equal(t, expectedDamage, damage, "ダメージが期待値と異なります")

	// Charmander のHPが正しく減少しているか確認
	expectedHP := 0
	assert.Equal(t, expectedHP, charmander.HP, "CharmanderのHPが期待値と異なります")
}

func TestCanHitWithFixedHitCheck(t *testing.T) {
	// テスト対象の技を作成
	move := NewMove("Thunderbolt", 90, 95, "Electric")

	// 命中判定を固定化（常に命中する）
	move.hitCheckFunc = func() bool {
		return true // 常に命中する
	}

	// CanHitの結果をテスト（命中するはず）
	if !move.CanHit() {
		t.Errorf("命中するはずのCanHitがfalseを返しました")
	}

	// 命中判定を固定化（常にミスする）
	move.hitCheckFunc = func() bool {
		return false // 常にミスする
	}

	// CanHitの結果をテスト（ミスするはず）
	if move.CanHit() {
		t.Errorf("ミスするはずのCanHitがtrueを返しました")
	}
}

// // 技の命中判定をテスト
// func TestMoveCanHit(t *testing.T) {
// 	thunderbolt := NewMove("Thunderbolt", 90, 95, "Electric")

// 	hitCount := 0
// 	trials := 1000

// 	for i := 0; i < trials; i++ {
// 		if thunderbolt.CanHit() {
// 			hitCount++
// 		}
// 	}

// 	hitRate := float64(hitCount) / float64(trials)
// 	expectedHitRate := 0.95

// 	// 実際の命中率が期待範囲内か確認
// 	if hitRate < expectedHitRate-0.05 || hitRate > expectedHitRate+0.05 {
// 		t.Errorf("技の命中率が期待値と異なります。期待: %.2f, 実際: %.2f", expectedHitRate, hitRate)
// 	}
// }
