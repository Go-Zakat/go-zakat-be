package repository

import "go-zakat/internal/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
	FindByGoogleID(googleID string) (*entity.User, error)
	Update(user *entity.User) error
}
