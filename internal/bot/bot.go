package bot

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/huntergood/tbot/internal/repository"

	"github.com/huntergood/tbot/pkg/parser"
)

// Bot структура отвечающая за бота
type Bot struct {
	url        string
	LastUpdate int
	repo       *repository.Repo
}

// NewBot Инициализация бота
func NewBot(url, dbname string) *Bot {
	r := repository.NewRepo(dbname)
	r.Connect()
	return &Bot{
		url:  url,
		repo: r,
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

// Migrate ...
func (b *Bot) Migrate() {
	b.repo.Migration()
}

// HandlerMessage обработчик сообщений
func (b *Bot) HandlerMessage(resp ResultRT) {
	switch resp.Message.Text {
	case "/start":
		u := b.repo.GetUserByID(resp.Message.From.ID)
		if u == nil {
			b.repo.AddUser(resp.Message.From.ID, resp.Message.From.Username)
			b.sendMessageStart(resp.Message.From.ID)
			return
		}
	case "/random":
		fallthrough
	case "Случайный пост":
		b.sendRandom(resp.Message.From.ID)
		return
	}
	b.SendMessage(resp.Message.From.ID, resp.Message.Text)
}

func (b *Bot) sendMessageStart(id int) {
	var sUser = StartMessage{
		ChatID: id,
		Text:   "Приветствую",
		ReplayMT: ReplayMT{
			ResizeKeyboard: true,
			Keyboard:       [][]KeyboardT{{KeyboardT{Text: "Случайный пост"}}, {KeyboardT{Text: "Подписаться на тег"}}},
		},
	}

	buf, err := json.Marshal(sUser)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = http.Post(b.url+"/sendMessage", "application/json", bytes.NewBuffer(buf)); err != nil {
		log.Fatal(err)
	}

}

// SendMessage //
func (b *Bot) SendMessage(id int, str string) {
	var mUser = MessageUserT{
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

func (b *Bot) getStrings(url, reg string) []string {
	html := parser.GetHTML(url)
	return parser.GetObject(html, reg)
}

func (b *Bot) sendRandom(id int) {
	var strs []string
	var wg = new(sync.WaitGroup)
	for len(strs) == 0 {
		strs = b.getStrings("http://reactor.cc/random", `<div\sclass="image".*?href\s*=\s*['"]([^\s'"]+)[\s'"]`)
	}
	for _, str := range strs {
		var res = JSONReact{
			Type: "Image",
			Image: ImageReactor{
				URL: str,
			},
		}
		f := res
		wg.Add(1)
		go func(wgf *sync.WaitGroup) {
			b.sendUsers(f, []int{id})
			wgf.Done()
		}(wg)
	}
	wg.Wait()
}

// SendMessageReactor парсит и отправляет сообщения с reactor
func (b *Bot) SendMessageReactor(url string) {
	var idusers []int = b.repo.GetIDUsers()
	var wg = new(sync.WaitGroup)
	rand.Seed(time.Now().Unix())
	page := strconv.Itoa(rand.Intn(2500) + 1)
	strs := b.getStrings(url+page, `<script\stype="application/ld\+json"[^>]*>(.+?)</script>`)
	for i, str := range strs {
		if len(str) == 0 {
			continue
		}
		log.Println(url+page, strconv.Itoa(i))
		var res = JSONReact{}
		if err := json.Unmarshal([]byte(str), &res); err != nil {
			log.Fatal(err)
		}
		f := res
		wg.Add(1)
		go func(wgf *sync.WaitGroup) {
			b.sendUsers(f, idusers)
			wgf.Done()
		}(wg)
	}
	wg.Wait()
}

func (b *Bot) sendUsers(res JSONReact, users []int) {
	var wg = new(sync.WaitGroup)
	for _, user := range users {
		wg.Add(1)
		go func(wgf *sync.WaitGroup, user int) {
			b.sendMedia(res, user)
			wgf.Done()
		}(wg, user)
	}
	wg.Wait()
}

func (b *Bot) sendMedia(res JSONReact, id int) {
	var (
		postURL      = "/sendPhoto"
		buff         []byte
		err          error
		replauMarkup = ReplayMarkupT{
			[][]InlineKeyboardT{
				{InlineKeyboardT{Text: "Открыть пост", URL: "http://joyreactor.cc" + res.MainEntity.ID}},
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
