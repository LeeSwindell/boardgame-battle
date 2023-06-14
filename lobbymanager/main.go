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

var appEnv string

// Keep global mutex or attach to types?
var globalMu sync.Mutex
var hub = newHub()
var lobbies = map[int]Lobby{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Lobby struct {
	ID      int           `json:"id"`
	Name    string        `json:"name"`
	Host    string        `json:"host"`
	Players []LobbyPlayer `json:"players"`
}

type LobbyPlayer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Character string `json:"character"`
}

func getUniquePlayerId() uuid.UUID {
	return uuid.New()
}

func getUniqueLobbyId() int {
	return int(uuid.New().ID())
}

func main() {
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
			"https://hogwartsbattle.fly.dev",
			"https://hogwartsbackend.fly.dev",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	go hub.run()
	go messageBroadcaster.Broadcast()

	r.HandleFunc("/sessionid", sessionidHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/lobbies", GetLobbiesHandler)
	r.HandleFunc("/connectsocket", AddClientHandler)
	r.HandleFunc("/connectsocket/{username}", AddClientWithUsernameHandler)
	r.HandleFunc("/lobby/create", CreateLobbyHandler)
	r.HandleFunc("/lobby/{id}/join", JoinLobbyHandler)
	r.HandleFunc("/lobby/{id}/refresh", RefreshLobbyHandler)
	r.HandleFunc("/lobby/{id}/setchar", SetCharHandler)
	r.HandleFunc("/lobby/{id}/leave", LeaveLobbyHandler)
	r.HandleFunc("/lobby/{id}/startgame", StartGameHandler)
	r.HandleFunc("/game/{id}/refreshgamestate", RefreshGamestateHandler)
	r.HandleFunc("/game/{id}/getuserinput/{user}", GetUserInputHandler)
	r.HandleFunc("/game/{id}/askusertodiscard/{user}", AskUserToDiscardHandler)
	r.HandleFunc("/game/{id}/askusertoselectplayer/{user}", AskUserToSelectPlayerHandler)
	r.HandleFunc("/game/{id}/submituserchoice", SubmitUserChoiceHandler)

	// Start the game server
	// go game.RunGameServer()

	handler := c.Handler(r)
	log.Println("running lobby manager on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
