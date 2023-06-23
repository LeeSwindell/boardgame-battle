package main

import (
	"log"
)

type DamageAllPlayers struct {
	Amount int
}

func (effect DamageAllPlayers) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		stunned := ChangePlayerHealth(user, -effect.Amount, gs)
		if stunned {
			StunPlayer(user, gs)
		}
	}
}

type DamageCurrentPlayer struct {
	Amount int
}

func (effect DamageCurrentPlayer) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	stunned := ChangePlayerHealth(user, -effect.Amount, gs)
	if stunned {
		StunPlayer(user, gs)
	}
}

type DamageAllPlayersButCurrent struct {
	Amount int
}

func (effect DamageAllPlayersButCurrent) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		if user != gs.CurrentTurn {
			stunned := ChangePlayerHealth(user, -effect.Amount, gs)
			if stunned {
				StunPlayer(user, gs)
			}
		}
	}
}

type GainMoney struct {
	Amount      int
	Description string
}

// Gives the current player Amount of money.
func (effect GainMoney) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	player := gs.Players[user]
	player.Money += effect.Amount
	gs.Players[user] = player
}

type ChooseOne struct {
	Effects []Effect

	// Options is the description given to user. The index of it should be the same as the Effect that it triggers.
	Options     []string `json:"options"`
	Description string   `json:"description"`
}

func (effect ChooseOne) Trigger(gs *Gamestate) {
	choice := getUserInput(gs.gameid, gs.CurrentTurn, effect)

	for i, option := range effect.Options {
		if choice == option {
			effect.Effects[i].Trigger(gs)
		}
	}
}

type GainDamage struct {
	Amount      int
	Description string
}

// Gives the current player Amount of damage
func (effect GainDamage) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	player := gs.Players[user]
	player.Damage += effect.Amount
	gs.Players[user] = player
}

type GainHealth struct {
	Amount int
}

// Gives the current player Amount of health
func (effect GainHealth) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	ChangePlayerHealth(user, effect.Amount, gs)
}

type GainDamagePerAllyPlayed struct{}

// Gives the current player damage per ally already played this turn.
func (effect GainDamagePerAllyPlayed) Trigger(gs *Gamestate) {
	amount := gs.turnStats.AlliesPlayed

	user := gs.CurrentTurn
	player := gs.Players[user]
	player.Damage += amount
	gs.Players[user] = player
}

type MoneyIfVillainKilled struct {
	Id     int
	Amount int
}

// Gives current player money (this turn) if a villain is killed.
func (effect MoneyIfVillainKilled) Trigger(gs *Gamestate) {
	// check if villain already killed
	if gs.turnStats.VillainsKilled > 0 {
		for name := range gs.Players {
			player, ok := gs.Players[name]
			if !ok {
				return
			}
			player.Damage += effect.Amount
			gs.Players[name] = player
		}
	} else {
		// Check to see if the turn has changed before this can take the lock.
		currentTurn := gs.CurrentTurn

		sub := Subscriber{
			id:              effect.Id,
			messageChan:     make(chan string),
			conditionMet:    "villain killed",
			conditionFailed: "end turn",
			unsubChan:       eventBroker.Messages,
		}

		go func() {
			eventBroker.Subscribe(sub)
			resChan := make(chan bool)
			go sub.Receive(resChan)

			for {
				res := <-resChan
				if !res {
					break
				}

				Logger("money if villain killed wants lock")
				gs.mu.Lock()
				if currentTurn == gs.CurrentTurn {
					user := gs.CurrentTurn
					player := gs.Players[user]
					player.Money += effect.Amount
					gs.Players[user] = player

					SendLobbyUpdate(gs.gameid, gs)
				}
				gs.mu.Unlock()
				Logger("money if villain killed releases lock")
			}
		}()

	}
}

type AllPlayersGainHealth struct {
	Amount int
}

func (effect AllPlayersGainHealth) Trigger(gs *Gamestate) {
	for _, p := range gs.Players {
		p.Health += effect.Amount
		gs.Players[p.Name] = p
	}
}

type AllPlayersGainMoney struct {
	Amount int
}

func (effect AllPlayersGainMoney) Trigger(gs *Gamestate) {
	for _, p := range gs.Players {
		p.Money += effect.Amount
		gs.Players[p.Name] = p
	}
}

type ActivePlayerDiscards struct {
	Amount int
}

