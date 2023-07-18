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

func HarryStartingDeck() []Card {
	deck := []Card{}

	deck = append(deck, invisibilityCloak())
	deck = append(deck, hedwig())
	deck = append(deck, firebolt())
	for i := 0; i < 7; i++ {
		deck = append(deck, alohamora())
	}

	deck = ShuffleCards(deck)

	return deck
}

func HermioneStartingDeck() []Card {
	deck := []Card{}

	deck = append(deck, crookshanks())
	deck = append(deck, talesOfBeedleTheBard())
	deck = append(deck, timeTurner())
	for i := 0; i < 7; i++ {
		deck = append(deck, alohamora())
	}

	deck = ShuffleCards(deck)

	return deck
}

func NevilleStartingDeck() []Card {
	deck := []Card{}

	deck = append(deck, trevor())
	deck = append(deck, remembrall())
	deck = append(deck, mandrake())
	for i := 0; i < 7; i++ {
		deck = append(deck, alohamora())
	}

	deck = ShuffleCards(deck)

	return deck
}

func LunaStartingDeck() []Card {
	deck := []Card{}

	deck = append(deck, lionHat())
	deck = append(deck, spectrespecs())
	deck = append(deck, crumpleHornedSnorkack())
	for i := 0; i < 7; i++ {
		deck = append(deck, alohamora())
	}

	deck = ShuffleCards(deck)

	return deck
}

// Returns array of active villains, array of villain deck
// func CreateVillains() ([]Villain, []Villain) {
// 	villainDeck := []Villain{
// 		draco(),
// 		quirrell(),
// 		crabbeAndGoyle(),
// 		bartyCrouchJr(),
// 		basilisk(),
// 		bellatrixLestrange(),
// 		cornishPixies(),
// 		dementor(),
// 		fenrirGreyback(),
// 		doloresUmbridge(),
// 		fluffy(),
// 		luciusMalfoy(),
// 		tomRiddle(),
// 		peterPettigrew(),
// 		troll(),
// 		norbert(),
// 	}

// 	villainDeck = ShuffleVillains(villainDeck)
// 	villains := villainDeck[:3]
// 	villainDeck = villainDeck[3:]
// 	villainDeck = append(villainDeck, voledmortFive())

// 	for i := range villains {
// 		villains[i].Active = true
// 	}

// 	return villains, villainDeck
// }

