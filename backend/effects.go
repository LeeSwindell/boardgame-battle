package game

import "math/rand"

type DamageAllPlayers struct {
	amount int
}

func (effect DamageAllPlayers) Trigger(gs *Gamestate) {
	for _, player := range gs.Players {
		player.Health -= effect.amount
	}
}

func Damage(p *Player, amount int) {
	p.Health -= amount
}

func DamageAll(gs *Gamestate, amount int) {
	for _, p := range gs.Players {
		p.Health -= amount
	}
}

func spendMoney(p *Player, amount int) {
	p.Money -= amount
}

func gainMoney(p *Player, amount int) {
	p.Money += amount
}

func giveAllMoney(gs *Gamestate, amount int) {
	for _, p := range gs.Players {
		p.Money += amount
	}
}

func changeDamage(p *Player, amount int) {
	p.Damage += amount
}

func giveAllDamage(gs *Gamestate, amount int) {
	for _, p := range gs.Players {
		p.Damage += amount
	}
}

func buyCard(gs *Gamestate, p *Player, card Card, cost int) {
	spendMoney(p, cost)
	p.Discard.Cards = append(p.Discard.Cards, card)
	// Remove from market FIX
}

func addControl(gs *Gamestate) {
	l := gs.Locations[0]
	l.CurControl += 1
	if l.CurControl == l.MaxControl && len(gs.Locations) > 1 {
		gs.Locations = gs.Locations[1:]
	} else if l.CurControl == l.MaxControl && len(gs.Locations) == 1 {
		gameOver(gs, "no")
	}
}

func damageVillain(v *Villain) {
	v.CurDamage += 1
	if v.CurDamage == v.MaxHp {
		// FIX trigger death effect, draw new villain
	}
}

func healVillain(v *Villain) {
	if v.CurDamage > 0 {
		v.CurDamage -= 1
	}
}

func rollSlytherinDie(p *Player) {
	n := rand.Intn(6)

	switch {
	case n < 3:
		// damage
	case n == 3:
		//health
	case n == 4:
		//draw card
	case n == 5:
		//money
	}
}

func gameOver(gs *Gamestate, didYaWin string) {
	// You Lose!!!!
}
