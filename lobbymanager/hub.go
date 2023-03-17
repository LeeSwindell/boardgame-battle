package main

import (
	"log"
)

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
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
		println("get lobbies request")
		RefreshLobby(c)
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
