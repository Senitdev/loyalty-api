package repository

import (
	"loyalty-api/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user models.User) models.User
	FindAll() []models.User
	GetUserByEmail(email string) (models.User, error)
	DeleteUserById(id uint) error
	UpdateUser(id uint, user models.User) (models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

// DeleteUserById implements UserRepository.
func (u *userRepository) DeleteUserById(id uint) error {
	result := u.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindAll implements UserRepository.
func (u *userRepository) FindAll() []models.User {
	var user []models.User
	result := u.DB.Find(&user)
	if result.Error != nil {
		return user
	}
	return user
}

// GetUserByEmail implements UserRepository.
func (u *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := u.DB.Where("email", email).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

// Save implements UserRepository.
func (u *userRepository) Save(user models.User) models.User {
	result := u.DB.Save(&user)
	if result.Error != nil {
		return user
	}
	return user
}

// UpdateUser implements UserRepository.
func (u *userRepository) UpdateUser(id uint, user models.User) (models.User, error) {
	panic("unimplemented")
}

// Constructeur
func NewUserRepository(repo *gorm.DB) UserRepository {
	return &userRepository{
		DB: repo,
	}
}
