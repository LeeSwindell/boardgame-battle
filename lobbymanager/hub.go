package main

import (
	"log"
)

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) handleMessage(c *Client, msgType int, message Message) {
	switch message.Type {
	case "RefreshLobby":
		println("get lobbies request")
		RefreshLobby(c)
	case "AddPlayer":
		println("adding player")
		AddPlayer(message)
	default:
		log.Printf("unknown message type %s", message.Type)
	}
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
