package interfaces

import (
	"context"

	"Users/internal/models/entity"
)

type Repository interface {
	Get(ctx context.Context) ([]*entity.UserEntity, error)
	GetOneById(ctx context.Context, id string) (*entity.UserEntity, error)
	Create(ctx context.Context, user *entity.UserEntity) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, user *entity.UserEntity) error
}
