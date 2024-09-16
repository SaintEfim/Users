package controller

import (
	"Users/internal/models/entity"
	"Users/internal/models/interfaces"
	"fmt"
)

type Controller struct {
	rep interfaces.UserRepository
}

func InitController(rep interfaces.UserRepository) *Controller {
	return &Controller{rep: rep}
}

func (c *Controller) Get() ([]*entity.UserEntity, error) {
	users, err := c.rep.Get()
	if err != nil {
		return nil, fmt.Errorf("error retrieving users: %v", err) // Ошибка при получении пользователей
	}
	return users, nil
}

func (c *Controller) GetOneById(id string) (*entity.UserEntity, error) {
	user, err := c.rep.GetOneById(id)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user with id %s: %v", id, err) // Ошибка при получении пользователя по ID
	}
	return user, nil
}

func (c *Controller) Create(user *entity.UserEntity) error {
	err := c.rep.Create(user)
	if err != nil {
		return fmt.Errorf("error creating user: %v", err) // Ошибка при создании пользователя
	}
	return nil
}

func (c *Controller) Delete(id string) error {
	err := c.rep.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting user with id %s: %v", id, err) // Ошибка при удалении пользователя по ID
	}
	return nil
}

func (c *Controller) Update(id string, user *entity.UserEntity) error {
	err := c.rep.Update(id, user)
	if err != nil {
		return fmt.Errorf("error updating user with id %s: %v", id, err) // Ошибка при обновлении пользователя
	}
	return nil
}
