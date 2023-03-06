package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/lobbies", getLobbiesHandler)
	r.HandleFunc("/lobby/{id}", lobbyHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
