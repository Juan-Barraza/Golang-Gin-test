package usecases

import (
	"fmt"
	"golang-test/internal/application/dtos"
	"golang-test/internal/domain/entities"
	"golang-test/internal/domain/errors"
	"golang-test/internal/domain/repositories"
	"time"
)

var validHolidayTypes = map[string]bool{
	"Civil":     true,
	"Religioso": true,
}

type GetAllHolidayUseCase struct {
	holidayRepository repositories.IHolidayRepository
}

func NewGetAllHolidayUseCase(repository repositories.IHolidayRepository) *GetAllHolidayUseCase {
	return &GetAllHolidayUseCase{holidayRepository: repository}
}

func (g *GetAllHolidayUseCase) Execute(filters *entities.HolidayFilter) ([]dtos.ResponseHolidateDto, error) {
	holidays, err := g.holidayRepository.GetAll()
	if err != nil {
		return nil, err
	}
	holidays, err = g.filterHolidays(holidays, filters)
	if err != nil {
		return nil, err
	}

	response := dtos.NewListResponseHolidateDto(holidays)

	return response, nil
}

func (g *GetAllHolidayUseCase) filterHolidays(holidays []entities.Holiday, filters *entities.HolidayFilter) ([]entities.Holiday, error) {
	if filters == nil {
		return holidays, nil
	}

	if filters.Type != "" && !g.isValidType(filters.Type) {
		return nil, fmt.Errorf("%w: %q", errors.ErrInvalidType, filters.Type)
	}

	var startDate, endDate time.Time
	var err error
	if filters.StartDate != "" {
		startDate, err = time.Parse("2006-01-02", filters.StartDate)
		if err != nil {
			return nil, fmt.Errorf("%w: %q", errors.ErrInvalidDate, filters.StartDate)
		}
	}
	if filters.EndDate != "" {
		endDate, err = time.Parse("2006-01-02", filters.EndDate)
		if err != nil {
			return nil, fmt.Errorf("%w: %q", errors.ErrInvalidDate, filters.EndDate)
		}
	}

	filteredHolidays := make([]entities.Holiday, 0)
	for _, holiday := range holidays {

		if filters.Type != "" && holiday.Type != filters.Type {
			continue
		}

		if !startDate.IsZero() && holiday.Date.Before(startDate) {
			continue
		}
		if !endDate.IsZero() && holiday.Date.After(endDate) {
			continue
		}

		filteredHolidays = append(filteredHolidays, holiday)
	}
	return filteredHolidays, nil
}

func (g *GetAllHolidayUseCase) isValidType(t string) bool {
	return validHolidayTypes[t]
}
