package main

import (
	"log"
)

type DamageAllPlayers struct {
	Amount int
}

func (effect DamageAllPlayers) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		stunned := ChangePlayerHealth(user, -effect.Amount, gs)
		if stunned {
			StunPlayer(user, gs)
		}
	}
}

type DamageCurrentPlayer struct {
	Amount int
}

func (effect DamageCurrentPlayer) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	stunned := ChangePlayerHealth(user, -effect.Amount, gs)
	if stunned {
		StunPlayer(user, gs)
	}
}

type DamageAllPlayersButCurrent struct {
	Amount int
}

func (effect DamageAllPlayersButCurrent) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		if user != gs.CurrentTurn {
			stunned := ChangePlayerHealth(user, -effect.Amount, gs)
			if stunned {
				StunPlayer(user, gs)
			}
		}
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
	Effects []Effect

	// Options is the description given to user. The index of it should be the same as the Effect that it triggers.
	Options     []string `json:"options"`
	Description string   `json:"description"`
}

func (effect ChooseOne) Trigger(gs *Gamestate) {
	choice := getUserInput(gs.gameid, gs.CurrentTurn, effect)

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
	Amount int
}

// Gives the current player Amount of health
func (effect GainHealth) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	ChangePlayerHealth(user, effect.Amount, gs)
	// player := gs.Players[user]
	// player.Health += effect.Amount
	// gs.Players[user] = player
}

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
			resChan := make(chan bool)
			go sub.Receive(resChan)

			for {
				res := <-resChan
				if !res {
					break
				}

				Logger("money if villain killed wants lock")
				gs.mu.Lock()
				if currentTurn == gs.CurrentTurn {
					user := gs.CurrentTurn
					player := gs.Players[user]
					player.Money += effect.Amount
					gs.Players[user] = player

					SendLobbyUpdate(gs.gameid, gs)
				}
				gs.mu.Unlock()
				Logger("money if villain killed releases lock")
			}
		}()

	}
}

type AllPlayersGainHealth struct {
	Amount int
}

func (effect AllPlayersGainHealth) Trigger(gs *Gamestate) {
	for _, p := range gs.Players {
		p.Health += effect.Amount
		gs.Players[p.Name] = p
	}
}

type AllPlayersGainMoney struct {
	Amount int
}

func (effect AllPlayersGainMoney) Trigger(gs *Gamestate) {
	for _, p := range gs.Players {
		p.Money += effect.Amount
		gs.Players[p.Name] = p
	}
}

type ActivePlayerDiscards struct {
	Amount int
}

func (effect ActivePlayerDiscards) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	player := gs.Players[user]
	cardName := AskUserToDiscard(gs.gameid, user, player.Hand, "")

	for i, c := range player.Hand {
		if c.Name == cardName {
			player.Hand = RemoveCardAtIndex(player.Hand, i)
			player.Discard = append(player.Discard, c)
			break
		}
	}

	gs.Players[user] = player

	event := Event{senderId: -1, message: "player discarded", data: user}
	eventBroker.Messages <- event

	// update turnstats
}

type AddToLocation struct {
	Amount int
}

func (effect AddToLocation) Trigger(gs *Gamestate) {
	loc := gs.Locations[gs.CurrentLocation]
	loc.CurControl += effect.Amount
	gs.Locations[gs.CurrentLocation] = loc
	Logger("sending location added event")
	// This is blocking! no one currently listening for event?
	eventBroker.Messages <- LocationAddedEvent
	Logger("sent location added event")

	if loc.CurControl >= loc.MaxControl {
		switch gs.CurrentLocation {
		case len(gs.Locations) - 1:
			log.Println("game over, loser!!!")
		default:
			gs.CurrentLocation += 1
		}
	}
}

type RemoveFromLocation struct {
	Amount int
}

func (effect RemoveFromLocation) Trigger(gs *Gamestate) {
	loc := gs.Locations[gs.CurrentLocation]
	loc.CurControl -= effect.Amount
	if loc.CurControl < 0 {
		loc.CurControl = 0
	}
	gs.Locations[gs.CurrentLocation] = loc
}

type DrawCards struct {
	Amount int
}

func (effect DrawCards) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	DrawXCards(user, gs, effect.Amount)
}

type AllDrawCards struct {
	Amount int
}

func (effect AllDrawCards) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		DrawXCards(user, gs, effect.Amount)
	}
}

type SendGameUpdateEffect struct{}

func (effect SendGameUpdateEffect) Trigger(gs *Gamestate) {
	SendLobbyUpdate(gs.gameid, gs)
}

type HealAnyIfVillainKilled struct {
	Id     int
	Amount int
}

func (effect HealAnyIfVillainKilled) Trigger(gs *Gamestate) {
	healAnyPlayer := HealAnyPlayer{Amount: effect.Amount}

	// check if villain already killed
	if gs.turnStats.VillainsKilled > 0 {
		healAnyPlayer.Trigger(gs)
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
			resChan := make(chan bool)
			go sub.Receive(resChan)

			for {
				res := <-resChan
				if !res {
					break
				}

				Logger("heal if villain killed wants lock")
				gs.mu.Lock()
				if res && currentTurn == gs.CurrentTurn {
					healAnyPlayer.Trigger(gs)

					SendLobbyUpdate(gs.gameid, gs)
				}
				gs.mu.Unlock()
				Logger("heal if villain killed releases lock")
			}
		}()
	}
}

type HealAnyPlayer struct {
	Amount int
}

func (effect HealAnyPlayer) Trigger(gs *Gamestate) {
	playernames := []string{}
	for p := range gs.Players {
		playernames = append(playernames, p)
	}

	choice := AskUserToSelectPlayer(0, gs.CurrentTurn, playernames)
	ChangePlayerHealth(choice, effect.Amount, gs)

	// player := gs.Players[choice]
	// player.Health += effect.Amount
	// gs.Players[choice] = player
}
