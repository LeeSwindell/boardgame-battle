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
	r.HandleFunc("/{id}/buycard/{cardid}", func(w http.ResponseWriter, r *http.Request) {
		BuyCardHandler(w, r, gs)
	})

	handler := c.Handler(r)
	log.Println("starting game engine on port 8080!")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func StartGame(players map[string]Player, turnOrder []string) {
	gs := Gamestate{
		Players:         players,
		Villains:        CreateVillains(),
		Locations:       CreateLocations(),
		DarkArts:        CreateDarkArtDeck(),
		MarketDeck:      CreateMarketDeck(),
		Market:          CreateMarket(),
		CurrentDarkArt:  0,
		CurrentLocation: 0,
		DarkArtsPlayed:  []DarkArt{},
		CurrentTurn:     turnOrder[0],
		TurnOrder:       turnOrder,
		turnStats:       TurnStats{},
		mu:              sync.Mutex{},
	}

	for _, p := range gs.Players {
		user := p.Name
		DrawXCards(user, &gs, 5)
	}

	go eventBroker.StartPublishing()
	go RunGameServer(&gs)

	StartNewTurn(0, &gs)
}
