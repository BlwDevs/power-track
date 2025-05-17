package repository

import (
	"power-track/models"

	"gorm.io/gorm"
)

type InverterRepository struct {
	db *gorm.DB
}

func NewInverterRepository(db *gorm.DB) *InverterRepository {
	return &InverterRepository{
		db: db,
	}
}

// Create adiciona um novo inversor no banco de dados
func (r *InverterRepository) Create(inverter *models.Inverter) (*models.Inverter, error) {
	if err := r.db.Create(inverter).Error; err != nil {
		return nil, err
	}
	return inverter, nil
}

// GetByID busca um inversor pelo ID
func (r *InverterRepository) GetByID(id uint) (*models.Inverter, error) {
	var inverter models.Inverter
	result := r.db.First(&inverter, id)
	return &inverter, result.Error
}

// GetAll retorna todos os inversores
func (r *InverterRepository) GetAll() ([]models.Inverter, error) {
	var inverters []models.Inverter
	result := r.db.Find(&inverters)
	return inverters, result.Error
}

// Update atualiza os dados de um inversor
func (r *InverterRepository) Update(inverter *models.Inverter) error {
	return r.db.Save(inverter).Error
}

// Delete remove um inversor pelo ID
func (r *InverterRepository) Delete(id uint) error {
	return r.db.Delete(&models.Inverter{}, id).Error
}

// GetBySerialNumber busca um inversor pelo número de série
func (r *InverterRepository) GetBySerialNumber(serialNumber string) (*models.Inverter, error) {
	var inverter models.Inverter
	result := r.db.Where("serial_number = ?", serialNumber).First(&inverter)
	return &inverter, result.Error
}

// GetByStatus retorna todos os inversores com um determinado status
func (r *InverterRepository) GetByStatus(status models.Status) ([]models.Inverter, error) {
	var inverters []models.Inverter
	result := r.db.Where("status = ?", status).Find(&inverters)
	return inverters, result.Error
}
