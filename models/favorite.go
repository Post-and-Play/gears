package models

import "gopkg.in/validator.v2"

type Favorite struct {
	Id       uint `json:"id" gorm:"primaryKey;autoIncrement"`
	GameId   uint `json:"game_id" validate:"nonzero" gorm:"not null"`
	UserId   uint `json:"user_id" validate:"nonzero" gorm:"not null"`
}

func FavoriteValidator(favorite *Favorite) error {
	if err := validator.Validate(favorite); err != nil {
		return err
	}

	return nil
}

type FavoriteGame struct {
	Id       uint `json:"id" gorm:"primaryKey;autoIncrement"`
	GameId   uint `json:"game_id" validate:"nonzero" gorm:"not null"`
	UserId   uint `json:"user_id" validate:"nonzero" gorm:"not null"`
	Name        string  `json:"name" validate:"nonzero" gorm:"not null"`
	Genders     string  `json:"genders" validate:"nonzero" gorm:"not null"`
	Description string  `json:"description" validate:"nonzero" gorm:"not null"`
	CoverAdr    string  `json:"cover_adr" validate:"nonzero" gorm:"not null"`
	TopAdr      string  `json:"top_adr" validate:"nonzero" gorm:"not null"`
}

func FavoriteGameValidator(favorite *FavoriteGame) error {
	if err := validator.Validate(favorite); err != nil {
		return err
	}

	return nil
}