package routes

import "github.com/gin-gonic/gin"

func StartServerWeb() *gin.Engine {
	router := gin.Default()

	// router.GET("/books", controllers.Index)
	// router.POST("/books", controllers.Create)
	// router.GET("/books/:Id", controllers.Detail)
	// router.PUT("/books/:Id", controllers.Updated)
	// router.DELETE("/books/:Id", controllers.Deleted)

	return router
}
