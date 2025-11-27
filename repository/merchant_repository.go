package repository

import (
	"errors"
	"loyalty-api/internal/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type MerchantRepository interface {
	Save(merchant models.Merchant) error
	FindAll() []models.Merchant
	DeleteMerchantById(id int) error
	FindMerchantById(id int) (models.Merchant, error)
	FindMerchantByEmail(email string) (models.Merchant, error)
	UpdateMerchant(id int, merchant models.Merchant) (models.Merchant, error)
}
type merchantRepository struct {
	DB *gorm.DB
}

// DeleteMerchantById implements MerchantRepository.
func (m *merchantRepository) DeleteMerchantById(id int) error {
	if result := m.DB.Delete(&models.Merchant{}, id).Error; result != nil {
		return result
	}
	return nil

}

// FindAll implements MerchantRepository.
func (m *merchantRepository) FindAll() []models.Merchant {
	var merchant []models.Merchant
	result := m.DB.Find(&models.Merchant{})
	if result != nil {
		return merchant
	}
	return merchant
}

// FindMerchantByEmail implements MerchantRepository.
func (m *merchantRepository) FindMerchantByEmail(email string) (models.Merchant, error) {
	var merchant models.Merchant
	if result := m.DB.Where("email", email).Find(&merchant).Error; result != nil {
		return merchant, result
	}
	return merchant, nil

}

// FindMerchantById implements MerchantRepository.
func (m *merchantRepository) FindMerchantById(id int) (models.Merchant, error) {
	var merchant models.Merchant
	if result := m.DB.Where("id", id).Find(&merchant).Error; result != nil {
		return merchant, result
	}
	return merchant, nil
}

// Save implements MerchantRepository.
func (m *merchantRepository) Save(merchant models.Merchant) error {
	id := uuid.New()
	// Génère le hash du mot de passe avant de sauvegarder
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(merchant.Password), bcrypt.DefaultCost)
	if err != nil {
		// tu peux gérer l’erreur comme tu veux, par exemple log.Fatal ou retourner une erreur
		panic("Erreur lors du hachage du mot de passe : " + err.Error())
	}
	merchant.UserRef = id.String()
	var user models.User
	user.CreatedAt = time.Now()
	user.Role = "merchant"
	user.Email = merchant.Email
	user.Name = merchant.Name
	user.Ref = id.String()
	user.Password = string(hashedPassword)
	if err := m.DB.Save(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("erreur de doublons")

		}
	}
	result := m.DB.Save(&merchant)
	if result != nil {
		return result.Error
	}
	//On renseigne
	return nil
}

// UpdateMerchant implements MerchantRepository.
func (m *merchantRepository) UpdateMerchant(id int, merchant models.Merchant) (models.Merchant, error) {
	panic("unimplemented")
}

// contructeur
func NewMerchantRepository(repo *gorm.DB) MerchantRepository {
	return &merchantRepository{
		DB: repo,
	}
}
