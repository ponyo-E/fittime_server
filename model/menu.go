package model

import "github.com/google/uuid"

type ExercisePart struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	ExerciseID uuid.UUID `gorm:"size:30"`
	Exercise   Exercise
	MenuID     uuid.UUID `gorm:"size:30"`
	Menu       Menu
	No         int
	Time       int
}

type Menu struct {
	ID            uuid.UUID `gorm:"primaryKey"`
	Title         string
	UserID        uuid.UUID `gorm:"size:30"`
	User          User
	Body          string
	Path          string
	Nice          int
	Point         int
	ExerciseParts []ExercisePart
}

func SearchMenu(word string) ([]Menu, error) {
	var menus []Menu
	err := db.Where("body LIKE ?", "%"+word+"%").Find(&menus).Error
	if err != nil {
		return nil, err
	}

	if len(menus) > 5 {
		menus = menus[:5]
	}

	return menus, nil
}

func AddMenu(menu Menu) error {
	err := db.Create(&menu).Error
	if err != nil {
		return err
	}
	err = db.Create(&menu.ExerciseParts).Error
	if err != nil {
		return err
	}
	return nil
}
