package game

import (
	"encoding/json"
	"log"
	"net/http"
)

func PlayCardHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	id, user := getUserAndId(r)
	if gs.CurrentTurn.Name != user {
		log.Println("Not your turn!")
		return
	}

	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)
	cardname := data["cardname"]

	card := cards[cardname.(string)]

	gs.mu.Lock()
	defer gs.mu.Unlock()
	for _, e := range card.Effects {
		e.Trigger(gs)
	}

	SendLobbyUpdate(id, gs)
}
