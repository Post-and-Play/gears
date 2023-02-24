package models

import (
	"gopkg.in/validator.v2"
)

type Login struct {
	UserName string `json:"user_name"`
	Mail     string `json:"mail"`
	Password string `json:"password" validate:"nonzero"`
}

func LoginValidator(login *Login) error {
	if err := validator.Validate(login); err != nil {
		return err
	}

	return nil
}
