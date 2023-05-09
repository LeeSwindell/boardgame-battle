package game

import "github.com/google/uuid"

func alohamora() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Alohomora",
		ImgPath:  "/images/starters/alohomora.jpg",
		CardType: "spell",
		Effects:  []Effect{GainMoney{Amount: 1}},
	}
}

func pigwidgeon() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Pigwidgeon",
		ImgPath:  "/images/starters/pigwidgeon.jpg",
		CardType: "ally",
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
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Cleansweep 11",
		ImgPath:  "/images/starters/cleansweep.jpg",
		CardType: "item",
		Effects:  []Effect{},
	}
}

func bertieBottsEveryFlavourBeans() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Bertie Botts Every-Flavour Beans",
		ImgPath:  "/images/starters/bertiebottseveryflavourbeans.jpg",
		CardType: "item",
		Effects: []Effect{
			GainMoney{Amount: 1},
			GainDamagePerAllyPlayed{},
		},
	}
}
