package handler

import (
	"net/http"
	"power-track/models"
	"power-track/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StringpvHandler struct {
	stringpvService *service.StringpvService
}

func NewStringpvHandler(service *service.StringpvService) *StringpvHandler {
	return &StringpvHandler{
		stringpvService: service,
	}
}

// GetLatest retorna os dados mais recentes de uma string fotovoltaica
func (h *StringpvHandler) GetLatest(ctx *gin.Context) {
	inverterID := ctx.Param("inverterId")
	if inverterID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do inversor é obrigatório",
		})
		return
	}

	id, err := strconv.ParseUint(inverterID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do inversor inválido",
		})
		return
	}

	data, err := h.stringpvService.GetLatestData(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// GetHistorical retorna os dados históricos de uma string fotovoltaica
func (h *StringpvHandler) GetHistorical(ctx *gin.Context) {
	inverterID := ctx.Param("inverterId")
	if inverterID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do inversor é obrigatório",
		})
		return
	}
	id, err := strconv.ParseUint(inverterID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do inversor inválido",
		})
		return
	}
	startTime := ctx.Query("start_time")
	endTime := ctx.Query("end_time")
	if startTime == "" || endTime == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "start_time e end_time são obrigatórios",
		})
		return
	}
	data, err := h.stringpvService.GetHistoricalData(uint(id), startTime, endTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// GetByInverter retorna todos os dados de strings de um inversor
func (h *StringpvHandler) GetByInverter(ctx *gin.Context) {
	inverterID := ctx.Param("inverterId")
	if inverterID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do inversor é obrigatório",
		})
		return
	}

	id, err := strconv.ParseUint(inverterID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do inversor inválido",
		})
		return
	}

	data, err := h.stringpvService.GetDataByInverter(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

// Create cria um novo registro de string fotovoltaica
func (h *StringpvHandler) Create(ctx *gin.Context) {
	var stringpv models.Stringpv
	if err := ctx.ShouldBindJSON(&stringpv); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}

	createdStringpv, err := h.stringpvService.CreateStringData(&stringpv)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdStringpv)
}

// CreateMany cria múltiplos registros de strings fotovoltaicas
func (h *StringpvHandler) CreateMany(ctx *gin.Context) {
	var strings []models.Stringpv
	if err := ctx.ShouldBindJSON(&strings); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos",
		})
		return
	}

	if len(strings) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "Nenhum dado de string fotovoltaica fornecido",
		})
		return
	}

	createdStrings, err := h.stringpvService.CreateManyStringsData(strings)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdStrings)
}
