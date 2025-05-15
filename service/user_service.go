package service

import (
	"power-track/models"
	"power-track/repository"
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