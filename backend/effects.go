package game

import (
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

// func (effect DamageAllPlayers) Describe() string {
// 	return fmt.Sprintf("Damage all players for %d", effect.Amount)
// }

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

// func (effect GainMoney) Describe() string {
// 	return effect.Description
// }

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

// func (effect ChooseOne) Describe() string {
// 	return effect.Description
// }

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

// func (effect GainDamage) Describe() string {
// 	return effect.Description
// }

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

// func (effect GainHealth) Describe() string {
// 	return effect.Description
// }

type GainDamagePerAllyPlayed struct{}

func (effect GainDamagePerAllyPlayed) Trigger(gs *Gamestate) {
	damage := gs.turnStats.AlliesPlayed

	for name := range gs.Players {
		player, ok := gs.Players[name]
		if !ok {
			return
		}
		player.Damage += damage
		gs.Players[name] = player
	}
}

type MoneyIfVillainKilled struct {
	Id     int
	Amount int
}

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
		currentTurn := gs.CurrentTurn.Name

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
			if res && currentTurn == gs.CurrentTurn.Name {
				for name := range gs.Players {
					player, ok := gs.Players[name]
					if !ok {
						return
					}
					player.Money += effect.Amount
					gs.Players[name] = player
				}

				// FIX lobby id
				SendLobbyUpdate(0, gs)
			}
		}()

		// Current problem - the gs lock is being held by the playcard handler, but this condition needs to be resolved asynchronously when it is met. If this func changes the gamestate, it could create a race condition. should probably move all locks to within the triggers that change gamestate, or create a new type for functions that alter gamestate maybe?
	}
}
