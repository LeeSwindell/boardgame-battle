package game

func RonStartingDeck() []Card {
	deck := []Card{}

	deck = append(deck, pigwidgeon())
	deck = append(deck, bertieBottsEveryFlavourBeans())
	deck = append(deck, cleansweep())
	for i := 0; i < 7; i++ {
		deck = append(deck, alohamora())
	}

	deck = ShuffleCards(deck)

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
		castleGates(),
		hagridsHut(),
		greatHall(),
	}

	return locations
}

func CreateDarkArtDeck() []DarkArt {
	darkArts := []DarkArt{
		dementorsKiss(),
		heWhoMustNotBeNamed(),
		flipendo(),
	}

	return darkArts
}

func CreateMarketDeck() []Card {
	deck := []Card{
		crystalBall(),
		finite(),
		incendio(),
	}

	return deck
}

// IDs cause render errors with react currently since market cards aren't unique.
func CreateMarket() []Card {
	market := []Card{
		crystalBall(),
		finite(),
		incendio(),
	}

	return market
}
