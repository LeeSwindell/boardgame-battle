package game

import "math/rand"

func Damage(p *Player, amount int) {
	p.health -= amount
}

func DamageAll(gs *Gamestate, amount int) {
	for _, p := range gs.players {
		p.health -= amount
	}
}

func spendMoney(p *Player, amount int) {
	p.money -= amount
}

func gainMoney(p *Player, amount int) {
	p.money += amount
}

func giveAllMoney(gs *Gamestate, amount int) {
	for _, p := range gs.players {
		p.money += amount
	}
}

func changeDamage(p *Player, amount int) {
	p.damage += amount
}

func giveAllDamage(gs *Gamestate, amount int) {
	for _, p := range gs.players {
		p.damage += amount
	}
}

func buyCard(gs *Gamestate, p *Player, card Card, cost int) {
	spendMoney(p, cost)
	p.discard.cards = append(p.discard.cards, card)
	// Remove from market FIX
}

func addControl(gs *Gamestate) {
	l := gs.locations[0]
	l.curControl += 1
	if l.curControl == l.maxControl && len(gs.locations) > 1 {
		gs.locations = gs.locations[1:]
	} else if l.curControl == l.maxControl && len(gs.locations) == 1 {
		gameOver(gs, "no")
	}
}

func damageVillain(v *Villain) {
	v.curDamage += 1
	if v.curDamage == v.maxHp {
		// FIX trigger death effect, draw new villain
	}
}

func healVillain(v *Villain) {
	if v.curDamage > 0 {
		v.curDamage -= 1
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
