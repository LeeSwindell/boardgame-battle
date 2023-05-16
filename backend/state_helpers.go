package game

import (
	"log"
	"math/rand"
	"time"
)

// Removes a card from an array at given index, returns a new array of cards
func RemoveCardAtIndex(s []Card, index int) []Card {
	ret := make([]Card, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// Removes a card and moves it to the PlayArea given a cardid.
func MoveToPlayed(user string, id int, gs *Gamestate) {
	for index, c := range gs.Players[user].Hand {
		if c.Id == id {
			card := gs.Players[user].Hand[index]
			newHand := RemoveCardAtIndex(gs.Players[user].Hand, index)
			updatedPlayer := gs.Players[user]

			updatedPlayer.Hand = newHand
			updatedPlayer.PlayArea = append(updatedPlayer.PlayArea, card)

			gs.Players[user] = updatedPlayer
			return
		}
	}

	log.Println("error with MoveToPlayed")
}

// Moves a users cards from the PlayArea into the Discard
func MovePlayedToDiscard(user string, gs *Gamestate) {
	updatedPlayer := gs.Players[user]
	updatedPlayer.Discard = append(updatedPlayer.Discard, updatedPlayer.PlayArea...)
	updatedPlayer.PlayArea = []Card{}
	gs.Players[user] = updatedPlayer
}

// Moves a users cards from the PlayArea into the Discard
func MoveHandToDiscard(user string, gs *Gamestate) {
	updatedPlayer := gs.Players[user]
	updatedPlayer.Discard = append(updatedPlayer.Discard, updatedPlayer.Hand...)
	updatedPlayer.Hand = []Card{}
	gs.Players[user] = updatedPlayer
}

// updates values of a player - shuffling discard pile into their deck
func ShuffleDiscardToDeck(player *Player) {
	addToDeck := ShuffleCards(player.Discard)
	player.Deck = append(player.Deck, addToDeck...)
	player.Discard = []Card{}
}

// Shuffles a slice of cards, and returns the new ordering.
func ShuffleCards(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())

	// Create a new slice to store the shuffled cards
	shuffledCards := make([]Card, len(cards))
	copy(shuffledCards, cards)

	// Shuffle the cards
	for i := range shuffledCards {
		j := rand.Intn(i + 1)
		shuffledCards[i], shuffledCards[j] = shuffledCards[j], shuffledCards[i]
	}

	return shuffledCards
}

func ShuffleDarkArts(da []DarkArt) []DarkArt {
	rand.Seed(time.Now().UnixNano())

	// Create a new slice to store the shuffled cards
	shuffledDA := make([]DarkArt, len(da))
	copy(shuffledDA, da)

	// Shuffle the cards
	for i := range shuffledDA {
		j := rand.Intn(i + 1)
		shuffledDA[i], shuffledDA[j] = shuffledDA[j], shuffledDA[i]
	}

	return shuffledDA
}

// Removes the top card of a players deck, returns the card.
func PopFromDeck(player *Player) (Card, bool) {
	if len(player.Deck) == 0 {
		ShuffleDiscardToDeck(player)
		if len(player.Deck) == 0 {
			log.Println("deck is completely empty mate")
			return Card{}, false
		}
	}

	topCard := player.Deck[len(player.Deck)-1]
	player.Deck = player.Deck[:len(player.Deck)-1]
	return topCard, true
}

// Used to draw cards and shuffle deck if needed.
func DrawXCards(user string, gs *Gamestate, amount int) {
	updated := gs.Players[user]
	for i := 0; i < amount; i++ {
		card, ok := PopFromDeck(&updated)
		if ok {
			updated.Hand = append(updated.Hand, card)
		}
	}

	gs.Players[user] = updated
}

func MoneyDamageToZero(user string, gs *Gamestate) {
	updated := gs.Players[user]
	updated.Damage = 0
	updated.Money = 0
	gs.Players[user] = updated
}

func NextTurnInOrder(gs *Gamestate) {
	for i, name := range gs.TurnOrder {
		if name == gs.CurrentTurn {
			gs.CurrentTurn = gs.TurnOrder[(i+1)%len(gs.TurnOrder)]
			break
		}
	}
}

func StartNewTurn(gameid int, gs *Gamestate) {
	// Starting next turn actions.
	for _, v := range gs.Villains {
		if v.playBeforeDA {
			for _, e := range v.Effect {
				e.Trigger(gs)
			}
		}
	}
	gs.Locations[gs.CurrentLocation].Effect.Trigger(gs)
	for _, v := range gs.Villains {
		if !v.playBeforeDA {
			for _, e := range v.Effect {
				e.Trigger(gs)
			}
		}
	}

	SendLobbyUpdate(gameid, gs)
}
