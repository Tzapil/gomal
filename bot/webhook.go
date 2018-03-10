package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Webhook struct {
	URL            string   `json:"url"`
	Certificate    string   `json:"certificate,omitempty"`
	MaxConnections int32    `json:"max_connections,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

func SetWebhook(token string, url string) {
	sendURL := BaseUrl + token + "/setWebhook"

	s := Webhook{
		URL: url,
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(s)

	fmt.Println(sendURL)
	fmt.Println(b)

	http.Post(sendURL, "application/json; charset=utf-8", b)
}

func RemoveWebhook(token string) {
	SetWebhook(token, "")
}
