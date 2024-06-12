package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"order-service/src/controllers/user"
)

func InitializeUserRoutes(r *gin.Engine, db *gorm.DB) {
	userController := user.NewUserController(db)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/order", userController.GetAllUsers)
		// Add more user-related routes here
	}
}
