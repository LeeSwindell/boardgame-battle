package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CardId struct {
	Id int `json:"id"`
}

func PlayCardHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	gameid, user := getIdAndUser(r)

	gs.mu.Lock()
	defer gs.mu.Unlock()

	if gs.CurrentTurn != user {
		log.Println("Not your turn!")
		return
	}

	var cardId CardId
	json.NewDecoder(r.Body).Decode(&cardId)

	triggerDivination := false
	for _, c := range gs.Players[user].Hand {
		if c.Id == cardId.Id {
			MoveToPlayed(user, c.Id, gs)

			switch c.CardType {
			case "spell":
				eventBroker.Messages <- SpellPlayed
			case "item":
				eventBroker.Messages <- ItemPlayed
				if gs.Players[user].Proficiency == "Divination" {
					triggerDivination = true
				}
			case "ally":
				eventBroker.Messages <- AllyPlayed
			}

			for _, e := range c.effects {
				e.Trigger(gs)
			}

			switch c.CardType {
			case "spell":
				gs.turnStats.SpellsPlayed += 1
				if gs.turnStats.SpellsPlayed == 4 && gs.Players[user].Character == "Hermione" {
					SelectTwoPlayersToGainStats{AmountDamage: 1, Exclusive: true}.Trigger(gs)
				}
			case "item":
				gs.turnStats.ItemsPlayed += 1
			case "ally":
				gs.turnStats.AlliesPlayed += 1
			}

			if gs.Players[user].Proficiency == "Potions" && !gs.Players[user].proficiencyUsed {
				if gs.turnStats.AlliesPlayed >= 1 && gs.turnStats.ItemsPlayed >= 1 && gs.turnStats.SpellsPlayed >= 1 {
					SelectPlayerToGainStats{AmountHealth: 1, AmountDamage: 1}.Trigger(gs)
					p := gs.Players[user]
					p.proficiencyUsed = true
					gs.Players[user] = p
				}
			}

			if triggerDivination {
				Scry{User: user}.Trigger(gs)
			}
		}
	}

	SendLobbyUpdate(gameid, gs)
}

func GetGamestateHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	json.NewEncoder(w).Encode(gs)
}

func EndTurnHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	gameid, user := getIdAndUser(r)

	gs.mu.Lock()
	defer gs.mu.Unlock()

	if gs.CurrentTurn != user {
		log.Println("not your turn!")
		return
	}

	// Testing
	log.Println("Start end turn debug: ", gs.turnNumber)
	assertUniqueCards(gs)
	log.Println("Start end turn debug: ", gs.turnNumber)
	// Testing

	eventBroker.Messages <- EndTurnEvent
	HealStunned(gs)
	MovePlayedToDiscard(user, gs)
	MoveHandToDiscard(user, gs)
	MoneyDamageToZero(user, gs)
	gs.DarkArtsPlayed = []DarkArt{}
	RefillHand(user, gs)
	NextTurnInOrder(gs)
	gs.turnStats = TurnStats{AlliesHealed: map[string]int{}}
	ResetPlayerInfo(gs)

	SendLobbyUpdate(gameid, gs)

	// Starting next turn actions.
	gs.turnNumber++
	for i, v := range gs.Villains {
		if !v.Active {
			for _, v := range gs.Villains {
				if v.Active && v.Name == "Death Eater" {
					DamageAllPlayers{Amount: 1}.Trigger(gs)
				}
			}
			gs.Villains[i].Active = true
		}
		if v.playBeforeDA && gs.turnNumber >= v.BlockedUntil {
			for _, e := range v.effect {
				Logger("triggering " + v.Name)
				e.Trigger(gs)
			}
		}
	}
	Logger("triggering locations")
	gs.Locations[gs.CurrentLocation].effect.Trigger(gs)

	Logger("After DA villains")
	for _, v := range gs.Villains {
		if !v.playBeforeDA && gs.turnNumber >= v.BlockedUntil {
			for _, e := range v.effect {
				Logger("triggering " + v.Name)
				e.Trigger(gs)
			}
		}
	}

	// Testing
	log.Println("End of end turn debug: ", gs.turnNumber-1)
	assertUniqueCards(gs)
	log.Println("End of end turn debug: ", gs.turnNumber-1)
	// Testing

	// Add to statelog
	addToGamestateLog(gs)

	SendLobbyUpdate(gameid, gs)
}

func DamageVillainHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	gameid, user := getIdAndUser(r)
	vars := mux.Vars(r)
	villainid, err := strconv.Atoi(vars["villainid"])
	if err != nil {
		log.Println("err getting villainid:", err.Error())
	}

	gs.mu.Lock()
	defer gs.mu.Unlock()

	player := gs.Players[user]

	for i, v := range gs.Villains {
		if v.Id == villainid {
			// do nothing if villain is already dead.
			if gs.Villains[i].CurDamage >= v.MaxHp {
				return
			}

			// do nothing if its voldemort and there's others.
			if v.Name == "Voldemort" && len(gs.Villains) != 1 {
				return
			}

			if player.Damage <= 0 && v.Name != "Norbert" {
				log.Println("not enough damage tokens to do this")
				return
			}

			alreadyHit := false
			for _, hitId := range gs.turnStats.VillainsHit {
				if gs.Villains[i].Id == hitId {
					alreadyHit = true
				}
			}

			// do nothing if Tarantallegra! is played and villain has already been hit.
			for _, da := range gs.DarkArtsPlayed {
				if da.Name == "Tarantallegra!" && alreadyHit && v.Name != "Norbert" {
					return
				}
			}

			// spend money to damage norbert.
			if v.Name == "Norbert" && player.Money > 0 {
				player.Money -= 1
			} else if v.Name == "Norbert" && player.Money <= 0 {
				return
			} else {
				player.Damage -= 1
			}

			gs.Villains[i].CurDamage += 1

			// check for Rons character effect.
			gs.turnStats.DamageDealt++
			if gs.turnStats.DamageDealt == 3 && player.Character == "Ron" {
				AllPlayersGainHealth{Amount: 1}.Trigger(gs)
			}

			// check for Care of Magical Creatures proficiency
			if player.Proficiency == "Care of Magical Creatures" && !player.proficiencyUsed {
				player.proficiencyUsed = true
				gs.Players[user] = player
				HealAnyPlayer{Amount: 2}.Trigger(gs)
				player = gs.Players[user]
			}

			if !alreadyHit && gs.Villains[i].Name != "Norbert" {
				gs.turnStats.VillainsHit = append(gs.turnStats.VillainsHit, gs.Villains[i].Id)
				eventBroker.Messages <- NewVillainHitEvent
			}
			gs.Players[user] = player

			// check if villain is now dead.
			if gs.Villains[i].CurDamage >= v.MaxHp {
				gs.Villains[i].Active = false
				// trigger villain death effect.
				for _, effect := range gs.Villains[i].deathEffect {
					effect.Trigger(gs)
				}
				if player.Proficiency == "Care of Magical Creatures" && (v.villainType == "creature" || v.villainType == "villain-creature") {
					RemoveFromLocation{Amount: 1}.Trigger(gs)
				}

				// remove villain, get new one
				newVillains := RemoveVillainAtIndex(gs.Villains, i)
				newVillains = AddNewActiveVillain(newVillains, gs)
				gs.Villains = newVillains
				eventBroker.Messages <- Event{senderId: -1, message: "villain killed"}
			}
			break
		}
	}

	SendLobbyUpdate(gameid, gs)
}

func BuyCardHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	gameid, user := getIdAndUser(r)
	vars := mux.Vars(r)
	cardid, err := strconv.Atoi(vars["cardid"])
	if err != nil {
		log.Println("err getting cardid:", err.Error())
	}

	gs.mu.Lock()
	defer gs.mu.Unlock()

	if gs.CurrentTurn != user {
		log.Println("not your turn!")
		return
	}

	player := gs.Players[user]
	cond1 := player.Proficiency == "Arithmancy"
	for i, c := range gs.Market {
		if c.Id == cardid && (player.Money >= c.Cost || (cond1 && c.houseDice && player.Money >= c.Cost-1)) {
			PurchaseCard(c, user, gs)
			RefillMarket(i, gs)
		}
	}

	SendLobbyUpdate(gameid, gs)
}

func UseProficiencyHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	_, user := getIdAndUser(r)
	if gs.CurrentTurn != user {
		return
	}

	player := gs.Players[user]
	switch player.Proficiency {
	case "Charms":
		if player.proficiencyUsed {
			return
		}
		numSpells := 0
		for _, c := range player.Hand {
			if c.CardType == "spell" {
				numSpells++
			}
		}
		if numSpells < 2 {
			return
		}
		player.proficiencyUsed = true
		gs.Players[user] = player

		DiscardACard{Target: user, Prompt: "Discard 1st spell", Cardtype: "spell"}.Trigger(gs)
		DiscardACard{Target: user, Prompt: "Discard 2nd spell", Cardtype: "spell"}.Trigger(gs)
		AllDrawCards{Amount: 1}.Trigger(gs)
		AllPlayersGainMoney{Amount: 1}.Trigger(gs)

	case "Transfiguration":
		if player.proficiencyUsed {
			return
		}
		numItems := 0
		for _, c := range player.Hand {
			if c.CardType == "item" {
				numItems++
			}
		}
		if numItems == 0 {
			return
		}
		player.proficiencyUsed = true
		gs.Players[user] = player
		DiscardACard{Target: user, Prompt: "Discard an item", Cardtype: "item"}.Trigger(gs)
		ActivePlayerSearchesDeckForX{CardType: "any", Target: user, CostRestraint: 5}.Trigger(gs)

	case "Flying Lessons":
		if player.proficiencyUsed || player.Money < 5 {
			return
		}
		player.Money -= 5
		player.proficiencyUsed = true
		gs.Players[user] = player
		RemoveFromLocation{Amount: 1}.Trigger(gs)
	}

	SendLobbyUpdate(gs.gameid, gs)
}
