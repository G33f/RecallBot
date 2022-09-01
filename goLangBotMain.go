package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"goLangBot/botInterface"
	"log"
	"time"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5779144203:AAHSjQh6r8TwC3mWYNqq7HuDcVnqSTLEZK4")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	//strs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}
	//butt := buttons.New(strs, 5, 7, 3)
	test := time.Now()
	bots := botInterface.New(test.Year(), int(test.Month()), test.Day())
	for update := range updates {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg = butt.OpenButtonsBar(msg)
		bots.OpenSelector(&msg)
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

		//msg = butt.CloseButtonsBar(msg)
		bots.CloseSelector(&msg)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "WRONG command!!!")
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}
