package main

import (
	"reflect"

	"github.com/google/uuid"
)

func alohamora() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Alohomora",
		SetId:    "game 1",
		ImgPath:  "/images/starters/alohomora.jpg",
		CardType: "spell",
		Cost:     0,
		effects:  []Effect{GainMoney{Amount: 1}},
	}
}

func pigwidgeon() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Pigwidgeon",
		SetId:    "game 1",
		ImgPath:  "/images/starters/pigwidgeon.jpg",
		CardType: "ally",
		Cost:     0,
		effects: []Effect{
			ChooseOne{
				Effects: []Effect{
					GainDamage{Amount: 1},
					GainHealth{Amount: 2},
				},
				Options:     []string{"Gain 1 Damage", "Gain 2 Health"},
				Description: "",
			},
		},
	}
}

func cleansweep() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Cleansweep 11",
		SetId:    "game 1",
		ImgPath:  "/images/starters/cleansweep.jpg",
		CardType: "item",
		Cost:     0,
		effects: []Effect{
			GainDamage{Amount: 1},
			MoneyIfVillainKilled{Id: id, Amount: 1},
		},
	}
}

func bertieBottsEveryFlavourBeans() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Bertie Botts Every-Flavour Beans",
		SetId:    "game 1",
		ImgPath:  "/images/starters/bertiebottseveryflavourbeans.jpg",
		CardType: "item",
		Cost:     0,
		effects: []Effect{
			GainMoney{Amount: 1},
			GainDamagePerAllyPlayed{},
		},
	}
}

func crookshanks() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Crookshanks",
		SetId:    "game 1",
		ImgPath:  "/images/starters/crookshanks.jpg",
		CardType: "ally",
		Cost:     0,
		effects: []Effect{
			ChooseOne{
				Effects: []Effect{
					GainDamage{Amount: 1},
					GainHealth{Amount: 2},
				},
				Options:     []string{"Gain 1 Damage", "Gain 2 Health"},
				Description: "Choose one:",
			},
		},
	}
}

func crumpleHornedSnorkack() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Crumple-horned Snorkack",
		SetId:    "box 1",
		ImgPath:  "/images/starters/crumplehornedsnorkack.jpg",
		CardType: "ally",
		Cost:     0,
		effects: []Effect{
			ChooseOne{
				Effects: []Effect{
					GainDamage{Amount: 1},
					GainHealth{Amount: 2},
				},
				Options:     []string{"Gain 1 Damage", "Gain 2 Health"},
				Description: "Choose one:",
			},
		},
	}
}

func firebolt() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Firebolt",
		SetId:    "game 1",
		ImgPath:  "/images/starters/firebolt.jpg",
		CardType: "item",
		Cost:     0,
		effects: []Effect{
			GainDamage{Amount: 1},
			MoneyIfVillainKilled{Amount: 1},
		},
	}
}

func hedwig() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Hedwig",
		SetId:    "game 1",
		ImgPath:  "/images/starters/hedwig.jpg",
		CardType: "ally",
		Cost:     0,
		effects: []Effect{
			ChooseOne{
				Effects: []Effect{
					GainDamage{Amount: 1},
					GainHealth{Amount: 2},
				},
				Options:     []string{"Gain 1 Damage", "Gain 2 Health"},
				Description: "Choose one:",
			},
		},
	}
}

func invisibilityCloak() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Invisibility Cloak",
		SetId:    "game 1",
		ImgPath:  "/images/starters/invisibilitycloak.jpg",
		CardType: "item",
		Cost:     0,
		effects: []Effect{
			GainMoney{Amount: 1},
		},
	}
}

func lionHat() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Lion Hat",
		SetId:    "box 1",
		ImgPath:  "/images/starters/lionhat.jpg",
		CardType: "item",
		Cost:     0,
		effects: []Effect{
			GainMoney{Amount: 1},
			LionHatEffect{},
		},
	}
}

type LionHatEffect struct{}

func (effect LionHatEffect) Log(gs *Gamestate) {
	gs.EffectLog = append(gs.EffectLog, reflect.Type.Name(reflect.TypeOf(effect)))
}

func (effect LionHatEffect) Trigger(gs *Gamestate) {
	brooms := make(map[string]bool)
	brooms["Quidditch Gear"] = true
	brooms["Firebolt"] = true
	brooms["Cleansweep 11"] = true
	brooms["Nimbus Two Thousand And One"] = true

	cur := gs.CurrentTurn
	for _, p := range gs.Players {
		if p.Name != cur {
			for _, c := range p.Hand {
				if brooms[c.Name] {
					ChangeStats{Target: gs.CurrentTurn, AmountDamage: 1}.Trigger(gs)
					return
				}
			}
		}
	}
}

