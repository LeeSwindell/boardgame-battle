package game

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

}
