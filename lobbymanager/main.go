package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// to start the server
// go run github.com/leeswindell/boardgame-battle/lobbymanager

var hub = newHub()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Lobby struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Host string `json:"host"`
}

type Lobbies struct {
	Lobbies []Lobby `json:"lobbies"`
}

type PlayersInLobby struct {
	Players []Player `json:"players"`
}

type Player struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Character string `json:"character"`
}

func GetLobbiesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	res := Lobbies{
		[]Lobby{
			{
				ID:   1,
				Name: "Casual Lobby",
				Host: "pico paco",
			},
			{
				ID:   2,
				Name: "superior lobby",
				Host: "leeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
			},
			{
				ID:   3,
				Name: "fun time Lobby",
				Host: "boilly bill",
			},
		},
	}

	l, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}

	w.Write(l)
}

func LobbyHandler(w http.ResponseWriter, r *http.Request, players *PlayersInLobby) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	playerId := getUniquePlayerId()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	if conn != nil {
		println("player ", playerId, " has connected")
		conn.WriteJSON(players)
	}

}

func getUniquePlayerId() uuid.UUID {
	return uuid.New()
}

func main() {
	r := mux.NewRouter()

	players := PlayersInLobby{
		[]Player{
			{
				ID:        1,
				Name:      "pico",
				Character: "Harry",
			},
			{
				ID:        2,
				Name:      "paco",
				Character: "Ron",
			},
		},
	}

	r.HandleFunc("/lobbies", GetLobbiesHandler)
	r.HandleFunc("/lobby/{id}", func(w http.ResponseWriter, r *http.Request) {
		LobbyHandler(w, r, &players)
	})

	log.Fatal(http.ListenAndServe(":8000", r))
}
