package repository

import (
	"encoding/json"
	"power-track/models"

	"gorm.io/gorm"
)

type UserParserInverterRepository struct {
	db *gorm.DB
}

func NewUserParserInverterRepository(db *gorm.DB) *UserParserInverterRepository {
	return &UserParserInverterRepository{
		db: db,
	}
}

// Create adiciona um novo cliente parser no banco de dados
func (r *UserParserInverterRepository) Create(client *models.UserParserInverter) (*models.UserParserInverter, error) {
	if err := r.db.Create(client).Error; err != nil {
		return nil, err
	}

	// Recarrega o cliente com o relacionamento do usuário, inverter e parserworker preenchido
	if err := r.db.Preload("User").Preload("Inverter").Preload("ParserWorker").First(client, client.ID).Error; err != nil {
		return nil, err
	}

	return client, nil
}

// GetByID busca um cliente parser pelo ID
func (r *UserParserInverterRepository) GetByID(id uint) (*models.UserParserInverter, error) {
	var client models.UserParserInverter
	result := r.db.Preload("User").First(&client, id)
	return &client, result.Error
}

// GetAll retorna todos os clientes parser
func (r *UserParserInverterRepository) GetAll() ([]models.UserParserInverter, error) {
	var clients []models.UserParserInverter
	result := r.db.Preload("User").Find(&clients)
	return clients, result.Error
}

// Update atualiza os dados de um cliente parser
func (r *UserParserInverterRepository) Update(client *models.UserParserInverter) error {
	return r.db.Save(client).Error
}

// Delete remove um cliente parser pelo ID
func (r *UserParserInverterRepository) Delete(id uint) error {
	return r.db.Delete(&models.UserParserInverter{}, id).Error
}

// GetByParserKey busca um cliente parser pela chave do parser
func (r *UserParserInverterRepository) GetByParserKey(parserKey string) (*models.UserParserInverter, error) {
	var client models.UserParserInverter
	result := r.db.Where("parser_key = ?", parserKey).First(&client)
	return &client, result.Error
}

// GetByUserID retorna todos os clientes parser de um determinado usuário
func (r *UserParserInverterRepository) GetByUserID(userID uint) ([]models.UserParserInverter, error) {
	var clients []models.UserParserInverter
	result := r.db.Preload("User").Where("user_id = ?", userID).Find(&clients)
	return clients, result.Error
}

// GetActive retorna todos os clientes parser ativos
func (r *UserParserInverterRepository) GetActive() ([]models.UserParserInverter, error) {
	var clients []models.UserParserInverter
	result := r.db.Where("active = ?", true).Find(&clients)
	return clients, result.Error
}

// GetInactive retorna todos os clientes parser inativos
func (r *UserParserInverterRepository) GetInactive() ([]models.UserParserInverter, error) {
	var clients []models.UserParserInverter
	result := r.db.Where("active = ?", false).Find(&clients)
	return clients, result.Error
}

// GetGrowattData retorna todos os clientes parser ativos do fabricante Growatt
func (r *UserParserInverterRepository) GetGrowattData() ([]map[string]interface{}, error) {
	var userParsers []models.UserParserInverter

	err := r.db.
		Preload("ParserWorker").
		Preload("Inverter").
		Joins("JOIN parser_workers ON parser_workers.id = user_parser_inverters.parser_worker_id").
		Where("parser_workers.manufacturer = ?", "Growatt").
		Where("user_parser_inverters.active = ?", true).
		Find(&userParsers).Error

	if err != nil {
		return nil, err
	}

	var growattUsers []map[string]interface{}

	for _, up := range userParsers {
		// Decodifica ParserParams
		var params map[string]interface{}
		if err := json.Unmarshal([]byte(up.ParserParams), &params); err != nil {
			continue // Ignora registros com JSON inválido
		}

		growattUser := map[string]interface{}{
			"id":            up.ID,
			"api_token":     "", // Preencher quando implementar
			"sn":            params["sn"],
			"device_type":   params["device_type"],
			"inverter_id":   up.InverterID,
			"growatt_token": params["growatt_token"],
			"stringsNum":    params["stringsNum"],
		}
		growattUsers = append(growattUsers, growattUser)
	}

	return growattUsers, nil
}
