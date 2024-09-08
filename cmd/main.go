package main

import (
	"github.com/kouji0705/pokemon-battle/internal/core"
)

func main() {

	// ポケモンの作成
	pikachu := &core.Pokemon{
		Name:    "ピカチュウ",
		HP:      100,
		Attack:  55,
		Defense: 40,
		Speed:   90,
		Moves: []core.Move{
			*core.NewMove("10まんボルト", 90, 95, "でんき"),
		},
	}

	charmander := &core.Pokemon{
		Name:    "ヒトカゲ",
		HP:      100,
		Attack:  52,
		Defense: 43,
		Speed:   65,
		Moves: []core.Move{
			*core.NewMove("かえんほうしゃ", 90, 95, "ほのお"),
		},
	}

	// バトルの開始
	battle := &core.Battle{Pokemon1: pikachu, Pokemon2: charmander}
	result := battle.Start()
	println(result)
}
