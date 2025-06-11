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

// Login autentica um usuário e retorna um token JWT
func (h *UserHandler) Login(c *gin.Context) {

	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	email := loginData.Email
	password := loginData.Password

	println("Email:", email)
	println("Password:", password)

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	authUser, err := h.userService.Authenticate(email, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	// Generate JWT token
	token, err := h.userService.GenerateToken(authUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		println("Error generating token:", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Logout invalidates the user's session
func (h *UserHandler) Logout(c *gin.Context) {
	userId, exists := c.Get("userID")
	if !exists || userId == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userID, ok := userId.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}
	err := h.userService.InvalidateToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not invalidate token"})
		println("Error invalidating token:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
