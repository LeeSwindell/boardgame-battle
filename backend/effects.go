package main

import (
	"log"
	"math/rand"
)

// Wrapper for changing stats of a given player
type ChangeStats struct {
	Target          string
	AmountHealth    int
	AmountDamage    int
	AmountMoney     int
	AmountCards     int
	AmountToDiscard int
	DiscardPrompt   string
	Cause           string
}

func (effect ChangeStats) Trigger(gs *Gamestate) {
	player := gs.Players[effect.Target]
	player.Damage += effect.AmountDamage
	player.Money += effect.AmountMoney
	gs.Players[effect.Target] = player

	stunned := ChangePlayerHealth(effect.Target, effect.AmountHealth, gs)
	if stunned {
		StunPlayer(effect.Target, gs)
	}
	DrawXCards(effect.Target, gs, effect.AmountCards)
	if effect.AmountToDiscard != 0 {
		GivenPlayerDiscards{User: effect.Target, Prompt: effect.DiscardPrompt, Amount: effect.AmountToDiscard}.Trigger(gs)
		cond := effect.Cause == "villain" || effect.Cause == "darkart"
		if player.Proficiency == "Defense Against the Dark Arts" && cond {
			ChangeStats{Target: effect.Target, AmountDamage: 1, AmountHealth: 1}.Trigger(gs)
		}
	}
}

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
	SendLobbyUpdate(gs.gameid, gs)
	choice := getUserInput(gs.gameid, gs.CurrentTurn, effect.Options, effect.Description)

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

// Active player triggers the effect GainX if a villain is killed.
type GainXIfVillainKilled struct {
	GainX ChangeStats
	Id    int
}

// Gives current player X (this turn) if a villain is killed.
func (effect GainXIfVillainKilled) Trigger(gs *Gamestate) {
	effect.GainX.Target = gs.CurrentTurn

	// check if villain already killed
	if gs.turnStats.VillainsKilled > 0 {
		for name := range gs.Players {
			player, ok := gs.Players[name]
			if !ok {
				return
			}
			effect.GainX.Trigger(gs)
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

				gs.mu.Lock()
				if currentTurn == gs.CurrentTurn {
					effect.GainX.Trigger(gs)

					SendLobbyUpdate(gs.gameid, gs)
				}
				gs.mu.Unlock()
			}
		}()

	}
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
	for user := range gs.Players {
		ChangePlayerHealth(user, effect.Amount, gs)
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

type GivenPlayerDiscards struct {
	User   string
	Amount int
	Prompt string
}

func (effect GivenPlayerDiscards) Trigger(gs *Gamestate) {
	user := effect.User
	player := gs.Players[user]

	cards := player.Hand
	if len(cards) == 0 {
		return
	}

	if effect.Prompt == "" {
		effect.Prompt = "Discard a card"
	}

	for i := 0; i < effect.Amount; i++ {
		SendLobbyUpdate(gs.gameid, gs)
		discardCardId := AskUserToSelectCard(user, gs.gameid, cards, effect.Prompt)
		DiscardFromId(user, discardCardId, gs)
		cards = gs.Players[user].Hand
		if len(cards) == 0 {
			return
		}
	}
}

type ActivePlayerDiscards struct {
	Amount int
	Prompt string
	Cause  string
}

func (effect ActivePlayerDiscards) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	player := gs.Players[user]

	cards := player.Hand
	if len(cards) == 0 {
		return
	}

	if effect.Prompt == "" {
		effect.Prompt = "Discard a card"
	}

	for i := 0; i < effect.Amount; i++ {
		SendLobbyUpdate(gs.gameid, gs)
		discardCardId := AskUserToSelectCard(user, gs.gameid, cards, effect.Prompt)
		DiscardFromId(user, discardCardId, gs)
		cards = gs.Players[user].Hand

		cond := effect.Cause == "villain" || effect.Cause == "darkart"
		if player.Proficiency == "Defense Against the Dark Arts" && cond {
			ChangeStats{Target: user, AmountDamage: 1, AmountHealth: 1}.Trigger(gs)
		}

		if len(cards) == 0 {
			return
		}
	}
}

type AddToLocation struct {
	Amount int
}

