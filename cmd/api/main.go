package main

import (
	usecases "golang-test/internal/application/use-cases"
	"golang-test/internal/config"
	"golang-test/internal/infrastructure/adapters"
	"golang-test/internal/infrastructure/http/router"
	"golang-test/internal/infrastructure/storage"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := config.Load()

	repo := storage.NewFileHolidayRepository(cfg.DataFilePath, logger)
	saveUseCase := usecases.NewSaveHolidayUseCase(repo)

	if !repo.Exists() {
		logger.Info("cache no encontrado, consultando API externa")
		client := adapters.NewBoostrClient(cfg.BoostrapApiUrl)
		raw, err := client.FetchHolidays()
		if err != nil {
			logger.Error("error consultando API externa", "error", err)
			os.Exit(1)
		}
		if _, err := saveUseCase.Execute(raw); err != nil {
			logger.Error("error guardando cache", "error", err)
			os.Exit(1)
		}
		logger.Info("cache creado exitosamente", "count", len(raw))
	} else {
		logger.Info("cache existente encontrado, se omite llamada a API externa")
	}

	r := router.NewRouter(repo)

	logger.Info("servidor iniciado", "port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		logger.Error("error iniciando servidor", "error", err)
		os.Exit(1)
	}
}
