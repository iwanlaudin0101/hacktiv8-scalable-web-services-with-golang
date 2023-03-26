package routes

import (
	"rest-api-postgres/controllers"

	"github.com/gin-gonic/gin"
)

func StartServerApi() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.AllBook)
	router.POST("/books", controllers.CreatedBook)
	router.GET("/books/:Id", controllers.DetailBook)
	router.PUT("/books/:Id", controllers.UpdatedBook)
	router.DELETE("/books/:Id", controllers.DeletedBook)

	return router
}
