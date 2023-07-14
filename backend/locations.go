package main

import (
	"github.com/google/uuid"
)

type RevealDarkArts struct {
	Amount int
}

func (effect RevealDarkArts) Trigger(gs *Gamestate) {
	user := gs.CurrentTurn
	for _, c := range gs.Players[user].Hand {
		if c.Name == "Finite Incantatem!" {
			if len(gs.DarkArtsPlayed) == 0 {
				curDarkArtIndex := gs.CurrentDarkArt

				// Play current dark art.
				curDarkArt := gs.DarkArts[curDarkArtIndex]
				Logger("playing dark art: " + curDarkArt.Name)
				gs.DarkArtsPlayed = append(gs.DarkArtsPlayed, curDarkArt)
				for _, e := range curDarkArt.effect {
					e.Trigger(gs)
				}

				LoadNewDarkArt(gs)
			}
			return
		}
	}

	for i := 0; i < effect.Amount; i++ {
		curDarkArtIndex := gs.CurrentDarkArt

		// Play current dark art.
		curDarkArt := gs.DarkArts[curDarkArtIndex]
		Logger("playing dark art: " + curDarkArt.Name)
		gs.DarkArtsPlayed = append(gs.DarkArtsPlayed, curDarkArt)
		for _, e := range curDarkArt.effect {
			e.Trigger(gs)
		}

		// Load next
		// newIndex := (curDarkArtIndex + 1) % len(gs.DarkArts)
		// if newIndex == 0 {
		// 	gs.DarkArts = ShuffleDarkArts(gs.DarkArts)
		// }

		// gs.CurrentDarkArt = newIndex
		LoadNewDarkArt(gs)
	}
}

func LoadNewDarkArt(gs *Gamestate) {
	curDarkArtIndex := gs.CurrentDarkArt
	newIndex := (curDarkArtIndex + 1) % len(gs.DarkArts)
	if newIndex == 0 {
		gs.DarkArts = ShuffleDarkArts(gs.DarkArts)
	}

	gs.CurrentDarkArt = newIndex
}

func greatHall() Location {
	return Location{
		Name:       "Great Hall",
		Id:         int(uuid.New().ID()),
		SetId:      "Box 1",
		ImgPath:    "/images/locations/greathall.jpg",
		MaxControl: 7,
		CurControl: 0,
		effect:     RevealDarkArts{Amount: 3},
	}
}

func hagridsHut() Location {
	return Location{
		Name:       "Hagrid's Hut",
		Id:         int(uuid.New().ID()),
		SetId:      "Box 1",
		ImgPath:    "/images/locations/hagridshut.jpg",
		MaxControl: 6,
		CurControl: 0,
		effect:     RevealDarkArts{Amount: 2},
	}
}

func castleGates() Location {
	return Location{
		Name:       "Great Hall",
		Id:         int(uuid.New().ID()),
		SetId:      "Box 1",
		ImgPath:    "/images/locations/castlegates.jpg",
		MaxControl: 5,
		CurControl: 0,
		effect:     RevealDarkArts{Amount: 1},
	}
}
