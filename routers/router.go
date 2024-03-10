package routers

import (
	"github.com/ak-yudha/crud-gin/controllers"
	"github.com/ak-yudha/crud-gin/db"
	"github.com/ak-yudha/crud-gin/repositories"
	"github.com/ak-yudha/crud-gin/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// SetupRouter sets up the Gin router
func SetupRouter(oauthConfig *oauth2.Config) *gin.Engine {
	r := gin.Default()

	// Initialize OAuth2 configuration
	oauthConfig = &oauth2.Config{
		ClientID:     "717668172473-903cb57ludtt1eqrs53ehrfsgt8a45df.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/auth/callback",
		Scopes:       []string{"user", "task"},
		Endpoint:     google.Endpoint,
	}

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
		v1.GET("/users/:id", userController.GetUser)
		v1.PUT("/users/:id", userController.UpdateUser)
		v1.DELETE("/users/:id", userController.DeleteUser)
	}

	r.GET("/auth/callback", func(c *gin.Context) {
		// Handle the OAuth2 callback
		code := c.Query("code")
		token, err := oauthConfig.Exchange(oauth2.NoContext, code)
		if err != nil {
			c.String(500, "Failed to exchange token: %s", err.Error())
			return
		}

		c.JSON(200, gin.H{
			"access_token":  token.AccessToken,
			"token_type":    token.TokenType,
			"expiry":        token.Expiry,
			"refresh_token": token.RefreshToken,
		})
	})

	return r
}
