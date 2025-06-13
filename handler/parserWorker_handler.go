package handler

import (
	"net/http"
	"power-track/models"
	"power-track/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ParserWorkerHandler struct {
	parserWorkerService *service.ParserWorkerService
}

func NewParserWorkerHandler(service *service.ParserWorkerService) *ParserWorkerHandler {
	return &ParserWorkerHandler{
		parserWorkerService: service,
	}
}

// Create cria um novo programa de processamento de dados
func (h *ParserWorkerHandler) Create(ctx *gin.Context) {
	var worker models.ParserWorker
	if err := ctx.ShouldBindJSON(&worker); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos para criar programa de processamento de dados: " + err.Error(),
		})
		return
	}

	createdWorker, err := h.parserWorkerService.CreateParserWorker(&worker)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdWorker)
}

// GetAll retorna todos os programas de processamento de dados
func (h *ParserWorkerHandler) GetAll(ctx *gin.Context) {
	workers, err := h.parserWorkerService.GetAllParserWorkers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao buscar programas de processamento de dados",
		})
		return
	}

	ctx.JSON(http.StatusOK, workers)
}

// GetByID retorna um programa de processamento de dados pelo ID
func (h *ParserWorkerHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID inválido",
		})
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID inválido: " + err.Error(),
		})
		return
	}

	worker, err := h.parserWorkerService.GetParserWorkerByID(uint(idUint))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"erro": "ParserWorker não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, worker)
}

// Update atualiza os dados de um programa de processamento de dados
func (h *ParserWorkerHandler) Update(ctx *gin.Context) {
	var worker models.ParserWorker
	if err := ctx.ShouldBindJSON(&worker); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos para atualizar programa de processamento de dados: " + err.Error(),
		})
		return
	}

	if err := h.parserWorkerService.UpdateParserWorker(&worker); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao atualizar programa de processamento de dados: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, worker)
}

// Delete remove um programa de processamento de dados pelo ID
func (h *ParserWorkerHandler) Deactivate(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID inválido",
		})
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID inválido: " + err.Error(),
		})
		return
	}

	if err := h.parserWorkerService.DeactivateParserWorker(uint(idUint)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao remover programa de processamento de dados: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *ParserWorkerHandler) Activate(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID inválido",
		})
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID inválido: " + err.Error(),
		})
		return
	}

	if err := h.parserWorkerService.ActivateParserWorker(uint(idUint)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao ativar parser worker: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// GetByManufacturer busca um programa de processamento de dados pelo fabricante
func (h *ParserWorkerHandler) GetByManufacturer(ctx *gin.Context) {
	manufacturer := ctx.Param("manufacturer")
	if manufacturer == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "Fabricante inválido",
		})
		return
	}

	worker, err := h.parserWorkerService.GetParserWorkerByManufacturer(manufacturer)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"erro": "Fabricante não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, worker)
}

// RefreshAPIKey atualiza a chave de API de um programa de processamento de dados
func (h *ParserWorkerHandler) RefreshAPIKey(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID inválido",
		})
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID inválido: " + err.Error(),
		})
		return
	}

	newAPIKey, err := h.parserWorkerService.RefreshAPIKey(uint(idUint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao atualizar chave de API: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"new_api_key": newAPIKey})
}
