package models

import "gopkg.in/validator.v2"

type Follow struct {
	Id              uint `json:"id" gorm:"primaryKey;autoIncrement"`
	FollowingUserId uint `json:"following_id" validate:"nonzero"`
	FollowedUserId  uint `json:"followed_id" validate:"nonzero"`
}

func FollowValidator(follow *Follow) error {
	if err := validator.Validate(follow); err != nil {
		return err
	}

	return nil
}
