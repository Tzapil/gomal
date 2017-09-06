package xml

import (
	"log"
	"encoding/xml"
)

type XmlAnime struct {
    Id string `xml:"series_animedb_id"`
    Title  string `xml:"series_title"`
    Synonyms string `xml:"series_synonyms"`
    Type string `xml:"series_type"`
    Episodes string `xml:"series_episodes"`
    Status string `xml:"series_status"`
    Start string `xml:"series_start"`
    End string `xml:"series_end"`
}

type XmlUser struct {
    Id string `xml:"user_id"`
    Name string `xml:"user_name"`
    Watching string `xml:"user_watching"`
    Completed string `xml:"user_completed"`
    OnHold string `xml:"user_onhold"`
    Dropped string `xml:"user_dropped"`
    PlanToWatch string `xml:"user_plantowatch"`
    DaysSpent string `xml:"user_days_spent_watching"`
}

type Xml struct {
    User XmlUser `xml:"myinfo"`
    Anime []XmlAnime `xml:"anime"`
}

func Parse(body []byte) *Xml {
    v := &Xml{}
    err_parse := xml.Unmarshal(body, v)

    if err_parse != nil {
        log.Printf("Error caught while parse body:\n%s\n", err_parse.Error())
        return nil
    }

    return v
}