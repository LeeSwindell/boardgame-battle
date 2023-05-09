package game

import "github.com/google/uuid"

func draco() Villain {
	return Villain{
		Name:        "Draco Malfoy",
		Id:          int(uuid.New().ID()),
		ImgPath:     "/images/villains/dracomalfoy.jpg",
		SetId:       "Game 1",
		CurDamage:   0,
		MaxHp:       6,
		Effect:      nil,
		DeathEffect: nil,
	}
}

func quirrell() Villain {
	return Villain{
		Name:        "Quirinus Quirrell",
		Id:          int(uuid.New().ID()),
		ImgPath:     "/images/villains/quirrell.jpg",
		SetId:       "Game 1",
		CurDamage:   0,
		MaxHp:       6,
		Effect:      nil,
		DeathEffect: nil,
	}
}

func crabbeAndGoyle() Villain {
	return Villain{
		Name:        "Crabbe and Goyle",
		Id:          int(uuid.New().ID()),
		ImgPath:     "/images/villains/crabbeandgoyle.jpg",
		SetId:       "Game 1",
		CurDamage:   0,
		MaxHp:       5,
		Effect:      nil,
		DeathEffect: nil,
	}
}
