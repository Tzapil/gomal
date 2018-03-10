package xml

import (
	"encoding/xml"
	"log"
)

type XmlAnime struct {
	ID       string `xml:"series_animedb_id"`
	Title    string `xml:"series_title"`
	Synonyms string `xml:"series_synonyms"`
	Type     string `xml:"series_type"`
	Episodes string `xml:"series_episodes"`
	Status   string `xml:"series_status"`
	Watched  string `xml:"my_watched_episodes"`
	Score    string `xml:"my_score"`
	Start    string `xml:"series_start"`
	End      string `xml:"series_end"`
}

type XmlUser struct {
	ID          string `xml:"user_id"`
	Name        string `xml:"user_name"`
	Watching    string `xml:"user_watching"`
	Completed   string `xml:"user_completed"`
	OnHold      string `xml:"user_onhold"`
	Dropped     string `xml:"user_dropped"`
	PlanToWatch string `xml:"user_plantowatch"`
	DaysSpent   string `xml:"user_days_spent_watching"`
}

type Xml struct {
	User  *XmlUser   `xml:"myinfo"`
	Anime []XmlAnime `xml:"anime"`
}

func Parse(body []byte) *Xml {
	v := &Xml{}
	errParse := xml.Unmarshal(body, v)

	if errParse != nil {
		log.Printf("Error caught while parse body:\n%s\n", errParse.Error())
		return nil
	}

	return v
}
