package routing

import (
	"gateway_service/internal/config"
	"gateway_service/internal/middleware"
	"gateway_service/internal/proxy"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Пример публичного маршрута - не требует авторизации
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Авторизованные маршруты
	// Проверяем токен
	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		// Допустим, /recipes доступны для всех авторизованных
		authGroup.Any("/recipes/*path", func(c *gin.Context) {
			proxy.ProxyRequest(c, cfg.RecipeServiceURL)
		})

		// /users доступны только админам
		adminGroup := authGroup.Group("/users")
		adminGroup.Use(middleware.RoleMiddleware("admin"))
		{
			adminGroup.Any("/*path", func(c *gin.Context) {
				proxy.ProxyRequest(c, cfg.UserServiceURL)
			})
		}
	}

	return r
}
