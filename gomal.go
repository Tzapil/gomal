package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	b "github.com/tzapil/gomal/bot"
	x "github.com/tzapil/gomal/xml"
)

const bot_id = "285037035:AAEXfDvpfvAgpRaRKjBxGIQSwZU9Vn_sP5c"

const myanimelist_api = "http://myanimelist.net/malappinfo.php"

func getAnime(user string) *x.Xml {
	params := []string{"u=" + user, "status=all", "type=anime"}
	resp, errGet := http.Get(myanimelist_api + "?" + strings.Join(params, "&"))
	if errGet != nil {
		log.Printf("Error caught while taking page for user: %s\n%s\n", user, errGet.Error())
		return nil
	}

	// wait until conection and all transactions closed
	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)

	if errRead != nil {
		log.Printf("Error caught while taking page for user: %s\n%s\n", user, errRead.Error())
		return nil
	}

	return x.Parse(body)
}

func getMaxOffset(updates []b.Update) int32 {
	var result int32 = 0

	for i := 0; i < len(updates); i++ {
		if updates[i].ID > result {
			result = updates[i].ID
		}
	}

	return result
}

func AnswerUpdate(u b.Update) {
	if u.InlineQuery != nil {
		// answer
		s := []b.InlineQueryResultArticle{}
		// parse request
		query := strings.Trim(u.InlineQuery.Query, "")
		parts := strings.Split(query, " ")
		var user string = ""
		var anime string = ""
		switch len(parts) {
		case 0:
			return
		case 1:
			user = parts[0]
		default:
			user = parts[0]
			anime = parts[1]
		}

		animes := getAnime(user)
		if animes != nil {
			for j := 0; j < len(animes.Anime); j++ {
				a := animes.Anime[j]
				if !strings.Contains(strings.ToLower(a.Title), strings.ToLower(anime)) {
					continue
				}
				message := "User: " + animes.User.Name + "\n" +
					"<b>" + a.Title + "</b> " +
					a.Watched + "/" + a.Episodes + "\n" +
					"Status: " + a.Status + "\n" +
					"Score: " + a.Score + "\n" +
					"https://myanimelist.net/anime/" + a.ID

				s = append(s, *b.CreateResultArticle(a.Title, message, "", "HTML"))

				// cant return more than 10 results
				if len(s) >= 10 {
					break
				}
			}
		}

		b.SendAnswerInlineQuery(bot_id, u.InlineQuery.ID, s)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Mama, i have request")
	decoder := json.NewDecoder(r.Body)

	var update b.Update
	err := decoder.Decode(&update)
	if err != nil {
		log.Printf("Error caught while parsing updates\n%s\n", err.Error())
		return
	}

	defer r.Body.Close()
	log.Println(update)
	go AnswerUpdate(update)
}

func main() {
	// setup webhook
	b.SetWebhook(bot_id, "")
	time.Sleep(time.Second)
	b.SetWebhook(bot_id, "https://tzapil.tk:8443/mal/")

	fmt.Println("SETUP")

	// setup webhooks
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
