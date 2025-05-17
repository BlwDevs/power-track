package service

import (
	"errors"
	"power-track/models"
	"power-track/repository"
)

type StringpvService struct {
	repo *repository.StringpvRepository
}

func NewStringpvService(repo *repository.StringpvRepository) *StringpvService {
	return &StringpvService{
		repo: repo,
	}
}

// CreateStringData adds new photovoltaic string data
func (s *StringpvService) CreateStringData(stringpv *models.Stringpv) (*models.Stringpv, error) {
	if stringpv.InverterID == 0 {
		return nil, errors.New("inverter ID is required")
	}

	return s.repo.Create(stringpv)
}

// GetDataByInverter returns all string data for a specific inverter
func (s *StringpvService) GetDataByInverter(inverterID uint) ([]models.Stringpv, error) {
	if inverterID == 0 {
		return nil, errors.New("inverter ID is required")
	}

	return s.repo.GetByInverterID(inverterID)
}

// GetHistoricalData retrieves historical data for a photovoltaic string
func (s *StringpvService) GetHistoricalData(inverterID uint, startTime, endTime string) ([]models.Stringpv, error) {
	if inverterID == 0 {
		return nil, errors.New("inverter ID is required")
	}

	return s.repo.GetHistoricalDataByInverterID(inverterID, startTime, endTime)
}

// GetLatestData gets the most recent data for a photovoltaic string
func (s *StringpvService) GetLatestData(inverterID uint) (*models.Stringpv, error) {
	if inverterID == 0 {
		return nil, errors.New("inverter ID is required")
	}

	data, err := s.repo.GetLatestByInverterID(inverterID)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("no data found for this inverter")
	}

	return data, nil
}

// DeleteInverterData removes all string data for an inverter
func (s *StringpvService) DeleteInverterData(inverterID uint) error {
	if inverterID == 0 {
		return errors.New("inverter ID is required")
	}

	return s.repo.DeleteByInverterID(inverterID)
}
