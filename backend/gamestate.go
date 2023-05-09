package game

import "sync"

type Gamestate struct {
	Players     map[string]Player `json:"players"`
	Villains    []Villain         `json:"villains"`
	Locations   []Location        `json:"locations"`
	CurrentTurn Player            `json:"currentturn"`
	turnStats   TurnStats
	mu          sync.Mutex
}

type Player struct {
	Name      string
	Character string
	Health    int
	Damage    int
	Money     int
	Deck      []Card
	Hand      []Card
	Discard   []Card
	PlayArea  []Card
}

type Card struct {
	Id       int
	Name     string
	ImgPath  string
	CardType string
	Effects  []Effect
}

type Location struct {
	MaxControl int
	CurControl int
	Effect     []Effect
}

type Villain struct {
	Name        string
	Id          int
	ImgPath     string
	SetId       string
	CurDamage   int
	MaxHp       int
	Effect      Effect
	DeathEffect Effect
}

// Define an effect as something that changes the gamestate.
type Effect interface {
	Trigger(gs *Gamestate)
	// Describe() string
}

type TurnStats struct {
	AlliesPlayed   int
	ItemsPlayed    int
	SpellsPlayed   int
	VillainsKilled int
}
