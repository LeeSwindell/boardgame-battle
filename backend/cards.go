package game

import "github.com/google/uuid"

func alohamora() Card {
	return Card{
		Id:      int(uuid.New().ID()),
		Name:    "Alohomora",
		ImgPath: "/images/marketcards/alohomora.jpg",
		Effects: []Effect{GainMoney{amount: 1}},
	}
}
