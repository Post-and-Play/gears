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
