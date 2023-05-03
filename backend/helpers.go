package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// returns the lobbyid and username associated with a request
func getUserAndId(r *http.Request) (int, string) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("err getting id:", err.Error())
	}
	user := r.Header.Get("Authorization")
	if user == "" {
		log.Println("empty username")
	}

	return id, user
}

func SendLobbyUpdate(id int, gs *Gamestate) {
	url := fmt.Sprintf("http://localhost:8000/game/%d/refreshgamestate", id)
	data, err := json.Marshal(gs)
	if err != nil {
		log.Println("err marshaling gamestate:", err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Println("err sending lobby update:", err.Error())
	}

	client := http.Client{}
	client.Do(req)
}
