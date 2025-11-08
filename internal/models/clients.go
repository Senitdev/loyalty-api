package models

import (
	"time"
)

type Clients struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Email       string    `gorm:"unique" json:"email"`
	Birthday    string    `json:"birthday"`
	Mobile      string    `json:"mobile" gorm:"unique"`
	SoldePoints int       `json:"soldepoints"`
	CreatedAt   time.Time `json:"created_at"`
	Password    string
	UserRef     string
}
