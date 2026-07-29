package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"golang-test/internal/application/dtos"
	"golang-test/internal/application/ports"
	"golang-test/internal/domain/entities"
)

type HolidayHandler struct {
	getAllUseCase ports.IGetAllHolidayUseCase
}

func NewHolidayHandler(uc ports.IGetAllHolidayUseCase) *HolidayHandler {
	return &HolidayHandler{getAllUseCase: uc}
}

// GetHolidays godoc
// @Summary      Obtener feriados
// @Description  Obtiene la lista de feriados filtrados por tipo (Civil, Religioso) y/o rango de fechas (startDate, endDate). Soporta respuestas en JSON y XML.
// @Tags         holidays
// @Accept       json,xml
// @Produce      json,xml
// @Param        type       query     string  false  "Tipo de feriado (Civil, Religioso)"
// @Param        startDate  query     string  false  "Fecha inicial (Formato YYYY-MM-DD)"
// @Param        endDate    query     string  false  "Fecha final (Formato YYYY-MM-DD)"
// @Success      200        {object}  dtos.HolidaysResponse
// @Failure      400        {object}  map[string]string
// @Router       /holidays [get]
func (h *HolidayHandler) GetHolidays(c *gin.Context) {
	filters := &entities.HolidayFilter{
		Type:      c.Query("type"),
		StartDate: c.Query("startDate"),
		EndDate:   c.Query("endDate"),
	}

	result, err := h.getAllUseCase.Execute(filters)
	if err != nil {
		c.Negotiate(http.StatusBadRequest, gin.Negotiate{
			Offered: []string{gin.MIMEJSON, gin.MIMEXML},
			Data:    gin.H{"error": err.Error()},
		})
		return
	}

	response := dtos.HolidaysResponse{Holidays: result}

	c.Negotiate(http.StatusOK, gin.Negotiate{
		Offered: []string{gin.MIMEJSON, gin.MIMEXML},
		Data:    response,
	})
}
