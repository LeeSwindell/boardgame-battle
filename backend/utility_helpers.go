package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// returns the lobbyid and username associated with a request
func getIdAndUser(r *http.Request) (int, string) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("err getting id:", err.Error())
	}
	user := r.Header.Get("Authorization")
	if user == "" {
		log.Println("empty username")
	}

	return id, user
}

func SendLobbyUpdate(id int, gs *Gamestate) {
	url := fmt.Sprintf("http://localhost:8000/game/%d/refreshgamestate", id)
	if appEnv == "prod" || os.Getenv("APP_ENV") == "prod" {
		url = fmt.Sprintf("https://lobbymanager.fly.dev/game/%d/refreshgamestate", id)
	}

	data, err := json.Marshal(gs)
	if err != nil {
		log.Println("err marshaling gamestate:", err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Println("err sending lobby update:", err.Error())
	}

	client := http.Client{}
	client.Do(req)
}

func getUserInput(id int, user string, effect Effect) string {
	// pass result to card effect.
	url := fmt.Sprintf("http://localhost:8000/game/%d/getuserinput/%s", id, user)
	if appEnv == "prod" || os.Getenv("APP_ENV") == "prod" {
		url = fmt.Sprintf("https://lobbymanager.fly.dev/game/%d/getuserinput/%s", id, user)
	}
	data, err := json.Marshal(effect)
	if err != nil {
		log.Println("err marshaling options:", err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Println("err sending user input POST:", err.Error())
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("err reading response body:", err.Error())
	}

	return string(body)
}

// Sends a request to lobby manager for the cardId to discard from a users hand.
func AskUserToDiscard(gameid int, user string, hand []Card) string {
	url := fmt.Sprintf("http://localhost:8000/game/%d/askusertodiscard/%s", gameid, user)
	if appEnv == "prod" || os.Getenv("APP_ENV") == "prod" {
		url = fmt.Sprintf("https://lobbymanager.fly.dev/game/%d/askusertodiscard/%s", gameid, user)
	}
	data, err := json.Marshal(hand)
	if err != nil {
		log.Println("err marshaling options:", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Println("err sending user discard POST:", err.Error())
	}
	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("err reading response body:", err.Error())
	}

	return string(body)
}

func AskUserToSelectPlayer(gameid int, user string, players []string) string {
	endpoint := fmt.Sprintf("http://localhost:8000/game/%d/askusertoselectplayer/%s", gameid, user)
	if appEnv == "prod" || os.Getenv("APP_ENV") == "prod" {
		endpoint = fmt.Sprintf("https://lobbymanager.fly.dev/game/%d/askusertoselectplayer/%s", gameid, user)
	}

	// Encode the players slice as JSON
	playerJSON, err := json.Marshal(players)
	if err != nil {
		log.Println("err encoding players slice:", err)
		return ""
	}

	// Create a new http.Request object with the POST method and the encoded values.
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(string(playerJSON)))
	if err != nil {
		log.Println("err creating http.request in selectplayer", err)
	}

	// Set the Content-Type header to application/x-www-form-urlencoded.
	req.Header.Set("Content-Type", "application/json")

	// Submit the request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("err sending request in selectplayer", err)
	}
	defer resp.Body.Close()

	// Check the response status code.
	if resp.StatusCode != 200 {
		log.Println("status bad, select player", resp.StatusCode)
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err reading response body:", err.Error())
	}

	return string(body)
}

func stringifyCards(cards []Card) string {
	res := ""
	for _, c := range cards {
		res += c.Name + " "
	}

	return res
}

func getGsForGameID(r *http.Request) (*Gamestate, bool) {
	id, _ := getIdAndUser(r)
	globalMu.Lock()
	defer globalMu.Unlock()
	gs, ok := states[id]
	if !ok {
		// log.Println("error getting gs from id for handler")
		keys := []int{}
		for k := range states {
			keys = append(keys, k)
		}
		log.Println("error getting gs from id ** id: ", id, "** states keys: ", keys)
		return nil, false
	}

	return gs, true
}
