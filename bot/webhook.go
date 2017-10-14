package bot

import (
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
)

type Webhook struct {
	Url string `json:"url"`
	Certificate string `json:"certificate,omitempty"`
	MaxConnections int32 `json:"max_connections,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

func SetWebhook(token string, url string) {
	send_url := BaseUrl + token + "/setWebhook"

	s := Webhook{
		Url: url,
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(s)

	fmt.Println(send_url)
	fmt.Println(b)
	
	http.Post(send_url, "application/json; charset=utf-8", b)
}

func RemoveWebhook(token string) {
	SetWebhook(token, "")
}