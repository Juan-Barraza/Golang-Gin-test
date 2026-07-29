package usecases_test

import (
	"errors"
	"testing"
	"time"

	usecases "golang-test/internal/application/use-cases"
	"golang-test/internal/domain/entities"
	domainErrors "golang-test/internal/domain/errors"
)

type mockHolidayRepository struct {
	holidays []entities.Holiday
	err      error
}

func (m *mockHolidayRepository) GetAll() ([]entities.Holiday, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.holidays, nil
}

func (m *mockHolidayRepository) SaveAll(holidays []entities.Holiday) error {
	return nil
}

func setupMockHolidays() []entities.Holiday {
	d1, _ := time.Parse("2006-01-02", "2026-01-01")
	d2, _ := time.Parse("2006-01-02", "2026-04-03")
	d3, _ := time.Parse("2006-01-02", "2026-05-01")

	return []entities.Holiday{
		{Date: d1, Title: "Año Nuevo", Type: "Civil", Inalienable: true, Extra: "Civil e Irrenunciable"},
		{Date: d2, Title: "Viernes Santo", Type: "Religioso", Inalienable: false, Extra: "Religioso"},
		{Date: d3, Title: "Día del Trabajo", Type: "Civil", Inalienable: true, Extra: "Civil e Irrenunciable"},
	}
}

func TestGetAllHolidayUseCase_Execute_SuccessNoFilters(t *testing.T) {
	repo := &mockHolidayRepository{holidays: setupMockHolidays()}
	useCase := usecases.NewGetAllHolidayUseCase(repo)

	result, err := useCase.Execute(nil)
	if err != nil {
		t.Fatalf("se esperaba nil error, se obtuvo: %v", err)
	}

	if len(result) != 3 {
		t.Errorf("se esperaban 3 feriados, se obtuvieron: %d", len(result))
	}
}

func TestGetAllHolidayUseCase_Execute_FilterByType(t *testing.T) {
	repo := &mockHolidayRepository{holidays: setupMockHolidays()}
	useCase := usecases.NewGetAllHolidayUseCase(repo)

	filter := &entities.HolidayFilter{Type: "Religioso"}
	result, err := useCase.Execute(filter)
	if err != nil {
		t.Fatalf("se esperaba nil error, se obtuvo: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("se esperaba 1 feriado religioso, se obtuvieron: %d", len(result))
	}

	if result[0].Title != "Viernes Santo" {
		t.Errorf("se esperaba Viernes Santo, se obtuvo: %s", result[0].Title)
	}
}

func TestGetAllHolidayUseCase_Execute_FilterByDateRange(t *testing.T) {
	repo := &mockHolidayRepository{holidays: setupMockHolidays()}
	useCase := usecases.NewGetAllHolidayUseCase(repo)

	filter := &entities.HolidayFilter{
		StartDate: "2026-01-01",
		EndDate:   "2026-04-01",
	}

	result, err := useCase.Execute(filter)
	if err != nil {
		t.Fatalf("se esperaba nil error, se obtuvo: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("se esperaba 1 feriado en el rango, se obtuvieron: %d", len(result))
	}

	if result[0].Title != "Año Nuevo" {
		t.Errorf("se esperaba Año Nuevo, se obtuvo: %s", result[0].Title)
	}
}

func TestGetAllHolidayUseCase_Execute_InvalidType(t *testing.T) {
	repo := &mockHolidayRepository{holidays: setupMockHolidays()}
	useCase := usecases.NewGetAllHolidayUseCase(repo)

	filter := &entities.HolidayFilter{Type: "Invalido"}
	_, err := useCase.Execute(filter)

	if err == nil {
		t.Fatal("se esperaba error por tipo invalido, se obtuvo nil")
	}

	if !errors.Is(err, domainErrors.ErrInvalidType) {
		t.Errorf("se esperaba ErrInvalidType, se obtuvo: %v", err)
	}
}

func TestGetAllHolidayUseCase_Execute_InvalidDateFormat(t *testing.T) {
	repo := &mockHolidayRepository{holidays: setupMockHolidays()}
	useCase := usecases.NewGetAllHolidayUseCase(repo)

	filter := &entities.HolidayFilter{StartDate: "01-01-2026"}
	_, err := useCase.Execute(filter)

	if err == nil {
		t.Fatal("se esperaba error por formato de fecha invalido, se obtuvo nil")
	}

	if !errors.Is(err, domainErrors.ErrInvalidDate) {
		t.Errorf("se esperaba ErrInvalidDate, se obtuvo: %v", err)
	}
}
