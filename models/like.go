package models

import "gopkg.in/validator.v2"

type Like struct {
	ID       uint `json:"id" gorm:"primaryKey;autoIncrement"`
	ReviewId uint `json:"review_id" validate:"nonzero"`
	UserId   uint `json:"user_id" validate:"nonzero"`
}

func LikeValidator(like *Like) error {
	if err := validator.Validate(like); err != nil {
		return err
	}

	return nil
}
