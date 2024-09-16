package controller

import (
	"Users/internal/models/entity"
	"Users/internal/models/interfaces"
)

type Controller struct {
	rep interfaces.UserRepository
}

func InitController(rep interfaces.UserRepository) *Controller {
	return &Controller{rep: rep}
}

func (c *Controller) Get() ([]*entity.UserEntity, error) {
	return c.rep.Get()
}

func (c *Controller) GetOneById(id string) (*entity.UserEntity, error) {
	return c.rep.GetOneByID(id)
}

func (c *Controller) Create(user *entity.UserEntity) error {
	return c.rep.Create(user)
}

func (c *Controller) Delete(id string) error {
	return c.rep.Delete(id)
}

func (c *Controller) Update(id string, user *entity.UserEntity) error {
	return c.rep.Update(id, user)
}
