package bot

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/huntergood/tbot/pkg/parser"
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
func (b *Bot) GetUpdates() ResponseT {
	res, err := http.Get(b.url + "/getUpdates?offset=" + strconv.Itoa(b.LastUpdate))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	target := ResponseT{}
	/* Variant 2
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, targer); err != nil {
		log.Fatal(err)
	}
	*/
	if err := json.NewDecoder(res.Body).Decode(&target); err != nil {
		log.Fatal(err)
	}
	if len(target.Result) != 0 {
		b.LastUpdate = target.Result[len(target.Result)-1].UpdateID + 1
	}
	return target
}

// HandlerMessage обработчик сообщений
func (b *Bot) HandlerMessage(resp ResponseT) {

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

	if _, err = http.Post(b.url+"/sendMessage", "application/json", bytes.NewBuffer(buf)); err != nil {
		log.Fatal(err)
	}
	// Continue ...
}

// SendMessageReactor парсит и отправляет сообщения с reactor
func (b *Bot) SendMessageReactor(url string) {
	html := parser.GetHTML(url)
	strs := parser.GetObject(html, `<script\stype="application/ld\+json"[^>]*>(.+?)</script>`)
	for _, str := range strs {
		res := JSONReact{}
		if err := json.Unmarshal([]byte(str), &res); err != nil {
			log.Fatal(err)
		}
		f := res
		go b.sendMedia(f, 281115651)
	}
}

func (b *Bot) sendMedia(res JSONReact, id int) {
	var (
		postURL      = "/sendPhoto"
		buff         []byte
		err          error
		replauMarkup = ReplayMarkupT{
			[][]InlineKeyboardT{
				{
					InlineKeyboardT{
						Text: "Открыть пост",
						URL:  "http://joyreactor.cc" + res.MainEntity.ID,
					},
				},
			},
		}
	)
	if !b.checkPhoto(res.Image.URL) {
		postURL = "/sendAnimation"
		ams := &MessageUserAnimationT{
			ChatID:       id,
			Animation:    res.Image.URL,
			Caption:      getCaption(res.HeadLine),
			ReplayMarkup: replauMarkup,
		}
		buff, err = json.Marshal(ams)
		if err != nil {
			log.Fatal(err)
		}
	}

	if len(buff) == 0 {
		ams := &MessageUserPhotoT{
			ChatID:       id,
			Photo:        res.Image.URL,
			Caption:      getCaption(res.HeadLine),
			ReplayMarkup: replauMarkup,
		}
		buff, err = json.Marshal(ams)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := http.Post(b.url+postURL, "application/json", bytes.NewBuffer(buff)); err != nil {
		log.Fatal(err)
	}

}

func (b *Bot) checkPhoto(img string) bool {
	split := strings.Split(img, ".")
	if split[len(split)-1] == "gif" {
		return false
	}
	return true
}

func getCaption(s string) string {
	var result = make([]string, 0)
	split := strings.Split(s, " :: ")
	for _, str := range split {
		result = append(result, "#"+str)
	}
	return strings.Join(result, " ")
}
