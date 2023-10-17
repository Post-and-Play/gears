package models

import "gopkg.in/validator.v2"

type Review struct {
	Id       uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId   uint    `json:"user_id" validate:"nonzero" gorm:"not null"`
	GameId   uint    `json:"game_id" validate:"nonzero" gorm:"not null"`
	Grade    float64 `json:"grade" validate:"nonzero" gorm:"not null"`
	ImageAdr string  `json:"image_adr"`
	Opinion  string  `json:"opinion" validate:"nonzero" gorm:"not null"`
	Likes    int     `json:"likes" gorm:"default:0"`
}

func ReviewValidator(review *Review) error {
	if err := validator.Validate(review); err != nil {
		return err
	}

	return nil
}

type ReviewUser struct {
	Id       uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId   uint    `json:"user_id" validate:"nonzero" gorm:"not null"`
	GameId   uint    `json:"game_id" validate:"nonzero" gorm:"not null"`
	Grade    float64 `json:"grade" validate:"nonzero" gorm:"not null"`
	ImageAdr string  `json:"image_adr"`
	Opinion  string  `json:"opinion" validate:"nonzero" gorm:"not null"`
	Likes    int     `json:"likes" gorm:"default:0"`
	Name     string  `json:"name" validate:"nonzero"`
	PhotoAdr string  `json:"photo_adr"`
}
