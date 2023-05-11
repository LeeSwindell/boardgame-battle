package game

func greatHall() Location {
	return Location{
		Name:       "Great Hall",
		Id:         3,
		SetId:      "Box 1",
		ImgPath:    "/images/locations/greathall.jpg",
		MaxControl: 7,
		CurControl: 0,
		Effect:     RevealDarkArts{Amount: 3},
	}
}
