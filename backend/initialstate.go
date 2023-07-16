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
	villains := villainDeck[:2]
	// testing pettigrew
	villains = append(villains, peterPettigrew())
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
		obliviate(),
		crucio(),
		fiendfyre(),
		morsmordre(),
		blastended(),
		educationalDecree(),
		imperio(),
		legilimency(),
		oppugno(),
		petrification(),
		poison(),
		ragingTroll(),
		relashio(),
		sectumsempra(),
		slugulusEructo(),
		tarantallegra(),
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
		albusDumbledore(),
		arthurWeasley(),
		bezoar(),
		choChang(),
		deluminator(),
		expectoPatronum(),
		felixFelicis(),
		filiusFlitwick(),
		hogwartsAHistory(),
		mollyWeasley(),
		quidditchGear(),
		siriusBlack(),
		stupefy(),
		sybillTrelawney(),
		butterbeer(),
		dobbyTheHouseElf(),
		essenceOfDittany(),
		fang(),
		fleurDelacour(),
		goldenSnitch(),
		pensieve(),
		rubeusHagrid(),
		advancedPotionMaking(),
		alastorMadEyeMoody(),
		argusFilchAndMrsNorris(),
		cedricDiggory(),
		descendo(),
		expelliarmus(),
		ginnyWeasley(),
		horaceSlughorn(),
		kingsleyShacklebolt(),
		lumos(),
		lunaLovegood(),
		nimbusTwoThousandAndOne(),
		nymphadoraTonks(),
		polyjuicePotion(),
		pomonaSprout(),
		severusSnape(),
		swordOfGryffindor(),
		tergeo(),
		viktorKrum(),
		fawkesThePhoenix(),
		minervaMcgonagall(),
		remusLupin(),
		elderWand(),
		chocolateFrog(),
		gilderoyLockhart(),
		maraudersMap(),
		protego(),
		accio(),
		fredWeasley(),
		georgeWeasley(),
		owls(),
		sortingHat(),
		wingardiumLeviosa(),
		petrificusTotalus(),
		harp(),
		finiteIncantatem(),
		confundus(),
	}

	deck = ShuffleCards(deck)

	return deck
}

func CreateMarket() []Card {
	// deck := CreateMarketDeck()
	// market := deck[0:6]
	// return market

	return createTestMarket()
}

func RefillMarket(index int, gs *Gamestate) {
	// refill deck if empty
	if len(gs.MarketDeck) == 0 {
		gs.MarketDeck = CreateMarketDeck()
	}

	gs.Market[index] = gs.MarketDeck[0]
	gs.MarketDeck = gs.MarketDeck[1:]
}

func createTestMarket() []Card {
	m := []Card{
		hogwartsAHistory(),
		hogwartsAHistory(),
		hogwartsAHistory(),
		hogwartsAHistory(),
		hogwartsAHistory(),
		hogwartsAHistory(),
	}

	return m
}
