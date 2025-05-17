package handler

import (
	"net/http"
	"power-track/models"
	"power-track/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserParserHandler struct {
	UserParserService *service.UserParserService
}

func NewUserParserHandler(service *service.UserParserService) *UserParserHandler {
	return &UserParserHandler{
		UserParserService: service,
	}
}

// Create cria um novo cliente parser
func (h *UserParserHandler) Create(ctx *gin.Context) {
	var client models.UserParser
	if err := ctx.ShouldBindJSON(&client); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos para criar cliente parser: " + err.Error(),
		})
		return
	}

	createdClient, err := h.UserParserService.CreateUserParser(&client)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdClient)
}

// GetAll retorna todos os clientes parser
func (h *UserParserHandler) GetAll(ctx *gin.Context) {
	clients, err := h.UserParserService.GetAllUserParsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao buscar clientes parser",
		})
		return
	}

	ctx.JSON(http.StatusOK, clients)
}

// GetByID retorna um cliente parser pelo ID
func (h *UserParserHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do cliente parser é obrigatório",
		})
		return
	}

	clientID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do cliente parser inválido",
		})
		return
	}

	client, err := h.UserParserService.GetUserParserByID(uint(clientID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"erro": "Cliente parser não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, client)
}

// Update atualiza os dados de um cliente parser
func (h *UserParserHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do cliente parser é obrigatório",
		})
		return
	}

	clientID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do cliente parser inválido",
		})
		return
	}

	var client models.UserParser
	if err := ctx.ShouldBindJSON(&client); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos para atualizar cliente parser: " + err.Error(),
		})
		return
	}

	client.ID = uint(clientID)
	err = h.UserParserService.UpdateUserParser(&client)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"mensagem": "Cliente parser atualizado com sucesso",
	})
}

// Delete remove um cliente parser pelo ID
func (h *UserParserHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do cliente parser é obrigatório",
		})
		return
	}

	clientID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do cliente parser inválido",
		})
		return
	}

	err = h.UserParserService.DeleteUserParser(uint(clientID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"mensagem": "Cliente parser removido com sucesso",
	})
}

// GetByUserID retorna todos os clientes parser de um determinado usuário
func (h *UserParserHandler) GetByUserID(ctx *gin.Context) {
	userID := ctx.Param("userId")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do usuário é obrigatório",
		})
		return
	}

	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do usuário inválido",
		})
		return
	}

	clients, err := h.UserParserService.GetUserParsersByUserID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao buscar clientes parser do usuário",
		})
		return
	}

	ctx.JSON(http.StatusOK, clients)
}
