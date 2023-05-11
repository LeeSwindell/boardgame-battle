package game

func RonStartingDeck() []Card {
	deck := []Card{}

	deck = append(deck, pigwidgeon())
	deck = append(deck, bertieBottsEveryFlavourBeans())
	deck = append(deck, cleansweep())
	for i := 0; i < 7; i++ {
		deck = append(deck, alohamora())
	}

	return deck
}

func CreateVillains() []Villain {
	villains := []Villain{
		draco(),
		quirrell(),
		crabbeAndGoyle(),
	}

	return villains
}

func CreateLocations() []Location {
	locations := []Location{
		greatHall(),
	}

	return locations
}
