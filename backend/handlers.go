package game

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

			MoveToPlayed(user, c.Id, gs)
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

	gs.mu.Lock()
	defer gs.mu.Unlock()

	if gs.CurrentTurn != user {
		log.Println("not your turn!")
		return
	}

	//
	// Potentially refactor into 2 separate events, end turn and start turn.
	// Release lock between these two events to any ongoing effects like villains
	// Let them trigger things that may have happened during an end or start turn event
	// Label Villains to get played before dark arts or after - depending on type ?
	//

	eventBroker.Messages <- EndTurnEvent
	MovePlayedToDiscard(user, gs)
	MoveHandToDiscard(user, gs)
	MoneyDamageToZero(user, gs)
	gs.DarkArtsPlayed = []DarkArt{}
	DrawXCards(user, gs, 5)
	NextTurnInOrder(gs)
	gs.turnStats = TurnStats{}

	SendLobbyUpdate(gameid, gs)

	// Starting next turn actions.
	for _, v := range gs.Villains {
		if v.playBeforeDA {
			for _, e := range v.Effect {
				e.Trigger(gs)
			}
		}
	}
	gs.Locations[gs.CurrentLocation].Effect.Trigger(gs)
	for _, v := range gs.Villains {
		if !v.playBeforeDA {
			for _, e := range v.Effect {
				e.Trigger(gs)
			}
		}
	}

	SendLobbyUpdate(gameid, gs)

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
	if updatedPlayer.Damage <= 0 {
		log.Println("not enough damage tokens to do this")
		return
	}

	for i, v := range gs.Villains {
		if v.Id == villainid {
			if gs.Villains[i].CurDamage == v.MaxHp {
				// do nothing if villain is already dead.
				return
			}

			gs.Villains[i].CurDamage += 1
			updatedPlayer.Damage -= 1
			gs.Players[user] = updatedPlayer

			// check if villain is now dead.
			if gs.Villains[i].CurDamage == v.MaxHp {
				// remove villain, get new one
				eventBroker.Messages <- Event{senderId: -1, message: "villain killed"}
			}
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
			gs.Market[i] = RefillMarket(c.Name)
		}
	}

	gs.Players[user] = player
	// don't remove from market for now.

	SendLobbyUpdate(gameid, gs)
}
