package models

import "gopkg.in/validator.v2"

type Recaptcha struct {
	Secret     string    `json:"secret"`
	Response   string    `json:"response"`
	RemoteIp   string    `json:"remoteip"`
}

func RecaptchaValidator(recaptcha *Recaptcha) error {
	if err := validator.Validate(recaptcha); err != nil {
		return err
	}

	return nil
}