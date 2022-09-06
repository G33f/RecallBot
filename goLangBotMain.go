package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"goLangBot/botInterface"
	"log"
)

type tgStruct struct {
	bot    *tgbotapi.BotAPI
	msg    tgbotapi.MessageConfig
	update tgbotapi.Update
}

func (tg *tgStruct) SelectYear(bot botInterface.BotInterface) {
	bot.SelectYear(&tg.msg)
	tg.msg.Text = "Choose the Year"
	tg.msg.ReplyToMessageID = tg.update.Message.MessageID
	if _, err := tg.bot.Send(tg.msg); err != nil {
		log.Panic(err)
	}
}

func main() {
	//connStr := "postgresql://postgres:postgres@127.0.0.1:5432/test?sslmode=disable"
	//db, err := sql.Open("postgres", connStr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = db.Ping()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//rows, err := db.Query("INSERT INTO recalls(date, username, note)\n VALUES ('2022-09-08', 'g3333f', 'vchera bil DR')")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer rows.Close()
	//if err != nil {
	//	fmt.Println("its`nt work!")
	//}
	tg := new(tgStruct)
	var err error
	tg.bot, err = tgbotapi.NewBotAPI("5779144203:AAHSjQh6r8TwC3mWYNqq7HuDcVnqSTLEZK4")
	if err != nil {
		log.Panic(err)
	}

	tg.bot.Debug = true

	log.Printf("Authorized on account %s", tg.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tg.bot.GetUpdatesChan(u)
	bot := botInterface.New()
	for tg.update = range updates {
		tg.msg = tgbotapi.NewMessage(tg.update.Message.Chat.ID, tg.update.Message.Text)
		if tg.update.Message != nil {
			switch tg.update.Message.Text {
			case "/start":
				tg.SelectYear(bot)
			default:
				bot.GetInputRequest(&tg.msg, tg.update.Message.Text)
				tg.msg.Text = fmt.Sprintf("year whos coose %s", tg.update.Message.Text)
				tg.msg.ReplyToMessageID = tg.update.Message.MessageID
				if _, err := tg.bot.Send(tg.msg); err != nil {
					log.Panic(err)
				}
			}
		}
	}
}
