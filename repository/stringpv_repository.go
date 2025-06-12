package repository

import (
	"power-track/models"

	"gorm.io/gorm"
)

// StringpvRepository é a estrutura que encapsula as operações de banco de dados
// relacionadas aos dados de strings fotovoltaicas
type StringpvRepository struct {
	db *gorm.DB
}

// NewStringpvRepository cria uma nova instância do repositório
func NewStringpvRepository(db *gorm.DB) *StringpvRepository {
	return &StringpvRepository{
		db: db,
	}
}

// Create adiciona um novo registro de string fotovoltaica no banco de dados
// e retorna o objeto criado
func (r *StringpvRepository) Create(stringpv *models.Stringpv) (*models.Stringpv, error) {
	if err := r.db.Create(stringpv).Error; err != nil {
		return nil, err
	}
	return stringpv, nil
}

// GetByInverterID retorna todas as strings fotovoltaicas associadas a um inversor específico
func (r *StringpvRepository) GetByInverterID(inverterID uint) ([]models.Stringpv, error) {
	var stringpvs []models.Stringpv
	result := r.db.Where("inverter_id = ?", inverterID).Find(&stringpvs)
	return stringpvs, result.Error
}

// GetLatestByInverterID obtém o registro mais recente de uma string fotovoltaica
// para um determinado inversor
func (r *StringpvRepository) GetLatestByInverterID(inverterID uint) (*models.Stringpv, error) {
	var stringpv models.Stringpv
	result := r.db.Where("inverter_id = ?", inverterID).
		Order("timestamp desc").
		First(&stringpv)
	return &stringpv, result.Error
}

// DeleteByInverterID remove todos os registros de strings fotovoltaicas
// associados a um determinado inversor
func (r *StringpvRepository) DeleteByInverterID(inverterID uint) error {
	return r.db.Where("inverter_id = ?", inverterID).Delete(&models.Stringpv{}).Error
}

// GetHistoricalDataByInverterID retorna os dados históricos de strings fotovoltaicas a partir de um intervalo de tempo
func (r *StringpvRepository) GetHistoricalDataByInverterID(inverterID uint, startTime, endTime string) ([]models.Stringpv, error) {
	var stringpvs []models.Stringpv
	result := r.db.Where("inverter_id = ? AND timestamp BETWEEN ? AND ?", inverterID, startTime, endTime).Find(&stringpvs)
	return stringpvs, result.Error
}

// CreateMany cria múltiplos registros de strings fotovoltaicas
func (r *StringpvRepository) CreateMany(stringpvs []models.Stringpv) ([]models.Stringpv, error) {
	if len(stringpvs) == 0 {
		return nil, nil // Retorna nil se não houver dados para inserir
	}

	if err := r.db.Create(&stringpvs).Error; err != nil {
		return nil, err
	}
	return stringpvs, nil
}
