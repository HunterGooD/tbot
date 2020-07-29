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
		botT.SendMessageReactor("http://reactor.cc/")
		botT.SendMessageReactor("http://anime.reactor.cc/")
	}
}

/*
Parse Img
2020/07/29 10:49:20 http://reactor.cc/1551 0
2020/07/29 10:49:22 http://anime.reactor.cc/1775 0
2020/07/29 10:49:22 http://anime.reactor.cc/1775 1
2020/07/29 10:49:22 http://anime.reactor.cc/1775 2
2020/07/29 10:49:22 http://anime.reactor.cc/1775 3
Должно быть от 0 до 9
*/
