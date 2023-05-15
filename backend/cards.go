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
					GainDamage{Amount: 1, Description: "gain 1 damage"},
					GainHealth{Amount: 2, Description: "gain 2 health"},
				},
				Options:     []string{"Gain 1 Damage", "Gain 2 Health"},
				Description: "choose one of these effects, bub",
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
