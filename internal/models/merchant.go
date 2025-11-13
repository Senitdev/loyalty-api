package models

import (
	"time"
)

type Merchant struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" gorm:"unique"`
	Email     string    `gorm:"unique" json:"email"`
	LogoURL   string    `json:"logo_url"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Password  string    `json:"-"`
	UserRef   string
	// Relations
	Rewards []Reward `json:"rewards,omitempty"`
}
