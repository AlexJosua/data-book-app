package handlers

import (
	"database/sql"
	"net/http"

	"go-books/config"

	"github.com/gin-gonic/gin"
)

// ========================
// Get All Categories
// ========================
func GetCategories(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name FROM category")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var categories []gin.H
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		categories = append(categories, gin.H{"id": id, "name": name})
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// ========================
// Create Category
// ========================
func CreateCategory(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	_, err := config.DB.Exec(`
		INSERT INTO category (name, created_at, created_by)
		VALUES ($1, NOW(), 'system')
	`, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created"})
}

// ========================
// Get Category By ID
// ========================
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")

	var name string
	err := config.DB.QueryRow("SELECT name FROM category WHERE id=$1", id).Scan(&name)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "name": name})
}

// ========================
// Update Category
// ========================
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	result, err := config.DB.Exec(`
		UPDATE category
		SET name=$1, modified_at=NOW(), modified_by='system'
		WHERE id=$2
	`, req.Name, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated"})
}

// ========================
// Delete Category
// ========================
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	result, err := config.DB.Exec("DELETE FROM category WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}

// ========================
// Get Books by Category
// ========================
func GetBooksByCategory(c *gin.Context) {
	id := c.Param("id")

	rows, err := config.DB.Query(`
		SELECT id, title, description, price, release_year
		FROM books
		WHERE category_id=$1
	`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []gin.H
	for rows.Next() {
		var bookID, price, releaseYear int
		var title, description string
		if err := rows.Scan(&bookID, &title, &description, &price, &releaseYear); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, gin.H{
			"id":           bookID,
			"title":        title,
			"description":  description,
			"price":        price,
			"release_year": releaseYear,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}
