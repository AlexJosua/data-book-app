package models

import "time"

type Category struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	Name       string    `gorm:"size:255;not null"`
	CreatedAt  time.Time
	CreatedBy  string    `gorm:"size:255"`
	ModifiedAt time.Time
	ModifiedBy string    `gorm:"size:255"`
	Books      []Book    `gorm:"foreignKey:CategoryID"`
}
