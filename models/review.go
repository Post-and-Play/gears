package models

import "gopkg.in/validator.v2"

type Review struct {
	Id       uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId   uint    `json:"user_id" validate:"nonzero"`
	GameId   uint    `json:"game_id" validate:"nonzero"`
	Grade    float64 `json:"grade" validate:"nonzero"`
	ImageAdr string  `json:"image_adr" validate:"nonzero"`
	Opinion  string  `json:"opinion"`
}

func ReviewValidator(review *Review) error {
	if err := validator.Validate(review); err != nil {
		return err
	}

	return nil
}
