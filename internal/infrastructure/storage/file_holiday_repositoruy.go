package storage

import (
	"encoding/json"
	"golang-test/internal/domain/entities"
	"log/slog"
	"os"
)

type FileHolidayRepository struct {
	filePath string
	logger   *slog.Logger
}

func NewFileHolidayRepository(filePath string, logger *slog.Logger) *FileHolidayRepository {
	return &FileHolidayRepository{filePath: filePath, logger: logger}
}

func (r *FileHolidayRepository) GetAll() ([]entities.Holiday, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}
	var holidays []entities.Holiday
	if err := json.Unmarshal(data, &holidays); err != nil {
		return nil, err
	}

	r.logger.Info("cache local cargado", "count", len(holidays))
	return holidays, nil

}

func (r *FileHolidayRepository) SaveAll(holidays []entities.Holiday) error {
	data, err := json.MarshalIndent(holidays, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll("data", 0755); err != nil {
		return err
	}
	return os.WriteFile(r.filePath, data, 0644)
}

func (r *FileHolidayRepository) Exists() bool {
	_, err := os.Stat(r.filePath)
	return err == nil
}
