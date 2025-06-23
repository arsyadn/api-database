package routes

import (
	"auth-api/controllers"
	"auth-api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	protected := router.Group("/")
	nonProcted := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", userController.GetUser)
	}
	nonProcted.GET("/users/:id", userController.GetUserByIdController)
}