package main

import (
	"log"
	"net/http"
	"sync"
)

var StateLog []*Gamestate

func addToGamestateLog(gs *Gamestate) {

	playermap := make(map[string]Player)
	for user, p := range gs.Players {
		playermap[user] = p
	}

	copy := gs.DeepCopy()

	for user := range copy.Players {
		p := copy.Players[user]
		p.Health = playermap[user].Health
		p.Damage = playermap[user].Damage
		p.Money = playermap[user].Money
		p.Name = playermap[user].Name
		p.Character = playermap[user].Character
		copy.Players[user] = p
	}

	StateLog = append(StateLog, &copy)
}

func undoHandler(w http.ResponseWriter, r *http.Request, gs *Gamestate) {
	id, user := getIdAndUser(r)
	gs.mu.Lock()
	defer gs.mu.Unlock()

	log.Println("start undo:", states[id].turnNumber, states[id].Players[user].Name)

	if len(StateLog) == 0 {
		return
	}
	newGS := StateLog[len(StateLog)-1]
	StateLog = StateLog[:len(StateLog)-1]
	newGS.mu = sync.Mutex{}

	globalMu.Lock()
	defer globalMu.Unlock()

	states[id] = newGS

	log.Println("end undo:", states[id].turnNumber, states[id].Players[user].Name)
	SendLobbyUpdate(id, newGS)
}

// DeepCopy generates a deep copy of Gamestate
func (o Gamestate) DeepCopy() Gamestate {
	var cp Gamestate = o
	if o.Players != nil {
		cp.Players = make(map[string]Player, len(o.Players))
		for k2, v2 := range o.Players {
			var cp_Players_v2 Player
			if v2.Deck != nil {
				cp_Players_v2.Deck = make([]Card, len(v2.Deck))
				copy(cp_Players_v2.Deck, v2.Deck)
				for i4 := range v2.Deck {
					if v2.Deck[i4].effects != nil {
						cp_Players_v2.Deck[i4].effects = make([]Effect, len(v2.Deck[i4].effects))
						copy(cp_Players_v2.Deck[i4].effects, v2.Deck[i4].effects)
					}
				}
			}
			if v2.Hand != nil {
				cp_Players_v2.Hand = make([]Card, len(v2.Hand))
				copy(cp_Players_v2.Hand, v2.Hand)
				for i4 := range v2.Hand {
					if v2.Hand[i4].effects != nil {
						cp_Players_v2.Hand[i4].effects = make([]Effect, len(v2.Hand[i4].effects))
						copy(cp_Players_v2.Hand[i4].effects, v2.Hand[i4].effects)
					}
				}
			}
			if v2.Discard != nil {
				cp_Players_v2.Discard = make([]Card, len(v2.Discard))
				copy(cp_Players_v2.Discard, v2.Discard)
				for i4 := range v2.Discard {
					if v2.Discard[i4].effects != nil {
						cp_Players_v2.Discard[i4].effects = make([]Effect, len(v2.Discard[i4].effects))
						copy(cp_Players_v2.Discard[i4].effects, v2.Discard[i4].effects)
					}
				}
			}
			if v2.PlayArea != nil {
				cp_Players_v2.PlayArea = make([]Card, len(v2.PlayArea))
				copy(cp_Players_v2.PlayArea, v2.PlayArea)
				for i4 := range v2.PlayArea {
					if v2.PlayArea[i4].effects != nil {
						cp_Players_v2.PlayArea[i4].effects = make([]Effect, len(v2.PlayArea[i4].effects))
						copy(cp_Players_v2.PlayArea[i4].effects, v2.PlayArea[i4].effects)
					}
				}
			}
			cp.Players[k2] = cp_Players_v2
		}
	}
	if o.Villains != nil {
		cp.Villains = make([]Villain, len(o.Villains))
		copy(cp.Villains, o.Villains)
		for i2 := range o.Villains {
			if o.Villains[i2].effect != nil {
				cp.Villains[i2].effect = make([]Effect, len(o.Villains[i2].effect))
				copy(cp.Villains[i2].effect, o.Villains[i2].effect)
			}
			if o.Villains[i2].deathEffect != nil {
				cp.Villains[i2].deathEffect = make([]Effect, len(o.Villains[i2].deathEffect))
				copy(cp.Villains[i2].deathEffect, o.Villains[i2].deathEffect)
			}
		}
	}
	if o.Locations != nil {
		cp.Locations = make([]Location, len(o.Locations))
		copy(cp.Locations, o.Locations)
	}
	if o.DarkArts != nil {
		cp.DarkArts = make([]DarkArt, len(o.DarkArts))
		copy(cp.DarkArts, o.DarkArts)
		for i2 := range o.DarkArts {
			if o.DarkArts[i2].effect != nil {
				cp.DarkArts[i2].effect = make([]Effect, len(o.DarkArts[i2].effect))
				copy(cp.DarkArts[i2].effect, o.DarkArts[i2].effect)
			}
		}
	}
	if o.MarketDeck != nil {
		cp.MarketDeck = make([]Card, len(o.MarketDeck))
		copy(cp.MarketDeck, o.MarketDeck)
		for i2 := range o.MarketDeck {
			if o.MarketDeck[i2].effects != nil {
				cp.MarketDeck[i2].effects = make([]Effect, len(o.MarketDeck[i2].effects))
				copy(cp.MarketDeck[i2].effects, o.MarketDeck[i2].effects)
			}
		}
	}
	if o.Market != nil {
		cp.Market = make([]Card, len(o.Market))
		copy(cp.Market, o.Market)
		for i2 := range o.Market {
			if o.Market[i2].effects != nil {
				cp.Market[i2].effects = make([]Effect, len(o.Market[i2].effects))
				copy(cp.Market[i2].effects, o.Market[i2].effects)
			}
		}
	}
	if o.TurnOrder != nil {
		cp.TurnOrder = make([]string, len(o.TurnOrder))
		copy(cp.TurnOrder, o.TurnOrder)
	}
	if o.DarkArtsPlayed != nil {
		cp.DarkArtsPlayed = make([]DarkArt, len(o.DarkArtsPlayed))
		copy(cp.DarkArtsPlayed, o.DarkArtsPlayed)
		for i2 := range o.DarkArtsPlayed {
			if o.DarkArtsPlayed[i2].effect != nil {
				cp.DarkArtsPlayed[i2].effect = make([]Effect, len(o.DarkArtsPlayed[i2].effect))
				copy(cp.DarkArtsPlayed[i2].effect, o.DarkArtsPlayed[i2].effect)
			}
		}
	}
	if o.villainDeck != nil {
		cp.villainDeck = make([]Villain, len(o.villainDeck))
		copy(cp.villainDeck, o.villainDeck)
		for i2 := range o.villainDeck {
			if o.villainDeck[i2].effect != nil {
				cp.villainDeck[i2].effect = make([]Effect, len(o.villainDeck[i2].effect))
				copy(cp.villainDeck[i2].effect, o.villainDeck[i2].effect)
			}
			if o.villainDeck[i2].deathEffect != nil {
				cp.villainDeck[i2].deathEffect = make([]Effect, len(o.villainDeck[i2].deathEffect))
				copy(cp.villainDeck[i2].deathEffect, o.villainDeck[i2].deathEffect)
			}
		}
	}
	if o.turnStats.VillainsHit != nil {
		cp.turnStats.VillainsHit = make([]int, len(o.turnStats.VillainsHit))
		copy(cp.turnStats.VillainsHit, o.turnStats.VillainsHit)
	}
	return cp
}
