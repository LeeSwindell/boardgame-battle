package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	// game "github.com/LeeSwindell/boardgame-battle/backend"
	"github.com/gorilla/mux"
)

func GetLobbiesHandler(w http.ResponseWriter, r *http.Request) {
	// Convert the map to a slice
	var lobbiesSlice []Lobby
	for _, lobby := range lobbies {
		lobbiesSlice = append(lobbiesSlice, lobby)
	}

	l, err := json.Marshal(lobbiesSlice)
	if err != nil {
		log.Println("GetLobbies error: ")
		log.Println(err)
	}

	w.Write(l)
}

func AddClientHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	pid := getUniquePlayerId()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("AddClient error: ")
		log.Println(err)
		return
	}

	client := &Client{hub: hub, conn: conn, pid: pid, username: ""}
	client.hub.register <- client

	// println("player ", pid.String(), " is connecting")

	go client.readPump()
}

func AddClientWithUsernameHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	pid := getUniquePlayerId()
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		log.Println("error adding client with username, no username")
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("AddClient error: ")
		log.Println(err)
		return
	}

	client := &Client{hub: hub, conn: conn, pid: pid, username: username}
	client.hub.register <- client

	// println("player ", pid.String(), " is connecting")

	go client.readPump()
}

func CreateLobbyHandler(w http.ResponseWriter, r *http.Request) {
	globalMu.Lock()
	defer globalMu.Unlock()
	lobbyid := getUniqueLobbyId()
	// lobbyNumber++
	playerid := int(getUniquePlayerId().ID())

	hostname := r.Header.Get("Authorization")
	lobby := Lobby{
		ID:   lobbyid,
		Name: hostname + "'s lobby!",
		Host: hostname,
		Players: []LobbyPlayer{
			{
				ID:        playerid,
				Name:      hostname,
				Character: "Harry",
			},
		},
	}

	lobbies[lobbyid] = lobby

	_, err := io.WriteString(w, fmt.Sprint(lobbyid))
	if err != nil {
		log.Println("err creating lobby:", err)
	}
}

func JoinLobbyHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("Authorization")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("err joining lobby:", err.Error())
	}
	globalMu.Lock()
	defer globalMu.Unlock()

	newPlayerID := int(getUniquePlayerId().ID())
	// log.Println(newPlayerID, "<-- new player with id created")

	if username != "" {
		newPlayer := LobbyPlayer{
			ID:        newPlayerID,
			Name:      username,
			Character: "Harry",
		}
		lobby := lobbies[id]
		lobby.Players = append(lobby.Players, newPlayer)
		lobbies[id] = lobby
	} else {
		log.Println("err joining lobby, username empty")
	}

	// Trigger refresh for all clients in lobby.
	hub.SendRefreshRequest()
}

func RefreshLobbyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("error with lobbyid:", err.Error())
	}

	globalMu.Lock()
	defer globalMu.Unlock()
	lobby, ok := lobbies[id]
	if !ok {
		http.Error(w, "Lobby not found", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(lobby)
	if err != nil {
		log.Println(err)
	}

	w.Write(res)
}

func SetCharHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("err getting id in setchar:", err.Error())
	}
	user := r.Header.Get("Authorization")
	if user == "" {
		log.Println("empty username in setchar handler")
		return
	}
	var data map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newChar, ok := data["character"].(string)
	if !ok {
		log.Println("error selecting char")
	}

	globalMu.Lock()
	defer globalMu.Unlock()

	for i, p := range lobbies[id].Players {
		if p.Name == user {
			lobbies[id].Players[i].Character = newChar
			break
		}
	}

	hub.SendRefreshRequest()
}

func LeaveLobbyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("err getting id in leavelobbyhandler:", err.Error())
	}
	user := r.Header.Get("Authorization")
	if user == "" {
		log.Println("empty username in leavelobby handler")
		return
	}

	globalMu.Lock()
	defer globalMu.Unlock()

	remove := func(slice []LobbyPlayer, s int) []LobbyPlayer {
		return append(slice[:s], slice[s+1:]...)
	}

	if len(lobbies[id].Players) <= 1 {
		// just delete the empty lobby.
		delete(lobbies, id)
		// lobbies = append(lobbies[:id], lobbies[id+1:]...)
	} else {
		for i, p := range lobbies[id].Players {
			if p.Name == user {
				log.Println(lobbies[id].Players)
				lobby := lobbies[id]
				lobby.Players = remove(lobby.Players, i)
				lobbies[id] = lobby
				// lobbies[id].Players = remove(lobbies[id].Players, i)
				hub.SendRefreshRequest()
				break
			}
		}
	}
}

func StartGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("err getting id in startlobbyhandler:", err.Error())
	}

	startingPlayers := map[string]Player{}
	turnOrder := []string{}
	for _, p := range lobbies[id].Players {
		startingPlayers[p.Name] = Player{
			Name:      p.Name,
			Character: p.Character,
			Health:    10,
			Money:     0,
			Damage:    0,
			Deck:      []Card{},
			Hand:      []Card{},
			PlayArea:  []Card{},
			Discard:   []Card{},
		}

		turnOrder = append(turnOrder, p.Name)
	}

	// turn this function call into a post request
	// game.StartGame(startingPlayers, turnOrder, id)

	// Convert startingPlayers and turnOrder to JSON
	data := struct {
		StartingPlayers map[string]Player `json:"startingPlayers"`
		TurnOrder       []string          `json:"turnOrder"`
		ID              int               `json:"id"`
	}{
		StartingPlayers: startingPlayers,
		TurnOrder:       turnOrder,
		ID:              id,
	}

	// Convert data to JSON bytes
	payload, err := json.Marshal(data)
	if err != nil {
		log.Println("error marshaling data:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send POST request to the desired endpoint
	baseURL := "http://localhost:8080" // Replace with your actual base URL
	url := fmt.Sprintf("%s/startgame", baseURL)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Println("error sending POST request:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("received non-OK status code:", resp.StatusCode)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	hub.SendStartGame()
	// delete lobby info.
}

func RefreshGamestateHandler(w http.ResponseWriter, r *http.Request) {
	// Not needed until lobbies are grouped by id - currently only one game will work.
	// vars := mux.Vars(r)
	// id, err := strconv.Atoi(vars["id"])
	// if err != nil {
	// 	log.Println("err getting id in refreshgamestatehandler:", err.Error())
	// }

	var gs Gamestate
	json.NewDecoder(r.Body).Decode(&gs)

	hub.SendGameState(&gs)
}

var userInputChan = make(chan string)

func GetUserInputHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, ok := vars["user"]
	if !ok {
		log.Println("err getting username in getuserinputhandler")
	}

	var chooseOne ChooseOne
	json.NewDecoder(r.Body).Decode(&chooseOne)

	hub.askPlayerChoice(user, chooseOne.Options)

	// If a user submits multiple inputs somehow, this will block and be offset
	// Change to check for a submit id with each choice?
	choice := <-userInputChan
	w.Write([]byte(choice))
}

func AskUserToDiscardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, ok := vars["user"]
	if !ok {
		log.Println("err getting username in AskUserToDiscardHandler")
	}

	var hand []Card
	json.NewDecoder(r.Body).Decode(&hand)
	choices := []string{}
	for _, c := range hand {
		choices = append(choices, c.Name)
	}

	hub.askPlayerChoice(user, choices)

	// If a user submits multiple inputs somehow, this will block and be offset
	// Change to check for a submit id with each choice?
	choice := <-userInputChan
	w.Write([]byte(choice))
}

func AskUserToSelectPlayerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, ok := vars["user"]
	if !ok {
		log.Println("err getting username in AskUserToSelectPlayer")
	}

	var players []string
	json.NewDecoder(r.Body).Decode(&players)

	hub.askPlayerChoice(user, players)

	// If a user submits multiple inputs somehow, this will block and be offset
	// Change to check for a submit id with each choice?
	choice := <-userInputChan
	w.Write([]byte(choice))
}

func SubmitUserChoiceHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Choice string `json:"choice"`
	}
	json.NewDecoder(r.Body).Decode(&data)

	userInputChan <- data.Choice
}
