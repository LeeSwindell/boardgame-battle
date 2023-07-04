package main

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

// Returns array of active villains, array of villain deck
func CreateVillains() ([]Villain, []Villain) {
	villainDeck := []Villain{
		draco(),
		quirrell(),
		crabbeAndGoyle(),
		bartyCrouchJr(),
		basilisk(),
		bellatrixLestrange(),
		cornishPixies(),
		dementor(),
		fenrirGreyback(),
		doloresUmbridge(),
		fluffy(),
		luciusMalfoy(),
		tomRiddle(),
		peterPettigrew(),
		troll(),
		norbert(),
	}
	// for testing the latest villains.
	// villains := villainDeck[len(villainDeck)-3:]

	villainDeck = ShuffleVillains(villainDeck)
	villains := villainDeck[:3]
	villainDeck = villainDeck[3:]
	villainDeck = append(villainDeck, voledmortFive())

	for i := range villains {
		villains[i].Active = true
	}

	return villains, villainDeck
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
		avadaKedavra(),
		expulso(),
		handOfGlory(),
		heirOfSlytherin(),
		inquisitorialSquad(),
		menacingGrowl(),
		regeneration(),
	}

	return darkArts
}

func CreateMarketDeck() []Card {
	deck := []Card{
		crystalBall(),
		finite(),
		incendio(),
		oliverWood(),
		reparo(),
		triwizardCup(),
	}

	return deck
}

func CreateMarket() []Card {
	market := []Card{
		crystalBall(),
		finite(),
		incendio(),
		oliverWood(),
		reparo(),
		triwizardCup(),
	}

	return market
}

func RefillMarket(cardname string) Card {
	cards := make(map[string]func() Card)
	cards["Crystal Ball"] = crystalBall
	cards["Finite!"] = finite
	cards["Incendio!"] = incendio
	cards["Oliver Wood"] = oliverWood
	cards["Reparo"] = reparo
	cards["Triwizard Cup"] = triwizardCup

	return cards[cardname]()
}
