package bot

import (
	"strconv"
	"errors"
	"log"
	"net/http"
	"encoding/json"
)

func GetUpdates(token string, offset int32) (result []Update, err error) {
	url := BaseUrl + token + "/getUpdates?offset=" + strconv.Itoa(int(offset))
	resp, err_get := http.Get(url)
    if err_get != nil {
        log.Printf("Error caught while taking updates for bot: %s\n%s\n", token, err_get.Error())
        return nil, err_get
    }

    // wait until conection and all transactions closed
    defer resp.Body.Close()

	var updates Answer
    err_read := json.NewDecoder(resp.Body).Decode(&updates)

    if (err_read != nil) {
        log.Printf("Error caught while taking updates for bot: %s\n%s\n", token, err_read.Error())
        return nil, err_read
	}

	if updates.Success {
		return updates.Result, nil
	}
	
	return nil, errors.New("Telegram service error")
}