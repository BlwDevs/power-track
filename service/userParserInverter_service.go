package service

import (
	"power-track/models"
	"power-track/repository"
)

type UserParserInverterService struct {
	UserParserInverterRepo *repository.UserParserInverterRepository
}

func NewUserParserInverterService(repo *repository.UserParserInverterRepository) *UserParserInverterService {
	return &UserParserInverterService{
		UserParserInverterRepo: repo,
	}
}

// CreateUserParserInverter cria um novo cliente parser
func (s *UserParserInverterService) CreateUserParserInverter(client *models.UserParserInverter) (*models.UserParserInverter, error) {
	return s.UserParserInverterRepo.Create(client)
}

// GetAllUserParserInverters retorna todos os clientes parser
func (s *UserParserInverterService) GetAllUserParserInverters() ([]models.UserParserInverter, error) {
	return s.UserParserInverterRepo.GetAll()
}

// GetUserParserInverterByID retorna um cliente parser pelo ID
func (s *UserParserInverterService) GetUserParserInverterByID(id uint) (*models.UserParserInverter, error) {
	return s.UserParserInverterRepo.GetByID(id)
}

// UpdateUserParserInverter atualiza os dados de um cliente parser
func (s *UserParserInverterService) UpdateUserParserInverter(client *models.UserParserInverter) error {
	return s.UserParserInverterRepo.Update(client)
}

// DeleteUserParserInverter remove um cliente parser pelo ID
func (s *UserParserInverterService) DeleteUserParserInverter(id uint) error {
	return s.UserParserInverterRepo.Delete(id)
}

// GetUserParserInverterByKey busca um cliente parser pela chave
func (s *UserParserInverterService) GetUserParserInverterByKey(key string) (*models.UserParserInverter, error) {
	return s.UserParserInverterRepo.GetByParserKey(key)
}

// GetUserParserInvertersByUserID retorna todos os clientes parser de um determinado usuário
func (s *UserParserInverterService) GetUserParserInvertersByUserID(userID uint) ([]models.UserParserInverter, error) {
	return s.UserParserInverterRepo.GetByUserID(userID)
}

// GetActiveUserParserInverters retorna todos os clientes parser ativos
func (s *UserParserInverterService) GetActiveUserParserInverters() ([]models.UserParserInverter, error) {
	return s.UserParserInverterRepo.GetActive()
}

// GetInactiveUserParserInverters retorna todos os clientes parser inativos
func (s *UserParserInverterService) GetInactiveUserParserInverters() ([]models.UserParserInverter, error) {
	return s.UserParserInverterRepo.GetInactive()
}
