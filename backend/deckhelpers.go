package game

// Removes a card from an array at given index, returns a new array of cards
func RemoveCardAtIndex(s []Card, index int) []Card {
	ret := make([]Card, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func MoveToPlayed(user string, index int, gs *Gamestate) {
	card := gs.Players[user].Hand.Cards[index]
	newHand := RemoveCardAtIndex(gs.Players[user].Hand.Cards, index)
	updatedPlayer := gs.Players[user]

	updatedPlayer.Hand.Cards = newHand
	updatedPlayer.PlayArea.Cards = append(updatedPlayer.PlayArea.Cards, card)

	gs.Players[user] = updatedPlayer
}

func MoveToDiscard() {

}

func ShuffleDiscardToDeck() {

}

func Draw5Cards() {

}
