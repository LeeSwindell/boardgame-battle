package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type Users struct {
	mu    sync.Mutex
	Users []string
}

var users Users

// checks if users sessionid is a valid logged in user
func sessionidHandler(w http.ResponseWriter, r *http.Request) {
	sessionid := r.Header.Get("Authorization")
	if sessionid == "" {
		log.Println("no session id")
	}

	for _, u := range users.Users {
		if sessionid == u {
			io.WriteString(w, "true")
			return
		}
	}

	io.WriteString(w, "false")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get the username from the JSON data
	username, ok := data["username"].(string)
	if !ok {
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	if verifyCredentials(username) {
		users.mu.Lock()
		users.Users = append(users.Users, username)
		log.Println(username)
		log.Println(users.Users)
		users.mu.Unlock()

		expiration := time.Now().Add(7 * 24 * time.Hour)
		http.SetCookie(w, &http.Cookie{
			Name:    "session-name",
			Value:   username,
			Expires: expiration,
		})
	}
}

func verifyCredentials(username string) bool {
	// Check username is unique
	// if it exists already, take over that session?
	return true
}
