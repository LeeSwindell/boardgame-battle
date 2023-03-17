package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

// to start the server
// go run github.com/leeswindell/boardgame-battle/lobbymanager

var hub = newHub()

var players = PlayersInLobby{
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

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// func GetLobbiesHandler(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Access-Control-Allow-Origin", "*")

// 	res := Lobbies{
// 		[]Lobby{
// 			{
// 				ID:   1,
// 				Name: "Casual Lobby",
// 				Host: "pico paco",
// 			},
// 			{
// 				ID:   2,
// 				Name: "superior lobby",
// 				Host: "leeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
// 			},
// 			{
// 				ID:   3,
// 				Name: "fun time Lobby",
// 				Host: "boilly bill",
// 			},
// 		},
// 	}

// 	l, err := json.Marshal(res)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	w.Write(l)
// }

// func AddClientHandler(w http.ResponseWriter, r *http.Request) {
// 	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
// 	pid := getUniquePlayerId()
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	client := &Client{hub: hub, conn: conn, pid: pid}
// 	client.hub.register <- client

// 	println("player ", pid.String(), " is connecting")
// 	conn.WriteJSON(players)

// 	// start go routines for the client
// 	go client.readPump()
// 	// go client.writePump()
// }

// func RefreshLobbyHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	res, err := json.Marshal(players)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	w.Write(res)
// }

// func AddPlayerHandler(w http.ResponseWriter, r *http.Request) {
// 	var player Player
// 	err := json.NewDecoder(r.Body).Decode(&player)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	players.Players = append(players.Players, player)
// 	hub.SendRefreshRequest()

// 	w.WriteHeader(http.StatusCreated)
// }

func getUniquePlayerId() uuid.UUID {
	return uuid.New()
}

func main() {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	go hub.run()

	r.HandleFunc("/lobbies", GetLobbiesHandler)
	r.HandleFunc("/lobby/{id}", func(w http.ResponseWriter, r *http.Request) {
		AddClientHandler(w, r)
	})
	r.HandleFunc("/lobby/{id}/refresh", RefreshLobbyHandler)
	r.HandleFunc("/lobby/{id}/addplayer", AddPlayerHandler)

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
