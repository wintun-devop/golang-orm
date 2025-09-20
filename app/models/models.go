package models

import (
	"time"
)

type User struct {
	ID        string    `gorm:"type:text;primaryKey"`
	Username  string    `gorm:"type:varchar(255);not null;unique"`
	Email     string    `gorm:"type:varchar(255);not null;unique"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Role      string    `gorm:"type:varchar(50);default:user"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
