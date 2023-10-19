package models

import "gopkg.in/validator.v2"

type Game struct {
	Id          uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string  `json:"name" validate:"nonzero" gorm:"not null"`
	Genders     string  `json:"genders" validate:"nonzero" gorm:"not null"`
	Description string  `json:"description" validate:"nonzero" gorm:"not null"`
	CoverAdr    string  `json:"cover_adr" validate:"nonzero" gorm:"not null"`
	TopAdr      string  `json:"top_adr" validate:"nonzero" gorm:"not null"`
	Rating      float64 `json:"rating"`
	Reviews     int     `json:"reviews"`
}

func GameValidator(game *Game) error {
	if err := validator.Validate(game); err != nil {
		return err
	}

	return nil
}


type EditGame struct {
	Name        string  `json:"name" validate:"nonzero" gorm:"not null"`
	Genders     string  `json:"genders" validate:"nonzero" gorm:"not null"`
	Description string  `json:"description" validate:"nonzero" gorm:"not null"`
	CoverAdr    string  `json:"cover_adr" validate:"nonzero" gorm:"not null"`
	TopAdr      string  `json:"top_adr" validate:"nonzero" gorm:"not null"`
	Rating      float64 `json:"rating"`
	Reviews     int     `json:"reviews"`
}

func EditGameValidator(game *EditGame) error {
	if err := validator.Validate(game); err != nil {
		return err
	}

	return nil
}