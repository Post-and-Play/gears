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
	Subject string `json:"subject" validate:"nonzero"`
	Body	string `json:"body" validate:"nonzero"`
	OK		bool   `json:"ok" gorm:"default:false"`
}

func MailRequestValidator(mail *MailRequest) error {
	if err := validator.Validate(mail); err != nil {
		return err
	}

	return nil
}
