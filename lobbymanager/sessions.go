package main

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("your-secret-key-here"))

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if verifyCredentials(username, password) {
		session, err := store.New(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["username"] = username
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// TODO: handle successful login
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
}

func verifyCredentials(username, password string) bool {
	// TODO: verify username and password against a database or other authentication system
	return true
}

// SET SESSION ID IN RESPONSE
// session, err := store.Get(r, "session-name")
// if err != nil {
//   http.Error(w, err.Error(), http.StatusInternalServerError)
//   return
// }
// if session.Values["username"] == nil {
//   http.Error(w, "Unauthorized", http.StatusUnauthorized)
//   return
// }
// // Set session ID as cookie in response
// http.SetCookie(w, &http.Cookie{
//   Name:    "session-name",
//   Value:   session.ID,
//   Expires: time.Now().Add(time.Hour * 24),
