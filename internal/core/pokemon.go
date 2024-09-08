package core

type Pokemon struct {
	Name    string
	HP      int
	Attack  int
	Defense int
	Speed   int
	Moves   []Move // 使用可能な技のリスト
}

// ポケモンが技を使用し、相手にダメージを与える
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

// ポケモンが戦闘不能かどうかを判定
func (p *Pokemon) IsFainted() bool {
	return p.HP <= 0
}
