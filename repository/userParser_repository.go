package repository

import (
	"power-track/models"

	"gorm.io/gorm"
)

type UserParserRepository struct {
	db *gorm.DB
}

func NewUserParserRepository(db *gorm.DB) *UserParserRepository {
	return &UserParserRepository{
		db: db,
	}
}

// Create adiciona um novo cliente parser no banco de dados
func (r *UserParserRepository) Create(client *models.UserParser) (*models.UserParser, error) {
	if err := r.db.Create(client).Error; err != nil {
		return nil, err
	}

	// Recarrega o cliente com o relacionamento do usuário preenchido
	if err := r.db.Preload("User").First(client, client.ID).Error; err != nil {
		return nil, err
	}

	return client, nil
}

// GetByID busca um cliente parser pelo ID
func (r *UserParserRepository) GetByID(id uint) (*models.UserParser, error) {
	var client models.UserParser
	result := r.db.Preload("User").First(&client, id)
	return &client, result.Error
}

// GetAll retorna todos os clientes parser
func (r *UserParserRepository) GetAll() ([]models.UserParser, error) {
	var clients []models.UserParser
	result := r.db.Preload("User").Find(&clients)
	return clients, result.Error
}

// Update atualiza os dados de um cliente parser
func (r *UserParserRepository) Update(client *models.UserParser) error {
	return r.db.Save(client).Error
}

// Delete remove um cliente parser pelo ID
func (r *UserParserRepository) Delete(id uint) error {
	return r.db.Delete(&models.UserParser{}, id).Error
}

// GetByParserKey busca um cliente parser pela chave do parser
func (r *UserParserRepository) GetByParserKey(parserKey string) (*models.UserParser, error) {
	var client models.UserParser
	result := r.db.Where("parser_key = ?", parserKey).First(&client)
	return &client, result.Error
}

// GetByUserID retorna todos os clientes parser de um determinado usuário
func (r *UserParserRepository) GetByUserID(userID uint) ([]models.UserParser, error) {
	var clients []models.UserParser
	result := r.db.Preload("User").Where("user_id = ?", userID).Find(&clients)
	return clients, result.Error
}

// GetActive retorna todos os clientes parser ativos
func (r *UserParserRepository) GetActive() ([]models.UserParser, error) {
	var clients []models.UserParser
	result := r.db.Where("active = ?", true).Find(&clients)
	return clients, result.Error
}

// GetInactive retorna todos os clientes parser inativos
func (r *UserParserRepository) GetInactive() ([]models.UserParser, error) {
	var clients []models.UserParser
	result := r.db.Where("active = ?", false).Find(&clients)
	return clients, result.Error
}
