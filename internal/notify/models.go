package notify

type SendMessageRequest struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

type SendMessageResponse struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	Result      Result `json:"result"`
}

type EditMessageRequest struct {
	ChatID    int    `json:"chat_id"`
	Text      string `json:"text"`
	MessageID int    `json:"message_id"`
}

type Result struct {
	MessageID int `json:"message_id"`
}
