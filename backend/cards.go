package game

var TestCard = Card{
	Id:      0,
	Name:    "TestCard",
	Effects: []Effect{DamageAllPlayers{amount: 1}},
}

var cards = map[string]Card{
	"TestCard": TestCard,
}
