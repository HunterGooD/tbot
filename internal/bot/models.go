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

// MessageUserT Отправка обычного сообщения пользователю
type MessageUserT struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

// StartMessage
type StartMessage struct {
	ChatID   int      `json:"chat_id"`
	Text     string   `json:"text"`
	ReplayMT ReplayMT `json:"reply_markup"`
}

type ReplayMT struct {
	ResizeKeyboard bool          `json:"resize_keyboard"`
	Keyboard       [][]KeyboardT `json:"keyboard"`
}
type KeyboardT struct {
	Text string `json:"text"`
}

// MessageUserPhotoT Отправка сообщения с фото
type MessageUserPhotoT struct {
	ChatID       int           `json:"chat_id"`
	Photo        string        `json:"photo"`
	Caption      string        `json:"caption"`
	ReplayMarkup ReplayMarkupT `json:"replay_markup"`
}

// MessageUserAnimationT Отправка сообщения с анимацией
type MessageUserAnimationT struct {
	ChatID       int           `json:"chat_id"`
	Animation    string        `json:"animation"`
	Caption      string        `json:"caption"`
	ReplayMarkup ReplayMarkupT `json:"replay_markup"`
}

// ReplayMarkupT ..
type ReplayMarkupT struct {
	InlineKeyboard [][]InlineKeyboardT `json:"inline_keyboard"`
}

// InlineKeyboardT ..
type InlineKeyboardT struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

// JSONReact для распарсенных данных Reactora
type JSONReact struct {
	Type       string          `json:"@type"`
	MainEntity MainEntityReact `json:"mainEntityOfPage"`
	HeadLine   string          `json:"headline"`
	Image      ImageReactor    `json:"image"`
}

// MainEntityReact ..
type MainEntityReact struct {
	Type string `json:"@type"`
	ID   string `json:"@id"`
}

//ImageReactor ..
type ImageReactor struct {
	Type string `json:"@type"`
	URL  string `json:"url"`
}
