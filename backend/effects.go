package game

import (
	"fmt"
	"log"
)

type DamageAllPlayers struct {
	Amount      int
	Description string
}

func (effect DamageAllPlayers) Trigger(gs *Gamestate) {
	for name, _ := range gs.Players {
		player, ok := gs.Players[name]
		if !ok {
			log.Println("error getting player in DamageAllPlayers effect")
			return
		}
		player.Health -= effect.Amount
	}
}

func (effect DamageAllPlayers) Describe() string {
	return fmt.Sprintf("Damage all players for %d", effect.Amount)
}

type GainMoney struct {
	Amount      int
	Description string
}

// FIX - give only current turn player
func (effect GainMoney) Trigger(gs *Gamestate) {
	for name := range gs.Players {
		player, ok := gs.Players[name]
		if !ok {
			return
		}
		player.Money += effect.Amount
		gs.Players[name] = player
	}
}

func (effect GainMoney) Describe() string {
	return effect.Description
}

type ChooseOne struct {
	Effects     []Effect
	Options     []string `json:"options"`
	Description string   `json:"description"`
}

// FIX Lobbyid!
func (effect ChooseOne) Trigger(gs *Gamestate) {
	choice := getUserInput(0, gs.CurrentTurn.Name, effect)

	for i, option := range effect.Options {
		if choice == option {
			effect.Effects[i].Trigger(gs)
		}
	}
}

func (effect ChooseOne) Describe() string {
	return effect.Description
}

type GainDamage struct {
	Amount      int
	Description string
}

func (effect GainDamage) Trigger(gs *Gamestate) {
	for name := range gs.Players {
		player, ok := gs.Players[name]
		if !ok {
			return
		}
		player.Damage += effect.Amount
		gs.Players[name] = player
	}
}

func (effect GainDamage) Describe() string {
	return effect.Description
}

type GainHealth struct {
	Amount      int
	Description string
}

func (effect GainHealth) Trigger(gs *Gamestate) {
	for name := range gs.Players {
		player, ok := gs.Players[name]
		if !ok {
			return
		}
		player.Health += effect.Amount
		gs.Players[name] = player
	}
}

func (effect GainHealth) Describe() string {
	return effect.Description
}
