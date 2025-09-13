package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByField(field string, value interface{}) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db};
}

func (r *userRepository) FindByField(field string, value interface{}) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, field + " = ?", value).Error; err != nil {
		return nil, err
	}
	return &user, nil
}