package models

import "gopkg.in/validator.v2"

type Game struct {
	Id       uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string  `json:"name" validate:"nonzero"`
	Genders  string  `json:"genders" validate:"nonzero"`
	CoverAdr string  `json:"cover_adr" validate:"nonzero"`
	TopAdr   string  `json:"top_adr" validate:"nonzero"`
	Rating   float64 `json:"rating"`
	Reviews  int     `json:"reviews"`
}

func GameValidator(game *Game) error {
	if err := validator.Validate(game); err != nil {
		return err
	}

	return nil
}
