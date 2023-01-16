package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"name" validate:"nonzero"`
	UserName  string `json:"user_name" validate:"nonzero"`
	Mail      string `json:"mail" validate:"nonzero"`
	BirthDate string `json:"birth_date" validate:"nonzero"`
}

func UserValidator(user *User) error {
	if err := validator.Validate(user); err != nil {
		return err
	}

	return nil
}
