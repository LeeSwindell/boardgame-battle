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

// Removes a card at index and moves it to the PlayArea.
func MoveToPlayed(user string, index int, gs *Gamestate) {
	card := gs.Players[user].Hand[index]
	newHand := RemoveCardAtIndex(gs.Players[user].Hand, index)
	updatedPlayer := gs.Players[user]

	updatedPlayer.Hand = newHand
	updatedPlayer.PlayArea = append(updatedPlayer.PlayArea, card)

	gs.Players[user] = updatedPlayer
}

// Moves a users cards from the PlayArea into the Discard
func MoveToDiscard(user string, gs *Gamestate) {
	updatedPlayer := gs.Players[user]
	updatedPlayer.Discard = append(updatedPlayer.Discard, updatedPlayer.PlayArea...)
	updatedPlayer.PlayArea = []Card{}
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

// Removes the top card of a players deck, returns the card.
func PopFromDeck(player *Player) Card {
	if len(player.Deck) == 0 {
		ShuffleDiscardToDeck(player)
	}

	topCard := player.Deck[len(player.Deck)-1]
	player.Deck = player.Deck[:len(player.Deck)-1]
	return topCard
}

func Draw5Cards(user string, gs *Gamestate) {
	updated := gs.Players[user]
	for i := 0; i < 5; i++ {
		log.Println("deck:", stringifyCards(updated.Deck), "hand:", stringifyCards(updated.Hand))
		updated.Hand = append(updated.Hand, PopFromDeck(&updated))
	}

	gs.Players[user] = updated
}

func MoneyDamageToZero(user string, gs *Gamestate) {
	updated := gs.Players[user]
	updated.Damage = 0
	updated.Money = 0
	gs.Players[user] = updated
}
