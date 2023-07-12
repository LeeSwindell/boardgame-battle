package main

import "sync"

type Gamestate struct {
	Players         map[string]Player `json:"players"`
	Villains        []Villain         `json:"villains"`
	Locations       []Location        `json:"locations"`
	DarkArts        []DarkArt         `json:"darkarts"`
	MarketDeck      []Card            `json:"marketdeck"`
	Market          []Card            `json:"market"`
	CurrentTurn     string            `json:"currentturn"`
	TurnOrder       []string          `json:"turnorder"`
	CurrentDarkArt  int               `json:"currentdarkart"`
	CurrentLocation int               `json:"currentlocation"`
	DarkArtsPlayed  []DarkArt         `json:"darkartsplayed"`
	villainDeck     []Villain
	turnNumber      int
	turnStats       TurnStats
	started         bool
	mu              sync.Mutex
	gameid          int
}

type Player struct {
	Name         string
	Character    string
	Health       int
	Damage       int
	Money        int
	Deck         []Card
	Hand         []Card
	Discard      []Card
	PlayArea     []Card
	spellsToDeck bool
	itemsToDeck  bool
	alliesToDeck bool
}

type Card struct {
	Id        int      `json:"Id"`
	Name      string   `json:"Name"`
	SetId     string   `json:"SetId"`
	ImgPath   string   `json:"ImgPath"`
	CardType  string   `json:"CardType"`
	Cost      int      `json:"Cost"`
	Effects   []Effect `json:"Effects"`
	onDiscard func(target string, gs *Gamestate)
}

type Location struct {
	Name       string
	Id         int
	SetId      string
	ImgPath    string
	MaxControl int
	CurControl int
	Effect     Effect
}

type Villain struct {
	Name         string
	Id           int
	ImgPath      string
	SetId        string
	CurDamage    int
	MaxHp        int
	Active       bool
	BlockedUntil int
	villainType  string
	Effect       []Effect
	DeathEffect  []Effect

	// true if this villain should be played before DA events - for triggered events.
	playBeforeDA bool
}

type DarkArt struct {
	Name    string
	Id      int
	ImgPath string
	SetId   string
	Effects []Effect
}

type TurnStats struct {
	AlliesPlayed   int
	ItemsPlayed    int
	SpellsPlayed   int
	VillainsKilled int
	VillainsHit    []int
}

// Define an effect as something that changes the gamestate.
type Effect interface {
	Trigger(gs *Gamestate)
}

type TargetedEffect struct {
	Target string // a playername.
	Effect Effect
}
