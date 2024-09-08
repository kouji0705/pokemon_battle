package core

// Pokemon represents a Pokemon with basic stats.
type Pokemon struct {
	Name    string
	HP      int
	Attack  int
	Defense int
	Speed   int
	Moves   []Move // 使用可能な技のリスト
}

// AttackPokemon applies a selected move from this Pokemon to the target Pokemon.
func (p *Pokemon) AttackPokemon(target *Pokemon, move Move) (damage int, err error) {
	if !move.CanHit() {
		return 0, nil // 技が外れた場合
	}

	// ダメージ計算
	damage = move.CalculateDamage(p, target)
	target.HP -= damage
	if target.HP < 0 {
		target.HP = 0
	}
	return damage, nil
}

// IsFainted checks if the Pokemon has fainted (HP <= 0).
func (p *Pokemon) IsFainted() bool {
	return p.HP <= 0
}
