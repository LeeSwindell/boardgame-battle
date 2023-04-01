package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

// to start the server
// go run github.com/leeswindell/boardgame-battle/lobbymanager

// Keep global mutex or attach to types?
var globalMu sync.Mutex
var hub = newHub()
var lobbyNumber int
var lobbies Lobbies

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Lobby struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Host    string   `json:"host"`
	Players []Player `json:"players"`
}

type Lobbies struct {
	Lobbies []Lobby `json:"lobbies"`
}

type Player struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Character string `json:"character"`
}

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func getUniquePlayerId() uuid.UUID {
	return uuid.New()
}

func main() {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	go hub.run()

	r.HandleFunc("/sessionid", sessionidHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/lobbies", GetLobbiesHandler)
	r.HandleFunc("/connectsocket", AddClientHandler)
	r.HandleFunc("/lobby/create", CreateLobbyHandler)
	r.HandleFunc("/lobby/{id}/join", JoinLobbyHandler)
	r.HandleFunc("/lobby/{id}/refresh", RefreshLobbyHandler)
	r.HandleFunc("/lobby/{id}/setchar", SetCharHandler)
	r.HandleFunc("/lobby/{id}/addplayer", AddPlayerHandler)

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
