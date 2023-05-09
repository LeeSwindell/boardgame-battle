package game

import "log"

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

func ShuffleDiscardToDeck() {

}

// Removes the top card from a deck, returns the card.
func PopFromDeck(deck *[]Card) Card {
	if len(*deck) == 0 {
		// handle empty deck case
		// shuffleDiscardToDeck()
		return Card{}
	}

	topCard := (*deck)[len(*deck)-1]
	*deck = (*deck)[:len(*deck)-1]
	return topCard
}

func Draw5Cards(user string, gs *Gamestate) {
	updated := gs.Players[user]
	for i := 0; i < 5; i++ {
		log.Println("deck:", updated.Deck, "hand:", updated.Hand)
		updated.Hand = append(updated.Hand, PopFromDeck(&updated.Deck))
	}

	gs.Players[user] = updated
}
