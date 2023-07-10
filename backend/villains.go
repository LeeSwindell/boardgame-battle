package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func draco() Villain {
	id := int(uuid.New().ID())
	return Villain{
		Name:         "Draco Malfoy",
		Id:           id,
		ImgPath:      "/images/villains/dracomalfoy.jpg",
		SetId:        "Game 1",
		CurDamage:    0,
		MaxHp:        6,
		Effect:       []Effect{DamageActiveIfLocationAdded{Amount: 2, Id: id}},
		DeathEffect:  []Effect{RemoveFromLocation{Amount: 1}},
		playBeforeDA: true,
	}
}

type DamageActiveIfLocationAdded struct {
	Amount int
	Id     int
}

func (effect DamageActiveIfLocationAdded) Trigger(gs *Gamestate) {
	log.Println("calling damageactiveiflocationadded")
	// find player who was active when location got added.
	currentTurn := gs.CurrentTurn

	sub := Subscriber{
		id:              effect.Id,
		messageChan:     make(chan string),
		conditionMet:    "location added",
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

			Logger("draco wants lock")
			gs.mu.Lock()
			Logger("draco gets lock")
			if currentTurn == gs.CurrentTurn {
				user := gs.CurrentTurn
				stunned := ChangePlayerHealth(user, -effect.Amount, gs)
				if stunned {
					StunPlayer(user, gs)
				}

				SendLobbyUpdate(gs.gameid, gs)
			}
			gs.mu.Unlock()
			Logger("draco releases lock")
		}
	}()
}

func quirrell() Villain {
	return Villain{
		Name:      "Quirinus Quirrell",
		Id:        int(uuid.New().ID()),
		ImgPath:   "/images/villains/quirrell.jpg",
		SetId:     "Game 1",
		CurDamage: 0,
		MaxHp:     6,
		Effect:    []Effect{DamageCurrentPlayer{Amount: 1}},
		DeathEffect: []Effect{
			AllPlayersGainMoney{Amount: 1},
			AllPlayersGainHealth{Amount: 1},
		},
		playBeforeDA: false,
	}
}

func crabbeAndGoyle() Villain {
	id := int(uuid.New().ID())
	return Villain{
		Name:      "Crabbe and Goyle",
		Id:        id,
		ImgPath:   "/images/villains/crabbeandgoyle.jpg",
		SetId:     "Game 1",
		CurDamage: 0,
		MaxHp:     5,
		Effect: []Effect{
			DamageIfDiscard{Amount: 1, Id: id},
		},
		DeathEffect:  []Effect{AllDrawCards{Amount: 1}},
		playBeforeDA: true,
	}
}

type DamageIfDiscard struct {
	Amount int
	Id     int
}

// only damages active player. not player who discarded
func (effect DamageIfDiscard) Trigger(gs *Gamestate) {
	// find player who was active when location got added.
	currentTurn := gs.CurrentTurn

	sub := Subscriber{
		id:              effect.Id,
		messageChan:     make(chan string),
		conditionMet:    "player discarded",
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

			Logger("c&g want lock")
			gs.mu.Lock()
			Logger("c&g gets lock")
			if res && currentTurn == gs.CurrentTurn {
				user := gs.CurrentTurn
				// log.Println("c&g deal 1 dmg")
				Logger("cg calling stunned")
				stunned := ChangePlayerHealth(user, -effect.Amount, gs)
				if stunned {
					StunPlayer(user, gs)
				}
				Logger("cg sending lobby update ")

				SendLobbyUpdate(gs.gameid, gs)
			}
			gs.mu.Unlock()
			Logger("c&g release lock")
		}
	}()
}

func bartyCrouchJr() Villain {
	return Villain{
		Name:      "Barty Crouch Jr.",
		Id:        int(uuid.New().ID()),
		ImgPath:   "/images/villains/bartycrouchjr.jpg",
		SetId:     "Game 4",
		CurDamage: 0,
		MaxHp:     7,
		// Set Barty Crouch to Active=false before triggering death effect.
		Active: false,
		// This hero just prevents location from being removed, build into
		// remove from location handler.
		Effect:       []Effect{},
		DeathEffect:  []Effect{RemoveFromLocation{2}},
		playBeforeDA: true,
	}
}

