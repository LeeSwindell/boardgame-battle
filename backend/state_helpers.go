package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Removes a card from an array at given index, returns a new array of cards
func RemoveCardAtIndex(s []Card, index int) []Card {
	ret := make([]Card, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// Removes a card and moves it to the PlayArea given a cardid.
func MoveToPlayed(user string, id int, gs *Gamestate) {
	for index, c := range gs.Players[user].Hand {
		if c.Id == id {
			card := gs.Players[user].Hand[index]
			newHand := RemoveCardAtIndex(gs.Players[user].Hand, index)
			updatedPlayer := gs.Players[user]

			updatedPlayer.Hand = newHand
			updatedPlayer.PlayArea = append(updatedPlayer.PlayArea, card)

			gs.Players[user] = updatedPlayer
			return
		}
	}

	log.Println("error with MoveToPlayed")
}

// Moves a users cards from the PlayArea into the Discard
func MovePlayedToDiscard(user string, gs *Gamestate) {
	updatedPlayer := gs.Players[user]
	updatedPlayer.Discard = append(updatedPlayer.Discard, updatedPlayer.PlayArea...)
	updatedPlayer.PlayArea = []Card{}
	gs.Players[user] = updatedPlayer
}

// Moves a users cards from the Hand into the Discard
func MoveHandToDiscard(user string, gs *Gamestate) {
	updatedPlayer := gs.Players[user]
	updatedPlayer.Discard = append(updatedPlayer.Discard, updatedPlayer.Hand...)
	updatedPlayer.Hand = []Card{}
	gs.Players[user] = updatedPlayer
}

// Moves the card with cardId from the users discard to their hand.
func MoveCardFromDiscardToHand(user string, cardId int, gs *Gamestate) {
	player := gs.Players[user]
	for i, c := range player.Discard {
		if c.Id == cardId {
			player.Discard = RemoveCardAtIndex(player.Discard, i)
			player.Hand = append(player.Hand, c)
			gs.Players[user] = player
		}
	}
}

// Moves the card with cardId from the users discard to their hand.
func MoveCardFromDeckToHand(user string, cardId int, gs *Gamestate) {
	player := gs.Players[user]
	for i, c := range player.Deck {
		if c.Id == cardId {
			player.Deck = RemoveCardAtIndex(player.Deck, i)
			player.Hand = append(player.Hand, c)
			gs.Players[user] = player
		}
	}
}

// updates values of a player - shuffling discard pile into their deck
func ShuffleDiscardToDeck(player *Player) {
	addToDeck := ShuffleCards(player.Discard)
	player.Deck = append(player.Deck, addToDeck...)
	player.Discard = []Card{}
}

// Shuffles a slice of cards, and returns the new ordering.
func ShuffleCards(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())

	// Create a new slice to store the shuffled cards
	shuffledCards := make([]Card, len(cards))
	copy(shuffledCards, cards)

	// Shuffle the cards
	for i := range shuffledCards {
		j := rand.Intn(i + 1)
		shuffledCards[i], shuffledCards[j] = shuffledCards[j], shuffledCards[i]
	}

	return shuffledCards
}

func ShuffleDarkArts(da []DarkArt) []DarkArt {
	rand.Seed(time.Now().UnixNano())

	// Create a new slice to store the shuffled cards
	shuffledDA := make([]DarkArt, len(da))
	copy(shuffledDA, da)

	// Shuffle the cards
	for i := range shuffledDA {
		j := rand.Intn(i + 1)
		shuffledDA[i], shuffledDA[j] = shuffledDA[j], shuffledDA[i]
	}

	return shuffledDA
}

// Shuffles a slice of cards, and returns the new ordering.
func ShuffleVillains(villains []Villain) []Villain {
	rand.Seed(time.Now().UnixNano())

	// Create a new slice to store the shuffled cards
	shuffledVillains := make([]Villain, len(villains))
	copy(shuffledVillains, villains)

	// Shuffle the cards
	for i := range shuffledVillains {
		j := rand.Intn(i + 1)
		shuffledVillains[i], shuffledVillains[j] = shuffledVillains[j], shuffledVillains[i]
	}

	return shuffledVillains
}

// Removes the top card of a players deck, returns the card.
func PopFromDeck(player *Player) (Card, bool) {
	if len(player.Deck) == 0 {
		ShuffleDiscardToDeck(player)
		if len(player.Deck) == 0 {
			log.Println("deck is completely empty mate")
			return Card{}, false
		}
	}

	topCard := player.Deck[len(player.Deck)-1]
	player.Deck = player.Deck[:len(player.Deck)-1]
	return topCard, true
}

// Used to draw cards and shuffle deck if needed. Draws from end of slice.
func DrawXCards(user string, gs *Gamestate, amount int) {
	// Basilisk prevents players from drawing extra cards.
	for _, v := range gs.Villains {
		if v.Name == "Basilisk" && v.Active {
			return
		}
	}

	// do nothing if petrification has been played.
	for _, da := range gs.DarkArtsPlayed {
		if da.Name == "Petrification" {
			return
		}
	}

	if amount <= 0 {
		return
	}

	updated := gs.Players[user]
	for i := 0; i < amount; i++ {
		card, ok := PopFromDeck(&updated)
		if ok {
			updated.Hand = append(updated.Hand, card)
		}
	}

	gs.Players[user] = updated
	gs.turnStats.CardsDrawn++
	if gs.Players[user].Character == "Luna" && gs.turnStats.CardsDrawn == 1 {
		HealAnyPlayer{Amount: 2}.Trigger(gs)
	}
}

// Used to draw 5 cards at end of turn, shuffle deck if needed.
func RefillHand(user string, gs *Gamestate) {
	updated := gs.Players[user]
	for i := 0; i < 5; i++ {
		card, ok := PopFromDeck(&updated)
		if ok {
			updated.Hand = append(updated.Hand, card)
		}
	}

	gs.Players[user] = updated
}

// Optionally check the return for cardtype of banished card for some effects.
func BanishCard(user string, cardId int, gs *Gamestate) string {
	// Find card in hand/discard/played/deck
	player := gs.Players[user]
	cardtype := ""
	for i, c := range player.Hand {
		if c.Id == cardId {
			player.Hand = RemoveCardAtIndex(player.Hand, i)
			cardtype = c.CardType
		}
	}
	for i, c := range player.Discard {
		if c.Id == cardId {
			player.Discard = RemoveCardAtIndex(player.Discard, i)
			cardtype = c.CardType
		}
	}
	for i, c := range player.PlayArea {
		if c.Id == cardId {
			player.PlayArea = RemoveCardAtIndex(player.PlayArea, i)
			cardtype = c.CardType
		}
	}
	for i, c := range player.Deck {
		if c.Id == cardId {
			player.Deck = RemoveCardAtIndex(player.Deck, i)
			cardtype = c.CardType
		}
	}

	gs.Players[user] = player
	return cardtype
}

func MoneyDamageToZero(user string, gs *Gamestate) {
	updated := gs.Players[user]
	updated.Damage = 0
	updated.Money = 0
	gs.Players[user] = updated
}

func NextTurnInOrder(gs *Gamestate) {
	for i, name := range gs.TurnOrder {
		if name == gs.CurrentTurn {
			gs.CurrentTurn = gs.TurnOrder[(i+1)%len(gs.TurnOrder)]
			break
		}
	}
}

func StartNewTurn(gameid int, gs *Gamestate) {
	// Starting next turn actions.
	Logger("stuck at 1")
	for _, v := range gs.Villains {
		if v.playBeforeDA {
			for _, e := range v.effect {
				e.Trigger(gs)
			}
		}
	}
	Logger("stuck at 2")
	gs.Locations[gs.CurrentLocation].effect.Trigger(gs)
	Logger("stuck at 3")
	for _, v := range gs.Villains {
		if !v.playBeforeDA {
			Logger(v.Name)
			for _, e := range v.effect {
				e.Trigger(gs)
			}
		}
	}
	Logger("stuck at 4")

	gs.started = true
	SendLobbyUpdate(gameid, gs)
}

// return true if player stunned, limit health between 0 and 10.
func ChangePlayerHealth(user string, change int, gs *Gamestate) bool {
	player := gs.Players[user]

	// Do nothing if player already has no health.
	if player.Health <= 0 {
		return false
	}

	// Check if Fenrir is active
	for _, v := range gs.Villains {
		if v.Name == "Fenrir Greyback" && v.Active && change >= 0 {
			// Do nothing, they won't be stunned so return false.
			return false
		}
	}

	// Check for Sectumsempra
	for _, da := range gs.DarkArtsPlayed {
		if da.Name == "Sectumsempra!" && change >= 0 {
			return false
		}
	}

	startingHealed, alreadyHealed := gs.turnStats.AlliesHealed[user]
	if change > 0 {
		gs.turnStats.AlliesHealed[user] += change
	}

	// Check for Neville
	// change already healed to map[user]int
	if change > 0 && gs.Players[gs.CurrentTurn].Character == "Neville" && !alreadyHealed {
		log.Println("give a choice then!")
		GivenPlayerChooseOneTargeted{
			User: user,
			EffectTargeting: []func(target string, effect Effect) Effect{
				TargetCreateStats,
				TargetCreateStats,
			},
			Effects: []Effect{
				ChangeStats{AmountHealth: 1},
				ChangeStats{AmountMoney: 1},
			},
			Options:     []string{"Gain 1 Health", "Gain 1 Money"},
			Description: "Neville saves the day! Choose one:",
		}.Trigger(gs)
	}

	cond := gs.turnStats.AlliesHealed[user] >= 3 && startingHealed < 3
	if gs.Players[gs.CurrentTurn].Proficiency == "Herbology" && cond {
		DrawXCards(user, gs, 1)
		p := gs.Players[gs.CurrentTurn]
		gs.Players[gs.CurrentTurn] = p
	}

	// Check for invisibility cloak
	if change < 0 {
		for _, c := range player.Hand {
			if c.Name == "Invisibility Cloak" {
				change = -1
			}
		}
	}

	player = gs.Players[user]
	player.Health += change

	if player.Health <= 0 {
		player.Health = 0
		gs.Players[user] = player
		return true
	} else if player.Health > 10 {
		player.Health = 10
	}

	gs.Players[user] = player
	return false
}

// change it so that players at 0 life go to 10 at start of turn.
func StunPlayer(user string, gs *Gamestate) {
	Logger("BEGIN STUN: CHECK FOR UNIQUE CARDS")
	assertUniqueCards(gs)
	Logger("good at start? ^^^")
	player := gs.Players[user]
	player.Money = 0
	player.Damage = 0
	gs.Players[user] = player
	discardAmount := len(player.Hand) / 2
	for i := 0; i < discardAmount; i++ {
		desc := fmt.Sprintf("Stunned! Discard a card: %d of %d", i+1, discardAmount)
		SendLobbyUpdate(gs.gameid, gs)
		cardName := AskUserToSelectCard(user, gs.gameid, player.Hand, desc)

		testids := []int{}
		for _, c := range player.Hand {
			testids = append(testids, c.Id)
		}
		log.Println("TURN:", gs.turnNumber, "HAND IDS:", testids)

		for i, c := range player.Hand {
			if c.Id == cardName {
				player.Hand = RemoveCardAtIndex(player.Hand, i)
				player.Discard = append(player.Discard, c)

				// Wrap the player mapping around onDiscard since it mutates the state directly.
				if c.onDiscard != nil {
					gs.Players[user] = player
					c.onDiscard(user, gs)
					player = gs.Players[user]
				}

				break
			}
		}

		gs.Players[user] = player

		event := Event{senderId: -1, message: "player discarded", data: user}
		Logger("stunned, sending event")
		eventBroker.Messages <- event
		Logger("stunned, sent event")
	}

	AddToLocation{Amount: 1}.Trigger(gs)
	log.Println("good at end? vvvv")
	assertUniqueCards(gs)
	Logger("ending stun")
}

func HealStunned(gs *Gamestate) {
	for user := range gs.Players {
		if gs.Players[user].Health <= 0 {
			player := gs.Players[user]
			player.Health = 10
			gs.Players[user] = player
		}
	}
}

// Removes a card from an array at given index, returns a new array of villains
func RemoveVillainAtIndex(vs []Villain, index int) []Villain {
	ret := make([]Villain, 0)
	ret = append(ret, vs[:index]...)
	return append(ret, vs[index+1:]...)
}

func AddNewActiveVillain(villains []Villain, gs *Gamestate) []Villain {
	// Get new villain
	if len(gs.villainDeck) == 0 {
		return villains
	} else if len(gs.villainDeck) == 1 {
		newVillain := gs.villainDeck[0]
		gs.villainDeck = []Villain{}
		return append(villains, newVillain)
	} else {
		newVillain := gs.villainDeck[0]
		gs.villainDeck = gs.villainDeck[1:]
		return append(villains, newVillain)
	}
}

func ResetPlayerInfo(gs *Gamestate) {
	for user := range gs.Players {
		player := gs.Players[user]
		player.alliesToDeck = false
		player.itemsToDeck = false
		player.spellsToDeck = false
		player.proficiencyUsed = false
		gs.Players[user] = player
	}
}

func PurchaseCard(c Card, user string, gs *Gamestate) {
	player := gs.Players[user]
	if c.houseDice && player.Proficiency == "Arithmancy" {
		player.Money -= c.Cost - 1
	} else {
		player.Money -= c.Cost
	}
	gs.Players[user] = player

	switch c.CardType {
	case "item":
		if player.itemsToDeck {
			ChooseOne{
				Effects: []Effect{
					GainCardToTopDeck{user: user, card: c},
					GainCardToDiscard{user: user, card: c},
				},
				Options:     []string{"Yes", "No"},
				Description: "Gain to top of deck?",
			}.Trigger(gs)
		} else {
			GainCardToDiscard{user: user, card: c}.Trigger(gs)
		}
	case "spell":
		if player.spellsToDeck {
			ChooseOne{
				Effects: []Effect{
					GainCardToTopDeck{user: user, card: c},
					GainCardToDiscard{user: user, card: c},
				},
				Options:     []string{"Yes", "No"},
				Description: "Gain to top of deck?",
			}.Trigger(gs)
		} else {
			GainCardToDiscard{user: user, card: c}.Trigger(gs)
		}
		if player.Proficiency == "History of Magic" {
			HealAnyPlayer{Amount: 1}.Trigger(gs)
		}
	case "ally":
		if player.alliesToDeck {
			ChooseOne{
				Effects: []Effect{
					GainCardToTopDeck{user: user, card: c},
					GainCardToDiscard{user: user, card: c},
				},
				Options:     []string{"Yes", "No"},
				Description: "Gain to top of deck?",
			}.Trigger(gs)
		} else {
			GainCardToDiscard{user: user, card: c}.Trigger(gs)
		}
	default:
		GainCardToDiscard{user: user, card: c}.Trigger(gs)
	}

	if c.Cost >= 4 {
		eventBroker.Messages <- DoloresUmbridgeTrigger
	}
}

func getPreviousUser(gs *Gamestate) string {
	currUser := gs.CurrentTurn
	currTurn := 0
	for i, u := range gs.TurnOrder {
		if u == currUser {
			currTurn = i
		}
	}

	prevTurn := currTurn - 1
	if prevTurn == -1 {
		prevTurn = len(gs.TurnOrder) - 1
	}

	return gs.TurnOrder[prevTurn]
}

func getNextUser(gs *Gamestate) string {
	currUser := gs.CurrentTurn
	currTurn := 0
	for i, u := range gs.TurnOrder {
		if u == currUser {
			currTurn = i
		}
	}

	nextTurn := (currTurn + 1) % len(gs.TurnOrder)
	return gs.TurnOrder[nextTurn]
}

func assertUniqueCards(gs *Gamestate) bool {
	ret := true
	seen := make(map[int]bool)
	name := ""
	dupe := 0
	for _, p := range gs.Players {
		for _, c := range p.Deck {
			if seen[c.Id] {
				ret = false
				dupe = c.Id
				name = c.Name
			}
			seen[c.Id] = true
		}

		for _, c := range p.Hand {
			if seen[c.Id] {
				ret = false
				dupe = c.Id
				name = c.Name
			}
			seen[c.Id] = true
		}

		for _, c := range p.Discard {
			if seen[c.Id] {
				ret = false
				dupe = c.Id
				name = c.Name
			}
			seen[c.Id] = true
		}
	}

	if !ret {
		log.Println("****************** DUPLICATE CARD IDS:", dupe, name)
	}

	return ret
}

func DiscardFromId(user string, cardId int, gs *Gamestate) {
	log.Println("vvvvvvvvvvvv STARTING DISCARDFROMID:", cardId)
	assertUniqueCards(gs)
	log.Println("OKAY AT START? ^")

	player := gs.Players[user]

	for i, c := range player.Hand {
		if c.Id == cardId {
			player.Hand = RemoveCardAtIndex(player.Hand, i)
			player.Discard = append(player.Discard, c)
			gs.Players[user] = player

			// Wrap the player mapping around onDiscard since it mutates the state directly.
			if c.onDiscard != nil {
				c.onDiscard(user, gs)
			}
		}
	}
	for i, c := range player.Deck {
		if c.Id == cardId {
			player.Deck = RemoveCardAtIndex(player.Deck, i)
			player.Discard = append(player.Discard, c)
			gs.Players[user] = player

			// Wrap the player mapping around onDiscard since it mutates the state directly.
			if c.onDiscard != nil {
				c.onDiscard(user, gs)
			}
		}
	}

	event := Event{senderId: -1, message: "player discarded", data: user}
	eventBroker.Messages <- event
	// update turnstats

	log.Println("^^^^^^^^^^^^^^^^^ ENDING DISCARDFROMID:", cardId)
	assertUniqueCards(gs)
	log.Println("OKAY AT END? ^")
}