func (effect AddToLocation) Trigger(gs *Gamestate) {
	loc := gs.Locations[gs.CurrentLocation]
	loc.CurControl += effect.Amount
	gs.Locations[gs.CurrentLocation] = loc
	eventBroker.Messages <- LocationAddedEvent

	if loc.CurControl >= loc.MaxControl {
		switch gs.CurrentLocation {
		case len(gs.Locations) - 1:
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
	log.Println("removing from location: start amount", loc.CurControl)
	if loc.CurControl == 0 {
		return
	}

	startingControl := loc.CurControl
	loc.CurControl -= effect.Amount
	if loc.CurControl < 0 {
		loc.CurControl = 0
	}

	log.Println("removing from location: 391 amount", loc.CurControl)

	// For Lucius effect - happens only when location control Actually changes.
	for i := 0; i < gs.Locations[gs.CurrentLocation].CurControl-loc.CurControl; i++ {
		eventBroker.Messages <- LocationRemovedEvent
	}

	log.Println("removing from location: 399 amount", loc.CurControl)

	// For Harry's character effect.
	for _, p := range gs.Players {
		if p.Character == "Harry" {
			HealAmount := startingControl - loc.CurControl
			AllPlayersGainHealth{Amount: HealAmount}.Trigger(gs)
		}
	}
	log.Println("removing from location: 407 amount", loc.CurControl)

	gs.Locations[gs.CurrentLocation] = loc
}

// current player draws Amount of cards.
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

	if len(playernames) == 0 {
		return
	}
	SendLobbyUpdate(gs.gameid, gs)
	choice := AskUserToSelectPlayer(gs.gameid, gs.CurrentTurn, playernames)
	ChangePlayerHealth(choice, effect.Amount, gs)
}

type SelectPlayerToGainStats struct {
	AmountHealth int
	AmountMoney  int
	AmountDamage int
	AmountCards  int
	ExcludeUser  string
}

func (effect SelectPlayerToGainStats) Trigger(gs *Gamestate) {
	playernames := []string{}
	for p := range gs.Players {
		if p != effect.ExcludeUser {
			playernames = append(playernames, p)
		}
	}

	if len(playernames) == 0 {
		return
	}
	SendLobbyUpdate(gs.gameid, gs)
	choice := AskUserToSelectPlayer(gs.gameid, gs.CurrentTurn, playernames)

	// Use helpers to change player these values before getting/setting player
	ChangePlayerHealth(choice, effect.AmountHealth, gs)
	DrawXCards(choice, gs, effect.AmountCards)

	player := gs.Players[choice]
	player.Damage += effect.AmountDamage
	player.Money += effect.AmountMoney
	gs.Players[choice] = player
}

type SelectTwoPlayersToGainStats struct {
	AmountHealth int
	AmountMoney  int
	AmountDamage int
	AmountCards  int
	Exclusive    bool
}

func (effect SelectTwoPlayersToGainStats) Trigger(gs *Gamestate) {
	playernames := []string{}
	for p := range gs.Players {
		playernames = append(playernames, p)
	}
	if len(playernames) == 0 {
		return
	}
	SendLobbyUpdate(gs.gameid, gs)
	choice := AskUserToSelectPlayer(gs.gameid, gs.CurrentTurn, playernames)

	ChangePlayerHealth(choice, effect.AmountHealth, gs)
	DrawXCards(choice, gs, effect.AmountCards)

	player := gs.Players[choice]
	player.Damage += effect.AmountDamage
	player.Money += effect.AmountMoney
	gs.Players[choice] = player

	if len(playernames) <= 1 {
		return
	}

	secondEffect := SelectPlayerToGainStats{
		AmountHealth: effect.AmountHealth,
		AmountMoney:  effect.AmountMoney,
		AmountDamage: effect.AmountDamage,
		AmountCards:  effect.AmountCards,
	}
	if effect.Exclusive {
		secondEffect.ExcludeUser = choice
	}

	secondEffect.Trigger(gs)
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
			SendLobbyUpdate(gs.gameid, gs)
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
		SendLobbyUpdate(gs.gameid, gs)
		choice := getUserInput(gs.gameid, user, effect.Options, effect.Description)

		for i, option := range effect.Options {
			if choice == option {
				effect.Effects[i].Trigger(gs)
			}
		}
	}
}

