package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pedroafonso18/Disparador-GO/internal/config"
)

type MessageRequest_WUZ struct {
	Number string `json:"Phone"`
	Text   string `json:"Body"`
}

func SendMessageWuz(num string, template string, token string) error {
	reqBody := MessageRequest_WUZ{
		Number: num,
		Text:   template,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("error marshaling request body: %v", err)
	}
	url := fmt.Sprintf("%s/chat/send/text", config.WUZURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("token", token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	return nil

}
