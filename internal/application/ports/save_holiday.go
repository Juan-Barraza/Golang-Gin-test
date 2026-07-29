package ports

import "golang-test/internal/application/dtos"

type ISaveHolidayUseCase interface {
	Execute(requestz []dtos.HolidayDto) ([]dtos.ResponseHolidateDto, error)
}
