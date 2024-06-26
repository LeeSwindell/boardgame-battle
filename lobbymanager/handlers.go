package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	log.Println("received connect socket request")
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
	getId, ok := vars["id"]
	if !ok {
		log.Println("error adding client with username, no id")
	}
	id, err := strconv.Atoi(getId)
	if err != nil {
		log.Println("addclientwithusername err, ", err)
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("AddClient error: ")
		log.Println(err)
		return
	}

	log.Println(id)

	client := &Client{hub: hub, conn: conn, pid: pid, username: username, lobbygroup: id}
	client.hub.register <- client

	// println("player ", pid.String(), " is connecting")

	go client.readPump()
}

func CreateLobbyHandler(w http.ResponseWriter, r *http.Request) {
	globalMu.Lock()
	defer globalMu.Unlock()
	lobbyid := getUniqueLobbyId()
	playerid := int(getUniquePlayerId().ID())

	hostname := r.Header.Get("Authorization")
	lobby := Lobby{
		ID:   lobbyid,
		Name: hostname + "'s lobby!",
		Host: hostname,
		Players: []LobbyPlayer{
			{
				ID:          playerid,
				Name:        hostname,
				Character:   "Harry",
				Proficiency: "Arithmancy",
			},
		},
	}

	log.Println("creating lobby:", lobby.Players[0].Character, lobby.Players[0].Proficiency)
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

	if username != "" {
		newPlayer := LobbyPlayer{
			ID:          newPlayerID,
			Name:        username,
			Character:   "Harry",
			Proficiency: "Arithmancy",
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

func SetProfHandler(w http.ResponseWriter, r *http.Request) {
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
	log.Println(data)
	newProf, ok := data["proficiency"].(string)
	if !ok {
		log.Println("error selecting proficiency")
		return
	}

	globalMu.Lock()
	defer globalMu.Unlock()

	for i, p := range lobbies[id].Players {
		if p.Name == user {
			lobbies[id].Players[i].Proficiency = newProf
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
			Name:        p.Name,
			Character:   p.Character,
			Health:      10,
			Money:       0,
			Damage:      0,
			Deck:        []Card{},
			Hand:        []Card{},
			PlayArea:    []Card{},
			Discard:     []Card{},
			Proficiency: p.Proficiency,
		}
		log.Println("checking prof!!!!1", p.Proficiency)
		log.Println("checking character: ", p.Character)
		turnOrder = append(turnOrder, p.Name)
	}

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
	baseURL := "http://localhost:8080"
	if appEnv == "prod" || os.Getenv("APP_ENV") == "prod" {
		baseURL = "https://hogwartsbackend.fly.dev"
	}
	// API UPDATE
	baseURL = config.BackendURL
	url := fmt.Sprintf("%s/startgame", baseURL)

	log.Println("sending start game post to:", url)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Println("error sending POST request:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	log.Println("wating for response to start game?")

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
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("err getting id in refreshgamestatehandler:", err.Error())
	}

	var gs Gamestate
	json.NewDecoder(r.Body).Decode(&gs)

	hub.SendGameState(&gs, id)
}

func GetUserInputHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, ok := vars["user"]
	if !ok {
		log.Println("err getting username in getuserinputhandler")
	}

	// var chooseOne ChooseOne
	var data struct {
		Options []string `json:"options"`
		Prompt  string   `json:"prompt"`
	}

	json.NewDecoder(r.Body).Decode(&data)

	listenID := 0
	if data.Prompt != "" {
		listenID = hub.askPlayerChoice(user, data.Options, data.Prompt)
	} else {
		listenID = hub.askPlayerChoice(user, data.Options, "Choose One")
	}

	listenChan := messageBroadcaster.RegisterListener(listenID)

	choice := <-listenChan
	w.Write([]byte(choice))
}

func AskUserToDiscardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, ok := vars["user"]
	if !ok {
		log.Println("err getting username in AskUserToDiscardHandler")
	}

	var data struct {
		Hand   []Card `json:"hand"`
		Prompt string `json:"prompt"`
	}
	json.NewDecoder(r.Body).Decode(&data)

	choices := []string{}
	for _, c := range data.Hand {
		choices = append(choices, c.Name)
	}
	if data.Prompt == "" {
		data.Prompt = "Discard a card"
	}

	listenID := hub.askPlayerChoice(user, choices, data.Prompt)
	listenChan := messageBroadcaster.RegisterListener(listenID)

	choice := <-listenChan
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

	listenID := hub.askPlayerChoice(user, players, "Select a player")
	listenChan := messageBroadcaster.RegisterListener(listenID)

	choice := <-listenChan
	w.Write([]byte(choice))
}

func AskUserToSelectCardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, ok := vars["user"]
	if !ok {
		log.Println("err getting username in AskUserToSelectCard")
	}

	var data struct {
		Cards  []Card `json:"cards"`
		Prompt string `json:"prompt"`
	}
	json.NewDecoder(r.Body).Decode(&data)

	listenID := hub.askPlayerToSelectCard(user, data.Cards, data.Prompt)
	listenChan := messageBroadcaster.RegisterListener(listenID)

	choice := <-listenChan
	log.Println("lb choice select card", choice)
	w.Write([]byte(choice))
}

func AskUserInputWithCardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, ok := vars["user"]
	if !ok {
		log.Println("err getting username in AskUserToSelectCard")
	}

	var data struct {
		Path    string   `json:"path"`
		Choices []string `json:"choices"`
		Prompt  string   `json:"prompt"`
	}
	json.NewDecoder(r.Body).Decode(&data)

	listenID := hub.askPlayerChoiceWithCard(user, data.Path, data.Choices, data.Prompt)
	listenChan := messageBroadcaster.RegisterListener(listenID)

	choice := <-listenChan
	log.Println("lb choice select card", choice)
	w.Write([]byte(choice))
}

func SubmitUserChoiceHandler(w http.ResponseWriter, r *http.Request) {
	var choice ChoiceMessage
	json.NewDecoder(r.Body).Decode(&choice)

	messageBroadcaster.InputChan <- choice
}
