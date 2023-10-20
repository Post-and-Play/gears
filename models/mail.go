package models

import "gopkg.in/validator.v2"

type Sender struct {
	SenderMail string
	SenderPass string
}

type Receiver struct {
	ReceiverMail string `json:"receiver" validate:"nonzero"`
}

type MailRequest struct {
	Subject		string `json:"subject" validate:"nonzero"`
	Title		string `json:"title" validate:"nonzero"`
	Message		string `json:"message" validate:"nonzero"`
	Link		string `json:"link" validate:"nonzero"`
	Footer		string `json:"footer" validate:"nonzero"`
	ButtonText  string `json:"button_text" validate:"nonzero"`
	OK			bool   `json:"ok" gorm:"default:false"`
}

func MailRequestValidator(mail *MailRequest) error {
	if err := validator.Validate(mail); err != nil {
		return err
	}

	return nil
}
