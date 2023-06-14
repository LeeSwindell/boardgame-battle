package main

import (
	"log"

	"github.com/google/uuid"
	// game "github.com/LeeSwindell/boardgame-battle/backend"
)

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
}

type Message struct {
	Type        string      `json:"type"`
	Data        interface{} `json:"data"`
	Description string      `json:"description"`
	ID          int         `json:"id"`
}

func newHub() *Hub {
	return &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) handleMessage(c *Client, msgType int, message Message) {
	switch message.Type {
	case "RefreshLobby":
		// println("get lobbies request")
		// RefreshLobby(c)
	default:
		log.Printf("unknown message type %s", message.Type)
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

func (h *Hub) SendStartGame() {
	message := Message{
		Type: "StartGame",
		Data: "",
	}

	for c := range h.clients {
		c.conn.WriteJSON(message)
	}
}

func (h *Hub) SendGameState(gs *Gamestate) {
	message := Message{
		Type: "Gamestate",
		Data: gs,
	}

	for c := range h.clients {
		c.conn.WriteJSON(message)
	}
}

func (h *Hub) askPlayerChoice(user string, choices []string, description string) int {
	// Include a display message field. rather than saying "choose one" for everything.

	id := int(uuid.New().ID())
	message := Message{
		Type:        "UserInput",
		Data:        choices,
		Description: description,
		ID:          id,
	}

	for c := range h.clients {
		if c.username == user {
			c.conn.WriteJSON(message)
			log.Println("asking for input: ", message.ID)
		}
	}

	return id
}

func (h *Hub) run() {
	for {
		select {
		case conn := <-h.register:
			h.clients[conn] = true
		case conn := <-h.unregister:
			delete(h.clients, conn)
		}
	}
}
