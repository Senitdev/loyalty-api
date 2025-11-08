package models

import "time"

type LoyaltyCard struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	ClientsID   int       `json:"clients_id"`
	MerchantID  int       `json:"merchant_id"`
	Points      int       `json:"points"`
	LastUpdated time.Time `json:"last_updated"`
	Clients     Clients   `json:"clients" gorm:"foreignKey:ClientsID"`
	Merchant    Merchant  `json:"merchant" gorm:"foreignKey:MerchantID"`
}
