package infra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Post-and-Play/gears/models"
)

type HttpClient struct {
	Client  *http.Client
	BaseUrl string
}

func NewHttpClient(url string) *HttpClient {
	return &HttpClient{
		BaseUrl: url,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (h *HttpClient) EdwigesPost(mail string) (string, error) {
	payload, err := json.Marshal(mail)
	if err != nil {
		log.Default().Printf("Marshal error: %+v", err)
		return "", err
	}

	var edwigesResponse models.Edwiges

	endpoint := fmt.Sprintf("%s/api/mail", h.BaseUrl)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
	if err != nil {
		log.Default().Printf("Request build error: %+v", err)
		return "", err
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		log.Default().Printf("Request do error: %+v", err)
		return "", err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&edwigesResponse)

	if resp.StatusCode != 200 {
		log.Default().Printf("Mail error: %+v", err)
		return "", err
	}

	return edwigesResponse.Mail, nil
}
