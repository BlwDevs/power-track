package service

import (
	"power-track/models"
	"power-track/repository"
)

type UserParserService struct {
	UserParserRepo *repository.UserParserRepository
}

func NewUserParserService(repo *repository.UserParserRepository) *UserParserService {
	return &UserParserService{
		UserParserRepo: repo,
	}
}

// CreateUserParser cria um novo cliente parser
func (s *UserParserService) CreateUserParser(client *models.UserParser) (*models.UserParser, error) {
	return s.UserParserRepo.Create(client)
}

// GetAllUserParsers retorna todos os clientes parser
func (s *UserParserService) GetAllUserParsers() ([]models.UserParser, error) {
	return s.UserParserRepo.GetAll()
}

// GetUserParserByID retorna um cliente parser pelo ID
func (s *UserParserService) GetUserParserByID(id uint) (*models.UserParser, error) {
	return s.UserParserRepo.GetByID(id)
}

// UpdateUserParser atualiza os dados de um cliente parser
func (s *UserParserService) UpdateUserParser(client *models.UserParser) error {
	return s.UserParserRepo.Update(client)
}

// DeleteUserParser remove um cliente parser pelo ID
func (s *UserParserService) DeleteUserParser(id uint) error {
	return s.UserParserRepo.Delete(id)
}

// GetUserParserByKey busca um cliente parser pela chave
func (s *UserParserService) GetUserParserByKey(key string) (*models.UserParser, error) {
	return s.UserParserRepo.GetByParserKey(key)
}

// GetUserParsersByUserID retorna todos os clientes parser de um determinado usuário
func (s *UserParserService) GetUserParsersByUserID(userID uint) ([]models.UserParser, error) {
	return s.UserParserRepo.GetByUserID(userID)
}

// GetActiveUserParsers retorna todos os clientes parser ativos
func (s *UserParserService) GetActiveUserParsers() ([]models.UserParser, error) {
	return s.UserParserRepo.GetActive()
}

// GetInactiveUserParsers retorna todos os clientes parser inativos
func (s *UserParserService) GetInactiveUserParsers() ([]models.UserParser, error) {
	return s.UserParserRepo.GetInactive()
}
