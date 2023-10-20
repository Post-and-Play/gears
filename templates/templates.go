package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

func BuildTemplate(subject string) (*template.Template, bytes.Buffer) {
	t, err := template.ParseFiles("templates/template.html")
	if err != nil {
		log.Panicf("Error on parse template")
		return nil, bytes.Buffer{}
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: " + subject + " \n%s\n\n", mimeHeaders)))

	return t, body
}
