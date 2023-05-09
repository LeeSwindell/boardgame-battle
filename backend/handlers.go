package game

import (
	"encoding/json"
	"log"
	"net/http"
)

type CardId struct {
	Id int `json:"id"`
}

func PlayCardHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	gameid, user := getIdAndUser(r)
	if gs.CurrentTurn.Name != user {
		log.Println("Not your turn!")
		return
	}

	var cardId CardId
	json.NewDecoder(r.Body).Decode(&cardId)

	gs.mu.Lock()
	defer gs.mu.Unlock()
	for i, c := range gs.Players[user].Hand {
		if c.Id == cardId.Id {
			card := c
			log.Println("playing card:", card.Name)
			for _, e := range card.Effects {
				log.Println("triggering an effect:", e)
				e.Trigger(gs)
			}
			MoveToPlayed(user, i, gs)
		}
	}

	SendLobbyUpdate(gameid, gs)
}

func GetGamestateHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	// id, _ := getUserAndId(r)

	gs.mu.Lock()
	defer gs.mu.Unlock()

	json.NewEncoder(w).Encode(gs)
}

func EndTurnHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	gameid, user := getIdAndUser(r)

	gs.mu.Lock()
	defer gs.mu.Unlock()

	MoveToDiscard(user, gs)
	MoneyDamageToZero(user, gs)
	Draw5Cards(user, gs)
	// change turn order

	SendLobbyUpdate(gameid, gs)
}
