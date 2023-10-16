package models

import "gopkg.in/validator.v2"

type Follow struct {
	Id              uint `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	FollowingUserId uint `json:"following_user_id" validate:"nonzero" gorm:"not null;foreignKey:users.id,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	FollowedUserId  uint `json:"followed_user_id" validate:"nonzero" gorm:"not null;foreignKey:users.id,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func FollowValidator(follow *Follow) error {
	if err := validator.Validate(follow); err != nil {
		return err
	}

	return nil
}
