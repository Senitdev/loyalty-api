package dto

import "time"

type MerchantDto struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" gorm:"unique"`
	Email       string    `gorm:"unique" json:"email"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	SoldePoints int       `json:"soldePoints"`
	CreatedAt   time.Time `json:"created_at"`
}
