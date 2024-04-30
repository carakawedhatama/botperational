package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SentViaWebhook(docType string, webhookUrl string, content *Discord) error {

	jsonMsg, err := json.Marshal(content)
	if err != nil {
		return err
	}

	_, err = http.Post(webhookUrl, "application/json", bytes.NewBuffer(jsonMsg))
	if err != nil {
		return err
	}

	fmt.Printf("\n%s sent via webhook", docType)

	return nil
}
