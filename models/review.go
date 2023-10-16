package models

import "gopkg.in/validator.v2"

type Review struct {
	Id       uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId   uint    `json:"user_id" validate:"nonzero" gorm:"not null;foreignKey:users.id,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	GameId   uint    `json:"game_id" validate:"nonzero" gorm:"not null;foreignKey:games.id,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Grade    float64 `json:"grade" validate:"nonzero" gorm:"not null"`
	ImageAdr string  `json:"image_adr"`
	Opinion  string  `json:"opinion" validate:"nonzero" gorm:"not null"`
	Likes    string  `json:"likes"`
}

func ReviewValidator(review *Review) error {
	if err := validator.Validate(review); err != nil {
		return err
	}

	return nil
}
