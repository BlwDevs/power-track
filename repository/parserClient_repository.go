package repository

import (
	"power-track/models"

	"gorm.io/gorm"
)

type ParserClientRepository struct {
	db *gorm.DB
}

func NewParserClientRepository(db *gorm.DB) *ParserClientRepository {
	return &ParserClientRepository{
		db: db,
	}
}

// Create adiciona um novo cliente parser no banco de dados
func (r *ParserClientRepository) Create(client *models.ParserClient) (*models.ParserClient, error) {
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
func (r *ParserClientRepository) GetByID(id uint) (*models.ParserClient, error) {
	var client models.ParserClient
	result := r.db.Preload("User").First(&client, id)
	return &client, result.Error
}

// GetAll retorna todos os clientes parser
func (r *ParserClientRepository) GetAll() ([]models.ParserClient, error) {
	var clients []models.ParserClient
	result := r.db.Preload("User").Find(&clients)
	return clients, result.Error
}

// Update atualiza os dados de um cliente parser
func (r *ParserClientRepository) Update(client *models.ParserClient) error {
	return r.db.Save(client).Error
}

// Delete remove um cliente parser pelo ID
func (r *ParserClientRepository) Delete(id uint) error {
	return r.db.Delete(&models.ParserClient{}, id).Error
}

// GetByParserKey busca um cliente parser pela chave do parser
func (r *ParserClientRepository) GetByParserKey(parserKey string) (*models.ParserClient, error) {
	var client models.ParserClient
	result := r.db.Where("parser_key = ?", parserKey).First(&client)
	return &client, result.Error
}

// GetByUserID retorna todos os clientes parser de um determinado usuário
func (r *ParserClientRepository) GetByUserID(userID uint) ([]models.ParserClient, error) {
	var clients []models.ParserClient
	result := r.db.Preload("User").Where("user_id = ?", userID).Find(&clients)
	return clients, result.Error
}

// GetActive retorna todos os clientes parser ativos
func (r *ParserClientRepository) GetActive() ([]models.ParserClient, error) {
	var clients []models.ParserClient
	result := r.db.Where("active = ?", true).Find(&clients)
	return clients, result.Error
}

// GetInactive retorna todos os clientes parser inativos
func (r *ParserClientRepository) GetInactive() ([]models.ParserClient, error) {
	var clients []models.ParserClient
	result := r.db.Where("active = ?", false).Find(&clients)
	return clients, result.Error
} 