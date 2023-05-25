package models

import "gopkg.in/validator.v2"

type Post struct {
	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId  uint   `json:"user_id" validate:"nonzero"`
	GameId  uint   `json:"game_id" validate:"nonzero"`
	Opinion string `json:"opinion" validate:"nonzero"`
}

func PostValidator(post *Post) error {
	if err := validator.Validate(post); err != nil {
		return err
	}

	return nil
}
