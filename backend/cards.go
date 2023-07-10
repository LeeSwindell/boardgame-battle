package main

import "github.com/google/uuid"

func alohamora() Card {
	return Card{
		Id:       int(uuid.New().ID()),
		Name:     "Alohomora",
		SetId:    "game 1",
		ImgPath:  "/images/starters/alohomora.jpg",
		CardType: "spell",
		Cost:     0,
		Effects:  []Effect{GainMoney{Amount: 100}},
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
			GainMoney{Amount: 1},
			GainDamagePerAllyPlayed{},
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
		Effects: []Effect{
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
		Effects:  []Effect{RemoveFromLocation{Amount: 1}},
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects:  []Effect{},
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
			DrawCards{Amount: 3},
			ActivePlayerDiscards{Amount: 2},
			RavenclawDice{},
		},
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
			GainDamage{Amount: 1},
			DrawCards{Amount: 1},
			RavenclawDice{},
		},
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
			DrawCards{Amount: 2},
			SybillDiscard{},
		},
	}
}

type SybillDiscard struct{}

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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
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
		Effects: []Effect{
			GainMoney{Amount: 2},
			GainStatIfXPlayed{AmountHealth: 2, Cardtype: "ally", Id: id},
		},
	}
}
