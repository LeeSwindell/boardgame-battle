package game

import "sync"

// players --> map[string]Player based on username
type Gamestate struct {
	Players     map[string]Player `json:"players"`
	Villains    []Villain         `json:"villains"`
	Locations   []Location        `json:"locations"`
	CurrentTurn Player            `json:"currentturn"`
	mu          sync.Mutex
}

type Player struct {
	Name      string
	Character string
	Health    int
	Damage    int
	Money     int
	Deck      Deck
	Hand      Hand
	Discard   Discard
	PlayArea  PlayArea
}

type Card struct {
	Id      int
	Name    string
	Effects []Effect
}

type Location struct {
	MaxControl int
	CurControl int
	Effect     []Effect
}

type PlayArea struct {
	Cards []Card
}

type Deck struct {
	Cards []Card
}

type Hand struct {
	Cards []Card
}

type Discard struct {
	Cards []Card
}

type Villain struct {
	Name        string
	CurDamage   int
	MaxHp       int
	Effect      Effect
	DeathEffect Effect
}

// Define an effect as something that changes the gamestate.
type Effect interface {
	Trigger(gs *Gamestate)
}