func (effect ActivePlayerDiscards) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	player := gs.Players[user]

	cards := player.Hand
	if len(cards) == 0 {
		return
	}

	discardCardId := AskUserToSelectCard(user, gs.gameid, cards, "Discard a card")
	for i, c := range cards {
		if c.Id == discardCardId {
			cards = RemoveCardAtIndex(cards, i)
			player.Discard = append(player.Discard, c)
		}
	}

	player.Hand = cards
	gs.Players[user] = player

	event := Event{senderId: -1, message: "player discarded", data: user}
	eventBroker.Messages <- event
	// update turnstats
}

type AddToLocation struct {
	Amount int
}

func (effect AddToLocation) Trigger(gs *Gamestate) {
	loc := gs.Locations[gs.CurrentLocation]
	loc.CurControl += effect.Amount
	gs.Locations[gs.CurrentLocation] = loc
	Logger("sending location added event")
	// This is blocking! no one currently listening for event?
	eventBroker.Messages <- LocationAddedEvent
	Logger("sent location added event")

	if loc.CurControl >= loc.MaxControl {
		switch gs.CurrentLocation {
		case len(gs.Locations) - 1:
			log.Println("game over, loser!!!")
		default:
			gs.CurrentLocation += 1
		}
	}
}

type RemoveFromLocation struct {
	Amount int
}

func (effect RemoveFromLocation) Trigger(gs *Gamestate) {
	// Barty Crouch prevents removal of locations.
	for _, v := range gs.Villains {
		if v.Name == "Barty Crouch Jr." && v.Active {
			return
		}
	}

	loc := gs.Locations[gs.CurrentLocation]
	if loc.CurControl == 0 {
		return
	}

	loc.CurControl -= effect.Amount
	if loc.CurControl < 0 {
		loc.CurControl = 0
	}

	// For Lucius effect - happens only when location control Actually changes.
	for i := 0; i < gs.Locations[gs.CurrentLocation].CurControl-loc.CurControl; i++ {
		eventBroker.Messages <- LocationRemovedEvent
	}

	gs.Locations[gs.CurrentLocation] = loc
}

type DrawCards struct {
	Amount int
}

func (effect DrawCards) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	DrawXCards(user, gs, effect.Amount)
}

type AllDrawCards struct {
	Amount int
}

func (effect AllDrawCards) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		DrawXCards(user, gs, effect.Amount)
	}
}

type SendGameUpdateEffect struct{}

func (effect SendGameUpdateEffect) Trigger(gs *Gamestate) {
	SendLobbyUpdate(gs.gameid, gs)
}

type HealAnyIfVillainKilled struct {
	Id     int
	Amount int
}

func (effect HealAnyIfVillainKilled) Trigger(gs *Gamestate) {
	healAnyPlayer := HealAnyPlayer{Amount: effect.Amount}

	// check if villain already killed
	if gs.turnStats.VillainsKilled > 0 {
		healAnyPlayer.Trigger(gs)
	} else {
		// Check to see if the turn has changed before this can take the lock.
		currentTurn := gs.CurrentTurn

		sub := Subscriber{
			id:              effect.Id,
			messageChan:     make(chan string),
			conditionMet:    "villain killed",
			conditionFailed: "end turn",
			unsubChan:       eventBroker.Messages,
		}

		go func() {
			eventBroker.Subscribe(sub)
			resChan := make(chan bool)
			go sub.Receive(resChan)

			for {
				res := <-resChan
				if !res {
					break
				}

				Logger("heal if villain killed wants lock")
				gs.mu.Lock()
				if res && currentTurn == gs.CurrentTurn {
					healAnyPlayer.Trigger(gs)

					SendLobbyUpdate(gs.gameid, gs)
				}
				gs.mu.Unlock()
				Logger("heal if villain killed releases lock")
			}
		}()
	}
}

type HealAnyPlayer struct {
	Amount int
}

func (effect HealAnyPlayer) Trigger(gs *Gamestate) {
	playernames := []string{}
	for p := range gs.Players {
		playernames = append(playernames, p)
	}

	choice := AskUserToSelectPlayer(0, gs.CurrentTurn, playernames)
	ChangePlayerHealth(choice, effect.Amount, gs)
}

type AllSearchDiscardPileForItem struct{}

