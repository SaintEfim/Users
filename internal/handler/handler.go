package handler

import (
	"fmt"
	"net/http"

	"Users/internal/models/dto"
	"Users/internal/models/entity"
	"Users/internal/models/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

type Handler struct {
	controller interfaces.Controller
}

func NewHandler(controller interfaces.Controller) interfaces.Handler {
	return &Handler{controller: controller}
}

func (h *Handler) ConfigureRoutes(r *gin.Engine) {
	r.GET("/api/v1/users", h.Get)
	r.GET("/api/v1/users/:id", h.GetOneById)
	r.POST("/api/v1/users", h.Create)
	r.DELETE("/api/v1/users/:id", h.Delete)
	r.PUT("/api/v1/users/:id", h.Update)
}

// Get - godoc
// @Summary List users
// @Description get users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} dto.Response{data=[]dto.UserDto} "Successful response"
// @Failure 500 {object} dto.Response
// @Router /api/v1/users [get]
func (h *Handler) Get(c *gin.Context) {
	users, err := h.controller.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error retrieving users: %v", err)})
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: users})
}

// GetOneById - godoc
// @Summary Get user by ID
// @Description get user by id
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dto.Response{data=dto.UserDto} "Successful response"
// @Failure 404 {object} dto.Response
// @Router /api/v1/users/{id} [get]
func (h *Handler) GetOneById(c *gin.Context) {
	id := c.Param("id")

	user, err := h.controller.GetOneById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Response{Message: "User not found"})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: user})
}

// Create - godoc
// @Summary Create a new user
// @Description create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDto true "User info"
// @Success 201 {object} dto.Response{data=dto.UserDto} "User created successfully"
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/users [post]
func (h *Handler) Create(c *gin.Context) {
	var userCreateDto dto.CreateUserDto
	var userEntity entity.UserEntity

	if err := c.ShouldBindJSON(&userCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: fmt.Sprintf("Error decoding request body: %v", err)})
		return
	}

	if err := deepcopier.Copy(&userCreateDto).To(&userEntity); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error mapping user: %v", err)})
		return
	}

	if err := h.controller.Create(&userEntity); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error creating user: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Message: "User created successfully",
		Data:    userEntity.Id,
	})
}

// Delete - godoc
// @Summary Delete user by ID
// @Description delete user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dto.Response "User deleted successfully"
// @Failure 404 {object} dto.Response
// @Router /api/v1/users/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.controller.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, dto.Response{Message: "User not found"})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Message: "User deleted successfully"})
}

// Update - godoc
// @Summary Update user by ID
// @Description update user`
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body dto.UpdateUserDto true "User info"
// @Success 200 {object} dto.Response{data=dto.UserDto} "User updated successfully"
// @Failure 400 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /api/v1/users/{id} [put]
func (h *Handler) Update(c *gin.Context) {

	id := c.Param("id")

	var userUpdateDto dto.UpdateUserDto
	var userEntity entity.UserEntity

	if err := c.ShouldBindJSON(&userUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Message: fmt.Sprintf("Error decoding request body: %v", err)})
		return
	}

	if err := deepcopier.Copy(&userUpdateDto).To(&userEntity); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error mapping user: %v", err)})
		return
	}

	if err := h.controller.Update(id, &userEntity); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Message: fmt.Sprintf("Error updating user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Message: "User updated successfully"})
}
