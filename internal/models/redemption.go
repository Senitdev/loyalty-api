package models

import "time"

type Redemption struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	UserID     int       `json:"user_id"`
	RewardID   int       `json:"reward_id"`
	RedeemedAt time.Time `json:"redeemed_at"`
	User       User      `json:"user" gorm:"foreignKey:UserID"`
	Reward     Reward    `json:"reward" gorm:"foreignKey:RewardID"`
}
