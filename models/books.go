package models

import "time"

type Book struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	ImageURL    string    `gorm:"size:255" json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	Price       int       `json:"price"`
	TotalPage   int       `json:"total_page"`
	Thickness   string    `gorm:"size:50" json:"thickness"`
	CategoryID  int       `json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   string    `gorm:"size:255" json:"created_by"`
	ModifiedAt  time.Time `json:"modified_at"`
	ModifiedBy  string    `gorm:"size:255" json:"modified_by"`
}
