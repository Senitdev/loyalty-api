package service

import (
	"loyalty-api/internal/models"
	"loyalty-api/repository"
)

type UserService interface {
	SaveUser(user models.User) models.User
	FindAllUser() []models.User
	GetUserByEmail(email string) (models.User, error)
	DeleteUserById(id uint) error
	UpdateUser(id uint, user models.User) (models.User, error)
}
type userService struct {
	service repository.UserRepository
}

// DeleteUserById implements UserService.
func (u *userService) DeleteUserById(id uint) error {
	if err := u.service.DeleteUserById(id); err != nil {
		return err
	}
	return nil
}

// FindAllUser implements UserService.
func (u *userService) FindAllUser() []models.User {
	return u.service.FindAll()
}

// GetUserByEmail implements UserService.
func (u *userService) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	user, err := u.service.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// SaveUser implements UserService.
func (u *userService) SaveUser(user models.User) models.User {
	return u.service.Save(user)
}

// UpdateUser implements UserService.
func (u *userService) UpdateUser(id uint, user models.User) (models.User, error) {
	panic("unimplemented")
}

// Constructeur
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		service: repo,
	}
}
