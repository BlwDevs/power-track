package handler

import (
	"net/http"
	"power-track/models"
	"power-track/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InverterHandler struct {
	inverterService *service.InverterService
}

func NewInverterHandler(service *service.InverterService) *InverterHandler {
	return &InverterHandler{
		inverterService: service,
	}
}

// GetLastData retorna os dados mais recentes do inversor
func (h *InverterHandler) GetLastData(ctx *gin.Context) {
	data, err := h.inverterService.GetLastData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao buscar últimos dados",
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// GetHistoricalData retorna dados históricos do inversor
func (h *InverterHandler) GetHistoricalData(ctx *gin.Context) {
	startDate := ctx.Query("dataInicio")
	endDate := ctx.Query("dataFim")
	inverterID := ctx.Query("inversorId")

	data, err := h.inverterService.GetHistoricalData(startDate, endDate, inverterID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// GetList retorna a lista de todos os inversores
func (h *InverterHandler) GetList(ctx *gin.Context) {
	inverters, err := h.inverterService.GetList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao buscar lista de inversores",
		})
		return
	}

	ctx.JSON(http.StatusOK, inverters)
}

// GetData retorna os dados de um inversor específico
func (h *InverterHandler) GetData(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do inversor é obrigatório",
		})
		return
	}

	inverterID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do inversor inválido",
		})
		return
	}

	data, err := h.inverterService.GetData(uint(inverterID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao buscar dados do inversor",
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}