func mandrake() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Mandrake",
		SetId:    "game 1",
		ImgPath:  "/images/starters/mandrake.jpg",
		CardType: "item",
		Cost:     0,
		effects: []Effect{
			ChooseOne{
				Effects: []Effect{
					GainDamage{Amount: 1},
					HealAnyPlayer{Amount: 2},
				},
				Options:     []string{"Gain 1 Damage", "Heal any player 2"},
				Description: "Choose one:",
			},
		},
	}
}

func remembrall() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Remembrall",
		SetId:    "game 1",
		ImgPath:  "/images/starters/remembrall.jpg",
		CardType: "item",
		Cost:     0,
		effects: []Effect{
			GainMoney{Amount: 1},
		},
		onDiscard: func(target string, gs *Gamestate) {
			ChangeStats{
				Target:      target,
				AmountMoney: 2,
			}.Trigger(gs)
		},
	}
}

func talesOfBeedleTheBard() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Tales of Beedle the Bard",
		SetId:    "game 1",
		ImgPath:  "/images/starters/talesofbeedlethebard.jpg",
		CardType: "item",
		Cost:     0,
		effects: []Effect{
			ChooseOne{
				Effects: []Effect{
					GainMoney{Amount: 2},
					AllPlayersGainMoney{Amount: 1},
				},
				Options:     []string{"Gain 2 Money", "All players gain 1 Money"},
				Description: "Choose one:",
			},
		},
	}
}

func spectrespecs() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Spectrespecs",
		SetId:    "box 1",
		ImgPath:  "/images/starters/spectrespecs.jpg",
		CardType: "item",
		Cost:     0,
		effects: []Effect{
			GainMoney{Amount: 1},
			ScryDarkarts{},
		},
	}
}

func timeTurner() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Time Turner",
		SetId:    "game 1",
		ImgPath:  "/images/starters/timeturner.jpg",
		CardType: "item",
		Cost:     0,
		effects: []Effect{
			GainMoney{Amount: 1},
			PurchasedXGoToDeck{"spell"},
		},
	}
}

func trevor() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Trevor",
		SetId:    "game 1",
		ImgPath:  "/images/starters/trevor.jpg",
		CardType: "ally",
		Cost:     0,
		effects: []Effect{
			ChooseOne{
				Effects: []Effect{
					GainDamage{Amount: 1},
					GainHealth{Amount: 2},
				},
				Options:     []string{"Gain 1 Damage", "Gain 2 Health"},
				Description: "Choose one:",
			},
		},
	}
}

func crystalBall() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Crystal Ball",
		SetId:    "game 3",
		ImgPath:  "/images/marketcards/crystalball.jpg",
		CardType: "item",
		Cost:     3,
		effects: []Effect{
			DrawCards{Amount: 2},
			SendGameUpdateEffect{},
			ActivePlayerDiscards{Amount: 1},
		},
	}
}

func finite() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Finite!",
		SetId:    "game 3",
		ImgPath:  "/images/marketcards/finite.jpg",
		CardType: "spell",
		Cost:     3,
		effects:  []Effect{RemoveFromLocation{Amount: 1}},
	}
}

func incendio() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Incendio!",
		SetId:    "game 1",
		ImgPath:  "/images/marketcards/incendio.jpg",
		CardType: "spell",
		Cost:     4,
		effects: []Effect{
			GainDamage{Amount: 1},
			DrawCards{Amount: 1},
		},
	}
}

func oliverWood() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Oliver Wood",
		SetId:    "game 1",
		ImgPath:  "/images/marketcards/oliverwood.jpg",
		CardType: "ally",
		Cost:     3,
		effects: []Effect{
			GainDamage{Amount: 1},
			HealAnyIfVillainKilled{Amount: 2, Id: id},
		},
	}
}

func reparo() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Reparo",
		SetId:    "game 1",
		ImgPath:  "/images/marketcards/reparo.jpg",
		CardType: "spell",
		Cost:     3,
		effects: []Effect{
			ChooseOne{
				Effects: []Effect{
					GainMoney{Amount: 2},
					DrawCards{Amount: 1},
				},
				Options: []string{"Gain 2 Money", "Draw 1 Card"},
			},
		},
	}
}

func triwizardCup() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Triwizard Cup",
		SetId:    "game 4",
		ImgPath:  "/images/marketcards/triwizardcup.jpg",
		CardType: "item",
		Cost:     5,
		effects: []Effect{
			GainDamage{Amount: 1},
			GainMoney{Amount: 1},
			GainHealth{Amount: 1},
		},
	}
}

