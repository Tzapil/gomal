package main

import (
    "fmt"
    "log"
    "time"
    "strings"
    "io/ioutil"
    "net/http"
    x "github.com/tzapil/gomal/xml"
    b "github.com/tzapil/gomal/bot"
	"encoding/json"
)

const bot_id = "285037035:AAEXfDvpfvAgpRaRKjBxGIQSwZU9Vn_sP5c"

const myanimelist_api = "http://myanimelist.net/malappinfo.php"

func get_anime(user string) *x.Xml {
    params := []string{"u=" + user, "status=all", "type=anime"}
    resp, err_get := http.Get(myanimelist_api + "?" + strings.Join(params, "&"))
    if err_get != nil {
        log.Printf("Error caught while taking page for user: %s\n%s\n", user, err_get.Error())
        return nil
    }

    // wait until conection and all transactions closed
    defer resp.Body.Close()

    body, err_read := ioutil.ReadAll(resp.Body)

    if (err_read != nil) {
        log.Printf("Error caught while taking page for user: %s\n%s\n", user, err_read.Error())
        return nil
    }

    return x.Parse(body)
}

func getMaxOffset(updates []b.Update) int32 {
    var result int32 = 0

    for i := 0; i < len(updates); i++ {
        if updates[i].Id > result {
            result = updates[i].Id
        }
    }

    return result
}

func AnswerUpdate(u b.Update) {
    if u.InlineQuery != nil {
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

        animes := get_anime(user)
        
        s := []b.InlineQueryResultArticle{}
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
                "https://myanimelist.net/anime/" + a.Id
            
            s = append(s, *b.CreateResultArticle(a.Title, message, "", "HTML"))

            // cant return more than 10 results
            if len(s) >= 10 {
                break
            }
        }

        b.SendAnswerInlineQuery(bot_id, u.InlineQuery.Id, s)
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
    // animes := get_anime("Tzapil")
    // if animes != nil {
    //     // TODO handle
    //     fmt.Println(animes)
    // }
    // r := b.CreateResultArticle("Title", "Message", "Description","HTML")
    // r1 := []b.InlineQueryResultArticle{*r}
    // r2 := b.CreateAnswerInlineQuery("cococo", r1)
    // fmt.Println(r2)

    // setup webhook
    b.SetWebhook(bot_id, "")
    time.Sleep(time.Second)
    b.SetWebhook(bot_id, "https://tzapil.tk:8443/mal/")

    fmt.Println("SETUP")

    // setup webhooks
    http.HandleFunc("/", handler)
    http.ListenAndServe(":7773", nil)

    // var maxOffset int32 = -1
    // for {
    //     updates, err := b.GetUpdates(bot_id, maxOffset + 1)
    //     if err != nil {
    //         log.Printf("Error caught while taking updates\n%s\n", err.Error())
    //         continue
    //     }


    //     // run answers in goroutine
    //     for i := 0; i < len(updates); i++ {
    //         go AnswerUpdate(updates[i])
    //     }

    //     // update maxOffset
    //     next := getMaxOffset(updates)
    //     if next > maxOffset {
    //         maxOffset = next
    //     }

    //     // wait one second
    //     time.Sleep(time.Second)
    // }

    // Create handlers
    // handler := func (w http.ResponseWriter, r *http.Request) {
    //     anime_id := r.FormValue("id")
    //     log.Printf("Request for anime with id %s\n", anime_id)
    //     if anime_id == "" {
    //         http.Error(w, "Anime id parameter required!", http.StatusBadRequest)
    //         return
    //     }

    //     res := ParseSite(anime_id)

    //     if res == nil {
    //         http.Error(w, "Anime not found!", http.StatusNotFound)
    //         return   
    //     }

    //     j, err := json.Marshal(res)
    //     if err != nil {
    //         http.Error(w, err.Error(), http.StatusInternalServerError)
    //         return   
    //     }

    //     w.Header().Set("Content-Type", "application/json; charset=utf-8")
    //     io.WriteString(w, string(j))
    // }
    // http.HandleFunc("/anime", handler)

    // // Start web server
    // log.Println("About to listen on 8080. Go to http://127.0.0.1:8080/anime?id=123")
    // log.Fatal(http.ListenAndServe(":8080", nil))
}
