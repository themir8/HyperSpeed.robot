package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var bot, _ = tb.NewBot(tb.Settings{
	// here token
	Token:  "1446049605:AAFU3Tzrp3SciOBQBS4dROnQFlsntxgy5to", // this token is not real
	Poller: &tb.LongPoller{Timeout: 10 * time.Second},
})

// log.Fatal(err)

func main() {

	bot.Handle("/start", startHandle)
	bot.Handle(tb.OnText, textHandle)
	log.Printf(
		"Bot connected! FirstName: %s, UserName: @%s, is bot: %t, Bot id: %v.\n",
		bot.Me.FirstName,
		bot.Me.Username,
		bot.Me.IsBot,
		bot.Me.ID)
	bot.Start()
}

func startHandle(m *tb.Message) {
	log.Println("Hello", m.Chat.Username)
	bot.Send(m.Sender, "Hello")
}

func textHandle(m *tb.Message) {
	// all the text messages that weren't
	// captured by existing handlers
}