func detention() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Detention!",
		SetId:    "box 1",
		ImgPath:  "/images/marketcards/detention.jpg",
		CardType: "item",
		Cost:     0,
		effects:  []Effect{},
		onDiscard: func(target string, gs *Gamestate) {
			ChangeStats{
				Target:       target,
				AmountHealth: -2,
			}.Trigger(gs)
		},
	}
}

func albusDumbledore() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Albus Dumbledore",
		SetId:    "game 1",
		ImgPath:  "/images/marketcards/albusdumbledore.jpg",
		CardType: "ally",
		Cost:     8,
		effects: []Effect{
			AllPlayersGainMoney{Amount: 1},
			AllPlayersGainHealth{Amount: 1},
			AllPlayersGainDamage{Amount: 1},
			AllDrawCards{Amount: 1},
		},
	}
}

func arthurWeasley() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Arthur Weasley",
		SetId:    "game 2",
		ImgPath:  "/images/marketcards/arthurweasley.jpg",
		CardType: "ally",
		Cost:     6,
		effects: []Effect{
			AllPlayersGainMoney{Amount: 2},
		},
	}
}

func bezoar() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Bezoar",
		SetId:    "game 6",
		ImgPath:  "/images/marketcards/bezoar.jpg",
		CardType: "item",
		Cost:     4,
		effects: []Effect{
			HealAnyPlayer{Amount: 3},
			DrawCards{Amount: 1},
		},
	}
}

func choChang() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Cho Chang",
		SetId:    "game 5",
		ImgPath:  "/images/marketcards/chochang.jpg",
		CardType: "ally",
		Cost:     4,
		effects: []Effect{
			DrawCards{Amount: 3},
			ActivePlayerDiscards{Amount: 2},
			RavenclawDice{},
		},
		houseDice: true,
	}
}

func deluminator() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Deluminator",
		SetId:    "game 6",
		ImgPath:  "/images/marketcards/deluminator.jpg",
		CardType: "item",
		Cost:     6,
		effects: []Effect{
			RemoveFromLocation{Amount: 2},
		},
	}
}

func expectoPatronum() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Expecto Patronum",
		SetId:    "game 3",
		ImgPath:  "/images/marketcards/expectopatronum.jpg",
		CardType: "spell",
		Cost:     5,
		effects: []Effect{
			GainDamage{Amount: 1},
			RemoveFromLocation{Amount: 1},
		},
	}
}

func felixFelicis() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Felix Felicis",
		SetId:    "game 6",
		ImgPath:  "/images/marketcards/felixfelicis.jpg",
		CardType: "item",
		Cost:     7,
		effects: []Effect{
			ChooseTwo{
				Exclusive: true,
				Effects: []Effect{
					GainDamage{Amount: 2},
					GainMoney{Amount: 2},
					GainHealth{Amount: 2},
					DrawCards{Amount: 2},
				},
				Options: []string{
					"Gain 2 Damage", "Gain 2 Money", "Gain 2 Health", "Draw 2 Cards",
				},
				Prompt: "Felix Felicis: Choose One",
			},
		},
	}
}

func filiusFlitwick() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Filius Flitwick",
		SetId:    "game 4",
		ImgPath:  "/images/marketcards/filiusflitwick.jpg",
		CardType: "ally",
		Cost:     6,
		effects: []Effect{
			GainDamage{Amount: 1},
			DrawCards{Amount: 1},
			RavenclawDice{},
		},
		houseDice: true,
	}
}

func hogwartsAHistory() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Hogwarts: A History",
		SetId:    "game 4",
		ImgPath:  "/images/marketcards/hogwartsahistory.jpg",
		CardType: "item",
		Cost:     4,
		effects: []Effect{
			ChooseOne{
				Effects: []Effect{
					RavenclawDice{},
					SlytherinDice{},
					GryffindorDice{},
					HufflepuffDice{},
				},
				Options: []string{
					"Ravenclaw", "Slytherin", "Gryffindor", "Hufflepuff",
				},
				Description: "Choose a House dice to roll",
			},
		},
		houseDice: true,
	}
}

func mollyWeasley() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Molly Weasley",
		SetId:    "game 2",
		ImgPath:  "/images/marketcards/mollyweasley.jpg",
		CardType: "ally",
		Cost:     6,
		effects: []Effect{
			AllPlayersGainMoney{Amount: 1},
			AllPlayersGainHealth{Amount: 2},
		},
	}
}

func quidditchGear() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Quidditch Gear",
		SetId:    "game 1",
		ImgPath:  "/images/marketcards/quidditchgear.jpg",
		CardType: "item",
		Cost:     3,
		effects: []Effect{
			GainDamage{Amount: 1},
			GainHealth{Amount: 1},
		},
	}
}

