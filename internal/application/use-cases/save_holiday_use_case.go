package usecases

import (
	"golang-test/internal/application/dtos"
	"golang-test/internal/domain/repositories"
)

type SaveHolidayUseCase struct {
	holidayRepository repositories.IHolidayRepository
}

func NewSaveHolidayUseCase(repository repositories.IHolidayRepository) *SaveHolidayUseCase {
	return &SaveHolidayUseCase{holidayRepository: repository}
}

func (s *SaveHolidayUseCase) Execute(hld []dtos.HolidayDto) ([]dtos.ResponseHolidateDto, error) {
	holidays, err := dtos.NewListHolidayDto(hld)
	if err != nil {
		return nil, err
	}

	err = s.holidayRepository.SaveAll(holidays)
	if err != nil {
		return nil, err
	}

	response := dtos.NewListResponseHolidateDto(holidays)

	return response, nil
}
