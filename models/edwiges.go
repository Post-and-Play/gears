package models

import "gopkg.in/validator.v2"

type Edwiges struct {
	Mail string `json:"mail" validate:"nonzero"`
}

func EdwigesValidator(edwiges *Edwiges) error {
	if err := validator.Validate(edwiges); err != nil {
		return err
	}

	return nil
}
