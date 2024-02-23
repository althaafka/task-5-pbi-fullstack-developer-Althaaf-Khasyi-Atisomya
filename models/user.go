package models

import (
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey; autoIncrement" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Photos    []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"photos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
