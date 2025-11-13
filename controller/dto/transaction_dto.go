package dto

import "time"

type TransactionDTO struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	Type       string    `json:"type"` // "earn" ou "spend"
	Points     int       `json:"points"`
	CreatedAt  time.Time `json:"created_at"`
	MerchantId int       `json:"merchant_id"`
	ClientId   int       `json:"client_id"`
	Merchant   string    `json:"merchant"`
}
