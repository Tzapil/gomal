package bot

import (
	"net/http"
	"bytes"
	"encoding/json"
)

func SendAnswerInlineQuery(token string, inline_query_id string, results []InlineQueryResultArticle) {
	url := BaseUrl + token + "/answerInlineQuery"

	s := CreateAnswerInlineQuery(inline_query_id, results)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(s)
	
	http.Post(url, "application/json; charset=utf-8", b)
}