package game

import "github.com/google/uuid"

func draco() Villain {
	return Villain{
		Name:         "Draco Malfoy",
		Id:           int(uuid.New().ID()),
		ImgPath:      "/images/villains/dracomalfoy.jpg",
		SetId:        "Game 1",
		CurDamage:    0,
		MaxHp:        6,
		Effect:       []Effect{DamageActiveIfLocationAdded{Amount: 2}},
		DeathEffect:  nil,
		playBeforeDA: true,
	}
}

type DamageActiveIfLocationAdded struct {
	Amount int
}

func (effect DamageActiveIfLocationAdded) Trigger(gs *Gamestate) {
	// find player who was active when location got added.
}

func quirrell() Villain {
	return Villain{
		Name:      "Quirinus Quirrell",
		Id:        int(uuid.New().ID()),
		ImgPath:   "/images/villains/quirrell.jpg",
		SetId:     "Game 1",
		CurDamage: 0,
		MaxHp:     6,
		Effect:    []Effect{DamageCurrentPlayer{Amount: 1}},
		DeathEffect: []Effect{
			AllPlayersGainMoney{Amount: 1},
			AllPlayersGainHealth{Amount: 1},
		},
		playBeforeDA: false,
	}
}

func crabbeAndGoyle() Villain {
	return Villain{
		Name:         "Crabbe and Goyle",
		Id:           int(uuid.New().ID()),
		ImgPath:      "/images/villains/crabbeandgoyle.jpg",
		SetId:        "Game 1",
		CurDamage:    0,
		MaxHp:        5,
		Effect:       nil,
		DeathEffect:  nil,
		playBeforeDA: false,
	}
}
