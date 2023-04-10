package game

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func endTurnHandler(w http.ResponseWriter, r *http.Request) {

}

func RunGameServer(gs *Gamestate) {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	r.HandleFunc("/{id}/endturn", endTurnHandler)
	r.HandleFunc("/{id}/playcard", func(w http.ResponseWriter, r *http.Request) {
		PlayCardHandler(w, r, gs)
	})

	handler := c.Handler(r)
	log.Println("starting game engine on port 8080!")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func StartGame(players []Player) {
	for i := range players {
		players[i].Health = 10
		players[i].Damage = 0
		players[i].Money = 0
	}

	gs := Gamestate{
		Players:     players,
		Villains:    []Villain{},
		Locations:   []Location{},
		CurrentTurn: players[0],
		mu:          sync.Mutex{},
	}

	go RunGameServer(&gs)
}
