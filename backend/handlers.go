package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CardId struct {
	Id int `json:"id"`
}

func PlayCardHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	gameid, user := getIdAndUser(r)

	gs.mu.Lock()
	defer gs.mu.Unlock()

	if gs.CurrentTurn != user {
		log.Println("Not your turn!")
		return
	}

	var cardId CardId
	json.NewDecoder(r.Body).Decode(&cardId)

	for _, c := range gs.Players[user].Hand {
		if c.Id == cardId.Id {
			card := c

			MoveToPlayed(user, c.Id, gs)

			for _, e := range card.Effects {
				e.Trigger(gs)
			}
			switch c.CardType {
			case "spell":
				gs.turnStats.SpellsPlayed += 1
			case "item":
				gs.turnStats.ItemsPlayed += 1
			case "ally":
				gs.turnStats.AlliesPlayed += 1
			}
		}
	}

	SendLobbyUpdate(gameid, gs)
}

func GetGamestateHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	json.NewEncoder(w).Encode(gs)
}

func EndTurnHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	gameid, user := getIdAndUser(r)

	Logger("end turn wants lock")
	gs.mu.Lock()
	defer gs.mu.Unlock()

	if gs.CurrentTurn != user {
		log.Println("not your turn!")
		return
	}

	Logger("end turn sending event")
	eventBroker.Messages <- EndTurnEvent
	Logger("end turn sent event")
	HealStunned(gs)
	MovePlayedToDiscard(user, gs)
	MoveHandToDiscard(user, gs)
	MoneyDamageToZero(user, gs)
	gs.DarkArtsPlayed = []DarkArt{}
	RefillHand(user, gs)
	NextTurnInOrder(gs)
	gs.turnStats = TurnStats{}

	SendLobbyUpdate(gameid, gs)

	// Starting next turn actions.
	Logger("new turn")
	Logger("Before DA villains")
	for i, v := range gs.Villains {
		if !v.Active {
			gs.Villains[i].Active = true
		}
		if v.playBeforeDA {
			for _, e := range v.Effect {
				e.Trigger(gs)
			}
		}
	}
	Logger("triggering locations")
	gs.Locations[gs.CurrentLocation].Effect.Trigger(gs)

	Logger("After DA villains")
	for _, v := range gs.Villains {
		if !v.playBeforeDA {
			for _, e := range v.Effect {
				Logger("triggering " + v.Name)
				e.Trigger(gs)
			}
		}
	}

	SendLobbyUpdate(gameid, gs)
	Logger("end turn releases lock")
}

func DamageVillainHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	gameid, user := getIdAndUser(r)
	vars := mux.Vars(r)
	villainid, err := strconv.Atoi(vars["villainid"])
	if err != nil {
		log.Println("err getting villainid:", err.Error())
	}

	gs.mu.Lock()
	defer gs.mu.Unlock()

	updatedPlayer := gs.Players[user]

	for i, v := range gs.Villains {
		if v.Id == villainid {
			// do nothing if villain is already dead.
			if gs.Villains[i].CurDamage >= v.MaxHp {
				return
			}

			// do nothing if its voldemort and there's others.
			if v.Name == "Voldemort" && len(gs.Villains) != 1 {
				return
			}

			if updatedPlayer.Damage <= 0 && v.Name != "Norbert" {
				log.Println("not enough damage tokens to do this")
				return
			}

			// spend money to damage norbert.
			if v.Name == "Norbert" && updatedPlayer.Money > 0 {
				updatedPlayer.Money -= 1
			} else {
				updatedPlayer.Damage -= 1
			}

			gs.Villains[i].CurDamage += 1
			gs.Players[user] = updatedPlayer

			// check if villain is now dead.
			if gs.Villains[i].CurDamage >= v.MaxHp {
				gs.Villains[i].Active = false
				// trigger villain death effect.
				for _, effect := range gs.Villains[i].DeathEffect {
					effect.Trigger(gs)
				}

				// remove villain, get new one
				newVillains := RemoveVillainAtIndex(gs.Villains, i)
				newVillains = AddNewActiveVillain(newVillains, gs)
				gs.Villains = newVillains
				eventBroker.Messages <- Event{senderId: -1, message: "villain killed"}
			}
			break
		}
	}

	SendLobbyUpdate(gameid, gs)
}

func BuyCardHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	gameid, user := getIdAndUser(r)
	vars := mux.Vars(r)
	cardid, err := strconv.Atoi(vars["cardid"])
	if err != nil {
		log.Println("err getting cardid:", err.Error())
	}

	gs.mu.Lock()
	defer gs.mu.Unlock()

	if gs.CurrentTurn != user {
		log.Println("not your turn!")
		return
	}

	player := gs.Players[user]
	for i, c := range gs.Market {
		if c.Id == cardid && player.Money >= c.Cost {
			player.Money -= c.Cost
			player.Discard = append(player.Discard, c)
			if c.Cost >= 4 {
				eventBroker.Messages <- DoloresUmbridgeTrigger
			}
			gs.Market[i] = RefillMarket(c.Name)
		}
	}

	gs.Players[user] = player
	// don't remove from market for now.

	SendLobbyUpdate(gameid, gs)
}
