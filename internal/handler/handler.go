package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/ulule/deepcopier"

	"Users/config"
	"Users/docs"
	"Users/internal/models/dto"
	"Users/internal/models/entity"
	"Users/internal/models/interfaces"
)

type Handler struct {
	cfg *config.Config
	rep interfaces.UserRepository
}

func InitServer(cfg *config.Config, rep interfaces.UserRepository) *Handler {
	return &Handler{
		cfg: cfg,
		rep: rep,
	}
}

func (handler *Handler) Run() error {
	router := gin.Default()
	handler.setGinMode()
	handler.configureSwagger(router)
	handler.configureRouter(router)
	return router.Run(handler.cfg.HTTPServer.Url)
}

func (handler *Handler) setGinMode() {
	mode := handler.cfg.EnvironmentVariables.Environment

	switch mode {
	case "development":
		gin.SetMode(gin.DebugMode)
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		log.Printf("Unknown server mode: %s, defaulting to 'development'", mode)
		gin.SetMode(gin.DebugMode)
	}
}

func (handler *Handler) configureRouter(r *gin.Engine) {
	r.GET("/api/v1/users", handler.HandleGet)
	r.GET("/api/v1/users/:id", handler.HandleGetOneById)
	r.POST("/api/v1/users", handler.HandleCreate)
	r.DELETE("/api/v1/users/:id", handler.HandleDelete)
	r.PUT("/api/v1/users/:id", handler.HandleUpdate)
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

func (handler *Handler) configureSwagger(router *gin.Engine) {
	docs.SwaggerInfo.Title = "Users Service API"
	docs.SwaggerInfo.Description = "This is a sample server Users server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
