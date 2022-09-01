package buttons

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//type Button interface {
//	OpenButtonsBar()
//	CloseButtonsBar()
//}

type Buttons struct {
	button       tgbotapi.ReplyKeyboardMarkup
	isButtonOpen bool
}

func New(buttonsName []string, linesCount int, buttonsOnEachLine int, spaceBeforeStart int) Buttons {
	var b Buttons
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
	return b
}

func (b *Buttons) OpenButtonsBar(msg *tgbotapi.MessageConfig) {
	b.isButtonOpen = true
	msg.ReplyMarkup = b.button
}

func (b *Buttons) CloseButtonsBar(msg *tgbotapi.MessageConfig) {
	b.isButtonOpen = false
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
}
