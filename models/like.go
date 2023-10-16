package models

import "gopkg.in/validator.v2"

type Like struct {
	Id       uint `json:"id" gorm:"primaryKey;autoIncrement"`
	ReviewId uint `json:"review_id" validate:"nonzero" gorm:"not null;foreignKey:reviews.id,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserId   uint `json:"user_id" validate:"nonzero" gorm:"not null;foreignKey:users.id,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func LikeValidator(like *Like) error {
	if err := validator.Validate(like); err != nil {
		return err
	}

	return nil
}
