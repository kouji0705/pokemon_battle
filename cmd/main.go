package main

import (
	"github.com/username/pokemon-battle/internal/core"
)

func main() {
	// ポケモンの作成
	pikachu := &core.Pokemon{
		Name:    "Pikachu",
		HP:      100,
		Attack:  55,
		Defense: 40,
		Speed:   90,
		Moves: []core.Move{
			{Name: "Thunderbolt", Power: 90, Accuracy: 95, Type: "Electric"},
		},
	}

	charmander := &core.Pokemon{
		Name:    "Charmander",
		HP:      100,
		Attack:  52,
		Defense: 43,
		Speed:   65,
		Moves: []core.Move{
			{Name: "Flamethrower", Power: 90, Accuracy: 95, Type: "Fire"},
		},
	}

	// バトルの開始
	battle := &core.Battle{Pokemon1: pikachu, Pokemon2: charmander}
	result := battle.Start()
	println(result)
}