func basilisk() Villain {
	return Villain{
		Name:      "Basilisk",
		Id:        int(uuid.New().ID()),
		ImgPath:   "/images/villains/basilisk.jpg",
		SetId:     "Game 2",
		CurDamage: 0,
		MaxHp:     8,
		Active:    false,
		// This hero just prevents players from drawing, build into
		// remove from draw handler.
		Effect: []Effect{},
		DeathEffect: []Effect{
			AllDrawCards{Amount: 1},
			RemoveFromLocation{1},
		},
		playBeforeDA: true,
	}
}

func bellatrixLestrange() Villain {
	return Villain{
		Name:      "Bellatrix Lestrange",
		Id:        int(uuid.New().ID()),
		ImgPath:   "/images/villains/bellatrixlestrange.jpg",
		SetId:     "Game 6",
		CurDamage: 0,
		MaxHp:     9,
		Active:    false,
		Effect: []Effect{
			RevealDarkArts{Amount: 1},
		},
		DeathEffect: []Effect{
			AllSearchDiscardPileForItem{},
			RemoveFromLocation{Amount: 2},
		},
		playBeforeDA: false,
	}
}

func cornishPixies() Villain {
	return Villain{
		Name:      "Cornish Pixies",
		Id:        int(uuid.New().ID()),
		ImgPath:   "/images/villains/cornishpixies.jpg",
		SetId:     "Box 1",
		CurDamage: 0,
		MaxHp:     6,
		Active:    false,
		Effect: []Effect{
			DamageActiveForEachEvenInHand{Amount: 2},
		},
		DeathEffect: []Effect{
			AllPlayersGainHealth{Amount: 2},
			AllPlayersGainMoney{Amount: 1},
		},
		playBeforeDA: false,
	}
}

type DamageActiveForEachEvenInHand struct {
	Amount int
}

func (effect DamageActiveForEachEvenInHand) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn

	numEvens := 0
	for _, c := range gs.Players[user].Hand {
		if c.Cost%2 == 0 && c.Cost != 0 {
			numEvens++
		}
	}

	damage := -1 * numEvens * effect.Amount
	stunned := ChangePlayerHealth(user, damage, gs)
	if stunned {
		StunPlayer(user, gs)
	}
}

func dementor() Villain {
	return Villain{
		Name:      "Dementor",
		Id:        int(uuid.New().ID()),
		ImgPath:   "/images/villains/dementor.jpg",
		SetId:     "Game 3",
		CurDamage: 0,
		MaxHp:     8,
		Active:    false,
		Effect: []Effect{
			DamageCurrentPlayer{Amount: 2},
		},
		DeathEffect: []Effect{
			AllPlayersGainHealth{Amount: 2},
			RemoveFromLocation{Amount: 1},
		},
		playBeforeDA: false,
	}
}

func fenrirGreyback() Villain {
	return Villain{
		Name:      "Fenrir Greyback",
		Id:        int(uuid.New().ID()),
		ImgPath:   "/images/villains/fenrirgreyback.jpg",
		SetId:     "Game 6",
		CurDamage: 0,
		MaxHp:     8,
		Active:    false,
		// Makes players unable to gain health, add a check in changeHealth
		Effect: []Effect{},
		DeathEffect: []Effect{
			AllPlayersGainHealth{Amount: 3},
			RemoveFromLocation{Amount: 2},
		},
		playBeforeDA: true,
	}
}

func doloresUmbridge() Villain {
	id := int(uuid.New().ID())
	return Villain{
		Name:      "Dolores Umbridge",
		Id:        id,
		ImgPath:   "/images/villains/doloresumbridge.jpg",
		SetId:     "Game 5",
		CurDamage: 0,
		MaxHp:     7,
		Active:    false,
		Effect:    []Effect{DoloresEffect{id}},
		DeathEffect: []Effect{
			AllPlayersGainMoney{Amount: 1},
			AllPlayersGainHealth{Amount: 2},
		},
		playBeforeDA: true,
	}
}

type DoloresEffect struct {
	id int
}

