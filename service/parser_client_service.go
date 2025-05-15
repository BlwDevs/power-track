package service

import (
	"power-track/models"
	"power-track/repository"
)

type ParserClientService struct {
	parserClientRepo *repository.ParserClientRepository
}

func NewParserClientService(repo *repository.ParserClientRepository) *ParserClientService {
	return &ParserClientService{
		parserClientRepo: repo,
	}
}

// CreateParserClient cria um novo cliente parser
func (s *ParserClientService) CreateParserClient(client *models.ParserClient) (*models.ParserClient, error) {
	return s.parserClientRepo.Create(client)
}

// GetAllParserClients retorna todos os clientes parser
func (s *ParserClientService) GetAllParserClients() ([]models.ParserClient, error) {
	return s.parserClientRepo.GetAll()
}

// GetParserClientByID retorna um cliente parser pelo ID
func (s *ParserClientService) GetParserClientByID(id uint) (*models.ParserClient, error) {
	return s.parserClientRepo.GetByID(id)
}

// UpdateParserClient atualiza os dados de um cliente parser
func (s *ParserClientService) UpdateParserClient(client *models.ParserClient) error {
	return s.parserClientRepo.Update(client)
}

// DeleteParserClient remove um cliente parser pelo ID
func (s *ParserClientService) DeleteParserClient(id uint) error {
	return s.parserClientRepo.Delete(id)
}

// GetParserClientByKey busca um cliente parser pela chave
func (s *ParserClientService) GetParserClientByKey(key string) (*models.ParserClient, error) {
	return s.parserClientRepo.GetByParserKey(key)
}

// GetParserClientsByUserID retorna todos os clientes parser de um determinado usuário
func (s *ParserClientService) GetParserClientsByUserID(userID uint) ([]models.ParserClient, error) {
	return s.parserClientRepo.GetByUserID(userID)
}

// GetActiveParserClients retorna todos os clientes parser ativos
func (s *ParserClientService) GetActiveParserClients() ([]models.ParserClient, error) {
	return s.parserClientRepo.GetActive()
}

// GetInactiveParserClients retorna todos os clientes parser inativos
func (s *ParserClientService) GetInactiveParserClients() ([]models.ParserClient, error) {
	return s.parserClientRepo.GetInactive()
} 