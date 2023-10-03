package models

import (
	"gopkg.in/validator.v2"
)

type User struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" validate:"nonzero"`
	UserName    string `json:"user_name" validate:"nonzero"`
	Password    string `json:"password" validate:"nonzero"`
	Mail        string `json:"mail" validate:"nonzero"`
	BirthDate   string `json:"birth_date" validate:"nonzero"`
	Description string `json:"description"`
	PhotoAdr    string `json:"photo_adr"`
	TopAdr      string `json:"top_adr"`
	Following   int    `json:"following"`
	Followed    int    `json:"followed"`
	EpicUser    string `json:"epic_user"`
	SteamUser   string `json:"steam_user"`
	DiscordUser string `json:"discord_user"`
	GithubUser  string `json:"github_user"`
	TwitchUser  string `json:"twitch_user"`
}

func UserValidator(user *User) error {
	if err := validator.Validate(user); err != nil {
		return err
	}

	return nil
}

/*
type EditUser struct {
	ID          uint   `json:"id" validate:"nonzero"`
	Name        string `json:"name"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	Mail        string `json:"mail"`
	Description string `json:"description"`
	PhotoAdr    string `json:"photo_adr"`
	TopAdr      string `json:"top_adr"`
	EpicUser    string `json:"epic_user"`
	SteamUser   string `json:"steam_user"`
	DiscordUser string `json:"discord_user"`
	GithubUser  string `json:"github_user"`
	TwitchUser  string `json:"twitch_user"`
}
*/

func EditUserValidator(user *EditUser) error {
	if err := validator.Validate(user); err != nil {
		return err
	}

	return nil
}
