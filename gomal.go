package main

import (
    "log"
    "fmt"
    "time"
    // "bytes"
    "strings"
    // "io"
    "io/ioutil"
    "net/http"

    // "encoding/xml"

    // "golang.org/x/net/html"

    // e "github.com/tzapil/gomal/entry"
    x "github.com/tzapil/gomal/xml"
    b "github.com/tzapil/gomal/bot"

    // "golang.org/x/text/encoding/charmap"
    // "golang.org/x/text/transform"

    // "encoding/json"
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

// func ParseSite(id string) *e.Entry {
//     log.Printf("Try to get page with id: %s\n", id)

//     answer, err_get := http.Get(site_adress + "base.php?id=" + id)
//     if err_get != nil {
//         log.Printf("Error caught while taking page with id: %s\n%s\n", id, err_get.Error())
//         return nil
//     }

//     defer answer.Body.Close()

//     body, err_read := ioutil.ReadAll(answer.Body)
//     if err_read != nil {
//         log.Printf("Error caught while reading body with id: %s\n%s\n", id, err_read.Error())
//         return nil
//     }

//     log.Printf("Try to parse page with id: %s\n", id)
//     z, err_parse := html.Parse(bytes.NewBuffer(Win1251ToUtf8(body)))

//     if err_parse != nil {
//         log.Printf("Error caught while parse body with id: %s\n%s\n", id, err_parse.Error())
//         return nil
//     }
  
//     log.Printf("Retrun information from page with id: %s\n", id)
//     return e.Parse(z)
// }

func getMaxOffset(updates []b.Update) int32 {
    var result int32 = 0

    for i := 0; i < len(updates); i++ {
        if updates[i].Id > result {
            result = updates[i].Id
        }
    }

    return result
}

func main() {
    animes := get_anime("Tzapil")
    if animes != nil {
        // TODO handle
        fmt.Println(animes)
    }
    r := b.CreateResultArticle("Title", "Message", "Description","HTML")
    r1 := []b.InlineQueryResultArticle{*r}
    r2 := b.CreateAnswerInlineQuery("cococo", r1)
    fmt.Println(r2)

    var maxOffset int32 = -1
    for {
        updates, err := b.GetUpdates(bot_id, maxOffset + 1)
        if err != nil {
            log.Printf("Error caught while taking updates\n%s\n", err.Error())
            continue
        }
        fmt.Println("UPDATED")
        fmt.Println(len(updates))

        for i := 0; i < len(updates); i++ {
            u := updates[i]
            if u.InlineQuery != nil {
                // parse request
                query := strings.Trim(u.InlineQuery.Query, "")
                parts := strings.Split(query, " ")
                var user string = ""
                var anime string = ""
                switch len(parts) {
                    case 0: 
                        continue
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
                }

                b.SendAnswerInlineQuery(bot_id, u.InlineQuery.Id, s)
            }
        }

        next := getMaxOffset(updates)
        if next > maxOffset {
            maxOffset = next
        }

        time.Sleep(time.Second)
    }

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
