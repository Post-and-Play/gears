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
	SecurityKey string `json:"security_key"`
}

func AdminValidator(admin *Admin) error {
	if err := validator.Validate(admin); err != nil {
		return err
	}

	return nil
}


type EditAdmin struct {
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
	Password    string `json:"password"`
}


func EditAdminPasswordValidator(admin *EditAdminPassword) error {
	if err := validator.Validate(admin); err != nil {
		return err
	}

	return nil
}


type ForgotAdmin struct {
	ID          uint   `json:"mail" validate:"nonzero"`
}

type RecoverPasswordAdmin struct {
	Password    string `json:"password"`
	SecurityKey string `json:"security_key"`
}

func RecoverPasswordAdminValidator(admin *RecoverPasswordAdmin) error {
	if err := validator.Validate(admin); err != nil {
		return err
	}
	return nil
}