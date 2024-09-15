package handler

import (
	"Users/internal/models/dto"
	"Users/internal/models/entity"
	"Users/internal/models/interfaces"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ulule/deepcopier"
)

type Handler struct {
	rep interfaces.UserRepository
}

func InitHandler(rep interfaces.UserRepository) *Handler {
	return &Handler{
		rep: rep,
	}
}

func (h *Handler) ConfigureRoutes(r *gin.Engine) {
	r.GET("/api/v1/users", h.HandleGet)
	r.GET("/api/v1/users/:id", h.HandleGetOneById)
	r.POST("/api/v1/users", h.HandleCreate)
	r.DELETE("/api/v1/users/:id", h.HandleDelete)
	r.PUT("/api/v1/users/:id", h.HandleUpdate)
}

// HandleGet - godoc
// @Summary List users
// @Description get users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} dto.UserDto
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/users [get]
func (handler *Handler) HandleGet(c *gin.Context) {
	users, err := handler.rep.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: fmt.Sprintf("Error retrieving users: %v", err)})
		return
	}
	c.JSON(http.StatusOK, users)
}

// HandleGetOneById - godoc
// @Summary Get user by ID
// @Description get user by id
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dto.UserDto
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/users/{id} [get]
func (handler *Handler) HandleGetOneById(c *gin.Context) {
	id := c.Param("id")

	user, err := handler.rep.GetOneByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// HandleCreate - godoc
// @Summary Create a new user
// @Description create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserDto true "User info"
// @Success 201 {string} string "User created successfully"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/users [post]
func (handler *Handler) HandleCreate(c *gin.Context) {
	var userCreateDto dto.CreateUserDto
	var userEntity entity.UserEntity

	if err := c.ShouldBindJSON(&userCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: fmt.Sprintf("Error decoding request body: %v", err)})
		return
	}

	if err := deepcopier.Copy(&userCreateDto).To(&userEntity); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: fmt.Sprintf("Error mapping user: %v", err)})
		return
	}

	if err := handler.rep.Create(&userEntity); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: fmt.Sprintf("Error creating user: %v", err)})
		return
	}

	// Возвращаем ответ
	c.String(http.StatusCreated, "User created successfully")
}

// HandleDelete - godoc
// @Summary Delete user by ID
// @Description delete user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string "User deleted successfully"
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/users/{id} [delete]
func (handler *Handler) HandleDelete(c *gin.Context) {
	id := c.Param("id")

	if err := handler.rep.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "User not found"})
		return
	}

	c.String(http.StatusOK, "User deleted successfully")
}

// HandleUpdate - godoc
// @Summary Update user by ID
// @Description update user`
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body dto.UpdateUserDto true "User info"
// @Success 200 {string} string "User updated successfully"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/users/{id} [put]
func (handler *Handler) HandleUpdate(c *gin.Context) {

	id := c.Param("id")

	var userUpdateDto dto.UpdateUserDto
	var userEntity entity.UserEntity

	if err := c.ShouldBindJSON(&userUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: fmt.Sprintf("Error decoding request body: %v", err)})
		return
	}

	if err := deepcopier.Copy(&userUpdateDto).To(&userEntity); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: fmt.Sprintf("Error mapping user: %v", err)})
		return
	}

	if err := handler.rep.Update(id, &userEntity); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: fmt.Sprintf("Error updating user: %v", err)})
		return
	}

	c.String(http.StatusOK, "User updated successfully")
}
