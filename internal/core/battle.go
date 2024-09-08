package core

// Battle represents a battle between two Pokemon.
type Battle struct {
	Pokemon1 *Pokemon
	Pokemon2 *Pokemon
}

// Start initiates the battle loop.
func (b *Battle) Start() string {
	for {
		// ポケモン1のターン
		move := b.Pokemon1.Moves[0] // 簡略化のため最初の技を使用
		damage, _ := b.Pokemon1.AttackPokemon(b.Pokemon2, move)
		println(b.Pokemon1.Name, "技：", move.Name, ", ", damage, "ダメージ！")
		if b.Pokemon2.IsFainted() {
			return b.Pokemon1.Name + " 勝利!"
		}

		// ポケモン2のターン
		move = b.Pokemon2.Moves[0] // 同様に、最初の技を使用
		damage, _ = b.Pokemon2.AttackPokemon(b.Pokemon1, move)
		println(b.Pokemon2.Name, "技：", move.Name, ", ", damage, "ダメージ!")
		if b.Pokemon1.IsFainted() {
			return b.Pokemon2.Name + " 勝利!"
		}
	}
}
