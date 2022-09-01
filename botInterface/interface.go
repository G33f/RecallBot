package botInterface

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"goLangBot/botInterface/buttons"
	"goLangBot/botInterface/date"
	"strconv"
)

type BotInterface struct {
	button buttons.Buttons
	date   date.Date
}

func New(year int, month int, day int) BotInterface {
	var tmp BotInterface
	tmp.date.SetYear(year)
	tmp.date.SetMonth(month)
	tmp.date.SetDay(day)
	tmp.CreateButtons()
	return tmp
}

func (bi BotInterface) countLine() int {
	days, spaces := bi.date.GetMaxDayInMonth(), bi.date.GetMonthStartWeekDay()
	if (days+spaces)%7 != 0 {
		return (days+spaces)/7 + 1
	} else {
		return (days + spaces) / 7
	}
}

func (bi BotInterface) makeMonthDayNumbersInStringArray() []string {
	days := bi.date.GetMaxDayInMonth()
	str := make([]string, days, days)
	for i := 0; i < days; i++ {
		str[i] = strconv.Itoa(i + 1)
	}
	return str
}

func (bi *BotInterface) CreateButtons() {
	bi.button = buttons.New(bi.makeMonthDayNumbersInStringArray(), bi.countLine(), bi.date.GetDayInWeek(), bi.date.GetMonthStartWeekDay())
}

func (bi *BotInterface) OpenSelector(msg *tgbotapi.MessageConfig) {
	bi.button.OpenButtonsBar(msg)
}

func (bi *BotInterface) CloseSelector(msg *tgbotapi.MessageConfig) {
	bi.button.OpenButtonsBar(msg)
}
