package bot

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

func GetUpdates(token string, offset int32) (result []Update, err error) {
	url := BaseUrl + token + "/getUpdates?offset=" + strconv.Itoa(int(offset))
	resp, errGet := http.Get(url)
	if errGet != nil {
		log.Printf("Error caught while taking updates for bot: %s\n%s\n", token, errGet.Error())
		return nil, errGet
	}

	// wait until conection and all transactions closed
	defer resp.Body.Close()

	var updates Answer
	errRead := json.NewDecoder(resp.Body).Decode(&updates)

	if errRead != nil {
		log.Printf("Error caught while taking updates for bot: %s\n%s\n", token, errRead.Error())
		return nil, errRead
	}

	if updates.Success {
		return updates.Result, nil
	}

	return nil, errors.New("Telegram service error")
}
