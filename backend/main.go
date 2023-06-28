package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// go run ./backend

var appEnv string

var eventBroker = EventBroker{
	mu:          sync.Mutex{},
	Subscribers: map[int]Subscriber{},
	Messages:    make(chan Event),
}

var globalMu = sync.Mutex{}
var states = map[int]*Gamestate{}

func main() {
	os.Setenv("LOG_LEVEL", "debug")

	config = NewConfiguration()
	println(appEnv)

	go eventBroker.StartPublishing()
	RunGameServer()
}

func RunGameServer() {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
			"http://localhost",
			"https://localhost",
			"https://hogwartsbattle.fly.dev",
			"https://lobbymanager.fly.dev",
			"https://www.gamewithyourfriends.dev",
			"https://lobby.gamewithyourfriends.dev",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	r.HandleFunc("/game/startgame", StartGameHandler)
	r.HandleFunc("/game/{id}/firstturn", func(w http.ResponseWriter, r *http.Request) {
		gs, ok := getGsForGameID(r)
		if !ok {
			return
		}
		gs.mu.Lock()
		defer gs.mu.Unlock()
		if !gs.started {
			StartNewTurn(gs.gameid, gs)
		}
	})
	r.HandleFunc("/game/{id}/endturn", func(w http.ResponseWriter, r *http.Request) {
		gs, ok := getGsForGameID(r)
		if !ok || !gs.started {
			return
		}
		EndTurnHandler(w, r, gs)
	})
	r.HandleFunc("/game/{id}/playcard", func(w http.ResponseWriter, r *http.Request) {
		gs, ok := getGsForGameID(r)
		if !ok || !gs.started {
			return
		}
		PlayCardHandler(w, r, gs)
	})
	r.HandleFunc("/game/{id}/getgamestate", func(w http.ResponseWriter, r *http.Request) {
		gs, ok := getGsForGameID(r)
		if !ok {
			return
		}
		GetGamestateHandler(w, r, gs)
	})
	r.HandleFunc("/game/{id}/damagevillain/{villainid}", func(w http.ResponseWriter, r *http.Request) {
		gs, ok := getGsForGameID(r)
		if !ok || !gs.started {
			return
		}
		DamageVillainHandler(w, r, gs)
	})
	r.HandleFunc("/game/{id}/buycard/{cardid}", func(w http.ResponseWriter, r *http.Request) {
		gs, ok := getGsForGameID(r)
		if !ok || !gs.started {
			return
		}
		BuyCardHandler(w, r, gs)
	})
	r.HandleFunc("/game/testserver", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world!")
		log.Println("hello world?")
	})

	handler := c.Handler(r)
	log.Println("starting game server")
	log.Fatal(http.ListenAndServe(":8080", handler))
	log.Println("closing game server")
}

func StartGameHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		StartingPlayers map[string]Player `json:"startingPlayers"`
		TurnOrder       []string          `json:"turnOrder"`
		ID              int               `json:"id"`
	}

	if r.Body == nil {
		log.Println("request body is empty")
		http.Error(w, "Empty request body", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("error decoding JSON:", err.Error(), "body:", r.Body)
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// log.Println("id:", data.ID, "***turn order:", data.TurnOrder, "****players:", data.StartingPlayers)

	villains, villainDeck := CreateVillains()
	gs := &Gamestate{
		Players:         data.StartingPlayers,
		Villains:        villains,
		Locations:       CreateLocations(),
		DarkArts:        CreateDarkArtDeck(),
		MarketDeck:      CreateMarketDeck(),
		Market:          CreateMarket(),
		CurrentDarkArt:  0,
		CurrentLocation: 0,
		DarkArtsPlayed:  []DarkArt{},
		CurrentTurn:     data.TurnOrder[0],
		TurnOrder:       data.TurnOrder,
		villainDeck:     villainDeck,
		turnStats:       TurnStats{},
		mu:              sync.Mutex{},
		gameid:          data.ID,
	}

	for user, p := range gs.Players {
		p.Deck = RonStartingDeck()
		p.Hand = []Card{}
		p.PlayArea = []Card{}
		p.Discard = []Card{}
		gs.Players[user] = p
		RefillHand(user, gs)
	}

	globalMu.Lock()
	defer globalMu.Unlock()

	states[data.ID] = gs

	gs.mu.Lock()
	defer gs.mu.Unlock()
	SendLobbyUpdate(gs.gameid, gs)
}
