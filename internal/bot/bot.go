package bot

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// Bot структура отвечающая за бота
type Bot struct {
	url        string
	LastUpdate int
}

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

// NewBot Инициализация бота
func NewBot(url string) *Bot {
	return &Bot{
		url: url,
	}
}

// GetUpdates Получение новых сообщений
func (b *Bot) GetUpdates() *ResponseT {
	res, err := http.Get(b.url + "/getUpdates?offset=" + strconv.Itoa(b.LastUpdate))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	target := &ResponseT{}
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		log.Fatal(err)
	}
	if len(target.Result) == 0 {
		return target
	}
	if b.LastUpdate != target.Result[len(target.Result)-1].UpdateID {
		b.LastUpdate = target.Result[len(target.Result)-1].UpdateID + 1
	}
	return target
}

// SendMessage //
func (b *Bot) SendMessage(id int, str string) {
	res, err := http.Get(b.url + "/sendMessage?chat_id=" + strconv.Itoa(id) + "&text=" + str)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	// Continue ...
}
