package main

import "github.com/google/uuid"

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
		res := sub.Receive()

		gs.mu.Lock()
		defer gs.mu.Unlock()
		if res && currentTurn == gs.CurrentTurn {
			user := gs.CurrentTurn
			player := gs.Players[user]
			player.Health -= effect.Amount
			gs.Players[user] = player

			// FIX lobby id
			SendLobbyUpdate(0, gs)
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
		DeathEffect:  []Effect{DrawCards{Amount: 1}},
		playBeforeDA: true,
	}
}

type DamageIfDiscard struct {
	Amount int
	Id     int
}

// Currently this will only trigger once per turn, not even for multiple characters.
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
		res := sub.Receive()

		gs.mu.Lock()
		defer gs.mu.Unlock()
		if res && currentTurn == gs.CurrentTurn {
			user := gs.CurrentTurn
			player := gs.Players[user]
			player.Health -= effect.Amount
			gs.Players[user] = player

			// FIX lobby id
			SendLobbyUpdate(0, gs)
		}
	}()
}
