package interfaces

import (
	"Users/internal/models/entity"
)

type UserRepository interface {
	Get() ([]*entity.UserEntity, error)
	GetOneById(id string) (*entity.UserEntity, error)
	Create(user *entity.UserEntity) error
	Delete(id string) error
	Update(id string, user *entity.UserEntity) error
}
