package models

import "time"

type User struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	Username   string    `gorm:"size:255;unique;not null"`
	Password   string    `gorm:"size:255;not null"`
	CreatedAt  time.Time
	CreatedBy  string    `gorm:"size:255"`
	ModifiedAt time.Time
	ModifiedBy string    `gorm:"size:255"`
}
