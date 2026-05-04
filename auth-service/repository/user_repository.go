package repository

import "unit_testku/auth-service/models"

// UserRepository adalah interface (kontrak)
type UserRepository interface {
	GetByEmail(email string) (*models.User, error)
	Create(user *models.User) error
}