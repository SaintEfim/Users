package controller

import (
	"context"
	"fmt"

	"Users/internal/models/entity"
	"Users/internal/models/interfaces"
)

type Controller struct {
	rep interfaces.Repository
}

func NewController(rep interfaces.Repository) interfaces.Controller {
	return &Controller{rep: rep}
}

func (c *Controller) Get(ctx context.Context) ([]*entity.UserEntity, error) {
	users, err := c.rep.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("error retrieving users: %v", err)
	}
	return users, nil
}

func (c *Controller) GetOneById(ctx context.Context, id string) (*entity.UserEntity, error) {
	user, err := c.rep.GetOneById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user with id %s: %v", id, err)
	}
	return user, nil
}

func (c *Controller) Create(ctx context.Context, user *entity.UserEntity) error {
	err := c.rep.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	return nil
}

func (c *Controller) Delete(ctx context.Context, id string) error {
	err := c.rep.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("error deleting user with id %s: %v", id, err)
	}
	return nil
}

func (c *Controller) Update(ctx context.Context, id string, user *entity.UserEntity) error {
	err := c.rep.Update(ctx, id, user)
	if err != nil {
		return fmt.Errorf("error updating user with id %s: %v", id, err)
	}
	return nil
}
