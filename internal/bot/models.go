package bot

// ResponseT Полный ответ
type ResponseT struct {
	Ok     bool       `json:"ok"`
	Result []ResultRT `json:"result"`
}

// ResultRT каждый ответ содержит
type ResultRT struct {
	UpdateID int       `json:"update_id"`
	Message  MessageRT `json:"message"`
}

// MessageRT элемент сообщения
type MessageRT struct {
	MessageID int    `json:"message_id"`
	From      FromRT `json:"from"`
	Chat      ChatRT `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
}

// FromRT Информация от кого
type FromRT struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// ChatRT Информация о чате
type ChatRT struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

// MessageUserT ...
type MessageUserT struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}
