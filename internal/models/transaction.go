package models

import "time"

// Transaction represents a points transaction (earning or spending points) on a loyalty card.
type Transaction struct {
	ID            int         `gorm:"primaryKey" json:"id"`
	LoyaltyCardID int         `json:"loyaltycard_id"`
	Type          string      `json:"type"` // "earn" ou "spend"
	Points        int         `json:"points"`
	Description   string      `json:"description"`
	CreatedAt     time.Time   `json:"created_at"`
	MerchantId    int         `json:"merchant_id"`
	LoyaltyCard   LoyaltyCard `json:"loyalty_card" gorm:"foreignKey:LoyaltyCardID"`
}