func CreateBox1Villains() ([]Villain, []Villain) {
	villainDeck := []Villain{
		basilisk(),
		cornishPixies(),
		dementor(),
		fluffy(),
		troll(),
		norbert(),
	}

	possVillains := []Villain{
		draco(),
		quirrell(),
		crabbeAndGoyle(),
		bartyCrouchJr(),
		bellatrixLestrange(),
		fenrirGreyback(),
		doloresUmbridge(),
		luciusMalfoy(),
		tomRiddle(),
		peterPettigrew(),
	}
	possVillains = ShuffleVillains(possVillains)[0:5]
	villainDeck = append(villainDeck, possVillains...)

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

// func CreateDarkArtDeck() []DarkArt {
// 	darkArts := []DarkArt{
// 		dementorsKiss(),
// 		heWhoMustNotBeNamed(),
// 		flipendo(),
// 		avadaKedavra(),
// 		expulso(),
// 		handOfGlory(),
// 		heirOfSlytherin(),
// 		inquisitorialSquad(),
// 		menacingGrowl(),
// 		regeneration(),
// 		obliviate(),
// 		crucio(),
// 		fiendfyre(),
// 		morsmordre(),
// 		blastended(),
// 		educationalDecree(),
// 		imperio(),
// 		legilimency(),
// 		oppugno(),
// 		petrification(),
// 		poison(),
// 		ragingTroll(),
// 		relashio(),
// 		sectumsempra(),
// 		slugulusEructo(),
// 		tarantallegra(),
// 	}

// 	return darkArts
// }

func createNumAccurateDarkArts() []DarkArt {
	deck := []DarkArt{
		// Year 1
		expulso(), expulso(), expulso(),
		petrification(), petrification(),
		flipendo(), flipendo(),
		heWhoMustNotBeNamed(), heWhoMustNotBeNamed(), heWhoMustNotBeNamed(),
		// Year 2
		handOfGlory(), handOfGlory(),
		relashio(),
		obliviate(),
		poison(),
		// Year 3
		oppugno(),
		tarantallegra(),
		dementorsKiss(), dementorsKiss(),
		// Year 4
		crucio(),
		regeneration(),
		avadaKedavra(),
		imperio(),
		morsmordre(), morsmordre(),
		heirOfSlytherin(), heirOfSlytherin(),
		// Year 5
		imperio(),
		educationalDecree(), educationalDecree(),
		legilimency(),
		crucio(),
		morsmordre(),
		avadaKedavra(),
		// Year 6
		sectumsempra(), sectumsempra(),
		morsmordre(),
		// Year 7
		avadaKedavra(),
		imperio(),
		fiendfyre(),
		crucio(),
		// Box 1
		menacingGrowl(), menacingGrowl(),
		blastended(),
		slugulusEructo(),
		ragingTroll(), ragingTroll(),
		inquisitorialSquad(), inquisitorialSquad(),
	}
	return ShuffleDarkArts(deck)
}

// func CreateMarketDeck() []Card {
// 	deck := []Card{
// 		crystalBall(),
// 		finite(),
// 		incendio(),
// 		oliverWood(),
// 		reparo(),
// 		triwizardCup(),
// 		albusDumbledore(),
// 		arthurWeasley(),
// 		bezoar(),
// 		choChang(),
// 		deluminator(),
// 		expectoPatronum(),
// 		felixFelicis(),
// 		filiusFlitwick(),
// 		hogwartsAHistory(),
// 		mollyWeasley(),
// 		quidditchGear(),
// 		siriusBlack(),
// 		stupefy(),
// 		sybillTrelawney(),
// 		butterbeer(),
// 		dobbyTheHouseElf(),
// 		essenceOfDittany(),
// 		fang(),
// 		fleurDelacour(),
// 		goldenSnitch(),
// 		pensieve(),
// 		rubeusHagrid(),
// 		advancedPotionMaking(),
// 		alastorMadEyeMoody(),
// 		argusFilchAndMrsNorris(),
// 		cedricDiggory(),
// 		descendo(),
// 		expelliarmus(),
// 		ginnyWeasley(),
// 		horaceSlughorn(),
// 		kingsleyShacklebolt(),
// 		lumos(),
// 		lunaLovegood(),
// 		nimbusTwoThousandAndOne(),
// 		nymphadoraTonks(),
// 		polyjuicePotion(),
// 		pomonaSprout(),
// 		severusSnape(),
// 		swordOfGryffindor(),
// 		tergeo(),
// 		viktorKrum(),
// 		fawkesThePhoenix(),
// 		minervaMcgonagall(),
// 		remusLupin(),
// 		elderWand(),
// 		chocolateFrog(),
// 		gilderoyLockhart(),
// 		maraudersMap(),
// 		protego(),
// 		accio(),
// 		fredWeasley(),
// 		georgeWeasley(),
// 		owls(),
// 		sortingHat(),
// 		wingardiumLeviosa(),
// 		petrificusTotalus(),
// 		harp(),
// 		finiteIncantatem(),
// 		confundus(),
// 		oldSock(),
// 	}

// 	deck = ShuffleCards(deck)

// 	return deck
// }

// returns (deck, current market).
func createNumAccurateMarketDeck() ([]Card, []Card) {
	deck := []Card{
		// Year 1
		wingardiumLeviosa(), wingardiumLeviosa(), wingardiumLeviosa(),
		descendo(), descendo(),
		reparo(), reparo(), reparo(), reparo(), reparo(), reparo(),
		lumos(), lumos(),
		incendio(), incendio(), incendio(), incendio(),
		essenceOfDittany(), essenceOfDittany(), essenceOfDittany(), essenceOfDittany(),
		quidditchGear(), quidditchGear(), quidditchGear(), quidditchGear(),
		goldenSnitch(),
		sortingHat(),
		oliverWood(),
		albusDumbledore(),
		rubeusHagrid(),
		// Year 2
		finite(), finite(),
		expelliarmus(), expelliarmus(),
		polyjuicePotion(), polyjuicePotion(),
		nimbusTwoThousandAndOne(), nimbusTwoThousandAndOne(),
		fawkesThePhoenix(),
		mollyWeasley(),
		dobbyTheHouseElf(),
		arthurWeasley(),
		ginnyWeasley(),
		gilderoyLockhart(),
		// Year 3
		petrificusTotalus(), petrificusTotalus(),
		expectoPatronum(), expectoPatronum(),
		maraudersMap(),
		crystalBall(), crystalBall(),
		butterbeer(), butterbeer(), butterbeer(),
		chocolateFrog(), chocolateFrog(), chocolateFrog(),
		remusLupin(), sybillTrelawney(), siriusBlack(),
		// Year 4
		protego(), protego(), protego(),
		accio(), accio(),
		pensieve(), triwizardCup(),
		hogwartsAHistory(), hogwartsAHistory(), hogwartsAHistory(), hogwartsAHistory(), hogwartsAHistory(), hogwartsAHistory(),
		fleurDelacour(),
		alastorMadEyeMoody(),
		filiusFlitwick(),
		pomonaSprout(),
		severusSnape(),
		minervaMcgonagall(),
		viktorKrum(),
		cedricDiggory(),
		// Year 5
		owls(), owls(),
		choChang(),
		georgeWeasley(),
		fredWeasley(),
		kingsleyShacklebolt(),
		nymphadoraTonks(),
		lunaLovegood(),
		stupefy(), stupefy(),
		// Year 6
		horaceSlughorn(),
		advancedPotionMaking(),
		bezoar(), bezoar(),
		deluminator(),
		elderWand(),
		felixFelicis(), felixFelicis(),
		confundus(), confundus(),
		// Year 7
		swordOfGryffindor(),
		// Box 1
		finiteIncantatem(), finiteIncantatem(),
		tergeo(), tergeo(), tergeo(), tergeo(), tergeo(), tergeo(),
		harp(),
		oldSock(), oldSock(),
		argusFilchAndMrsNorris(),
		fang(),
	}

	deck = ShuffleCards(deck)
	market := deck[0:6]
	deck = deck[6:]
	return deck, market
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
		gs.MarketDeck, _ = createNumAccurateMarketDeck()
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
