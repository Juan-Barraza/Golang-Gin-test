package router

import (
	usecases "golang-test/internal/application/use-cases"
	"golang-test/internal/domain/repositories"
	"golang-test/internal/infrastructure/http/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(repo repositories.IHolidayRepository) *gin.Engine {
	router := gin.Default()

	getAllUseCase := usecases.NewGetAllHolidayUseCase(repo)
	holidayHandler := handler.NewHolidayHandler(getAllUseCase)

	router.GET("/holidays", holidayHandler.GetHolidays)

	return router
}
