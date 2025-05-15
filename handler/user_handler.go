package handler

import (
	"net/http"
	"power-track/models"
	"power-track/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

// Create cria um novo usuário
func (h *UserHandler) Create(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos para criar usuário: " + err.Error(),
		})
		return
	}

	createdUser, err := h.userService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

// GetAll retorna todos os usuários
func (h *UserHandler) GetAll(ctx *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao buscar usuários",
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// GetByID retorna um usuário pelo ID
func (h *UserHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do usuário é obrigatório",
		})
		return
	}

	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do usuário inválido",
		})
		return
	}

	user, err := h.userService.GetUserByID(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"erro": "Usuário não encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// Update atualiza os dados de um usuário
func (h *UserHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do usuário é obrigatório",
		})
		return
	}

	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do usuário inválido",
		})
		return
	}

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "Dados inválidos para atualizar usuário: " + err.Error(),
		})
		return
	}

	user.ID = uint(userID)
	err = h.userService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"mensagem": "Usuário atualizado com sucesso",
	})
}

// Delete remove um usuário pelo ID
func (h *UserHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do usuário é obrigatório",
		})
		return
	}

	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": "ID do usuário inválido",
		})
		return
	}

	err = h.userService.DeleteUser(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"mensagem": "Usuário removido com sucesso",
	})
} 