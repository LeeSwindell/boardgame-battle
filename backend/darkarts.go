package main

import (
	"math/rand"

	"github.com/google/uuid"
)

func dementorsKiss() DarkArt {
	return DarkArt{
		Name:    "Dementor's Kiss",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/dementorskiss.jpg",
		SetId:   "game 3",
		Effects: []Effect{
			DamageCurrentPlayer{Amount: 2},
			DamageAllPlayersButCurrent{Amount: 1},
		},
	}
}

func flipendo() DarkArt {
	return DarkArt{
		Name:    "Flipendo",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/flipendo.jpg",
		SetId:   "game 1",
		Effects: []Effect{
			DamageCurrentPlayer{Amount: 1},
			ActivePlayerDiscards{Amount: 1},
		},
	}
}

func heWhoMustNotBeNamed() DarkArt {
	return DarkArt{
		Name:    "He-Who-Must-Not-Be-Named",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/hewhomustnotbenamed.jpg",
		SetId:   "game 1",
		Effects: []Effect{
			AddToLocation{Amount: 1},
		},
	}
}

func avadaKedavra() DarkArt {
	return DarkArt{
		Name:    "Avada Kedavra!",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/avadakedavra.jpg",
		SetId:   "game 4",
		Effects: []Effect{
			AvadaKedavraEffect{Damage: 3},
			RevealDarkArts{Amount: 1},
		},
	}
}

type AvadaKedavraEffect struct {
	Damage int
}

func (effect AvadaKedavraEffect) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	stunned := ChangePlayerHealth(user, -effect.Damage, gs)
	if stunned {
		StunPlayer(user, gs)
		AddToLocation{Amount: 1}.Trigger(gs)
	}
	LoadNewDarkArt(gs)
}

func expulso() DarkArt {
	return DarkArt{
		Name:    "Expulso!",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/expulso.jpg",
		SetId:   "game 1",
		Effects: []Effect{
			DamageCurrentPlayer{Amount: 2},
		},
	}
}

func handOfGlory() DarkArt {
	return DarkArt{
		Name:    "Hand of Glory",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/handofglory.jpg",
		SetId:   "game 2",
		Effects: []Effect{
			DamageCurrentPlayer{Amount: 1},
			AddToLocation{Amount: 1},
		},
	}
}

func heirOfSlytherin() DarkArt {
	return DarkArt{
		Name:    "Heir of Slytherin",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/heirofslytherin.jpg",
		SetId:   "game 4",
		Effects: []Effect{
			HeirOfSlytherinDiceEffect{Prompt: "Heir of Slytherin! Discard a card"},
		},
	}
}

type HeirOfSlytherinDiceEffect struct {
	Prompt string
}

func (effect HeirOfSlytherinDiceEffect) Trigger(gs *Gamestate) {
	n := rand.Intn(6)
	switch n {
	case 0:
		AddToLocation{Amount: 1}.Trigger(gs)
	case 1:
		HealAllVillains{Amount: 1}.Trigger(gs)
	case 2:
		AllDiscard{Amount: 1, Prompt: effect.Prompt}.Trigger(gs)
	default:
		DamageAllPlayers{Amount: 1}.Trigger(gs)
	}
}

func inquisitorialSquad() DarkArt {
	return DarkArt{
		Name:    "Inquisitorial Squad",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/inquisitorialsquad.jpg",
		SetId:   "box 1",
		Effects: []Effect{
			GainDetentionToHand{Active: true},
			DamageAllPerDetention{Amount: 1},
		},
	}
}

func menacingGrowl() DarkArt {
	return DarkArt{
		Name:    "Menacing Growl",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/menacinggrowl.jpg",
		SetId:   "box 1",
		Effects: []Effect{
			DamageAllPerMatchingCost{Cost: 3, Amount: 1},
		},
	}
}

func regeneration() DarkArt {
	return DarkArt{
		Name:    "Regeneration",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/regeneration.jpg",
		SetId:   "game 4",
		Effects: []Effect{
			HealAllVillains{Amount: 2},
		},
	}
}
