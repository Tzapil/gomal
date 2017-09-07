package bot

import "github.com/google/uuid"

func CreateResultArticle(title, message_text, description, parse_mode string) *InlineQueryResultArticle {
	input_message_content := &InputTextMessageContent{Text: message_text, ParseMode: parse_mode}

	uuid := uuid.Must(uuid.NewRandom()).String()

	return &InlineQueryResultArticle{
		Id: uuid,
		Type: "article", 
		Title: title, 
		Description: description, 
		InputMessageContent: input_message_content,
	}
}

func CreateAnswerInlineQuery(inline_query_id string, results[] InlineQueryResultArticle) *AnswerInlineQuery {
	return &AnswerInlineQuery{
		InlineQueryId: inline_query_id,
		Results: results,
	}
}