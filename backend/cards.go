package game

import "github.com/google/uuid"

func alohamora() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Alohomora",
		SetId:    "game 1",
		ImgPath:  "/images/starters/alohomora.jpg",
		CardType: "spell",
		Cost:     0,
		Effects:  []Effect{GainMoney{Amount: 1}},
	}
}

func pigwidgeon() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Pigwidgeon",
		SetId:    "game 1",
		ImgPath:  "/images/starters/pigwidgeon.jpg",
		CardType: "ally",
		Cost:     0,
		Effects: []Effect{
			ChooseOne{
				Effects: []Effect{
					GainDamage{Amount: 1},
					GainHealth{Amount: 2},
				},
				Options: []string{"Gain 1 Damage", "Gain 2 Health"},
			},
		},
	}
}

func cleansweep() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Cleansweep 11",
		SetId:    "game 1",
		ImgPath:  "/images/starters/cleansweep.jpg",
		CardType: "item",
		Cost:     0,
		Effects: []Effect{
			GainDamage{Amount: 1},
			MoneyIfVillainKilled{Id: id, Amount: 1},
		},
	}
}

func bertieBottsEveryFlavourBeans() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Bertie Botts Every-Flavour Beans",
		SetId:    "game 1",
		ImgPath:  "/images/starters/bertiebottseveryflavourbeans.jpg",
		CardType: "item",
		Cost:     0,
		Effects: []Effect{
			GainMoney{Amount: 1},
			GainDamagePerAllyPlayed{},
		},
	}
}

func crystalBall() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Crystal Ball",
		SetId:    "game 3",
		ImgPath:  "/images/marketcards/crystalball.jpg",
		CardType: "item",
		Cost:     3,
		Effects: []Effect{
			DrawCards{Amount: 2},
			SendGameUpdateEffect{},
			ActivePlayerDiscards{Amount: 1},
		},
	}
}

func finite() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Finite!",
		SetId:    "game 3",
		ImgPath:  "/images/marketcards/finite.jpg",
		CardType: "spell",
		Cost:     3,
		Effects:  []Effect{RemoveFromLocation{Amount: 1}},
	}
}

func incendio() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Incendio!",
		SetId:    "game 1",
		ImgPath:  "/images/marketcards/incendio.jpg",
		CardType: "spell",
		Cost:     4,
		Effects: []Effect{
			GainDamage{Amount: 1},
			DrawCards{Amount: 1},
		},
	}
}

func oliverWood() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Oliver Wood",
		SetId:    "game 1",
		ImgPath:  "/images/marketcards/oliverwood.jpg",
		CardType: "ally",
		Cost:     3,
		Effects: []Effect{
			GainDamage{Amount: 1},
			HealAnyIfVillainKilled{Amount: 2, Id: id},
		},
	}
}

func reparo() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Reparo",
		SetId:    "game 1",
		ImgPath:  "/images/marketcards/reparo.jpg",
		CardType: "spell",
		Cost:     3,
		Effects: []Effect{
			ChooseOne{
				Effects: []Effect{
					GainMoney{Amount: 2},
					DrawCards{Amount: 1},
				},
				Options: []string{"Gain 2 Money", "Draw 1 Card"},
			},
		},
	}
}

func triwizardCup() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Triwizard Cup",
		SetId:    "game 4",
		ImgPath:  "/images/marketcards/triwizardcup.jpg",
		CardType: "item",
		Cost:     5,
		Effects: []Effect{
			GainDamage{Amount: 1},
			GainMoney{Amount: 1},
			GainHealth{Amount: 1},
		},
	}
}
