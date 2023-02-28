package models

import "gopkg.in/validator.v2"

type Game struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	Name            string `json:"name" validate:"nonzero"`
	MainGender      string `json:"main_gender" validate:"nonzero"`
	SecondGender    string `json:"second_gender" validate:"nonzero"`
	ThirdGender     string `json:"third_gender" validate:"nonzero"`
	CoverAdr        string `json:"cover_adr" validate:"nonzero"`
	BackCoverAdr    string `json:"back_cover_adr" validate:"nonzero"`
	VideoAdr        string `json:"video_adr" validate:"nonzero"`
	Rating          string `json:"rating" validate:"nonzero"`
	FavoriteCount   string `json:"favorite_count" validate:"nonzero"`
	RankingPosition string `json:"ranking_position" validate:"nonzero"`
}

func GameValidator(game *Game) error {
	if err := validator.Validate(game); err != nil {
		return err
	}

	return nil
}
