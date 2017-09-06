package bot

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

func GetUpdates(token string) {
	url := BaseUrl + token + "/getUpdates"
	resp, err_get := http.Get(url)
    if err_get != nil {
        log.Printf("Error caught while taking updates for bot: %s\n%s\n", token, err_get.Error())
        return
    }

    // wait until conection and all transactions closed
    defer resp.Body.Close()

	var updates Answer
    err_read := json.NewDecoder(resp.Body).Decode(&updates)

    if (err_read != nil) {
        log.Printf("Error caught while taking updates for bot: %s\n%s\n", token, err_read.Error())
        return
	}
	
	fmt.Println(updates)
}