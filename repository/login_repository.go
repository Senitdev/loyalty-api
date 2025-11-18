package repository

import (
	"fmt"
	"loyalty-api/controller/dto"
	"loyalty-api/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRepository interface {
	GetUserByLoginAndPassword(email string, password string) (dto.UserDTO, bool)
}
type loginRepository struct {
	DB *gorm.DB
}

// GetUserByLoginAndPassword implements LoginRepository.
func (l *loginRepository) GetUserByLoginAndPassword(email string, password string) (dto.UserDTO, bool) {
	if email == "" || password == "" {
		return dto.UserDTO{}, false
	}
	var user models.User
	// 🔹 Récupération de l’utilisateur par email
	if err := l.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Utilisateur non trouvé")
			return dto.UserDTO{}, false
		}
		fmt.Println("Erreur DB :", err)
		var userDTO dto.UserDTO
		userDTO.Email = user.Email
		userDTO.Role = user.Role
		return userDTO, false
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		// si err != nil => les mots de passe ne correspondent pas
		fmt.Println("Mot de passe invalide")
		return dto.UserDTO{}, false
	}
	fmt.Println("Authentification réussie !")
	return dto.UserDTO{
		Email: user.Email,
		Role:  user.Role,
	}, true
}

func NewLoginRepository(db *gorm.DB) LoginRepository {
	return &loginRepository{
		DB: db,
	}
}
