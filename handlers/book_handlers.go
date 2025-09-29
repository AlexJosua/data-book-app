package handlers

import (
	"database/sql"
	"net/http"

	"go-books/config"
	"go-books/models"

	"github.com/gin-gonic/gin"
)

// ========================
// Get All Books
// ========================
func GetBooks(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id
		FROM books`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(
			&b.ID, &b.Title, &b.Description, &b.ImageURL,
			&b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, b)
	}
	c.JSON(http.StatusOK, books)
}

// ========================
// Create Book
// ========================
func CreateBook(c *gin.Context) {
	var req models.Book
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validasi release year
	if req.ReleaseYear < 1980 || req.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release_year must be between 1980 and 2024"})
		return
	}

	// Konversi thickness
	if req.TotalPage > 100 {
		req.Thickness = "tebal"
	} else {
		req.Thickness = "tipis"
	}

	_, err := config.DB.Exec(`
		INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,NOW(),'system')`,
		req.Title, req.Description, req.ImageURL, req.ReleaseYear,
		req.Price, req.TotalPage, req.Thickness, req.CategoryID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created"})
}

// ========================
// Get Book By ID
// ========================
func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	var b models.Book
	err := config.DB.QueryRow(`
		SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id
		FROM books WHERE id=$1`, id).Scan(
		&b.ID, &b.Title, &b.Description, &b.ImageURL,
		&b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, b)
}

// ========================
// Delete Book
// ========================
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	result, err := config.DB.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

// ========================
// Update Book
// ========================
func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var req models.Book
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validasi release year
	if req.ReleaseYear < 1980 || req.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release_year must be between 1980 and 2024"})
		return
	}

	// Konversi thickness
	if req.TotalPage > 100 {
		req.Thickness = "tebal"
	} else {
		req.Thickness = "tipis"
	}

	result, err := config.DB.Exec(`
		UPDATE books SET title=$1, description=$2, image_url=$3, release_year=$4,
		price=$5, total_page=$6, thickness=$7, category_id=$8, modified_at=NOW(), modified_by='system'
		WHERE id=$9`,
		req.Title, req.Description, req.ImageURL, req.ReleaseYear,
		req.Price, req.TotalPage, req.Thickness, req.CategoryID, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated"})
}
