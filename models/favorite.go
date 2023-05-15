package models

import "gopkg.in/validator.v2"

type Favorite struct {
	ID              uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string `json:"name" validate:"nonzero"`
	Genders         string `json:"genders" validate:"nonzero"`
	CoverAdr        string `json:"cover_adr" validate:"nonzero"`
	Rating          string `json:"rating" validate:"nonzero"`
	FavoriteCount   string `json:"favorite_count" validate:"nonzero"`
	RankingPosition string `json:"ranking_position" validate:"nonzero"`
}

func FavoriteValidator(favorite *Favorite) error {
	if err := validator.Validate(favorite); err != nil {
		return err
	}

	return nil
}