type GivenPlayerChooseOneTargeted struct {
	// funcs that create an effect with a given target.
	User            string
	EffectTargeting []func(target string, effect Effect) Effect

	Effects []Effect
	// Options is the description given to user. The index of it should be the same as the Effect that it triggers.
	Options     []string `json:"options"`
	Description string   `json:"description"`
}

func (effect GivenPlayerChooseOneTargeted) Trigger(gs *Gamestate) {
	user := effect.User
	SendLobbyUpdate(gs.gameid, gs)
	choice := getUserInput(gs.gameid, user, effect.Options, effect.Description)

	for i, option := range effect.Options {
		if choice == option {
			e := effect.EffectTargeting[i](user, effect.Effects[i])
			e.Trigger(gs)
		}
	}
}

type AllChooseOneTargeted struct {
	// funcs that create an effect with a given target.
	EffectTargeting []func(target string, effect Effect) Effect

	Effects []Effect
	// Options is the description given to user. The index of it should be the same as the Effect that it triggers.
	Options     []string `json:"options"`
	Description string   `json:"description"`
}

func (effect AllChooseOneTargeted) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		SendLobbyUpdate(gs.gameid, gs)
		choice := getUserInput(gs.gameid, user, effect.Options, effect.Description)

		for i, option := range effect.Options {
			if choice == option {
				e := effect.EffectTargeting[i](user, effect.Effects[i])
				e.Trigger(gs)
			}
		}
	}
}

// return a new version, that way the cards effect isn't changed every other time it gets triggered.
func TargetCreateStats(target string, effect Effect) Effect {
	changeStats, ok := effect.(ChangeStats)
	if !ok {
		log.Println("Type assertion failed: TargetChangeStats")
	}
	changeStats.Target = target
	return changeStats
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
		SendLobbyUpdate(gs.gameid, gs)
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
			SendLobbyUpdate(gs.gameid, gs)
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
	stunned := ChangePlayerHealth(effect.Playername, effect.Amount, gs)
	if stunned {
		StunPlayer(effect.Playername, gs)
	}
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

type GainDetentionToHand struct {
	// whether to give the active player the detention, or not.
	Active bool
}

func (effect GainDetentionToHand) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	player := gs.Players[user]
	player.Hand = append(player.Hand, detention())
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
			SendLobbyUpdate(gs.gameid, gs)
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
			SendLobbyUpdate(gs.gameid, gs)
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

type DamageAllPerDetention struct {
	Amount int
}

func (effect DamageAllPerDetention) Trigger(gs *Gamestate) {
	for user := range gs.Players {
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
}

type AllDiscard struct {
	Amount int // doesn't do anything.
	Prompt string
	Cause  string
}

// Only discards one card atm.
func (effect AllDiscard) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		player := gs.Players[user]

		cards := player.Hand
		if len(cards) == 0 {
			return
		}

		SendLobbyUpdate(gs.gameid, gs)
		discardCardId := AskUserToSelectCard(user, gs.gameid, cards, effect.Prompt)
		for i, c := range cards {
			if c.Id == discardCardId {
				cards = RemoveCardAtIndex(cards, i)
				player.Discard = append(player.Discard, c)

				// Wrap the player mapping around onDiscard since it mutates the state directly.
				if c.onDiscard != nil {
					gs.Players[user] = player
					c.onDiscard(user, gs)
					player = gs.Players[user]
				}
			}
		}

		player.Hand = cards
		gs.Players[user] = player

		cond := effect.Cause == "villain" || effect.Cause == "darkart"
		if player.Proficiency == "Defense Against the Dark Arts" && cond {
			ChangeStats{Target: user, AmountDamage: 1, AmountHealth: 1}.Trigger(gs)
		}

		event := Event{senderId: -1, message: "player discarded", data: user}
		eventBroker.Messages <- event
	}
}

type DamageAllPerMatchingCost struct {
	Cost   int
	Amount int
}

func (effect DamageAllPerMatchingCost) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		damage := 0
		for _, c := range gs.Players[user].Hand {
			if c.Cost == effect.Cost {
				damage += effect.Amount
			}
		}
		if damage > 0 {
			stunned := ChangePlayerHealth(user, -damage, gs)
			if stunned {
				StunPlayer(user, gs)
			}
		}
	}
}

