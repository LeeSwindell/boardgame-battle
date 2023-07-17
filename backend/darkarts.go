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
		effect: []Effect{
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
		effect: []Effect{
			DamageCurrentPlayer{Amount: 1},
			ActivePlayerDiscards{Amount: 1, Prompt: "Flipendo! Discard a card", Cause: "darkart"},
		},
	}
}

func heWhoMustNotBeNamed() DarkArt {
	return DarkArt{
		Name:    "He-Who-Must-Not-Be-Named",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/hewhomustnotbenamed.jpg",
		SetId:   "game 1",
		effect: []Effect{
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
		effect: []Effect{
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
		effect: []Effect{
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
		effect: []Effect{
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
		effect: []Effect{
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
		AllDiscard{Amount: 1, Prompt: effect.Prompt, Cause: "darkart"}.Trigger(gs)
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
		effect: []Effect{
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
		effect: []Effect{
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
		effect: []Effect{
			HealAllVillains{Amount: 2},
		},
	}
}

func crucio() DarkArt {
	return DarkArt{
		Name:    "Crucio!",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/crucio.jpg",
		SetId:   "game 5",
		effect: []Effect{
			DamageCurrentPlayer{Amount: 1},
			LoadDarkArtEffect{},
			RevealDarkArts{Amount: 1},
		},
	}
}

type LoadDarkArtEffect struct{}

func (effect LoadDarkArtEffect) Trigger(gs *Gamestate) {
	LoadNewDarkArt(gs)
}

func fiendfyre() DarkArt {
	return DarkArt{
		Name:    "Fiendfyre",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/fiendfyre.jpg",
		SetId:   "game 7",
		effect: []Effect{
			DamageAllPlayers{Amount: 3},
		},
	}
}

func morsmordre() DarkArt {
	return DarkArt{
		Name:    "Morsmordre!",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/morsmordre.jpg",
		SetId:   "game 4",
		effect: []Effect{
			DamageAllPlayers{Amount: 1},
			AddToLocation{Amount: 1},
			CheckForDeathEater{},
		},
	}
}

type CheckForDeathEater struct{}

func (effect CheckForDeathEater) Trigger(gs *Gamestate) {
	for _, v := range gs.Villains {
		if v.Active && v.Name == "Death Eater" {
			DamageAllPlayers{Amount: 1}.Trigger(gs)
		}
	}
}

func blastended() DarkArt {
	return DarkArt{
		Name:    "Blast-ended",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/blastended.jpg",
		SetId:   "box 1",
		effect: []Effect{
			PreviousHeroDoesX{
				ChangeStats{
					AmountHealth:    -1,
					AmountToDiscard: 1,
					DiscardPrompt:   "Blast-ended Skrewt! Discard a card:",
					Cause:           "darkart",
				},
			},
		},
	}
}

func educationalDecree() DarkArt {
	return DarkArt{
		Name:    "Educational Decree",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/educationaldecree.jpg",
		SetId:   "game 5",
		effect: []Effect{
			DamageActivePerCardGreaterThanCost{
				Amount: 1,
				Cost:   4,
			},
		},
	}
}

func imperio() DarkArt {
	return DarkArt{
		Name:    "Imperio!",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/imperio.jpg",
		SetId:   "game 5",
		effect: []Effect{
			ActivePlayerSelectsOtherPlayerToDoX{
				ChangeStats{AmountHealth: -2},
			},
			LoadDarkArtEffect{},
			RevealDarkArts{Amount: 1},
		},
	}
}

func legilimency() DarkArt {
	return DarkArt{
		Name:    "Legilimency",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/legilimency.jpg",
		SetId:   "game 5",
		effect: []Effect{
			AllRevealTopCardAndX{
				X: legilimencyEffect,
			},
		},
	}
}

// Only for a revealed card on top of deck.
func legilimencyEffect(card Card, user string, gs *Gamestate) {
	if card.CardType != "spell" {
		return
	}
	ChangeStats{Target: user, AmountHealth: -2}.Trigger(gs)
	player := gs.Players[user]
	if card.onDiscard != nil {
		gs.Players[user] = player
		card.onDiscard(user, gs)
		player = gs.Players[user]
	}
	player.Discard = append(player.Discard, card)
	player.Deck = player.Deck[:len(player.Deck)-1]
	gs.Players[user] = player
	if player.Proficiency == "Defense Against the Dark Arts" {
		ChangeStats{Target: user, AmountDamage: 1, AmountHealth: 1}.Trigger(gs)
	}
	eventBroker.Messages <- PlayerDiscarded
}

func obliviate() DarkArt {
	return DarkArt{
		Name:    "Obliviate!",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/obliviate.jpg",
		SetId:   "game 2",
		effect: []Effect{
			AllChooseOneTargeted{
				EffectTargeting: []func(target string, effect Effect) Effect{
					TargetDiscardACard,
					TargetCreateStats,
				},
				Effects: []Effect{
					DiscardACard{Cardtype: "spell", Prompt: "Discard a spell", Cause: "darkart"},
					ChangeStats{AmountHealth: -2},
				},
				Options:     []string{"Discard a spell", "Lose 2 Health"},
				Description: "Obliviate! All heroes choose one:",
			},
		},
	}
}

func oppugno() DarkArt {
	return DarkArt{
		Name:    "Oppugno!",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/oppugno.jpg",
		SetId:   "game 3",
		effect: []Effect{
			AllRevealTopCardAndX{
				X: OppugnoEffect,
			},
		},
	}
}

func OppugnoEffect(card Card, user string, gs *Gamestate) {
	if card.Cost == 0 {
		return
	}

	ChangeStats{Target: user, AmountHealth: -2}.Trigger(gs)
	player := gs.Players[user]
	if card.onDiscard != nil {
		gs.Players[user] = player
		card.onDiscard(user, gs)
		player = gs.Players[user]
	}
	player.Discard = append(player.Discard, card)
	player.Deck = player.Deck[:len(player.Deck)-1]
	gs.Players[user] = player
	if player.Proficiency == "Defense Against the Dark Arts" {
		ChangeStats{Target: user, AmountDamage: 1, AmountHealth: 1}.Trigger(gs)
	}

	eventBroker.Messages <- PlayerDiscarded
}

func petrification() DarkArt {
	return DarkArt{
		Name:    "Petrification",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/petrification.jpg",
		SetId:   "game 1",
		effect: []Effect{
			DamageAllPlayers{Amount: 1},
		},
	}
}

func poison() DarkArt {
	return DarkArt{
		Name:    "Poison",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/poison.jpg",
		SetId:   "game 2",
		effect: []Effect{
			AllChooseOneTargeted{
				EffectTargeting: []func(target string, effect Effect) Effect{
					TargetDiscardACard,
					TargetCreateStats,
				},
				Effects: []Effect{
					DiscardACard{Cardtype: "ally", Prompt: "Discard an ally", Cause: "darkart"},
					ChangeStats{AmountHealth: -2},
				},
				Options:     []string{"Discard an ally", "Lose 2 Health"},
				Description: "Poison (which somehow everyone drank)! All heroes choose one:",
			},
		},
	}
}

func ragingTroll() DarkArt {
	return DarkArt{
		Name:    "Raging Troll",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/ragingtroll.jpg",
		SetId:   "box 1",
		effect: []Effect{
			NextHeroDoesX{
				ChangeStats{
					AmountHealth: -2,
				},
			},
			AddToLocation{Amount: 1},
		},
	}
}

func relashio() DarkArt {
	return DarkArt{
		Name:    "Relashio!",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/relashio.jpg",
		SetId:   "game 2",
		effect: []Effect{
			AllChooseOneTargeted{
				EffectTargeting: []func(target string, effect Effect) Effect{
					TargetDiscardACard,
					TargetCreateStats,
				},
				Effects: []Effect{
					DiscardACard{Cardtype: "item", Prompt: "Discard an item", Cause: "darkart"},
					ChangeStats{AmountHealth: -2},
				},
				Options:     []string{"Discard an item", "Lose 2 Health"},
				Description: "Relashio! All heroes choose one:",
			},
		},
	}
}

func sectumsempra() DarkArt {
	return DarkArt{
		Name:    "Sectumsempra!",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/sectumsempra.jpg",
		SetId:   "game 6",
		effect: []Effect{
			DamageAllPlayers{Amount: 2},
		},
	}
}

func slugulusEructo() DarkArt {
	return DarkArt{
		Name:    "Slugulus Eructo!",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/sluguluseructo.jpg",
		SetId:   "game 6",
		effect: []Effect{
			DamageAllPerCreature{Amount: 1},
		},
	}
}

func tarantallegra() DarkArt {
	return DarkArt{
		Name:    "Tarantallegra!",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/tarantallegra.jpg",
		SetId:   "game 3",
		effect: []Effect{
			DamageCurrentPlayer{Amount: 1},
		},
	}
}
