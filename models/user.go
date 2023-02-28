package models

import (
	"gopkg.in/validator.v2"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" validate:"nonzero"`
	UserName  string `json:"user_name" validate:"nonzero"`
	Password  string `json:"password" validate:"nonzero"`
	Mail      string `json:"mail" validate:"nonzero"`
	BirthDate string `json:"birth_date" validate:"nonzero"`
}

func UserValidator(user *User) error {
	if err := validator.Validate(user); err != nil {
		return err
	}

	return nil
}
