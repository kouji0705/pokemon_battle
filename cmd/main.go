package main

import (
	"github.com/kouji0705/pokemon-battle/internal/core"
)

func main() {

	// ポケモンの作成
	pikachu := &core.Pokemon{
		Name:    "ピカチュウ", // 日本語の名前
		HP:      100,
		Attack:  55,
		Defense: 40,
		Speed:   90,
		// 技の名前も日本語
		Moves: []core.Move{
			*core.NewMove("10まんボルト", 90, 95, "でんき"), // 「Thunderbolt」を日本語化
		},
	}

	charmander := &core.Pokemon{
		Name:    "ヒトカゲ", // 日本語の名前
		HP:      100,
		Attack:  52,
		Defense: 43,
		Speed:   65,
		// 技の名前も日本語
		Moves: []core.Move{
			*core.NewMove("かえんほうしゃ", 90, 95, "ほのお"), // 「Flamethrower」を日本語化
		},
	}

	// バトルの開始
	battle := &core.Battle{Pokemon1: pikachu, Pokemon2: charmander}
	result := battle.Start()
	println(result)
}
