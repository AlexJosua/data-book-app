package models

import "time"

type Book struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"size:255;not null"`
	Description string    `gorm:"type:text"`
	ImageURL    string    `gorm:"size:255"`
	ReleaseYear int
	Price       int
	TotalPage   int
	Thickness   string    `gorm:"size:50"`
	CategoryID  int
	Category    Category  `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time
	CreatedBy   string    `gorm:"size:255"`
	ModifiedAt  time.Time
	ModifiedBy  string    `gorm:"size:255"`
}
