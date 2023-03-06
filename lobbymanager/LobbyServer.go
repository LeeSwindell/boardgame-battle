package main

import (
	"log"
	"net/http"
)

func getLobbiesHandler(w http.ResponseWriter, r *http.Request) {
	println("help")
}

func lobbyHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	print(conn)
}
