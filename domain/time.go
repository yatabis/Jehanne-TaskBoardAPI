package domain

import "time"

type Date time.Time
type Datetime time.Time
type Minutes time.Duration

var FixedTime time.Time

func Now() Datetime {
	if !FixedTime.IsZero() {
		return Datetime(FixedTime)
	}
	return Datetime(time.Now())
}

func Today() Date {
	year, month, day := time.Time(Now()).Date()
	return Date(time.Date(year, month, day, 0, 0, 0, 0, time.Local))
}

func (d Date) IsZero() bool {
	return time.Time(d).IsZero()
}

func (d *Date) Init(init Date) {
	if d.IsZero() {
		*d = init
	}
}

func NewDate(s string) (Date, error) {
	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return Date{}, err
	}
	return Date(date), err
}