type DamageActivePerMatchingCost struct {
	Cost   int
	Amount int
}

func (effect DamageActivePerMatchingCost) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	damage := 0
	for _, c := range gs.Players[user].Hand {
		if c.Cost == effect.Cost {
			damage += effect.Amount
		}
	}
	if damage > 0 {
		stunned := ChangePlayerHealth(user, -damage, gs)
		if stunned {
			StunPlayer(user, gs)
		}
	}
}

type DamageActivePerCardGreaterThanCost struct {
	Cost   int
	Amount int
}

func (effect DamageActivePerCardGreaterThanCost) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	damage := 0
	for _, c := range gs.Players[user].Hand {
		if c.Cost >= effect.Cost {
			damage += effect.Amount
		}
	}
	if damage > 0 {
		stunned := ChangePlayerHealth(user, -damage, gs)
		if stunned {
			StunPlayer(user, gs)
		}
	}
}

type AllPlayersGainDamage struct {
	Amount int
}

func (effect AllPlayersGainDamage) Trigger(gs *Gamestate) {
	for _, p := range gs.Players {
		p.Damage += effect.Amount
		gs.Players[p.Name] = p
	}
}

type RavenclawDice struct{}

func (effect RavenclawDice) Trigger(gs *Gamestate) {
	n := rand.Intn(6)

	if gs.Players[gs.CurrentTurn].Proficiency == "Arithmancy" {
		options := []string{"Yes", "No"}
		switch n {
		case 0:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Money, would you like to reroll?")
			if choice == "No" {
				AllPlayersGainMoney{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		case 1:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Health, would you like to reroll?")
			if choice == "No" {
				AllPlayersGainHealth{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		case 2:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Damage, would you like to reroll?")
			if choice == "No" {
				AllPlayersGainDamage{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		default:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Cards, would you like to reroll?")
			if choice == "No" {
				AllDrawCards{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		}
	}

	switch n {
	case 0:
		AllPlayersGainMoney{Amount: 1}.Trigger(gs)
	case 1:
		AllPlayersGainHealth{Amount: 1}.Trigger(gs)
	case 2:
		AllPlayersGainDamage{Amount: 1}.Trigger(gs)
	default:
		AllDrawCards{Amount: 1}.Trigger(gs)
	}
}

type SlytherinDice struct{}

func (effect SlytherinDice) Trigger(gs *Gamestate) {
	n := rand.Intn(6)

	if gs.Players[gs.CurrentTurn].Proficiency == "Arithmancy" {
		options := []string{"Yes", "No"}
		switch n {
		case 0:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Money, would you like to reroll?")
			if choice == "No" {
				AllPlayersGainMoney{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		case 1:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Health, would you like to reroll?")
			if choice == "No" {
				AllPlayersGainHealth{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		case 2:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Cards, would you like to reroll?")
			if choice == "No" {
				AllDrawCards{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		default:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Damage, would you like to reroll?")
			if choice == "No" {
				AllPlayersGainDamage{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		}
	}

	switch n {
	case 0:
		AllPlayersGainMoney{Amount: 1}.Trigger(gs)
	case 1:
		AllPlayersGainHealth{Amount: 1}.Trigger(gs)
	case 2:
		AllDrawCards{Amount: 1}.Trigger(gs)
	default:
		AllPlayersGainDamage{Amount: 1}.Trigger(gs)
	}
}

type GryffindorDice struct{}

func (effect GryffindorDice) Trigger(gs *Gamestate) {
	n := rand.Intn(6)

	if gs.Players[gs.CurrentTurn].Proficiency == "Arithmancy" {
		options := []string{"Yes", "No"}
		switch n {
		case 0:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Damage, would you like to reroll?")
			if choice == "No" {
				AllPlayersGainDamage{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		case 1:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Health, would you like to reroll?")
			if choice == "No" {
				AllPlayersGainHealth{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		case 2:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Cards, would you like to reroll?")
			if choice == "No" {
				AllDrawCards{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		default:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Money, would you like to reroll?")
			if choice == "No" {
				AllPlayersGainMoney{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		}
	}

	switch n {
	case 0:
		AllPlayersGainDamage{Amount: 1}.Trigger(gs)
	case 1:
		AllPlayersGainHealth{Amount: 1}.Trigger(gs)
	case 2:
		AllDrawCards{Amount: 1}.Trigger(gs)
	default:
		AllPlayersGainMoney{Amount: 1}.Trigger(gs)
	}
}

type HufflepuffDice struct{}

func (effect HufflepuffDice) Trigger(gs *Gamestate) {
	n := rand.Intn(6)

	if gs.Players[gs.CurrentTurn].Proficiency == "Arithmancy" {
		options := []string{"Yes", "No"}
		switch n {
		case 0:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Money, would you like to reroll?")
			if choice == "No" {
				AllPlayersGainMoney{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		case 1:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Damage, would you like to reroll?")
			if choice == "No" {
				AllPlayersGainDamage{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		case 2:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Cards, would you like to reroll?")
			if choice == "No" {
				AllDrawCards{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		default:
			SendLobbyUpdate(gs.gameid, gs)
			choice := getUserInput(gs.gameid, gs.CurrentTurn, options, "You rolled Health, would you like to reroll?")
			if choice == "No" {
				AllPlayersGainHealth{Amount: 1}.Trigger(gs)
				return
			}
			n = rand.Intn(6)
		}
	}

	switch n {
	case 0:
		AllPlayersGainMoney{Amount: 1}.Trigger(gs)
	case 1:
		AllPlayersGainDamage{Amount: 1}.Trigger(gs)
	case 2:
		AllDrawCards{Amount: 1}.Trigger(gs)
	default:
		AllPlayersGainHealth{Amount: 1}.Trigger(gs)
	}
}

type ChooseTwo struct {
	Exclusive bool
	Effects   []Effect
	Options   []string
	Prompt    string
}

func (effect ChooseTwo) Trigger(gs *Gamestate) {
	firstChoice := ChooseOne{
		Effects:     effect.Effects,
		Options:     effect.Options,
		Description: effect.Prompt + "(1 of 2)",
	}
	SendLobbyUpdate(gs.gameid, gs)
	choice := getUserInput(gs.gameid, gs.CurrentTurn, firstChoice.Options, firstChoice.Description)

	secondChoice := ChooseOne{}
	for i, option := range firstChoice.Options {
		if choice == option {
			firstChoice.Effects[i].Trigger(gs)
			secondChoice.Effects = append(firstChoice.Effects[:i], firstChoice.Effects[i+1:]...)
			secondChoice.Options = append(firstChoice.Options[:i], firstChoice.Options[i+1:]...)
			secondChoice.Description = effect.Prompt + "(2 of 2)"
		}
	}
	SendLobbyUpdate(gs.gameid, gs)
	choice = getUserInput(gs.gameid, gs.CurrentTurn, secondChoice.Options, secondChoice.Description)
	for i, option := range secondChoice.Options {
		if choice == option {
			secondChoice.Effects[i].Trigger(gs)
		}
	}
}

// Gives current user health/money/dmg/cards if a cardtype X is played during their turn.
type GainStatIfXPlayed struct {
	AmountHealth int
	AmountDamage int
	AmountMoney  int
	AmountCards  int
	Cardtype     string
	Id           int
}

// Triggers off its own card, e.g. play an ally card which triggers when an ally is played.
func (effect GainStatIfXPlayed) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	triggerEffect := false
	changeStats := ChangeStats{
		Target:       user,
		AmountHealth: effect.AmountHealth,
		AmountDamage: effect.AmountDamage,
		AmountMoney:  effect.AmountMoney,
		AmountCards:  effect.AmountCards,
	}

	switch effect.Cardtype {
	case "ally":
		if gs.turnStats.AlliesPlayed > 0 {
			triggerEffect = true
		}
	case "item":
		if gs.turnStats.ItemsPlayed > 0 {
			triggerEffect = true
		}
	case "spell":
		if gs.turnStats.SpellsPlayed > 0 {
			triggerEffect = true
		}
	}

	if triggerEffect {
		changeStats.Trigger(gs)
		return
	}

	sub := Subscriber{
		id:              effect.Id,
		messageChan:     make(chan string),
		conditionMet:    effect.Cardtype + " played",
		conditionFailed: "end turn",
		unsubChan:       eventBroker.Messages,
	}

	// Only trigger once.
	go func() {
		eventBroker.Subscribe(sub)
		resChan := make(chan bool)
		go sub.Receive(resChan)

		res := <-resChan
		if !res {
			return
		}

		gs.mu.Lock()
		// check if turn has changed while waiting for lock.
		if res && user == gs.CurrentTurn {
			changeStats.Trigger(gs)

			SendLobbyUpdate(gs.gameid, gs)
		}
		gs.mu.Unlock()
	}()
}

type AllPlayersAtMaxHealthGainX struct {
	// AmountHealth int
	AmountMoney  int
	AmountDamage int
	AmountCards  int
}

func (effect AllPlayersAtMaxHealthGainX) Trigger(gs *Gamestate) {
	for user, p := range gs.Players {
		if p.Health >= 10 {
			DrawXCards(user, gs, effect.AmountCards)
			player := gs.Players[user]
			player.Damage += effect.AmountDamage
			player.Money += effect.AmountMoney
			gs.Players[user] = player
		}
	}
}

type ActivePlayerBanishes struct {
	Hand     bool
	Discard  bool
	Deck     bool
	PlayArea bool
}

func (effect ActivePlayerBanishes) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	player := gs.Players[user]

	cards := []Card{}
	if effect.Hand {
		cards = append(cards, player.Hand...)
	}
	if effect.Discard {
		cards = append(cards, player.Discard...)
	}
	if effect.Deck {
		cards = append(cards, player.Deck...)
	}
	if effect.PlayArea {
		cards = append(cards, player.PlayArea...)
	}

	if len(cards) == 0 {
		return
	}

	SendLobbyUpdate(gs.gameid, gs)
	choice := AskUserToSelectCard(user, gs.gameid, cards, "Banish a card:")
	BanishCard(user, choice, gs)
}

// FIX -- May banish a card
type ActivePlayerBanishAndGainXIfY struct {
	Hand     bool
	Discard  bool
	Deck     bool
	PlayArea bool

	CardType string // "item", "spell", "ally"
	GainX    Effect
}

func (effect ActivePlayerBanishAndGainXIfY) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	player := gs.Players[user]

	cards := []Card{}
	if effect.Hand {
		cards = append(cards, player.Hand...)
	}
	if effect.Discard {
		cards = append(cards, player.Discard...)
	}
	if effect.Deck {
		cards = append(cards, player.Deck...)
	}
	if effect.PlayArea {
		cards = append(cards, player.PlayArea...)
	}

	if len(cards) == 0 {
		return
	}

	SendLobbyUpdate(gs.gameid, gs)
	choice := AskUserToSelectCard(user, gs.gameid, cards, "Banish a card:")
	choiceType := BanishCard(user, choice, gs)

	if choiceType == effect.CardType {
		effect.GainX.Trigger(gs)
	}
}

type GainXPerSpellPlayed struct {
	X ChangeStats
}

func (effect GainXPerSpellPlayed) Trigger(gs *Gamestate) {
	effect.X.Target = gs.CurrentTurn
	numSpells := gs.turnStats.SpellsPlayed
	for i := 0; i < numSpells; i++ {
		effect.X.Trigger(gs)
	}
}

type ChoosePlayerToGainX struct {
	X ChangeStats
}

func (effect ChoosePlayerToGainX) Trigger(gs *Gamestate) {
	currentTurn := gs.CurrentTurn
	players := []string{}
	for name := range gs.Players {
		players = append(players, name)
	}

	if len(players) == 0 {
		return
	}
	SendLobbyUpdate(gs.gameid, gs)
	effect.X.Target = AskUserToSelectPlayer(gs.gameid, currentTurn, players)
	effect.X.Trigger(gs)
}

type ActivePlayerSearchesDiscardForX struct {
	CardType string // "any" if it doesn't matter what type.
}

func (effect ActivePlayerSearchesDiscardForX) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	choices := []Card{}
	for _, c := range gs.Players[user].Discard {
		if c.CardType == effect.CardType || effect.CardType == "any" {
			choices = append(choices, c)
		}
	}

	if len(choices) != 0 {
		prompt := "Choose a card from your discard to gain to your hand!"
		SendLobbyUpdate(gs.gameid, gs)
		cardId := AskUserToSelectCard(user, gs.gameid, choices, prompt)
		MoveCardFromDiscardToHand(user, cardId, gs)
	}
}

type ActivePlayerSearchesDeckForX struct {
	CardType      string // "any" if it doesn't matter what type.
	Target        string
	CostRestraint int
}

func (effect ActivePlayerSearchesDeckForX) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	if effect.Target != "" {
		user = effect.Target
	}

	choices := []Card{}
	for _, c := range gs.Players[user].Deck {
		if (c.CardType == effect.CardType || effect.CardType == "any") && (c.Cost <= effect.CostRestraint || effect.CostRestraint == 0) {
			choices = append(choices, c)
		}
	}

	if len(choices) != 0 {
		prompt := "Choose a card from your deck to gain to your hand"
		SendLobbyUpdate(gs.gameid, gs)
		cardId := AskUserToSelectCard(user, gs.gameid, choices, prompt)
		MoveCardFromDeckToHand(user, cardId, gs)
	}
}

type GainXIfYSpellsPlayed struct {
	Id int
	X  ChangeStats
	Y  int
}

func (effect GainXIfYSpellsPlayed) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	effect.X.Target = user
	numSpellsPlayed := gs.turnStats.SpellsPlayed

	if numSpellsPlayed >= effect.Y {
		effect.X.Trigger(gs)
	}

	sub := Subscriber{
		id:              effect.Id,
		messageChan:     make(chan string),
		conditionMet:    "spell played",
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

			gs.mu.Lock()
			if user == gs.CurrentTurn {
				numSpellsPlayed++
				if numSpellsPlayed >= effect.Y {
					effect.X.Trigger(gs)
					SendLobbyUpdate(gs.gameid, gs)
					gs.mu.Unlock()
					return
				}
			}
			gs.mu.Unlock()
		}
	}()
}

type PurchasedXGoToDeck struct {
	X string
}

func (effect PurchasedXGoToDeck) Trigger(gs *Gamestate) {
	player := gs.Players[gs.CurrentTurn]
	switch effect.X {
	case "item":
		player.itemsToDeck = true
	case "spell":
		player.spellsToDeck = true
	case "ally":
		player.alliesToDeck = true
	}

	gs.Players[gs.CurrentTurn] = player
}

type GainCardToTopDeck struct {
	user string
	card Card
}

func (effect GainCardToTopDeck) Trigger(gs *Gamestate) {
	player := gs.Players[effect.user]
	player.Deck = append(player.Deck, effect.card)
	gs.Players[effect.user] = player
}

type GainCardToDiscard struct {
	user string
	card Card
}

func (effect GainCardToDiscard) Trigger(gs *Gamestate) {
	player := gs.Players[effect.user]
	player.Discard = append(player.Discard, effect.card)
	gs.Players[effect.user] = player
}

type BlockVillainEffects struct {
	villain  bool
	creature bool
}

func (effect BlockVillainEffects) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	turnNumber := gs.turnNumber
	numTurns := len(gs.TurnOrder)
	unblockedAt := turnNumber + numTurns

	choices := []string{}
	for _, v := range gs.Villains {
		if v.Active && ((effect.villain && v.villainType == "villain") || (effect.creature && v.villainType == "creature") || (v.villainType == "villain-creature")) && (turnNumber >= v.BlockedUntil) {
			choices = append(choices, v.Name)
		}
	}

	if len(choices) > 0 {
		SendLobbyUpdate(gs.gameid, gs)
		choice := AskUserToSelectPlayer(gs.gameid, user, choices)
		for i, v := range gs.Villains {
			if v.Name == choice {
				gs.Villains[i].BlockedUntil = unblockedAt
			}
		}
	}
}

type PreviousHeroDoesX struct {
	X ChangeStats
}

func (effect PreviousHeroDoesX) Trigger(gs *Gamestate) {
	prevHero := getPreviousUser(gs)
	effect.X.Target = prevHero
	effect.X.Trigger(gs)
}

type NextHeroDoesX struct {
	X ChangeStats
}

func (effect NextHeroDoesX) Trigger(gs *Gamestate) {
	nextHero := getNextUser(gs)
	effect.X.Target = nextHero
	effect.X.Trigger(gs)
}

type ActivePlayerSelectsOtherPlayerToDoX struct {
	X ChangeStats
}

func (effect ActivePlayerSelectsOtherPlayerToDoX) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	otherPlayers := []string{}
	for name := range gs.Players {
		if name != user {
			otherPlayers = append(otherPlayers, name)
		}
	}

	if len(otherPlayers) == 0 {
		return
	}

	SendLobbyUpdate(gs.gameid, gs)
	effect.X.Target = AskUserToSelectPlayer(gs.gameid, user, otherPlayers)
	effect.X.Trigger(gs)
}

type AllRevealTopCardAndX struct {
	X func(card Card, user string, gs *Gamestate)
}

func (effect AllRevealTopCardAndX) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		player := gs.Players[user]

		if len(player.Deck) == 0 {
			ShuffleDiscardToDeck(&player)
			if len(player.Deck) == 0 {
				log.Println("deck is completely empty mate")
				return
			}
		}
		gs.Players[user] = player

		topCard := player.Deck[len(player.Deck)-1]
		effect.X(topCard, user, gs)
	}
}

type DiscardACard struct {
	Target   string
	Prompt   string
	Cardtype string // "any" for any card
	Cause    string
}

func (effect DiscardACard) Trigger(gs *Gamestate) {
	log.Println("starting DiscardACard vvvvvvv")
	assertUniqueCards(gs)
	user := effect.Target
	choices := []Card{}
	for _, c := range gs.Players[user].Hand {
		if c.CardType == effect.Cardtype || effect.Cardtype == "any" {
			choices = append(choices, c)
		}
	}

	if len(choices) == 0 {
		return
	}

	prompt := effect.Prompt
	if prompt == "" {
		prompt = "Discard a card"
	}

	SendLobbyUpdate(gs.gameid, gs)
	choice := AskUserToSelectCard(user, gs.gameid, choices, prompt)
	DiscardFromId(user, choice, gs)
	cond := effect.Cause == "villain" || effect.Cause == "darkart"
	if gs.Players[user].Proficiency == "Defense Against the Dark Arts" && cond {
		ChangeStats{Target: user, AmountDamage: 1, AmountHealth: 1}.Trigger(gs)
	}

	assertUniqueCards(gs)
	log.Println("ending DiscardACard ^^^^^ ONE ERROR MEANS BUG")
}

func TargetDiscardACard(target string, effect Effect) Effect {
	ret, ok := effect.(DiscardACard)
	if !ok {
		log.Println("Type Assertion failed: TargetDiscardACard")
	}
	ret.Target = target
	return ret
}

type DamageAllPerCreature struct {
	Amount int
}

func (effect DamageAllPerCreature) Trigger(gs *Gamestate) {
	damage := 0
	for _, v := range gs.Villains {
		if v.Active && (v.villainType == "creature" || v.villainType == "villain-creature") {
			damage += effect.Amount
		}
	}

	DamageAllPlayers{Amount: damage}.Trigger(gs)
}

type Scry struct {
	User string
}

func (effect Scry) Trigger(gs *Gamestate) {
	player := gs.Players[effect.User]
	if len(player.Deck) == 0 {
		ShuffleDiscardToDeck(&player)
		gs.Players[effect.User] = player
	}

	topCard := player.Deck[len(player.Deck)-1]
	path := topCard.ImgPath
	SendLobbyUpdate(gs.gameid, gs)
	choice := AskUserInputWithCard(gs.gameid, effect.User, path, "Discard or place card back on deck:", []string{"Discard", "Place on deck"})

	if choice == "Discard" {
		DiscardFromId(effect.User, topCard.Id, gs)
	}
}

type ScryDarkarts struct {
	User string
}

func (effect ScryDarkarts) Trigger(gs *Gamestate) {
	effect.User = gs.CurrentTurn
	path := gs.DarkArts[(gs.CurrentDarkArt+1)%len(gs.DarkArts)].ImgPath
	SendLobbyUpdate(gs.gameid, gs)
	choice := AskUserInputWithCard(gs.gameid, effect.User, path, "You look at the top of the Dark Arts deck:", []string{"Discard", "Place on deck"})

	if choice == "Discard" {
		LoadNewDarkArt(gs)
	}
}
