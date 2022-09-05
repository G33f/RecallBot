package buttons

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type buttons struct {
	button       tgbotapi.ReplyKeyboardMarkup
	isButtonOpen bool
}

type Buttons interface {
	OpenButtonsBar(msg *tgbotapi.MessageConfig)
	CloseButtonsBar(msg *tgbotapi.MessageConfig)
}

func New(buttonsName []string, linesCount int, buttonsOnEachLine int, spaceBeforeStart int) Buttons {
	b := buttons{}
	tmp := spaceBeforeStart
	var tmpButtonsForAPI = make([][]tgbotapi.KeyboardButton, linesCount, linesCount)
	for i := 0; i < linesCount; i++ {
		var buttonsLine = make([]tgbotapi.KeyboardButton, buttonsOnEachLine, buttonsOnEachLine)
		for j := 0; j < buttonsOnEachLine; j++ {
			if tmp != 0 {
				tmp--
				buttonsLine[j] = tgbotapi.NewKeyboardButton(" ")
			} else {
				if i*buttonsOnEachLine+j-spaceBeforeStart < len(buttonsName) {
					buttonsLine[j] = tgbotapi.NewKeyboardButton(buttonsName[i*buttonsOnEachLine+j-spaceBeforeStart])
				} else {
					buttonsLine[j] = tgbotapi.NewKeyboardButton(" ")
				}
			}
		}
		tmpButtonsForAPI[i] = tgbotapi.NewKeyboardButtonRow(buttonsLine...)
	}
	b.button = tgbotapi.NewReplyKeyboard(tmpButtonsForAPI...)
	return &b
}

func (b *buttons) OpenButtonsBar(msg *tgbotapi.MessageConfig) {
	b.isButtonOpen = true
	msg.ReplyMarkup = b.button
}

func (b *buttons) CloseButtonsBar(msg *tgbotapi.MessageConfig) {
	b.isButtonOpen = false
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
}
