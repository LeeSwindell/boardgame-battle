package main

import (
	"io"
	"log"
	"net/http"
	"sync"
)

type Users struct {
	mu    sync.Mutex
	Users []string
}

var users Users

// returns the session id if it exists
func sessionidHandler(w http.ResponseWriter, r *http.Request) {
	sessionid, err := r.Cookie("sessionid")
	if err != nil {
		log.Println(err)
	}
	res := sessionid.String()
	io.WriteString(w, res)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	if verifyCredentials(username) {
		users.mu.Lock()
		users.Users = append(users.Users, username)
		users.mu.Unlock()

		http.SetCookie(w, &http.Cookie{
			Name:  "session-name",
			Value: username,
		})
	}

}

func verifyCredentials(username string) bool {
	// Check username is unique
	// if it exists already, take over that session?
	return true
}