func (effect DoloresEffect) Trigger(gs *Gamestate) {
	// find player who was active when location got added.
	currentTurn := gs.CurrentTurn

	sub := Subscriber{
		id:              effect.id,
		messageChan:     make(chan string),
		conditionMet:    "umbridge condition",
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
			if res && currentTurn == gs.CurrentTurn {
				user := gs.CurrentTurn
				stunned := ChangePlayerHealth(user, -1, gs)
				if stunned {
					StunPlayer(user, gs)
				}

				SendLobbyUpdate(gs.gameid, gs)
			}
			gs.mu.Unlock()
		}
	}()
}

func fluffy() Villain {
	id := int(uuid.New().ID())
	return Villain{
		Name:      "Fluffy",
		Id:        id,
		ImgPath:   "/images/villains/fluffy.jpg",
		SetId:     "Box 1",
		CurDamage: 0,
		MaxHp:     8,
		Active:    false,
		Effect:    []Effect{FluffyEffect{}},
		DeathEffect: []Effect{
			AllPlayersGainHealth{Amount: 1},
			AllDrawCards{Amount: 1},
		},
		playBeforeDA: false,
	}
}

type FluffyEffect struct{}

// For each item, choose one: lose a life or discard.
func (effect FluffyEffect) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	numItems := 0
	for _, c := range gs.Players[user].Hand {
		if c.CardType == "item" {
			numItems++
		}
	}

	for i := 0; i < numItems; i++ {
		desc := fmt.Sprintf("Fluffy!!! Choose one: (%d of %d)", i+1, numItems)
		ChooseOne{
			Effects: []Effect{
				DamageCurrentPlayer{Amount: 1},
				ActivePlayerDiscards{Amount: 1},
			},
			Options:     []string{"Lose a life", "Discard a card"},
			Description: desc,
		}.Trigger(gs)
	}
}

// FIX!!! NOT WORKING!
func luciusMalfoy() Villain {
	id := int(uuid.New().ID())
	return Villain{
		Name:      "Lucius Malfoy",
		Id:        id,
		ImgPath:   "/images/villains/luciusmalfoy.jpg",
		SetId:     "Game 2",
		CurDamage: 0,
		MaxHp:     7,
		Active:    false,
		Effect:    []Effect{LuciusEffect{id: id}},
		DeathEffect: []Effect{
			AllPlayersGainMoney{Amount: 1},
			RemoveFromLocation{Amount: 1},
		},
		playBeforeDA: true,
	}
}

type LuciusEffect struct {
	id int
}

// CHECK IF THIS GETS TRIGGERED WHEN HE DIES -  CHECK IF VILLAIN STILL ACTIVE WHEN TRIGGERING.
func (effect LuciusEffect) Trigger(gs *Gamestate) {
	sub := Subscriber{
		id:              effect.id,
		messageChan:     make(chan string),
		conditionMet:    "location removed",
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
			if res {
				HealAllVillains{Amount: 1}.Trigger(gs)
				SendLobbyUpdate(gs.gameid, gs)
			}
			gs.mu.Unlock()
		}
	}()
}

func tomRiddle() Villain {
	id := int(uuid.New().ID())
	return Villain{
		Name:         "Tom Riddle",
		Id:           id,
		ImgPath:      "/images/villains/tomriddle.jpg",
		SetId:        "Game 2",
		CurDamage:    0,
		MaxHp:        6,
		Active:       false,
		Effect:       []Effect{TomRiddleEffect{}},
		DeathEffect:  []Effect{TomRiddleDeathEffect{}},
		playBeforeDA: false,
	}
}

type TomRiddleEffect struct{}

// For each ally, choose one: lose 2 life or discard.
func (effect TomRiddleEffect) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	numItems := 0
	for _, c := range gs.Players[user].Hand {
		if c.CardType == "ally" {
			numItems++
		}
	}

	for i := 0; i < numItems; i++ {
		desc := fmt.Sprintf("Tom Riddle! Choose one: (%d of %d)", i+1, numItems)
		ChooseOne{
			Effects: []Effect{
				DamageCurrentPlayer{Amount: 2},
				ActivePlayerDiscards{Amount: 1},
			},
			Options:     []string{"Lose 2 life", "Discard a card"},
			Description: desc,
		}.Trigger(gs)
	}
}

