package game

import "github.com/google/uuid"

func alohamora() Card {
	return Card{
		Id:      int(uuid.New().ID()),
		Name:    "Alohomora",
		ImgPath: "/images/marketcards/alohomora.jpg",
		Effects: []Effect{GainMoney{Amount: 1}},
	}
}

func pigwidgeon() Card {
	return Card{
		Id:      int(uuid.New().ID()),
		Name:    "Pigwidgeon",
		ImgPath: "/images/marketcards/pigwidgeon.jpg",
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