func (effect AllSearchDiscardPileForItem) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		choices := []Card{}
		for _, c := range gs.Players[user].Discard {
			if c.CardType == "item" {
				choices = append(choices, c)
			}
		}

		if len(choices) != 0 {
			prompt := "Choose an item from your discard to gain to your hand!"
			cardId := AskUserToSelectCard(user, gs.gameid, choices, prompt)
			MoveCardFromDiscardToHand(user, cardId, gs)
		}
	}
}

type HealAllVillains struct {
	Amount int
}

func (effect HealAllVillains) Trigger(gs *Gamestate) {
	for i, v := range gs.Villains {
		// check active to avoid healing dead heroes with their own death effect
		if v.Active && v.CurDamage > 0 {
			gs.Villains[i].CurDamage -= effect.Amount
			if gs.Villains[i].CurDamage < 0 {
				gs.Villains[i].CurDamage = 0
			}
		}
	}
}

type AllChooseOne struct {
	Effects []Effect

	// Options is the description given to user. The index of it should be the same as the Effect that it triggers.
	Options     []string `json:"options"`
	Description string   `json:"description"`
}

func (effect AllChooseOne) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		choice := getUserInput(gs.gameid, user, effect)

		for i, option := range effect.Options {
			if choice == option {
				effect.Effects[i].Trigger(gs)
			}
		}
	}
}

type ChosenPlayerSearchesDiscardForX struct {
	SearchType string
	Playername string
}

func (effect ChosenPlayerSearchesDiscardForX) Trigger(gs *Gamestate) {
	choices := []Card{}
	for _, c := range gs.Players[effect.Playername].Discard {
		if c.CardType == effect.SearchType {
			choices = append(choices, c)
		}
	}

	if len(choices) != 0 {
		prompt := "Choose one to gain to your hand (from discard)"
		cardId := AskUserToSelectCard(effect.Playername, gs.gameid, choices, prompt)
		MoveCardFromDiscardToHand(effect.Playername, cardId, gs)
	}
}

type AllSearchHandOrDiscardForX struct {
	SearchType string
	prompt     string
}

func (effect AllSearchHandOrDiscardForX) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		player := gs.Players[user]
		choices := []Card{}
		for _, c := range player.Hand {
			if c.CardType == effect.SearchType {
				choices = append(choices, c)
			}
		}
		for _, c := range player.Discard {
			if c.CardType == effect.SearchType {
				choices = append(choices, c)
			}
		}
		if len(choices) != 0 {
			cardId := AskUserToSelectCard(user, gs.gameid, choices, effect.prompt)
			MoveCardFromDiscardToHand(user, cardId, gs)
		}
	}
}

type ChosenPlayerGainsHealth struct {
	Playername string
	Amount     int
}

func (effect ChosenPlayerGainsHealth) Trigger(gs *Gamestate) {
	ChangePlayerHealth(effect.Playername, effect.Amount, gs)
}

type GainDetentionToDiscard struct {
	// whether to give the active player the detention, or not.
	Active bool
}

func (effect GainDetentionToDiscard) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	player := gs.Players[user]
	player.Discard = append(player.Discard, detention())
	gs.Players[user] = player
}

type AllBanishItem struct{}

func (effect AllBanishItem) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		player := gs.Players[user]
		choices := []Card{}
		for _, c := range player.Hand {
			if c.CardType == "item" {
				choices = append(choices, c)
			}
		}
		for _, c := range player.Discard {
			if c.CardType == "item" {
				choices = append(choices, c)
			}
		}
		if len(choices) != 0 {
			cardId := AskUserToSelectCard(user, gs.gameid, choices, "Choose a card to banish")
			BanishCard(user, cardId, gs)
		}
	}
}

type AllBanishCard struct{}

func (effect AllBanishCard) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		player := gs.Players[user]
		choices := append(player.Hand, player.Discard...)

		if len(choices) != 0 {
			cardId := AskUserToSelectCard(user, gs.gameid, choices, "Choose a card to banish")
			BanishCard(user, cardId, gs)
		}
	}
}

type DamageActivePerDetention struct {
	Amount int
}

func (effect DamageActivePerDetention) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	damage := 0
	for _, c := range gs.Players[user].Hand {
		if c.Name == "Detention!" {
			damage--
		}
	}
	if damage > 0 {
		stunned := ChangePlayerHealth(user, damage, gs)
		if stunned {
			StunPlayer(user, gs)
		}
	}
}
