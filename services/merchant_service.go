package service

import (
	"loyalty-api/internal/models"
	"loyalty-api/repository"
	"time"
)

type MerchantService interface {
	SaveMerchant(merchant models.Merchant) models.Merchant
	FindAllMerchant() []models.Merchant
	DeleteMerchantById(id int) error
	FindMerchantById(id int) (models.Merchant, error)
	FindMerchantByEmail(email string) (models.Merchant, error)
	UpdateMerchant(id int, merchant models.Merchant) (models.Merchant, error)
}
type merchantService struct {
	service repository.MerchantRepository
}

// DeleteMerchantById implements MerchantService.
func (m *merchantService) DeleteMerchantById(id int) error {
	if err := m.service.DeleteMerchantById(id); err != nil {
		return err
	}
	return nil
}

// FindAllMerchant implements MerchantService.
func (m *merchantService) FindAllMerchant() []models.Merchant {
	return m.service.FindAll()
}

// FindMerchantByEmail implements MerchantService.
func (m *merchantService) FindMerchantByEmail(email string) (models.Merchant, error) {
	return m.service.FindMerchantByEmail(email)
}

// FindMerchantById implements MerchantService.
func (m *merchantService) FindMerchantById(id int) (models.Merchant, error) {
	return m.service.FindMerchantById(id)
}

// SaveMerchant implements MerchantService.
func (m *merchantService) SaveMerchant(merchant models.Merchant) models.Merchant {
	merchant.CreatedAt = time.Now()
	merchant.UpdatedAt = time.Now()
	return m.service.Save(merchant)
}

// UpdateMerchant implements MerchantService.
func (m *merchantService) UpdateMerchant(id int, merchant models.Merchant) (models.Merchant, error) {
	panic("unimplemented")
}

// Constructeur
func NewMerchantService(repo repository.MerchantRepository) MerchantService {
	return &merchantService{
		service: repo,
	}
}
