package date

import (
	dt "github.com/rickb777/date"
	"time"
)

func GetAllMonths() []string {
	return []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
}

type date struct {
	year              int
	month             int
	day               int
	maxDayInMonth     int
	monthStartWeekDay int
	dayInWeek         int
}

type Date interface {
	setDayInWeek()
	setMaxDayInMonth()
	setMonthStartWeekDay()
	SetYear(year int)
	SetMonth(month int)
	SetDay(day int)
	GetYear() int
	GetMonth() int
	GetMonthStartWeekDay() int
	GetMaxDayInMonth() int
	GetDayInWeek() int
}

func New() Date {
	return &date{}
}

func (d date) GetYear() int {
	return d.year
}
func (d date) GetMonth() int {
	return d.month
}

func (d *date) setDayInWeek() {
	d.dayInWeek = 7
}

func (d *date) setMaxDayInMonth() {
	d.maxDayInMonth = dt.DaysIn(d.year, time.Month(d.month))
}

func (d *date) setMonthStartWeekDay() {
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

func (d *date) SetYear(year int) {
	d.year = year
}

func (d *date) SetMonth(month int) {
	d.month = month
	d.setMaxDayInMonth()
	d.setMonthStartWeekDay()
	d.setDayInWeek()
}

func (d *date) SetDay(day int) {
	d.day = day
}

func (d *date) GetMonthStartWeekDay() int {
	return d.monthStartWeekDay
}
func (d *date) GetMaxDayInMonth() int {
	return d.maxDayInMonth
}

func (d *date) GetDayInWeek() int {
	return d.dayInWeek
}
