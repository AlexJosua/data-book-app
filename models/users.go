package models

import "time"

type User struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username   string    `gorm:"size:255;unique;not null" json:"username"`
	Password   string    `gorm:"size:255;not null" json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `gorm:"size:255" json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `gorm:"size:255" json:"modified_by"`
}
