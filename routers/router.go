package routers

import (
	"github.com/ak-yudha/crud-gin/controllers"
	"github.com/ak-yudha/crud-gin/db"
	"github.com/ak-yudha/crud-gin/repositories"
	"github.com/ak-yudha/crud-gin/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

// SetupRouter sets up the Gin router
func SetupRouter(oauthConfig *oauth2.Config) *gin.Engine {
	r := gin.Default()

	// Initialize database connection
	db, err := db.SetupDB()
	if err != nil {
		panic("Failed to connect to the database")
	}

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", userController.RegisterUser)
		v1.GET("/users", userController.GetUsers)
	}

	return r
}
