package game

func RonStartingDeck() []Card {
	deck := []Card{}

	deck = append(deck, pigwidgeon())
	for i := 0; i < 7; i++ {
		deck = append(deck, alohamora())
	}

	return deck
}
