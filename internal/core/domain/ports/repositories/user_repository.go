package repositories

import "github.com/naphat/fiber-ecommerce-api/internal/core/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	GetByEmail(email string) (*entities.User, error)
	GetByID(id uint) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id uint) error
	GetAll() ([]entities.User, error)
}