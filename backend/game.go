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

var states = map[int]*Gamestate{}

func RunGameServer() {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	// states := map[int]*Gamestate{}

	r.HandleFunc("/{id}/endturn", func(w http.ResponseWriter, r *http.Request) {
		id, _ := getIdAndUser(r)
		gs := states[id]
		EndTurnHandler(w, r, gs)
	})
	r.HandleFunc("/{id}/playcard", func(w http.ResponseWriter, r *http.Request) {
		id, _ := getIdAndUser(r)
		gs := states[id]
		PlayCardHandler(w, r, gs)
	})
	r.HandleFunc("/{id}/getgamestate", func(w http.ResponseWriter, r *http.Request) {
		id, _ := getIdAndUser(r)
		gs := states[id]
		GetGamestateHandler(w, r, gs)
	})
	r.HandleFunc("/{id}/damagevillain/{villainid}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := getIdAndUser(r)
		gs := states[id]
		DamageVillainHandler(w, r, gs)
	})
	r.HandleFunc("/{id}/buycard/{cardid}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := getIdAndUser(r)
		gs := states[id]
		BuyCardHandler(w, r, gs)
	})

	handler := c.Handler(r)
	log.Println("starting game server")
	log.Fatal(http.ListenAndServe(":8080", handler))
	log.Println("closing game server")
}

func StartGame(players map[string]Player, turnOrder []string, lobbyid int) {
	gs := &Gamestate{
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
		DrawXCards(user, gs, 5)
	}

	states[lobbyid] = gs
	// log.Println("**************", lobbyid)

	go eventBroker.StartPublishing()
	// go RunGameServer(&gs)

	StartNewTurn(0, gs)
}
