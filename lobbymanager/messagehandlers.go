package main

import (
	"log"
	"math/rand"
)

func RefreshLobby(c *Client) {
	err := c.conn.WriteJSON(players)
	if err != nil {
		log.Println("error writing json in RefreshLobby, ", err)
	}
}

func AddPlayer(message Message) {
	var player Player
	if p, ok := message.Data.(map[string]interface{}); ok {
		player.Name = p["Name"].(string)
		player.Character = p["Character"].(string)
		player.ID = rand.Intn(1000)
	}

	players.Players = append(players.Players, player)
	hub.SendRefreshRequest()
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