func siriusBlack() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Sirius Black",
		SetId:    "game 3",
		ImgPath:  "/images/marketcards/siriusblack.jpg",
		CardType: "ally",
		Cost:     6,
		effects: []Effect{
			GainDamage{Amount: 2},
			GainMoney{Amount: 1},
		},
	}
}

func stupefy() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Stupefy!",
		SetId:    "game 5",
		ImgPath:  "/images/marketcards/stupefy.jpg",
		CardType: "spell",
		Cost:     6,
		effects: []Effect{
			GainDamage{Amount: 1},
			RemoveFromLocation{Amount: 1},
			DrawCards{Amount: 1},
		},
	}
}

func sybillTrelawney() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Sybill Trelawney",
		SetId:    "game 3",
		ImgPath:  "/images/marketcards/sybilltrelawney.jpg",
		CardType: "ally",
		Cost:     4,
		effects: []Effect{
			DrawCards{Amount: 2},
			SybillDiscard{},
		},
	}
}

type SybillDiscard struct{}

func (effect SybillDiscard) Log(gs *Gamestate) {
	gs.EffectLog = append(gs.EffectLog, reflect.Type.Name(reflect.TypeOf(effect)))
}

func (effect SybillDiscard) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	player := gs.Players[user]

	cards := player.Hand
	if len(cards) == 0 {
		return
	}

	discardCardId := AskUserToSelectCard(user, gs.gameid, cards, "Discard a card (+2 money if it's a spell)")
	for i, c := range cards {
		if c.Id == discardCardId {
			if c.CardType == "spell" {
				player.Money += 2
			}
			// Wrap the player mapping around onDiscard since it mutates the state directly.
			if c.onDiscard != nil {
				gs.Players[user] = player
				c.onDiscard(user, gs)
				player = gs.Players[user]
			}
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

func butterbeer() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Butterbeer",
		SetId:    "game 3",
		ImgPath:  "/images/marketcards/butterbeer.jpg",
		CardType: "item",
		Cost:     3,
		effects: []Effect{
			SelectTwoPlayersToGainStats{
				AmountHealth: 1,
				AmountMoney:  1,
				Exclusive:    true,
			},
		},
	}
}

func dobbyTheHouseElf() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Dobby the House Elf",
		SetId:    "2",
		ImgPath:  "/images/marketcards/dobbythehouseelf.jpg",
		CardType: "ally",
		Cost:     4,
		effects: []Effect{
			RemoveFromLocation{Amount: 1},
			DrawCards{Amount: 1},
		},
	}
}

func essenceOfDittany() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Essence of Dittany",
		SetId:    "game 1",
		ImgPath:  "/images/marketcards/essenceofdittany.jpg",
		CardType: "item",
		Cost:     2,
		effects: []Effect{
			HealAnyPlayer{Amount: 2},
		},
	}
}

func fang() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Fang",
		SetId:    "box 1",
		ImgPath:  "/images/marketcards/fang.jpg",
		CardType: "ally",
		Cost:     3,
		effects: []Effect{
			SelectPlayerToGainStats{
				AmountHealth: 2,
				AmountMoney:  1,
			},
		},
	}
}

func goldenSnitch() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Golden Snitch",
		SetId:    "game 1",
		ImgPath:  "/images/marketcards/goldensnitch.jpg",
		CardType: "item",
		Cost:     5,
		effects: []Effect{
			GainMoney{Amount: 2},
			DrawCards{Amount: 1},
		},
	}
}

func pensieve() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Pensieve",
		SetId:    "game 4",
		ImgPath:  "/images/marketcards/pensieve.jpg",
		CardType: "item",
		Cost:     5,
		effects: []Effect{
			SelectTwoPlayersToGainStats{
				AmountMoney: 1,
				AmountCards: 1,
				Exclusive:   true,
			},
		},
	}
}

func rubeusHagrid() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Rubeus Hagrid",
		SetId:    "game 1",
		ImgPath:  "/images/marketcards/rubeushagrid.jpg",
		CardType: "ally",
		Cost:     4,
		effects: []Effect{
			GainDamage{Amount: 1},
			AllPlayersGainHealth{Amount: 1},
		},
	}
}

func fleurDelacour() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Fleur Delacour",
		SetId:    "game 4",
		ImgPath:  "/images/marketcards/fleurdelacour.jpg",
		CardType: "ally",
		Cost:     4,
		effects: []Effect{
			GainMoney{Amount: 2},
			GainStatIfXPlayed{AmountHealth: 2, Cardtype: "ally", Id: id},
		},
	}
}

