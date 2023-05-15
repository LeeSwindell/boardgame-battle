package game

import (
	"github.com/google/uuid"
)

type RevealDarkArts struct {
	Amount int
}

func (effect RevealDarkArts) Trigger(gs *Gamestate) {
	for i := 0; i < effect.Amount; i++ {
		curDarkArtIndex := gs.CurrentDarkArt

		// Play current dark art.
		curDarkArt := gs.DarkArts[curDarkArtIndex]
		for _, e := range curDarkArt.Effects {
			e.Trigger(gs)
		}

		// Load next
		newIndex := (curDarkArtIndex + 1) % len(gs.DarkArts)
		if newIndex == 0 {
			gs.DarkArts = ShuffleDarkArts(gs.DarkArts)
		}

		gs.CurrentDarkArt = newIndex
	}
}

func greatHall() Location {
	return Location{
		Name:       "Great Hall",
		Id:         int(uuid.New().ID()),
		SetId:      "Box 1",
		ImgPath:    "/images/locations/greathall.jpg",
		MaxControl: 7,
		CurControl: 0,
		Effect:     RevealDarkArts{Amount: 3},
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
		Effect:     RevealDarkArts{Amount: 2},
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
		Effect:     RevealDarkArts{Amount: 1},
	}
}
