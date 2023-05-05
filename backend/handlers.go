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
	lobbyId, user := getUserAndId(r)
	if gs.CurrentTurn.Name != user {
		log.Println("Not your turn!")
		return
	}

	var cardId CardId
	json.NewDecoder(r.Body).Decode(&cardId)

	gs.mu.Lock()
	defer gs.mu.Unlock()
	for _, c := range gs.Players[user].Hand.Cards {
		if c.Id == cardId.Id {
			card := c
			log.Println("playing card:", card.Name)
			for _, e := range card.Effects {
				log.Println("triggering an effect:", e)
				e.Trigger(gs)
			}
		}
	}

	SendLobbyUpdate(lobbyId, gs)
}

func GetGamestateHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	// id, _ := getUserAndId(r)

	gs.mu.Lock()
	defer gs.mu.Unlock()

	json.NewEncoder(w).Encode(gs)
}
