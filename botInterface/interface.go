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
	requestType string
}

type BotInterface interface {
	SelectYear(msg *tgbotapi.MessageConfig)
	GetInputRequest(msg *tgbotapi.MessageConfig, request string) error

	SetYear(int)
}

func New() BotInterface {
	tmp := new(botInterface)
	tmp.date = date.New()
	tmp.test = testInput.New()
	return tmp
}

func (bi *botInterface) SetYear(year int) {
	bi.date.SetYear(year)
}

func (bi *botInterface) SetMonth(month int) {
	bi.date.SetMonth(month)
}

func (bi *botInterface) SetDay(day int) {
	bi.date.SetDay(day)
}

func (bi *botInterface) CompleteRequest(Callback func(int), msg *tgbotapi.MessageConfig, request string) error {
	var err error
	if err = bi.test.TestInput(request); err != nil {
		return err
	}
	if i, err := strconv.Atoi(request); err == nil {
		Callback(i)
	}
	return err
}

func (bi *botInterface) years() (str []string) {
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
	bi.requestType = "year"
	bi.test.SetRequest(request)
	bi.button.OpenButtonsBar(msg)
}

func (bi *botInterface) Months() ([]string, int) {
	if bi.date.GetYear() == time.Now().Year() {
		now := time.Now().Month().String()
		monthList := date.GetAllMonths()
		for i, month := range monthList {
			if month == now {
				return monthList[i:], i
			}
		}
	}
	return date.GetAllMonths(), 0
}

func (bi *botInterface) SelectMonth(msg *tgbotapi.MessageConfig) {
	request, spaceBeforeStart := bi.Months()
	bi.button = buttons.New(request, 4, 3, spaceBeforeStart)
	bi.requestType = "month"
	bi.test.SetRequest(request)
	bi.button.OpenButtonsBar(msg)
}

func (bi *botInterface) MakeMonthDayNumbersInStringArray() []string {
	days := bi.date.GetMaxDayInMonth()
	str := make([]string, days, days)
	for i := 0; i < days; i++ {
		str[i] = strconv.Itoa(i + 1)
	}
	return str
}

func (bi *botInterface) CountLine() int {
	days, spaces := bi.date.GetMaxDayInMonth(), bi.date.GetMonthStartWeekDay()
	if (days+spaces)%7 != 0 {
		return (days+spaces)/7 + 1
	} else {
		return (days + spaces) / 7
	}
}

func (bi *botInterface) CreateDaysButtons() {
	bi.button = buttons.New(bi.MakeMonthDayNumbersInStringArray(), bi.CountLine(), bi.date.GetDayInWeek(), bi.date.GetMonthStartWeekDay())
}

func (bi *botInterface) Days() ([]string, int) {
	allDaysInMonth := bi.MakeMonthDayNumbersInStringArray()
	if time.Now().Month().String() == date.GetAllMonths()[bi.date.GetMonth()-1] && bi.date.GetYear() == time.Now().Year() {
		correctDay := strconv.Itoa(time.Now().Day())
		for i, day := range allDaysInMonth {
			if day == correctDay {
				return allDaysInMonth[i:], i + bi.date.GetMonthStartWeekDay()
			}
		}
	}
	return allDaysInMonth, bi.date.GetMonthStartWeekDay()
}

func (bi *botInterface) SelectDay(msg *tgbotapi.MessageConfig) {
	request, spaceBeforeStart := bi.Days()
	bi.button = buttons.New(request, bi.CountLine(), bi.date.GetDayInWeek(), spaceBeforeStart)
	bi.requestType = "day"
	bi.test.SetRequest(request)
	bi.button.OpenButtonsBar(msg)
}

func (bi *botInterface) GetInputRequest(msg *tgbotapi.MessageConfig, request string) error {
	var err error
	switch bi.requestType {
	case "year":
		tmp := bi.SetYear
		if err = bi.CompleteRequest(tmp, msg, request); err != nil {
			return err
		}
		bi.SelectMonth(msg)
		msg.Text = "Choose the Month"
	case "month":
		if err = bi.test.TestInput(request); err != nil {
			return err
		}
		for i, month := range date.GetAllMonths() {
			if month == request {
				bi.SetMonth(i + 1)
			}
		}
		bi.SelectDay(msg)
		msg.Text = "Select Day"
	case "day":
		tmp := bi.SetDay
		if err = bi.CompleteRequest(tmp, msg, request); err != nil {
			return err
		}
		bi.button.CloseButtonsBar(msg)
	}
	return err
}
