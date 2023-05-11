package game

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var eventBroker = EventBroker{
	mu:          sync.Mutex{},
	Subscribers: map[int]Subscriber{},
	Messages:    make(chan Event),
}

func RunGameServer(gs *Gamestate) {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	r.HandleFunc("/{id}/endturn", func(w http.ResponseWriter, r *http.Request) {
		EndTurnHandler(w, r, gs)
	})
	r.HandleFunc("/{id}/playcard", func(w http.ResponseWriter, r *http.Request) {
		PlayCardHandler(w, r, gs)
	})
	r.HandleFunc("/{id}/getgamestate", func(w http.ResponseWriter, r *http.Request) {
		GetGamestateHandler(w, r, gs)
	})
	r.HandleFunc("/{id}/damagevillain/{villainid}", func(w http.ResponseWriter, r *http.Request) {
		DamageVillainHandler(w, r, gs)
	})

	handler := c.Handler(r)
	log.Println("starting game engine on port 8080!")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// FIX turn order
func StartGame(players map[string]Player, turnOrder []string) {
	gs := Gamestate{
		Players:     players,
		Villains:    CreateVillains(),
		Locations:   []Location{},
		CurrentTurn: players[turnOrder[0]],
		turnStats:   TurnStats{},
		mu:          sync.Mutex{},
	}

	for _, p := range gs.Players {
		user := p.Name
		Draw5Cards(user, &gs)
	}

	go eventBroker.StartPublishing()
	go RunGameServer(&gs)
}
