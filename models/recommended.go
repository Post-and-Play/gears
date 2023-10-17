package models

import "gopkg.in/validator.v2"

type Recommended struct {
	Id          uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId      uint    `json:"user_id" validate:"nonzero" gorm:"not null"`
	Name        string  `json:"name" validate:"nonzero" gorm:"not null"`
	Genders     string  `json:"genders" validate:"nonzero" gorm:"not null"`
	Description string  `json:"description" validate:"nonzero" gorm:"not null"`
	Creator     string  `json:"creator" validate:"nonzero" gorm:"not null"`
	IsFree      bool    `json:"is_free" gorm:"not null"`
	Approved    bool    `json:"approved" gorm:"not null`
	CoverAdr    string  `json:"cover_adr" validate:"nonzero" gorm:"not null"`
	TopAdr      string  `json:"top_adr" validate:"nonzero" gorm:"not null"`
}

func RecommendedValidator(game *Recommended) error {
	if err := validator.Validate(game); err != nil {
		return err
	}

	return nil
}
