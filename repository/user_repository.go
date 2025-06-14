package repository

import (
	"power-track/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create adiciona um novo usuário no banco de dados
func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	//cria hash da senha
	if user.Password != "" {
		hashedPassword, err := user.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword

		if err := r.db.Create(user).Error; err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}

// GetByID busca um usuário pelo ID
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	return &user, result.Error
}

// GetAll retorna todos os usuários
func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	return users, result.Error
}

// Update atualiza os dados de um usuário
func (r *UserRepository) Update(user *models.User) error {
	return r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error
}

// Delete remove um usuário pelo ID
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// GetByCPF busca um usuário pelo CPF
func (r *UserRepository) GetByCPF(cpf string) (*models.User, error) {
	var user models.User
	result := r.db.Where("cpf = ?", cpf).First(&user)
	return &user, result.Error
}

// GetByEmail busca um usuário pelo email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}

// GetByPlan retorna todos os usuários com um determinado plano
func (r *UserRepository) GetByPlan(plan models.Plan) ([]models.User, error) {
	var users []models.User
	result := r.db.Where("plan = ?", plan).Find(&users)
	return users, result.Error
}
