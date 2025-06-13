package service

import (
	"crypto/rand"
	"encoding/hex"
	"power-track/models"
	"power-track/repository"
)

type ParserWorkerService struct {
	parserWorkerRepo *repository.ParserWorkerRepository
}

func NewParserWorkerService(repo *repository.ParserWorkerRepository) *ParserWorkerService {
	return &ParserWorkerService{
		parserWorkerRepo: repo,
	}
}

// CreateParserWorker cria um novo programa de processamento de dados
func (s *ParserWorkerService) CreateParserWorker(worker *models.ParserWorker) (*models.ParserWorker, error) {
	//gerar apikey unica
	var err error
	worker.ApiKey, err = generateAPIKey()
	if err != nil {
		return nil, err
	}
	return s.parserWorkerRepo.Create(worker)
}

// GetParserWorkerByID busca um programa de processamento de dados pelo ID
func (s *ParserWorkerService) GetParserWorkerByID(id uint) (*models.ParserWorker, error) {
	return s.parserWorkerRepo.GetByID(id)
}

// GetAllParserWorkers retorna todos os programas de processamento de dados
func (s *ParserWorkerService) GetAllParserWorkers() ([]models.ParserWorker, error) {
	return s.parserWorkerRepo.GetAll()
}

// UpdateParserWorker atualiza os dados de um programa de processamento de dados
func (s *ParserWorkerService) UpdateParserWorker(worker *models.ParserWorker) error {
	return s.parserWorkerRepo.Update(worker)
}

// DeleteParserWorker remove um programa de processamento de dados pelo ID
func (s *ParserWorkerService) DeactivateParserWorker(id uint) error {
	worker, err := s.parserWorkerRepo.GetByID(id)
	if err != nil {
		return err
	}
	worker.Active = false
	return s.parserWorkerRepo.Update(worker)
}

// Ativar parser worker
func (s *ParserWorkerService) ActivateParserWorker(id uint) error {
	worker, err := s.parserWorkerRepo.GetByID(id)
	if err != nil {
		return err
	}
	worker.Active = true
	return s.parserWorkerRepo.Update(worker)
}

// GetParserWorkerByManufacturer busca um programa de processamento de dados pelo fabricante
func (s *ParserWorkerService) GetParserWorkerByManufacturer(manufacturer string) (*models.ParserWorker, error) {
	return s.parserWorkerRepo.GetByManufacturer(manufacturer)
}

func generateAPIKey() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// RefreshAPIKey atualiza a chave de API de um programa de processamento de dados
func (s *ParserWorkerService) RefreshAPIKey(id uint) (string, error) {
	worker, err := s.parserWorkerRepo.GetByID(id)
	if err != nil {
		return "", err
	}
	newApiKey, err := generateAPIKey()
	if err != nil {
		return "", err
	}
	worker.ApiKey = newApiKey
	if err := s.parserWorkerRepo.Update(worker); err != nil {
		return "", err
	}
	return newApiKey, nil
}
