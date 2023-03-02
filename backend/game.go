package game

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CreateGame() Gamestate {
	player1 := Player{
		name:      "leerocks",
		character: "Ron",
		health:    10,
		damage:    0,
		money:     0,
	}
	player2 := Player{
		name:      "ollieisdabomb",
		character: "Neville",
		health:    10,
		damage:    2,
		money:     0,
	}
	villain1 := Villain{name: "moldyvort"}

	gs := Gamestate{
		players:  []Player{player1, player2},
		villains: []Villain{villain1},
	}

	return gs
}

func GameServer() {
	r := mux.NewRouter()
	r.HandleFunc("/initialize", StartGameHandler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func StartGameHandler(w http.ResponseWriter, r *http.Request) {
}