// player that kills tom gets to choose it all
type TomRiddleDeathEffect struct{}

func (effect TomRiddleDeathEffect) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		choice := ChooseOne{
			Effects: []Effect{
				ChosenPlayerGainsHealth{Amount: 2, Playername: user},
				ChosenPlayerSearchesDiscardForX{Playername: user, SearchType: "ally"},
			},
			Options:     []string{"Gain 2 health", "Search Discard Pile for Ally"},
			Description: "You killed Tom Riddle! Choose one:",
		}

		choice.Trigger(gs)
	}
}

func peterPettigrew() Villain {
	id := int(uuid.New().ID())
	return Villain{
		Name:      "Peter Pettigrew",
		Id:        id,
		ImgPath:   "/images/villains/peterpettigrew.jpg",
		SetId:     "Game 3",
		CurDamage: 0,
		MaxHp:     7,
		Active:    false,
		Effect:    []Effect{PeterPettigrewEffect{}},
		DeathEffect: []Effect{
			PeterPettigrewDeathEffect{},
			RemoveFromLocation{Amount: 1},
		},
		playBeforeDA: false,
	}
}

type PeterPettigrewEffect struct{}

// Don't "Reveal" the card for now, just discard it if it matches.
func (effect PeterPettigrewEffect) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	player := gs.Players[user]

	if len(player.Deck) == 0 {
		ShuffleDiscardToDeck(&player)
		if len(player.Deck) == 0 {
			log.Println("deck is completely empty mate")
			return
		}
	}

	topCard := player.Deck[len(player.Deck)-1]
	// If cost > 1, discard it and add 1 to location.
	if topCard.Cost >= 1 {
		Logger("Triggering pettigrews effect")
		player.Discard = append(player.Discard, topCard)
		player.Deck = player.Deck[:len(player.Deck)-1]
		AddToLocation{Amount: 1}.Trigger(gs)
	}

	gs.Players[user] = player
}

type PeterPettigrewDeathEffect struct{}

func (effect PeterPettigrewDeathEffect) Trigger(gs *Gamestate) {
	for user := range gs.Players {
		ChosenPlayerSearchesDiscardForX{Playername: user, SearchType: "spell"}.Trigger(gs)
	}
}

func voledmortFive() Villain {
	id := int(uuid.New().ID())
	return Villain{
		Name:      "Voldemort",
		Id:        id,
		ImgPath:   "/images/villains/voldemort5.jpg",
		SetId:     "Game 5",
		CurDamage: 0,
		MaxHp:     10,
		Active:    true,
		Effect: []Effect{
			DamageCurrentPlayer{Amount: 1},
			ActivePlayerDiscards{Amount: 1, Prompt: "Voldemort attacks! Discard a card"},
		},
		DeathEffect:  []Effect{},
		playBeforeDA: false,
	}
}

func troll() Villain {
	id := int(uuid.New().ID())
	return Villain{
		Name:      "Troll",
		Id:        id,
		ImgPath:   "/images/villains/troll.jpg",
		SetId:     "Box 1",
		CurDamage: 0,
		MaxHp:     7,
		Active:    false,
		Effect: []Effect{
			ChooseOne{
				Effects: []Effect{
					DamageCurrentPlayer{Amount: 2},
					GainDetentionToDiscard{Active: true},
				},
				Options: []string{
					"Lose 2 Life",
					"Gain a detention",
				},
				Description: "Troll attack!",
			},
		},
		DeathEffect: []Effect{
			AllPlayersGainHealth{Amount: 1},
			AllBanishItem{},
		},
		playBeforeDA: false,
	}
}

func norbert() Villain {
	id := int(uuid.New().ID())
	return Villain{
		Name:      "Norbert",
		Id:        id,
		ImgPath:   "/images/villains/norbert.jpg",
		SetId:     "Box 1",
		CurDamage: 0,
		MaxHp:     6,
		Active:    false,
		Effect: []Effect{
			DamageCurrentPlayer{Amount: 1},
			DamageActivePerDetention{Amount: 1},
		},
		DeathEffect: []Effect{
			AllBanishCard{},
		},
		playBeforeDA: false,
	}
}
