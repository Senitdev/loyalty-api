package models

import "time"

type Reward struct {
	ID             int       `gorm:"primaryKey" json:"id"`
	MerchantID     int       `json:"merchant_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	PointsRequired int       `json:"points_required"`
	CreatedAt      time.Time `json:"created_at"`
	Merchant       Merchant  `json:"merchant" gorm:"foreignKey:MerchantID"`
	IsActive       bool      `json:"isactive"`
}
