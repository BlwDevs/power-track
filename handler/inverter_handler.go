package handler

import (
	"fmt"
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

// Create cria um novo inversor
func (h *InverterHandler) Create(ctx *gin.Context) {
	var inverter models.Inverter
	if err := ctx.ShouldBindJSON(&inverter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos para criar inversor: " + err.Error(),
		})
		return
	}

	// Debug para ver o que está chegando
	fmt.Printf("Dados recebidos: %+v\n", inverter)

	createdInverter, err := h.inverterService.CreateInverter(&inverter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdInverter)
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

// DeleteById remove um inversor pelo ID
func (h *InverterHandler) DeleteById(ctx *gin.Context) {
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

	err = h.inverterService.DeleteInverter(uint(inverterID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"mensagem": "Inversor removido com sucesso",
	})
}

// Update atualiza um inversor existente pelo ID
func (h *InverterHandler) Update(ctx *gin.Context) {
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

	var inverter models.Inverter
	if err := ctx.ShouldBindJSON(&inverter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos para atualizar inversor: " + err.Error(),
		})
		return
	}

	// Garante que o ID do struct seja o mesmo da URL
	inverter.ID = uint(inverterID)

	updatedInverter := h.inverterService.UpdateInverter(&inverter)

	ctx.JSON(http.StatusOK, updatedInverter)
}
