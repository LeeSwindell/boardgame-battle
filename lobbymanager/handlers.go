package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func RefreshLobby(c *Client) {
	err := c.conn.WriteJSON(players)
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
		log.Print("sending refresh request")
		c.conn.WriteJSON(message)
	}
}

func GetLobbiesHandler(w http.ResponseWriter, r *http.Request) {
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
	conn.WriteJSON(players)

	go client.readPump()
}

func CreateLobbyHandler(w http.ResponseWriter, r *http.Request) {
	globalMu.Lock()
	lobbyid := lobbyNumber
	lobbyNumber++

	hostname := r.Header.Get("Authorization")
	lobby := Lobby{
		ID:   lobbyid,
		Name: hostname + "'s lobby!",
		Host: hostname,
	}

	lobbies.Lobbies = append(lobbies.Lobbies, lobby)
	globalMu.Unlock()

	_, err := io.WriteString(w, fmt.Sprint(lobbyid))
	if err != nil {
		log.Println("err creating lobby:", err)
	}
}

func JoinLobbyHandler(w http.ResponseWriter, r *http.Request) {
	// add client and player id to lobby group
	// redirect to lobby.
}

func RefreshLobbyHandler(w http.ResponseWriter, r *http.Request) {
	res, err := json.Marshal(players)
	if err != nil {
		log.Println(err)
	}

	w.Write(res)
}

func AddPlayerHandler(w http.ResponseWriter, r *http.Request) {
	var player Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	player.ID = rand.Intn(102030)

	players.Players = append(players.Players, player)
	hub.SendRefreshRequest()

	w.WriteHeader(http.StatusCreated)
}
