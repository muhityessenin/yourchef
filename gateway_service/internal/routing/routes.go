package routing

import (
	"fmt"
	"gateway_service/internal/config"
	"gateway_service/internal/middleware"
	"gateway_service/internal/proxy"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Пример публичного маршрута - не требует авторизации
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Публичные маршруты для регистрации и логина
	// Эти маршруты проксируем в user_service без проверки токена
	r.POST("/register", func(c *gin.Context) {
		// Обновляем путь, если необходимо. Предположим, в user_service маршрут /register без /users
		// Если в user_service маршрут "/register", то можно не менять путь.
		c.Request.URL.Path = "/register"
		fmt.Println(cfg.UserServiceURL)
		proxy.ProxyRequest(c, cfg.UserServiceURL)
	})

	r.POST("/login", func(c *gin.Context) {
		c.Request.URL.Path = "/login"
		fmt.Println(cfg.UserServiceURL)
		proxy.ProxyRequest(c, cfg.UserServiceURL)
	})
	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthMiddleware("ieqriebqFNEIvbv9ewvnon3u543v34248jnveibviewpvb"))
	fmt.Println(cfg.JWTSecret)
	{
		authGroup.Any("recipe/*path", func(c *gin.Context) {
			path := c.Param("path")
			c.Request.URL.Path = path
			fmt.Println(cfg.RecipeServiceURL)
			proxy.ProxyRequest(c, cfg.RecipeServiceURL)
		})

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
