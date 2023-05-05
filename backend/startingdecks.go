package game

func RonStartingDeck() Deck {
	deck := Deck{}

	deck.Cards = append(deck.Cards, pigwidgeon())
	for i := 0; i < 7; i++ {
		deck.Cards = append(deck.Cards, alohamora())
	}

	return deck
}