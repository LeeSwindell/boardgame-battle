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

var config = NewConfiguration()

type Configuration struct {
	LobbyManagerURL string
}

func NewConfiguration() *Configuration {
	mode := appEnv
	switch mode {
	case "dev":
		return &Configuration{
			LobbyManagerURL: "http://localhost:8000/lm",
		}
	case "prod":
		println("setting env to prod")
		return &Configuration{
			LobbyManagerURL: "https://www.gamewithyourfriends.dev/lm",
		}
	default:
		println("************* NO APP ENV PROVIDED ***********")
		// use dev env by default.
		return &Configuration{
			LobbyManagerURL: "http://localhost:8000/lm",
		}
	}
}

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
	url := fmt.Sprintf("%s/game/%d/refreshgamestate", config.LobbyManagerURL, id)

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

func getUserInput(id int, user string, options []string, prompt string) string {
	// pass result to card effect.
	url := fmt.Sprintf("%s/game/%d/getuserinput/%s", config.LobbyManagerURL, id, user)

	var dataToSend = struct {
		Options []string `json:"options"`
		Prompt  string   `json:"prompt"`
	}{Options: options, Prompt: prompt}

	data, err := json.Marshal(dataToSend)
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

// FIX ADD PROMPT
func AskUserToSelectPlayer(gameid int, user string, players []string) string {
	url := fmt.Sprintf("%s/game/%d/askusertoselectplayer/%s", config.LobbyManagerURL, gameid, user)

	// Encode the players slice as JSON
	playerJSON, err := json.Marshal(players)
	if err != nil {
		log.Println("err encoding players slice:", err)
		return ""
	}

	// Create a new http.Request object with the POST method and the encoded values.
	req, err := http.NewRequest("POST", url, strings.NewReader(string(playerJSON)))
	if err != nil {
		log.Println("err creating http.request in selectplayer", err)
	}

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

func AskUserToSelectCard(user string, gameid int, choices []Card, prompt string) int {
	url := fmt.Sprintf("%s/game/%d/askusertoselectcard/%s", config.LobbyManagerURL, gameid, user)

	var dataToSend = struct {
		Cards  []Card `json:"cards"`
		Prompt string `json:"prompt"`
	}{Cards: choices, Prompt: prompt}

	data, err := json.Marshal(dataToSend)
	if err != nil {
		log.Println("err marshaling options:", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Println("err sending AskUserToSelectCard POST:", err.Error())
	}
	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	// Should receive a CardId of with a selection in response.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("err reading response body:", err.Error())
	}

	Logger(string(body))

	selectionId, err := strconv.Atoi(string(body))
	if err != nil {
		log.Println("err, didn't receive cardId properly", err.Error())
	}

	return selectionId
}

// Ask the user for input, and display a card. returns the string representing a choice.
func AskUserInputWithCard(gameid int, user string, path string, prompt string, choices []string) string {
	url := fmt.Sprintf("%s/game/%d/userchoicewithcard/%s", config.LobbyManagerURL, gameid, user)

	var dataToSend = struct {
		Path    string   `json:"path"`
		Choices []string `json:"choices"`
		Prompt  string   `json:"prompt"`
	}{Path: path, Prompt: prompt, Choices: choices}

	data, err := json.Marshal(dataToSend)
	if err != nil {
		log.Println("err marshaling options:", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Println("err sending AskUserInputWithCard POST:", err.Error())
	}
	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	// Should receive a string in response.
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("err reading response body:", err.Error())
	}

	Logger(string(body))

	selection := string(body)

	return selection
}

func getGsForGameID(r *http.Request) (*Gamestate, bool) {
	id, _ := getIdAndUser(r)
	globalMu.Lock()
	defer globalMu.Unlock()
	gs, ok := states[id]
	if !ok {
		keys := []int{}
		for k := range states {
			keys = append(keys, k)
		}
		log.Println("error getting gs from id ** id: ", id, "** states keys: ", keys)
		return nil, false
	}

	return gs, true
}

func Logger(s string) {
	if os.Getenv("LOG_LEVEL") == "debug" {
		log.Println(s)
	}
}
