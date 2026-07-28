package entities

import "time"

type Holiday struct {
	Date        time.Time
	Title       string
	Phone       string
	Type        string
	Inalienable bool
	Extra       string
}

type HolidayFilter struct {
	Type      string
	StartDate string
	EndDate   string
}
