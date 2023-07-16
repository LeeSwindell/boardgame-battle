package main

import (
	"fmt"
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

var config = NewConfiguration()

type Configuration struct {
	BackendURL  string
	FrontendURL string
}

func NewConfiguration() *Configuration {
	mode := appEnv
	switch mode {
	case "dev":
		return &Configuration{
			BackendURL:  "http://localhost:8080/game",
			FrontendURL: "http://localhost:5173",
		}
	case "prod":
		return &Configuration{
			BackendURL:  "https://www.gamewithyourfriends.dev/game",
			FrontendURL: "https://www.gamewithyourfriends.dev",
		}
	default:
		log.Println("************* NO APP ENV PROVIDED ***********")
		// use dev env by default.
		return &Configuration{
			BackendURL:  "http://localhost:8080/game",
			FrontendURL: "http://localhost:5173",
		}
	}
}

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
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Character   string `json:"character"`
	Proficiency string `json:"proficiency"`
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
			"http://localhost",
			"https://localhost",
			"https://hogwartsbattle.fly.dev",
			"https://hogwartsbackend.fly.dev",
			"https://www.gamewithyourfriends.dev",
			"https://game.gamewithyourfriends.dev",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	go hub.run()
	go messageBroadcaster.Broadcast()

	r.HandleFunc("/lm/login", loginHandler)
	r.HandleFunc("/lm/sessionid", sessionidHandler)
	r.HandleFunc("/lm/lobbies", GetLobbiesHandler)
	r.HandleFunc("/lm/connectsocket", AddClientHandler)
	r.HandleFunc("/lm/connectsocket/{username}", AddClientWithUsernameHandler)
	r.HandleFunc("/lm/lobby/create", CreateLobbyHandler)
	r.HandleFunc("/lm/lobby/{id}/join", JoinLobbyHandler)
	r.HandleFunc("/lm/lobby/{id}/refresh", RefreshLobbyHandler)
	r.HandleFunc("/lm/lobby/{id}/setchar", SetCharHandler)
	r.HandleFunc("/lm/lobby/{id}/setprof", SetProfHandler)
	r.HandleFunc("/lm/lobby/{id}/leave", LeaveLobbyHandler)
	r.HandleFunc("/lm/lobby/{id}/startgame", StartGameHandler)
	r.HandleFunc("/lm/game/{id}/refreshgamestate", RefreshGamestateHandler)
	r.HandleFunc("/lm/game/{id}/getuserinput/{user}", GetUserInputHandler)
	r.HandleFunc("/lm/game/{id}/askusertodiscard/{user}", AskUserToDiscardHandler)
	r.HandleFunc("/lm/game/{id}/askusertoselectplayer/{user}", AskUserToSelectPlayerHandler)
	r.HandleFunc("/lm/game/{id}/askusertoselectcard/{user}", AskUserToSelectCardHandler)
	r.HandleFunc("/lm/game/{id}/submituserchoice", SubmitUserChoiceHandler)
	r.HandleFunc("/lm/testserver", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from the lobby manager!")
		log.Println("working!")
	})

	// Start the game server
	// go game.RunGameServer()

	handler := c.Handler(r)
	log.Println("running lobby manager on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
