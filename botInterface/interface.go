package botInterface

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"goLangBot/botInterface/buttons"
	"goLangBot/botInterface/date"
	"goLangBot/botInterface/testInput"
	"strconv"
	"time"
)

type botInterface struct {
	button      buttons.Buttons
	date        date.Date
	test        testInput.TestInput
	requestType map[int64]string
}

func (bi *botInterface) SetYear(year int) {
	bi.date.SetYear(year)
}

func (bi botInterface) years() (str []string) {
	now := time.Now().Year()
	str = make([]string, 10, 10)
	for i := 0; i < 10; i++ {
		str[i] = strconv.Itoa(now + i)
	}
	return str
}

func (bi *botInterface) SelectYear(msg *tgbotapi.MessageConfig) {
	request := bi.years()
	bi.button = buttons.New(request, 2, 5, 0)
	bi.requestType[msg.ChatID] = "year"
	bi.test.SetRequest(msg.ChatID, request)
	bi.button.OpenButtonsBar(msg)
}

func (bi *botInterface) GetInputRequest(msg *tgbotapi.MessageConfig, request string) {
	switch bi.requestType[msg.ChatID] {
	case "year":
		if i, err := strconv.Atoi(request); err == nil {
			bi.SetYear(i)
		}
	}
}

type BotInterface interface {
	SelectYear(msg *tgbotapi.MessageConfig)
	GetInputRequest(msg *tgbotapi.MessageConfig, request string)

	SetYear(int)
}

func New() BotInterface {
	tmp := new(botInterface)
	tmp.date = date.New()
	tmp.test = testInput.New()
	tmp.requestType = make(map[int64]string)
	return tmp
}

func (bi botInterface) countLine() int {
	days, spaces := bi.date.GetMaxDayInMonth(), bi.date.GetMonthStartWeekDay()
	if (days+spaces)%7 != 0 {
		return (days+spaces)/7 + 1
	} else {
		return (days + spaces) / 7
	}
}

func (bi botInterface) makeMonthDayNumbersInStringArray() []string {
	days := bi.date.GetMaxDayInMonth()
	str := make([]string, days, days)
	for i := 0; i < days; i++ {
		str[i] = strconv.Itoa(i + 1)
	}
	return str
}

func (bi *botInterface) CreateDaysButtons() {
	bi.button = buttons.New(bi.makeMonthDayNumbersInStringArray(), bi.countLine(), bi.date.GetDayInWeek(), bi.date.GetMonthStartWeekDay())
}

func (bi *botInterface) OpenSelector(msg *tgbotapi.MessageConfig) {
	bi.button.OpenButtonsBar(msg)
}

func (bi *botInterface) CloseSelector(msg *tgbotapi.MessageConfig) {
	bi.button.OpenButtonsBar(msg)
}
