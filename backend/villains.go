package main

import (
	"log"

	"github.com/google/uuid"
)

func draco() Villain {
	id := int(uuid.New().ID())
	return Villain{
		Name:         "Draco Malfoy",
		Id:           id,
		ImgPath:      "/images/villains/dracomalfoy.jpg",
		SetId:        "Game 1",
		CurDamage:    0,
		MaxHp:        6,
		Effect:       []Effect{DamageActiveIfLocationAdded{Amount: 2, Id: id}},
		DeathEffect:  []Effect{RemoveFromLocation{Amount: 1}},
		playBeforeDA: true,
	}
}

type DamageActiveIfLocationAdded struct {
	Amount int
	Id     int
}

func (effect DamageActiveIfLocationAdded) Trigger(gs *Gamestate) {
	log.Println("calling damageactiveiflocationadded")
	// find player who was active when location got added.
	currentTurn := gs.CurrentTurn

	sub := Subscriber{
		id:              effect.Id,
		messageChan:     make(chan string),
		conditionMet:    "location added",
		conditionFailed: "end turn",
		unsubChan:       eventBroker.Messages,
	}

	go func() {
		eventBroker.Subscribe(sub)
		resChan := make(chan bool)
		go sub.Receive(resChan)
		for {
			res := <-resChan
			if !res {
				break
			}

			Logger("draco wants lock")
			gs.mu.Lock()
			Logger("draco gets lock")
			if currentTurn == gs.CurrentTurn {
				user := gs.CurrentTurn
				stunned := ChangePlayerHealth(user, -effect.Amount, gs)
				if stunned {
					StunPlayer(user, gs)
				}

				SendLobbyUpdate(gs.gameid, gs)
			}
			gs.mu.Unlock()
			Logger("draco releases lock")
		}
	}()
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
	id := int(uuid.New().ID())
	return Villain{
		Name:      "Crabbe and Goyle",
		Id:        id,
		ImgPath:   "/images/villains/crabbeandgoyle.jpg",
		SetId:     "Game 1",
		CurDamage: 0,
		MaxHp:     5,
		Effect: []Effect{
			DamageIfDiscard{Amount: 1, Id: id},
		},
		DeathEffect:  []Effect{AllDrawCards{Amount: 1}},
		playBeforeDA: true,
	}
}

type DamageIfDiscard struct {
	Amount int
	Id     int
}

// only damages active player. not player who discarded
func (effect DamageIfDiscard) Trigger(gs *Gamestate) {
	// find player who was active when location got added.
	currentTurn := gs.CurrentTurn

	sub := Subscriber{
		id:              effect.Id,
		messageChan:     make(chan string),
		conditionMet:    "player discarded",
		conditionFailed: "end turn",
		unsubChan:       eventBroker.Messages,
	}

	go func() {
		eventBroker.Subscribe(sub)
		resChan := make(chan bool)
		go sub.Receive(resChan)

		for {
			res := <-resChan
			if !res {
				break
			}

			Logger("c&g want lock")
			gs.mu.Lock()
			Logger("c&g gets lock")
			if res && currentTurn == gs.CurrentTurn {
				user := gs.CurrentTurn
				// log.Println("c&g deal 1 dmg")
				Logger("cg calling stunned")
				stunned := ChangePlayerHealth(user, -effect.Amount, gs)
				if stunned {
					StunPlayer(user, gs)
				}
				Logger("cg sending lobby update ")

				SendLobbyUpdate(gs.gameid, gs)
			}
			gs.mu.Unlock()
			Logger("c&g release lock")
		}
	}()
}

func bartyCrouchJr() Villain {
	return Villain{
		Name:      "Barty Crouch Jr.",
		Id:        int(uuid.New().ID()),
		ImgPath:   "/images/villains/bartycrouchjr.jpg",
		SetId:     "Game 4",
		CurDamage: 0,
		MaxHp:     7,
		// Set Barty Crouch to Active=false before triggering death effect.
		Active: false,
		// This hero just prevents location from being removed, build into
		// remove from location handler.
		Effect:       []Effect{},
		DeathEffect:  []Effect{RemoveFromLocation{2}},
		playBeforeDA: true,
	}
}

func basilisk() Villain {
	return Villain{
		Name:      "Basilisk",
		Id:        int(uuid.New().ID()),
		ImgPath:   "/images/villains/basilisk.jpg",
		SetId:     "Game 2",
		CurDamage: 0,
		MaxHp:     8,
		Active:    false,
		// This hero just prevents players from drawing, build into
		// remove from draw handler.
		Effect: []Effect{},
		DeathEffect: []Effect{
			AllDrawCards{Amount: 1},
			RemoveFromLocation{1},
		},
		playBeforeDA: true,
	}
}

func bellatrixLestrange() Villain {
	return Villain{
		Name:      "Bellatrix Lestrange",
		Id:        int(uuid.New().ID()),
		ImgPath:   "/images/villains/bellatrixlestrange.jpg",
		SetId:     "Game 6",
		CurDamage: 0,
		MaxHp:     9,
		Active:    false,
		Effect: []Effect{
			RevealDarkArts{Amount: 1},
		},
		DeathEffect: []Effect{
			AllSearchDiscardPileForItem{},
			RemoveFromLocation{Amount: 2},
		},
		playBeforeDA: false,
	}
}

func cornishPixies() Villain {
	return Villain{
		Name:      "Cornish Pixies",
		Id:        int(uuid.New().ID()),
		ImgPath:   "/images/villains/cornishpixies.jpg",
		SetId:     "Box 1",
		CurDamage: 0,
		MaxHp:     6,
		Active:    false,
		Effect: []Effect{
			DamageActiveForEachEvenInHand{Amount: 2},
		},
		DeathEffect: []Effect{
			AllPlayersGainHealth{Amount: 2},
			AllPlayersGainMoney{Amount: 1},
		},
		playBeforeDA: false,
	}
}

type DamageActiveForEachEvenInHand struct {
	Amount int
}

func (effect DamageActiveForEachEvenInHand) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn

	numEvens := 0
	for _, c := range gs.Players[user].Hand {
		if c.Cost%2 == 0 && c.Cost != 0 {
			numEvens++
		}
	}

	damage := -1 * numEvens * effect.Amount
	ChangePlayerHealth(user, damage, gs)
}

func dementor() Villain {
	return Villain{
		Name:      "Dementor",
		Id:        int(uuid.New().ID()),
		ImgPath:   "/images/villains/dementor.jpg",
		SetId:     "Game 3",
		CurDamage: 0,
		MaxHp:     8,
		Active:    false,
		Effect: []Effect{
			DamageCurrentPlayer{Amount: 2},
		},
		DeathEffect: []Effect{
			AllPlayersGainHealth{Amount: 2},
			RemoveFromLocation{Amount: 1},
		},
		playBeforeDA: false,
	}
}

func fenrirGreyback() Villain {
	return Villain{
		Name:      "Fenrir Greyback",
		Id:        int(uuid.New().ID()),
		ImgPath:   "/images/villains/fenrirgreyback.jpg",
		SetId:     "Game 6",
		CurDamage: 0,
		MaxHp:     8,
		Active:    false,
		// Makes players unable to gain health, add a check in changeHealth
		Effect: []Effect{},
		DeathEffect: []Effect{
			AllPlayersGainHealth{Amount: 3},
			RemoveFromLocation{Amount: 2},
		},
		playBeforeDA: false,
	}
}
