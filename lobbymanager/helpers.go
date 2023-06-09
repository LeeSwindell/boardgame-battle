package main

// types for handlers.

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
	Id       int      `json:"Id"`
	Name     string   `json:"Name"`
	SetId    string   `json:"SetId"`
	ImgPath  string   `json:"ImgPath"`
	CardType string   `json:"CardType"`
	Cost     int      `json:"Cost"`
	Effects  []Effect `json:"Effects"`
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
	Name        string
	Id          int
	ImgPath     string
	SetId       string
	CurDamage   int
	MaxHp       int
	Active      bool
	Effect      []Effect
	DeathEffect []Effect

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
}

// Define an effect as something that changes the gamestate.
type Effect interface {
	Trigger(gs *Gamestate)
}

type ChooseOne struct {
	Effects []Effect

	// Options is the description given to user. The index of it should be the same as the Effect that it triggers.
	Options     []string `json:"options"`
	Description string   `json:"description"`
}
