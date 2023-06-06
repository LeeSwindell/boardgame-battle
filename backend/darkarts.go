package main

import "github.com/google/uuid"

func dementorsKiss() DarkArt {
	return DarkArt{
		Name:    "Dementor's Kiss",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/dementorskiss.jpg",
		SetId:   "game 3",
		Effects: []Effect{
			DamageCurrentPlayer{Amount: 2},
			DamageAllPlayersButCurrent{Amount: 1},
		},
	}
}

func flipendo() DarkArt {
	return DarkArt{
		Name:    "Flipendo",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/flipendo.jpg",
		SetId:   "game 1",
		Effects: []Effect{
			DamageCurrentPlayer{Amount: 1},
			ActivePlayerDiscards{Amount: 1},
		},
	}
}

func heWhoMustNotBeNamed() DarkArt {
	return DarkArt{
		Name:    "He-Who-Must-Not-Be-Named",
		Id:      int(uuid.New().ID()),
		ImgPath: "/images/darkarts/hewhomustnotbenamed.jpg",
		SetId:   "game 1",
		Effects: []Effect{
			AddToLocation{Amount: 1},
		},
	}
}
