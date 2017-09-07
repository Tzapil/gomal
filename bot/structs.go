package bot

const BaseUrl = "https://api.telegram.org/bot"

type Answer struct {
	Success bool `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	Id int32 `json:"update_id"`
	Message *Message `json:"message"`
	InlineQuery *InlineQuery `json:"inline_query"`
}

type User struct {
	Id int32 `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"username"`
	Language string `json:"language_code"`
}

type Chat struct {
	Id int32 `json:"id"`
	Type string `json:"type"`
}

type Message struct {
	Id int32 `json:"id"`
	From *User `json:"from"`
	// unix time
	Date int32 `json:"date"`
	Text string `json:"text"`
	Chat *Chat `json:"chat"`
	ReplyToMessage string `json:"reply_to_message"`
}

type CallbackQuery struct {
	Id string `json:"id"`
	From *User `json:"from"`
	Message *Message `json:"message"`
	InlineMessageId string `json:"inline_message_id"`
	ChatInstance string `json:"chat_instance"`
	Data string `json:"data"`
	GameShortName string `json:"game_short_name"`
}

type InlineQuery struct {
	Id string `json:"id"`
	From *User `json:"from"`
	Query string `json:"query"`
	Offset string `json:"offset"`
}

type AnswerInlineQuery struct {
	InlineQueryId string `json:"inline_query_id"`
	Results []InlineQueryResultArticle `json:"results"`
}

type InputTextMessageContent struct {
	Text string `json:"message_text"`
	ParseMode string `json:"parse_mode"`
}

type InlineQueryResultArticle struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	InputMessageContent *InputTextMessageContent `json:"input_message_content"`
}