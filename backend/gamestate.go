package game

type Gamestate struct {
	players  []Player
	villains []Villain
}

type Player struct {
	name      string
	character string
	health    int
	damage    int
	money     int
	deck      Deck
	hand      Hand
	discard   Discard
	playArea  PlayArea
}

type Card struct {
	id      int
	name    string
	effects []Effect
}

type PlayArea struct {
	cards []Card
}

type Deck struct {
	cards []Card
}

type Hand struct {
	cards []Card
}

type Discard struct {
	cards []Card
}

type Villain struct {
	name string
}

type Effect struct {
}
