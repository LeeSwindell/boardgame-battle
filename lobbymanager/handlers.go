package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetLobbiesHandler(w http.ResponseWriter, r *http.Request) {
	l, err := json.Marshal(lobbies)
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

	client := &Client{hub: hub, conn: conn, pid: pid}
	client.hub.register <- client

	println("player ", pid.String(), " is connecting")

	go client.readPump()
}

func CreateLobbyHandler(w http.ResponseWriter, r *http.Request) {
	globalMu.Lock()
	defer globalMu.Unlock()
	lobbyid := lobbyNumber
	lobbyNumber++

	hostname := r.Header.Get("Authorization")
	lobby := Lobby{
		ID:   lobbyid,
		Name: hostname + "'s lobby!",
		Host: hostname,
		Players: []Player{
			{
				ID:        0,
				Name:      hostname,
				Character: "Harry",
			},
		},
	}

	lobbies = append(lobbies, lobby)

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

	newPlayerID := len(lobbies[id].Players)
	log.Println(newPlayerID, "<-- new player with id created")

	if username != "" {
		newPlayer := Player{
			ID:        newPlayerID,
			Name:      username,
			Character: "Harry",
		}
		lobbies[id].Players = append(lobbies[id].Players, newPlayer)
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
	lobby := lobbies[id]

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

	remove := func(slice []Player, s int) []Player {
		return append(slice[:s], slice[s+1:]...)
	}

	if len(lobbies[id].Players) <= 1 {
		lobbies[id].Players = []Player{}
	} else {
		for i, p := range lobbies[id].Players {
			if p.Name == user {
				log.Println(lobbies[id].Players)
				lobbies[id].Players = remove(lobbies[id].Players, i)
				hub.SendRefreshRequest()
				break
			}
		}
	}
}