func advancedPotionMaking() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Advanced Potion Making",
		SetId:    "game 6",
		ImgPath:  "/images/marketcards/advancedpotionmaking.jpg",
		CardType: "item",
		Cost:     6,
		effects: []Effect{
			AllPlayersGainHealth{Amount: 2},
			AllPlayersAtMaxHealthGainX{
				AmountDamage: 1,
				AmountCards:  1,
			},
		},
	}
}

func alastorMadEyeMoody() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Alastor Mad-Eye Moody",
		SetId:    "game 4",
		ImgPath:  "/images/marketcards/alastormadeyemoody.jpg",
		CardType: "ally",
		Cost:     6,
		effects: []Effect{
			GainMoney{Amount: 2},
			RemoveFromLocation{Amount: 1},
		},
	}
}

func argusFilchAndMrsNorris() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Argus Filch & Mrs. Norris",
		SetId:    "box 1",
		ImgPath:  "/images/marketcards/argusfilchandmrsnorris.jpg",
		CardType: "ally",
		Cost:     4,
		effects: []Effect{
			DrawCards{Amount: 2},
			ChooseOne{
				Effects: []Effect{
					ActivePlayerDiscards{Amount: 1},
					ActivePlayerBanishes{Hand: true},
				},
				Options:     []string{"Discard a card", "Banish a card from hand"},
				Description: "Argus Finally helps out! Choose one:",
			},
		},
	}
}

func cedricDiggory() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Cedric Diggory",
		SetId:    "game 4",
		ImgPath:  "/images/marketcards/cedricdiggory.jpg",
		CardType: "ally",
		Cost:     4,
		effects: []Effect{
			GainDamage{Amount: 1},
			HufflepuffDice{},
		},
		houseDice: true,
	}
}

func descendo() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Descendo!",
		SetId:    "game 1",
		ImgPath:  "/images/marketcards/descendo.jpg",
		CardType: "spell",
		Cost:     5,
		effects: []Effect{
			GainDamage{Amount: 2},
		},
	}
}

func expelliarmus() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Expelliarmus!",
		SetId:    "game 2",
		ImgPath:  "/images/marketcards/expelliarmus.jpg",
		CardType: "spell",
		Cost:     6,
		effects: []Effect{
			GainDamage{Amount: 2},
			DrawCards{Amount: 1},
		},
	}
}

func ginnyWeasley() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Ginny Weasley",
		SetId:    "game 2",
		ImgPath:  "/images/marketcards/ginnyweasley.jpg",
		CardType: "ally",
		Cost:     4,
		effects: []Effect{
			GainDamage{Amount: 1},
			GainMoney{Amount: 1},
		},
	}
}

func horaceSlughorn() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Horace Slughorn",
		SetId:    "game 6",
		ImgPath:  "/images/marketcards/horaceslughorn.jpg",
		CardType: "ally",
		Cost:     6,
		effects: []Effect{
			AllChooseOne{
				Effects: []Effect{
					GainMoney{Amount: 1},
					GainHealth{Amount: 1},
				},
				Options:     []string{"Gain 1 money", "Gain 1 health"},
				Description: "Horace Slughorn: Choose one",
			},
			SlytherinDice{},
		},
		houseDice: true,
	}
}

func kingsleyShacklebolt() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Kingsley Shacklebolt",
		SetId:    "game 5",
		ImgPath:  "/images/marketcards/kingsleyshacklebolt.jpg",
		CardType: "ally",
		Cost:     7,
		effects: []Effect{
			GainDamage{Amount: 2},
			GainHealth{Amount: 1},
			RemoveFromLocation{Amount: 1},
		},
	}
}

func lumos() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Lumos",
		SetId:    "game 3",
		ImgPath:  "/images/marketcards/lumos.jpg",
		CardType: "spell",
		Cost:     4,
		effects: []Effect{
			AllDrawCards{Amount: 1},
		},
	}
}

func lunaLovegood() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Luna Lovegood",
		SetId:    "game 5",
		ImgPath:  "/images/marketcards/lunalovegood.jpg",
		CardType: "ally",
		Cost:     5,
		effects: []Effect{
			GainMoney{Amount: 1},
			GainStatIfXPlayed{AmountDamage: 1, Cardtype: "item"},
			RavenclawDice{},
		},
		houseDice: true,
	}
}

func nimbusTwoThousandAndOne() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Nimbus Two Thousand And One",
		SetId:    "game 2",
		ImgPath:  "/images/marketcards/nimbustwothousandandone.jpg",
		CardType: "item",
		Cost:     5,
		effects: []Effect{
			GainDamage{Amount: 2},
			MoneyIfVillainKilled{Amount: 2, Id: id},
		},
	}
}

