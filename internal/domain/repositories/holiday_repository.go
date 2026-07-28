package repositories

import "golang-test/internal/domain/entities"

type IHolidayRepository interface {
	GetAll() ([]entities.Holiday, error)
	SaveAll(holidays []entities.Holiday) error
}
