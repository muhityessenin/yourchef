package routes

import (
	"github.com/gin-gonic/gin"
	"user_service/internal/api/handler"
	"user_service/internal/domain/user"
)

func InitRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler) {
	router.POST("/register", userHandler.CreateUser)
	router.POST("/login", userHandler.Login)
	authorized := router.Group("/")
	authorized.Use(user.AuthMiddleware())
	{

		// authorized.GET("/profile", UserProfile)
		premium := authorized.Group("/premium")
		premium.Use(user.RoleMiddleware("premium", "admin"))
		{
			// premium.GET("/premium-content", PremiumContent)
		}

		admin := authorized.Group("/admin")
		admin.Use(user.RoleMiddleware("admin"))
		{
			admin.GET("/users/", userHandler.GetUsers)
			admin.PUT("/users/:id", userHandler.UpdateUser)
			admin.DELETE("/users/:id", userHandler.DeleteUser)
			admin.GET("/users/:id", userHandler.GetUserById)
			admin.GET("/users/search", userHandler.SearchUser)
		}
	}

	guest := router.Group("/")
	guest.Use(user.GuestMiddleware())
	{
		// guest.GET("/public-content", PublicContent)
	}
}
