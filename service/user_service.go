package service

import (
	"errors"
	"os"
	"power-track/models"
	"power-track/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

// CreateUser cria um novo usuário
func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	return s.userRepo.Create(user)
}

// GetAllUsers retorna todos os usuários
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAll()
}

// GetUserByID retorna um usuário pelo ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

// UpdateUser atualiza os dados de um usuário
func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

// DeleteUser remove um usuário pelo ID
func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

// GetUserByCPF busca um usuário pelo CPF
func (s *UserService) GetUserByCPF(cpf string) (*models.User, error) {
	return s.userRepo.GetByCPF(cpf)
}

// GetUserByEmail busca um usuário pelo email
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetByEmail(email)
}

// GetUsersByPlan retorna todos os usuários com um determinado plano
func (s *UserService) GetUsersByPlan(plan models.Plan) ([]models.User, error) {
	return s.userRepo.GetByPlan(plan)
}

// Realiza a autenticação do usuário
func (s *UserService) Authenticate(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("credenciais inválidas")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil { // Compare a senha fornecida com a senha armazenada
		return nil, errors.New("credenciais inválidas")
	}

	return user, nil
}

// GenerateToken gera um token JWT para o usuário
func (s *UserService) GenerateToken(user *models.User) (string, error) {
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := generateToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Invalida o token JWT do usuário
func (s *UserService) InvalidateToken(tokenString string) error {
	// Aqui você pode implementar a lógica para invalidar o token, como armazená-lo em um banco de dados ou cache
	// Por simplicidade, vamos apenas retornar nil, pois a invalidação real depende da implementação do middleware
	return nil
}
