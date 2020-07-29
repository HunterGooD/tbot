package main

import (
	"fmt"
	"os"
	"time"

	"github.com/huntergood/tbot/internal/bot"

	"github.com/huntergood/tbot/internal/config"
)

func main() {
	var s *config.Special = config.NewSpecial()
	botT := bot.NewBot(s.URL(), s.DBname)

	// если это выполняется программа завершается КОСТЫЛЬ
	if len(os.Args) == 2 {
		if os.Args[1] == "-m" {
			// костыль
			botT.Migrate()
		}
	}

	timer := time.NewTicker(time.Second * 30)
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
		botT.SendMessageReactor("http://reactor.cc/")
		botT.SendMessageReactor("http://anime.reactor.cc/")
	}
}
