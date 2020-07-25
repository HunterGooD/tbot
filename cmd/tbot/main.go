package main

import (
	"github.com/huntergood/tbot/internal/bot"

	"github.com/huntergood/tbot/internal/config"
)

var (
	// URL ..
	URL string
)

func main() {
	URL = config.GetURL()
	botT := bot.NewBot(URL)
	for {
		res := botT.GetUpdates()
		// BUG: 3 message loop
		for _, resp := range res.Result {
			botT.SendMessage(resp.Message.From.ID, resp.Message.Text)
		}
	}

}
