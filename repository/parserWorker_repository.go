package repository

import (
	"power-track/models"

	"gorm.io/gorm"
)

type ParserWorkerRepository struct {
	db *gorm.DB
}

func NewParserWorkerRepository(db *gorm.DB) *ParserWorkerRepository {
	return &ParserWorkerRepository{
		db: db,
	}
}

// Create adiciona um novo programa de processamento de dados
func (r *ParserWorkerRepository) Create(parserWorker *models.ParserWorker) (*models.ParserWorker, error) {
	if err := r.db.Create(parserWorker).Error; err != nil {
		return nil, err
	}
	return parserWorker, nil
}

// GetByID busca um programa de processamento de dados pelo ID
func (r *ParserWorkerRepository) GetByID(id uint) (*models.ParserWorker, error) {
	var parserWorker models.ParserWorker
	result := r.db.First(&parserWorker, id)
	return &parserWorker, result.Error
}

// GetAll retorna todos os programas de processamento de dados
func (r *ParserWorkerRepository) GetAll() ([]models.ParserWorker, error) {
	var parserWorkers []models.ParserWorker
	result := r.db.Find(&parserWorkers)
	return parserWorkers, result.Error
}

// Update atualiza os dados de um programa de processamento de dados
func (r *ParserWorkerRepository) Update(parserWorker *models.ParserWorker) error {
	return r.db.Save(parserWorker).Error
}

// Delete remove um programa de processamento de dados pelo ID
func (r *ParserWorkerRepository) Delete(id uint) error {
	return r.db.Delete(&models.ParserWorker{}, id).Error
}

// GetByManufacturer busca um programa de processamento de dados pelo fabricante
func (r *ParserWorkerRepository) GetByManufacturer(manufacturer string) (*models.ParserWorker, error) {
	var parserWorker models.ParserWorker
	result := r.db.Where("manufacturer = ?", manufacturer).First(&parserWorker)
	return &parserWorker, result.Error
}

// Valida a chave de API de um programa de processamento de dados
func (r *ParserWorkerRepository) ValidateAPIKey(apiKey string) bool {
	var parserWorker models.ParserWorker
	if apiKey == "" {
		return false
	}
	result := r.db.Where("api_key = ?", apiKey).First(&parserWorker)
	return result.Error == nil && parserWorker.ID != 0
}
