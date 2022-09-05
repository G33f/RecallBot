package main

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"goLangBot/botInterface"
	"log"
	"time"
)

func main() {
	connStr := "postgresql://postgres:postgres@127.0.0.1:5432/test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("INSERT INTO recalls(date, username, note)\n VALUES ('2022-09-08', 'g3333f', 'vchera bil DR')")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	if err != nil {
		fmt.Println("its`nt work!")
	}
	bot, err := tgbotapi.NewBotAPI("5779144203:AAHSjQh6r8TwC3mWYNqq7HuDcVnqSTLEZK4")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	test := time.Now()
	bots := botInterface.New(test.Year(), int(test.Month()), test.Day())
	for update := range updates {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		bots.OpenSelector(&msg)
		if update.Message != nil {
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
