package game

import "github.com/google/uuid"

var TestCard = Card{
	Id:      0,
	Name:    "TestCard",
	Effects: []Effect{DamageAllPlayers{amount: 1}},
}

var cards = map[string]Card{
	"TestCard": TestCard,
}

func alohamora() Card {
	return Card{
		Id:      int(uuid.New().ID()),
		Name:    "Alohamora",
		ImgPath: "/images/marketcards/alohamora.jpg",
		Effects: []Effect{GainMoney{amount: 1}},
	}
}
