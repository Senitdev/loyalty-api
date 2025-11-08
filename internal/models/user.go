package models

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"` // client ou merchant
	CreatedAt time.Time `json:"created_at"`
	Ref       string
}
