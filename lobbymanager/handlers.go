package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RefreshLobby(c *Client) {
	globalMu.Lock()
	defer globalMu.Unlock()
	err := c.conn.WriteJSON(lobbies)
	if err != nil {
		log.Println("error writing json in RefreshLobby, ", err)
	}
}

func (h *Hub) SendRefreshRequest() {
	message := Message{
		Type: "RefreshRequest",
		Data: "",
	}

	for c := range h.clients {
		c.conn.WriteJSON(message)
	}
}

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

	lobbies.Lobbies = append(lobbies.Lobbies, lobby)

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

	newPlayerID := len(lobbies.Lobbies[id].Players)
	log.Println(newPlayerID, "<-- new player with id created")

	if username != "" {
		newPlayer := Player{
			ID:        newPlayerID,
			Name:      username,
			Character: "Harry",
		}
		lobbies.Lobbies[id].Players = append(lobbies.Lobbies[id].Players, newPlayer)
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
	lobby := lobbies.Lobbies[id]

	res, err := json.Marshal(lobby)
	if err != nil {
		log.Println(err)
	}

	log.Println("#############################")
	log.Println(lobbies.Lobbies[id].Players)

	w.Write(res)
}

// FIX, set the char that actually changed, not the one who sent the request.
// ??
func SetCharHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lobbyid, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("err getting lobbyid in setchar:", err.Error())
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
	// find the player
	for i, p := range lobbies.Lobbies[lobbyid].Players {
		if p.Name == user {
			log.Println("!@#$%$^%&* user", user, "setting char of", p.Name, "to ", newChar)
			lobbies.Lobbies[lobbyid].Players[i].Character = newChar
			break
		}
	}

	hub.SendRefreshRequest()

}

func AddPlayerHandler(w http.ResponseWriter, r *http.Request) {
	var player Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	player.ID = rand.Intn(102030)

	// players.Players = append(players.Players, player)
	hub.SendRefreshRequest()

	w.WriteHeader(http.StatusCreated)
}