func nymphadoraTonks() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Nymphadora Tonks",
		SetId:    "game 5",
		ImgPath:  "/images/marketcards/nymphadoratonks.jpg",
		CardType: "ally",
		Cost:     5,
		effects: []Effect{
			ChooseOne{
				Effects: []Effect{
					GainMoney{Amount: 3},
					GainDamage{Amount: 2},
					RemoveFromLocation{Amount: 1},
				},
				Options:     []string{"Gain 3 Money", "Gain 2 Damage", "Remove 1 from location"},
				Description: "Choose one:",
			},
		},
	}
}

func polyjuicePotion() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Polyjuice Potion",
		SetId:    "game 2",
		ImgPath:  "/images/marketcards/polyjuicepotion.jpg",
		CardType: "item",
		Cost:     3,
		effects: []Effect{
			PolyjuiceEffect{Id: id},
		},
	}
}

// active player selects an ally they've played an copies it's effect.
type PolyjuiceEffect struct {
	Id int
}

func (effect PolyjuiceEffect) Log(gs *Gamestate) {
	gs.EffectLog = append(gs.EffectLog, reflect.Type.Name(reflect.TypeOf(effect)))
}

func (effect PolyjuiceEffect) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn

	playArea := gs.Players[user].PlayArea
	alliesPlayed := []Card{}
	for _, c := range playArea {
		if c.CardType == "ally" {
			alliesPlayed = append(alliesPlayed, c)
		}
	}

	if len(alliesPlayed) == 0 {
		return
	}

	choice := AskUserToSelectCard(user, gs.gameid, alliesPlayed, "Polyjuice: Select an Ally to copy")

	for _, c := range playArea {
		if c.Id == choice {
			for _, e := range c.effects {
				e.Trigger(gs)
			}
		}
	}
}

func pomonaSprout() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Pomona Sprout",
		SetId:    "game 4",
		ImgPath:  "/images/marketcards/pomonasprout.jpg",
		CardType: "ally",
		Cost:     6,
		effects: []Effect{
			GainMoney{Amount: 1},
			HealAnyPlayer{Amount: 2},
			HufflepuffDice{},
		},
		houseDice: true,
	}
}

func severusSnape() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Severus Snape",
		SetId:    "game 4",
		ImgPath:  "/images/marketcards/severussnape.jpg",
		CardType: "ally",
		Cost:     6,
		effects: []Effect{
			GainDamage{Amount: 1},
			GainHealth{Amount: 2},
			SlytherinDice{},
		},
		houseDice: true,
	}
}

func swordOfGryffindor() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Sword of Gryffindor",
		SetId:    "game 7",
		ImgPath:  "/images/marketcards/swordofgryffindor.jpg",
		CardType: "item",
		Cost:     7,
		effects: []Effect{
			GainDamage{Amount: 2},
			GryffindorDice{},
			GryffindorDice{},
		},
		houseDice: true,
	}
}

func tergeo() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Tergeo",
		SetId:    "box 1",
		ImgPath:  "/images/marketcards/tergeo.jpg",
		CardType: "spell",
		Cost:     2,
		effects: []Effect{
			GainMoney{Amount: 1},
			ActivePlayerBanishAndGainXIfY{
				Hand:     true,
				CardType: "item",
				GainX:    DrawCards{Amount: 1},
			},
		},
	}
}

func viktorKrum() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Viktor Krum",
		SetId:    "game 4",
		ImgPath:  "/images/marketcards/viktorkrum.jpg",
		CardType: "ally",
		Cost:     5,
		effects: []Effect{
			GainDamage{Amount: 2},
			GainXIfVillainKilled{
				Id:    id,
				GainX: ChangeStats{AmountHealth: 1, AmountMoney: 1},
			},
		},
	}
}

func fawkesThePhoenix() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Fawkes the Phoenix",
		SetId:    "game 2",
		ImgPath:  "/images/marketcards/fawkesthephoenix.jpg",
		CardType: "ally",
		Cost:     5,
		effects: []Effect{
			ChooseOne{
				Effects: []Effect{
					GainDamage{Amount: 2},
					AllPlayersGainHealth{Amount: 2},
				},
				Options:     []string{"Gain 2 Damage", "All players gain 2 Health"},
				Description: "Fawkes brings you a sword and some tears, choose one:",
			},
		},
	}
}

func minervaMcgonagall() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Minerva McGonagall",
		SetId:    "game 4",
		ImgPath:  "/images/marketcards/minervamcgonagall.jpg",
		CardType: "ally",
		Cost:     6,
		effects: []Effect{
			GainMoney{Amount: 1},
			GainDamage{Amount: 1},
			GryffindorDice{},
		},
		houseDice: true,
	}
}

