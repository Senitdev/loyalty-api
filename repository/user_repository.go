package repository

import (
	"fmt"
	"loyalty-api/internal/dto"
	"loyalty-api/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user models.User) models.User
	FindAll() []models.User
	GetUserByEmail(email string) (dto.UserDTO, error)
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
func (u *userRepository) GetUserByEmail(email string) (dto.UserDTO, error) {
	var user models.User
	var userDto dto.UserDTO
	result := u.DB.Debug().Where("email=?", email).Find(&user)
	if result.Error != nil {
		return userDto, result.Error
	}
	//Si le role est merchant on cherchant sur la table merchant son id
	var merchant models.Merchant
	if user.Role == "merchant" {
		results := u.DB.Where("email=?", email).Find(&merchant)
		if results.Error != nil {
			return userDto, result.Error
		}
		userDto.ID = merchant.ID
	}
	if user.Role == "client" {
		var clients models.Clients
		resul := u.DB.Where("email=?", email).Find(&clients)
		if resul.Error != nil {
			return userDto, resul.Error
		}
		fmt.Println("Role client", clients)
		userDto.ID = clients.ID
	}
	return userDto, nil
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
