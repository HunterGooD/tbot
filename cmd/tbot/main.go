package main

import (
	"fmt"
	"os"
	"time"

	"github.com/huntergood/tbot/internal/bot"
	"github.com/huntergood/tbot/internal/repository"

	"github.com/huntergood/tbot/internal/config"
)

var (
	// URL ..
	URL string
)

func main() {
	var s *config.Special = new(config.Special)
	URL = s.GetURL()

	// если это выполняется программа завершается
	if len(os.Args) == 2 {
		if os.Args[1] == "-m" {
			db := repository.Connect(s.DbNAme)
			repository.Migration(db)
		}
	}

	botT := bot.NewBot(URL)
	timer := time.NewTicker(time.Minute * 10)
	go updateContent(timer.C, botT)
	for {
		res := botT.GetUpdates()
		for _, resp := range res.Result {
			go botT.HandlerMessage(resp)
			// botT.SendMessage(resp.Message.From.ID, resp.Message.Text)
		}
	}

}

func updateContent(c <-chan time.Time, botT *bot.Bot) {
	// Использовать время для добавления в БД хм
	for range c {
		//parse img
		fmt.Println("Parse Img")
		botT.SendMessageReactor("http://joyreactor.cc/")
	}
}