func remusLupin() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Remus Lupin",
		SetId:    "game 3",
		ImgPath:  "/images/marketcards/remuslupin.jpg",
		CardType: "ally",
		Cost:     4,
		effects: []Effect{
			GainDamage{Amount: 1},
			HealAnyPlayer{Amount: 3},
		},
	}
}

func elderWand() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Elder Wand",
		SetId:    "game 7",
		ImgPath:  "/images/marketcards/elderwand.jpg",
		CardType: "item",
		Cost:     7,
		effects: []Effect{
			GainXPerSpellPlayed{
				X: ChangeStats{
					AmountDamage: 1,
					AmountHealth: 1,
				},
			},
		},
	}
}

func chocolateFrog() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Chocolate Frog",
		SetId:    "game 3",
		ImgPath:  "/images/marketcards/chocolatefrog.jpg",
		CardType: "item",
		Cost:     2,
		effects: []Effect{
			ChoosePlayerToGainX{
				X: ChangeStats{
					AmountHealth: 1,
					AmountMoney:  1,
				},
			},
		},
		onDiscard: func(target string, gs *Gamestate) {
			ChangeStats{
				Target:       target,
				AmountHealth: 1,
				AmountMoney:  1,
			}.Trigger(gs)
		},
	}
}

func gilderoyLockhart() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Gilderoy Lockhart",
		SetId:    "game 2",
		ImgPath:  "/images/marketcards/gilderoylockhart.jpg",
		CardType: "ally",
		Cost:     2,
		effects: []Effect{
			DrawCards{Amount: 1},
			ActivePlayerDiscards{Amount: 1, Prompt: "You swoon at the sight of Lockhart, discard a card"},
		},
		onDiscard: func(target string, gs *Gamestate) {
			ChangeStats{
				Target:      target,
				AmountCards: 1,
			}.Trigger(gs)
		},
	}
}

func maraudersMap() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Marauder's Map",
		SetId:    "game 3",
		ImgPath:  "/images/marketcards/maraudersmap.jpg",
		CardType: "item",
		Cost:     5,
		effects: []Effect{
			DrawCards{Amount: 2},
		},
		onDiscard: func(target string, gs *Gamestate) {
			AllDrawCards{Amount: 1}.Trigger(gs)
		},
	}
}

func protego() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Protego!",
		SetId:    "game 4",
		ImgPath:  "/images/marketcards/protego.jpg",
		CardType: "spell",
		Cost:     5,
		effects: []Effect{
			GainDamage{Amount: 1},
			GainHealth{Amount: 1},
		},
		onDiscard: func(target string, gs *Gamestate) {
			ChangeStats{
				Target:       target,
				AmountDamage: 1,
				AmountHealth: 1,
			}.Trigger(gs)
		},
	}
}

func accio() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Accio!",
		SetId:    "game 4",
		ImgPath:  "/images/marketcards/accio.jpg",
		CardType: "spell",
		Cost:     4,
		effects: []Effect{
			ChooseOne{
				Effects: []Effect{
					GainMoney{Amount: 2},
					ActivePlayerSearchesDiscardForX{CardType: "item"},
				},
				Options:     []string{"Gain 2 Money", "Search discard for an item"},
				Description: "Accio! Choose one:",
			},
		},
	}
}

func fredWeasley() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Fred Weasley",
		SetId:    "game 5",
		ImgPath:  "/images/marketcards/fredweasley.jpg",
		CardType: "ally",
		Cost:     4,
		effects: []Effect{
			GainDamage{Amount: 1},
			WeasleyTwinsEffect{Money: 1},
			GryffindorDice{},
		},
		houseDice: true,
	}
}

func georgeWeasley() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "George Weasley",
		SetId:    "game 5",
		ImgPath:  "/images/marketcards/georgeweasley.jpg",
		CardType: "ally",
		Cost:     4,
		effects: []Effect{
			GainDamage{Amount: 1},
			WeasleyTwinsEffect{Health: 1},
			GryffindorDice{},
		},
		houseDice: true,
	}
}

var WeasleyNames = map[string]bool{"Fred Weasley": true, "George Weasley": true, "Molly Weasley": true, "Ginny Weasley": true, "Arthur Weasley": true}

type WeasleyTwinsEffect struct {
	Money  int
	Health int
}

func (effect WeasleyTwinsEffect) Log(gs *Gamestate) {
	gs.EffectLog = append(gs.EffectLog, reflect.Type.Name(reflect.TypeOf(effect)))
}

