package main

import "github.com/google/uuid"

type Encounter struct {
	Name          string
	id            int
	setId         string
	effects       []Effect
	rewardEffects []Effect
}

// Effect - draw to 4 instead of 5 if health <= 4
// Complete - play 2 even cost cards in a turn
// Reward - heal any player 1 after playing even cost card.
func peskipiksiPesternome() Encounter {
	id := int(uuid.New().ID())
	return Encounter{
		Name:  "Peskipiksi Pesternomi!",
		id:    id,
		setId: "box 1",
	}
}
