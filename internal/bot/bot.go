package bot

import (
	"bytes"
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
	/* Variant 2
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, targer); err != nil {
		log.Fatal(err)
	}
	*/
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		log.Fatal(err)
	}
	if len(target.Result) != 0 {
		b.LastUpdate = target.Result[len(target.Result)-1].UpdateID + 1
	}
	return target
}

// SendMessage //
func (b *Bot) SendMessage(id int, str string) {
	var mUser MessageUserT = MessageUserT{
		ChatID: id,
		Text:   str,
	}
	buf, err := json.Marshal(mUser)
	if err != nil {
		log.Fatal(err)
	}

	_, err = http.Post(b.url+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		log.Fatal(err)
	}
	// Continue ...
}