func (effect WeasleyTwinsEffect) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn

	for name := range gs.Players {
		if name != user {
			for _, c := range gs.Players[name].Hand {
				_, ok := WeasleyNames[c.Name]
				if ok {
					AllPlayersGainHealth{Amount: effect.Health}.Trigger(gs)
					AllPlayersGainMoney{Amount: effect.Money}.Trigger(gs)
					return
				}
			}
		}
	}
}

func oldSock() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Old Sock",
		SetId:    "box 1",
		ImgPath:  "/images/marketcards/oldsock.jpg",
		CardType: "item",
		Cost:     1,
		effects: []Effect{
			GainMoney{Amount: 1},
			OldSockEffect{},
		},
		onDiscard: func(target string, gs *Gamestate) {
			ChangeStats{
				Target:      target,
				AmountMoney: 2,
			}.Trigger(gs)
		},
	}
}

type OldSockEffect struct{}

func (effect OldSockEffect) Log(gs *Gamestate) {
	gs.EffectLog = append(gs.EffectLog, reflect.Type.Name(reflect.TypeOf(effect)))
}

func (effect OldSockEffect) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	matchingName := "Dobby the House Elf"

	for name := range gs.Players {
		if name != user {
			for _, c := range gs.Players[name].Hand {
				if c.Name == matchingName {
					GainDamage{Amount: 2}.Trigger(gs)
				}
			}
		}
	}
}

func owls() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "O.W.L.S",
		SetId:    "game 5",
		ImgPath:  "/images/marketcards/owls.jpg",
		CardType: "item",
		Cost:     4,
		effects: []Effect{
			GainMoney{Amount: 2},
			GainXIfYSpellsPlayed{
				X: ChangeStats{
					AmountDamage: 1,
					AmountHealth: 1,
				},
				Y: 2,
			},
		},
	}
}

func sortingHat() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Sorting Hat",
		SetId:    "game 1",
		ImgPath:  "/images/marketcards/sortinghat.jpg",
		CardType: "item",
		Cost:     4,
		effects: []Effect{
			GainMoney{Amount: 2},
			PurchasedXGoToDeck{X: "ally"},
		},
	}
}

func wingardiumLeviosa() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Wingardium Leviosa",
		SetId:    "game 1",
		ImgPath:  "/images/marketcards/wingardiumleviosa.jpg",
		CardType: "spell",
		Cost:     2,
		effects: []Effect{
			GainMoney{Amount: 1},
			PurchasedXGoToDeck{X: "item"},
		},
	}
}

func petrificusTotalus() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Petrificus Totalus!",
		SetId:    "game 3",
		ImgPath:  "/images/marketcards/petrificustotalus.jpg",
		CardType: "spell",
		Cost:     6,
		effects: []Effect{
			GainDamage{Amount: 1},
			BlockVillainEffects{villain: true},
		},
	}
}

func harp() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Harp",
		SetId:    "box 1",
		ImgPath:  "/images/marketcards/harp.jpg",
		CardType: "item",
		Cost:     6,
		effects: []Effect{
			GainDamage{Amount: 1},
			BlockVillainEffects{creature: true},
		},
	}
}

func finiteIncantatem() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Finite Incantatem!",
		SetId:    "box 1",
		ImgPath:  "/images/marketcards/finiteincantatem.jpg",
		CardType: "spell",
		Cost:     6,
		effects: []Effect{
			RemoveFromLocation{Amount: 1},
		},
	}
}

func confundus() Card {
	id := int(uuid.New().ID())
	return Card{
		Id:       id,
		Name:     "Confundus!",
		SetId:    "game 6",
		ImgPath:  "/images/marketcards/confundus.jpg",
		CardType: "spell",
		Cost:     3,
		effects: []Effect{
			GainDamage{Amount: 1},
			ConfundusEffect{Id: id},
		},
	}
}

type ConfundusEffect struct{ Id int }

func (effect ConfundusEffect) Log(gs *Gamestate) {
	gs.EffectLog = append(gs.EffectLog, reflect.Type.Name(reflect.TypeOf(effect)))
}

func (effect ConfundusEffect) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	confund := RemoveFromLocation{Amount: 1}

	numVillains := 0
	for _, v := range gs.Villains {
		if v.Active && v.Name != "Norbert" {
			numVillains++
		}
	}
	if len(gs.turnStats.VillainsHit) >= numVillains {
		confund.Trigger(gs)
		return
	}

	sub := Subscriber{
		id:              effect.Id,
		messageChan:     make(chan string),
		conditionMet:    "new villain hit",
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
				if len(gs.turnStats.VillainsHit) >= numVillains {
					confund.Trigger(gs)
					SendLobbyUpdate(gs.gameid, gs)
					gs.mu.Unlock()
					return
				}
			}
			gs.mu.Unlock()
		}
	}()
}
