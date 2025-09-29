package models

import "time"

type Category struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"size:255;not null" json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `gorm:"size:255" json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `gorm:"size:255" json:"modified_by"`
	Books      []Book    `gorm:"foreignKey:CategoryID" json:"books"`
}
