package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	UserName string `json:"user_name" validate:"nonzero"`
	Mail     string `json:"mail" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
}

func LoginValidator(login *Login) error {
	if err := validator.Validate(login); err != nil {
		return err
	}

	return nil
}
