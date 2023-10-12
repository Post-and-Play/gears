package models

import (
	"gopkg.in/validator.v2"
)

type Admin struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" validate:"nonzero"`
	Password    string `json:"password" validate:"nonzero"`
	Mail        string `json:"mail" validate:"nonzero"`
	PhotoAdr    string `json:"photo_adr"`
}

func AdminValidator(admin *Admin) error {
	if err := validator.Validate(admin); err != nil {
		return err
	}

	return nil
}


type EditAdmin struct {
	ID          uint   `json:"id" validate:"nonzero"`
	Name        string `json:"name"`
	PhotoAdr    string `json:"photo_adr"`
}


func EditAdminValidator(admin *EditAdmin) error {
	if err := validator.Validate(admin); err != nil {
		return err
	}

	return nil
}


type EditAdminPassword struct {
	ID          uint   `json:"id" validate:"nonzero"`
	Password    string `json:"password"`
}


func EditAdminPasswordValidator(admin *EditAdminPassword) error {
	if err := validator.Validate(admin); err != nil {
		return err
	}

	return nil
}