package controller

import (
	"loyalty-api/internal/dto"
	"loyalty-api/internal/models"
	service "loyalty-api/services"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	SaveUser(ctx *gin.Context) models.User
	FindAllUser() []models.User
	GetUserByEmail(email string) (dto.UserDTO, error)
	DeleteUserById(id uint) error
	UpdateUser(id uint, user models.User) (models.User, error)
}
type userController struct {
	controller service.UserService
}

// DeleteUserById implements UserController.
func (u *userController) DeleteUserById(id uint) error {
	if err := u.controller.DeleteUserById(id); err != nil {
		return err
	}
	return nil
}

// FindAllUser implements UserController.
func (u *userController) FindAllUser() []models.User {
	return u.controller.FindAllUser()
}

// GetUserByEmail implements UserController.
func (u *userController) GetUserByEmail(email string) (dto.UserDTO, error) {
	userDTO, err := u.controller.GetUserByEmail(email)
	if err != nil {
		return userDTO, err
	}
	return userDTO, nil
}

// SaveUser implements UserController.
func (u *userController) SaveUser(ctx *gin.Context) models.User {
	var user models.User
	ctx.BindJSON(&user)
	u.controller.SaveUser(user)
	return user
}

// UpdateUser implements UserController.
func (u *userController) UpdateUser(id uint, user models.User) (models.User, error) {
	panic("unimplemented")
}

// Contructeur
func NewUserController(service service.UserService) UserController {
	return &userController{
		controller: service,
	}
}
