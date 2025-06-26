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
func (r *InverterRepository) GetByStatus(status string) ([]models.Inverter, error) {
	var inverters []models.Inverter
	result := r.db.Where("status = ?", status).Find(&inverters)
	return inverters, result.Error
}

// GetByUserID busca todos os inversores associados a um usuário específico
func (r *InverterRepository) GetByUserID(userID uint) ([]models.Inverter, error) {
	var userParsers []models.UserParserInverter
	// Passo 1: Buscar todos os UserParserInverter do usuário
	if err := r.db.Where("user_id = ?", userID).Find(&userParsers).Error; err != nil {
		return nil, err
	}

	// Passo 2: Extrair os inverter_ids
	var inverterIDs []uint
	for _, up := range userParsers {
		inverterIDs = append(inverterIDs, up.InverterID)
	}

	if len(inverterIDs) == 0 {
		return []models.Inverter{}, nil // Nenhum inversor associado
	}

	// Passo 3: Buscar os inversores pelos IDs
	var inverters []models.Inverter
	if err := r.db.Where("id IN ?", inverterIDs).Find(&inverters).Error; err != nil {
		return nil, err
	}

	return inverters, nil
}
