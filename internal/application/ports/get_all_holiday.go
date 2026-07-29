package ports

import (
	"golang-test/internal/application/dtos"
	"golang-test/internal/domain/entities"
)

type IGetAllHolidayUseCase interface {
	Execute(filters *entities.HolidayFilter) ([]dtos.ResponseHolidateDto, error)
}
