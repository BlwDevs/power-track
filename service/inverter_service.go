package service

import (
	"errors"
	"strconv"
	"power-track/models"
	"power-track/repository"
)

type InverterService struct {
	repo *repository.InverterRepository
}

func NewInverterService(repo *repository.InverterRepository) *InverterService {
	return &InverterService{
		repo: repo,
	}
}

// GetLastData retorna os dados mais recentes do inversor
func (s *InverterService) GetLastData() (*models.Inverter, error) {
	return s.repo.GetLastData()
}

// GetHistoricalData retorna dados históricos do inversor
func (s *InverterService) GetHistoricalData(startDate, endDate string, inverterID string) ([]models.Inverter, error) {
	if startDate == "" || endDate == "" || inverterID == "" {
		return nil, errors.New("parâmetros de data e ID do inversor são obrigatórios")
	}

	// Convert string ID to uint
	id, err := strconv.ParseUint(inverterID, 10, 64)
	if err != nil {
		return nil, errors.New("ID do inversor inválido")
	}

	// Use the converted ID
	inverter, err := s.repo.GetByID(uint(id))
	if err != nil {
		return nil, err
	}

	return []models.Inverter{*inverter}, nil
}

// GetList retorna a lista de todos os inversores
func (s *InverterService) GetList() ([]models.Inverter, error) {
	return s.repo.GetAll()
}

// GetData retorna os dados de um inversor específico
func (s *InverterService) GetData(id uint) (*models.Inverter, error) {
	if id == 0 {
		return nil, errors.New("ID do inversor é obrigatório")
	}

	return s.repo.GetByID(id)
}

// CreateInverter cria um novo inversor
func (s *InverterService) CreateInverter(inverter *models.Inverter) (*models.Inverter, error) {
	if inverter.SerialNumber == "" {
		return nil, errors.New("número de série é obrigatório")
	}

	// Verifica se já existe um inversor com o mesmo número de série
	existing, err := s.repo.GetBySerialNumber(inverter.SerialNumber)
	if err == nil && existing != nil {
		return nil, errors.New("já existe um inversor com este número de série")
	}

	return s.repo.Create(inverter)
}

// UpdateInverter atualiza os dados de um inversor
func (s *InverterService) UpdateInverter(inverter *models.Inverter) error {
	if inverter.ID == 0 {
		return errors.New("ID do inversor é obrigatório")
	}

	// Verifica se o inversor existe
	existing, err := s.repo.GetByID(inverter.ID)
	if err != nil {
		return errors.New("inversor não encontrado")
	}

	// Se o número de série foi alterado, verifica se já existe outro com o mesmo número
	if inverter.SerialNumber != existing.SerialNumber {
		duplicate, err := s.repo.GetBySerialNumber(inverter.SerialNumber)
		if err == nil && duplicate != nil && duplicate.ID != inverter.ID {
			return errors.New("já existe outro inversor com este número de série")
		}
	}

	return s.repo.Update(inverter)
}

// DeleteInverter remove um inversor
func (s *InverterService) DeleteInverter(id uint) error {
	if id == 0 {
		return errors.New("ID do inversor é obrigatório")
	}

	// Verifica se o inversor existe
	_, err := s.repo.GetByID(id)
	if err != nil {
		return errors.New("inversor não encontrado")
	}

	return s.repo.Delete(id)
}