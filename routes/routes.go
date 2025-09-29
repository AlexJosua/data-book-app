package routes

import (
	"go-books/handlers"
	"go-books/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")


	api.POST("/users/login", handlers.UserLogin)


	cat := api.Group("/categories", middleware.JWTAuthMiddleware())
	{
		cat.GET("", handlers.GetCategories)           // get all
		cat.POST("", handlers.CreateCategory)         // create
		cat.GET("/:id", handlers.GetCategoryByID)     // get by id
		cat.PUT("/:id", handlers.UpdateCategory)      // update
		cat.DELETE("/:id", handlers.DeleteCategory)   // delete
		cat.GET("/:id/books", handlers.GetBooksByCategory) // get books in category
	}


	books := api.Group("/books", middleware.JWTAuthMiddleware())
	{
		books.GET("", handlers.GetBooks)         // get all
		books.POST("", handlers.CreateBook)      // create
		books.GET("/:id", handlers.GetBookByID)  // get by id
		books.PUT("/:id", handlers.UpdateBook)   // update
		books.DELETE("/:id", handlers.DeleteBook) // delete
	}
}
