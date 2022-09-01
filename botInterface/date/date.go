package date

import (
	dt "github.com/rickb777/date"
	"time"
)

//type Dates interface {
//	SetYear()
//	SetMonth()
//	SetDay()
//	GetMonthStartWeekDay()
//	GetMaxDayInMonth()
//	GetDayInWeek()
//
//	setMaxDayInMonth()
//	setMonthStartWeekDay()
//	setDayInWeek()
//}

type Date struct {
	year              int
	month             int
	day               int
	maxDayInMonth     int
	monthStartWeekDay int
	dayInWeek         int
}

func (d *Date) setDayInWeek() {
	d.dayInWeek = 7
}

func (d *Date) setMaxDayInMonth() {
	d.maxDayInMonth = dt.DaysIn(d.year, time.Month(d.month))
}

func (d *Date) setMonthStartWeekDay() {
	m := dt.New(d.year, time.Month(d.month), 1)
	switch m.Weekday().String() {
	case "Monday":
		d.monthStartWeekDay = 0
	case "Tuesday":
		d.monthStartWeekDay = 1
	case "Wednesday":
		d.monthStartWeekDay = 2
	case "Thursday":
		d.monthStartWeekDay = 3
	case "Friday":
		d.monthStartWeekDay = 4
	case "Saturday":
		d.monthStartWeekDay = 5
	case "Sunday":
		d.monthStartWeekDay = 6
	}
}

func (d *Date) SetYear(year int) {
	d.year = year
}

func (d *Date) SetMonth(month int) {
	d.month = month
	d.setMaxDayInMonth()
	d.setMonthStartWeekDay()
	d.setDayInWeek()
}

func (d *Date) SetDay(day int) {
	d.day = day
}

func (d Date) GetMonthStartWeekDay() int {
	return d.monthStartWeekDay
}
func (d Date) GetMaxDayInMonth() int {
	return d.maxDayInMonth
}

func (d Date) GetDayInWeek() int {
	return d.dayInWeek
}
