package game

import (
	"log"
)

type DamageAllPlayers struct {
	Amount      int
	Description string
}

// This definitely doesn't work - changes value in map during range
func (effect DamageAllPlayers) Trigger(gs *Gamestate) {
	for name := range gs.Players {
		player, ok := gs.Players[name]
		if !ok {
			log.Println("error getting player in DamageAllPlayers effect")
			return
		}
		player.Health -= effect.Amount
	}
}

type GainMoney struct {
	Amount      int
	Description string
}

// Gives the current player Amount of money.
func (effect GainMoney) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	player := gs.Players[user]
	player.Money += effect.Amount
	gs.Players[user] = player
}

type ChooseOne struct {
	Effects     []Effect
	Options     []string `json:"options"`
	Description string   `json:"description"`
}

// FIX Lobbyid!
func (effect ChooseOne) Trigger(gs *Gamestate) {
	choice := getUserInput(0, gs.CurrentTurn, effect)

	for i, option := range effect.Options {
		if choice == option {
			effect.Effects[i].Trigger(gs)
		}
	}
}

type GainDamage struct {
	Amount      int
	Description string
}

// Gives the current player Amount of damage
func (effect GainDamage) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	player := gs.Players[user]
	player.Damage += effect.Amount
	gs.Players[user] = player
}

type GainHealth struct {
	Amount      int
	Description string
}

// Gives the current player Amount of health
func (effect GainHealth) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	player := gs.Players[user]
	player.Health += effect.Amount
	gs.Players[user] = player
}

// func (effect GainHealth) Describe() string {
// 	return effect.Description
// }

type GainDamagePerAllyPlayed struct{}

// Gives the current player damage per ally already played this turn.
func (effect GainDamagePerAllyPlayed) Trigger(gs *Gamestate) {
	amount := gs.turnStats.AlliesPlayed

	user := gs.CurrentTurn
	player := gs.Players[user]
	player.Damage += amount
	gs.Players[user] = player
}

type MoneyIfVillainKilled struct {
	Id     int
	Amount int
}

// Gives current player money (this turn) if a villain is killed.
func (effect MoneyIfVillainKilled) Trigger(gs *Gamestate) {
	// check if villain already killed
	if gs.turnStats.VillainsKilled > 0 {
		for name := range gs.Players {
			player, ok := gs.Players[name]
			if !ok {
				return
			}
			player.Damage += effect.Amount
			gs.Players[name] = player
		}
	} else {
		// Check to see if the turn has changed before this can take the lock.
		currentTurn := gs.CurrentTurn

		sub := Subscriber{
			id:              effect.Id,
			messageChan:     make(chan string),
			conditionMet:    "villain killed",
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
				player.Money += effect.Amount
				gs.Players[user] = player

				// FIX lobby id
				SendLobbyUpdate(0, gs)
			}
		}()

	}
}

type RevealDarkArts struct {
	Amount int
}

func (effect RevealDarkArts) Trigger(gs *Gamestate) {

}
