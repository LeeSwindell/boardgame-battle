package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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
	url := fmt.Sprintf("http://localhost:8000/game/%d/getuserinput", id)
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

func stringifyCards(cards []Card) string {
	res := ""
	for _, c := range cards {
		res += c.Name + " "
	}

	return res
}
