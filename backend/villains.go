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
	v := Villain{}

	return v
}
