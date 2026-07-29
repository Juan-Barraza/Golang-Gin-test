package dtos

import (
	"fmt"
	"golang-test/internal/domain/entities"
	"golang-test/internal/domain/errors"
	"time"
)

type HolidayDto struct {
	Date        string
	Title       string
	Phone       string
	Type        string
	Inalienable bool
	Extra       string
}

type ResponseHolidateDto struct {
	Date        string `json:"date" xml:"date"`
	Title       string `json:"title" xml:"title"`
	Phone       string `json:"phone" xml:"phone"`
	Type        string `json:"type" xml:"type"`
	Inalienable bool   `json:"inalienable" xml:"inalienable"`
	Extra       string `json:"extra" xml:"extra"`
}

type HolidaysResponse struct {
	Holidays []ResponseHolidateDto `json:"holidays" xml:"holiday"`
}


func NewHolidayDto(holiday HolidayDto) (*entities.Holiday, error) {
	date, err := time.Parse("2006-01-02", holiday.Date)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrInvalidDate, holiday.Date)
	}
	return &entities.Holiday{
		Date:        date,
		Title:       holiday.Title,
		Phone:       holiday.Phone,
		Type:        holiday.Type,
		Inalienable: holiday.Inalienable,
		Extra:       holiday.Extra,
	}, nil
}

func NewListHolidayDto(requests []HolidayDto) ([]entities.Holiday, error) {
	holidays := make([]entities.Holiday, 0, len(requests))
	for _, dto := range requests {
		holiday, err := NewHolidayDto(dto)
		if err != nil {
			return nil, err
		}
		holidays = append(holidays, *holiday)
	}
	return holidays, nil
}

func NewResponseHolidateDto(holiday entities.Holiday) *ResponseHolidateDto {
	return &ResponseHolidateDto{
		Date:        holiday.Date.Format("2006-01-02"),
		Title:       holiday.Title,
		Phone:       holiday.Phone,
		Type:        holiday.Type,
		Inalienable: holiday.Inalienable,
		Extra:       holiday.Extra,
	}
}

func NewListResponseHolidateDto(holidays []entities.Holiday) []ResponseHolidateDto {
	dtos := make([]ResponseHolidateDto, len(holidays))
	for i, holiday := range holidays {
		dtos[i] = *NewResponseHolidateDto(holiday)
	}
	return dtos
}
