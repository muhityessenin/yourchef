package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"user_service/internal/api/handler"
	"user_service/internal/api/routes"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(userHandler *handler.UserHandler) *Server {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://your-frontend-domain.com"}, // Укажите домены фронтенда
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(MethodNotAllowedMiddleware())

	routes.InitRoutes(router.Group("/"), userHandler)

	return &Server{router}
}

func (s *Server) Run(infoLog *log.Logger, errorLog *log.Logger) {
	infoLog.Printf("starting server on: 8001")
	err := s.engine.Run(":8001")
	errorLog.Fatal(err)
}

func MethodNotAllowedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedMethods := map[string]bool{
			"GET":    true,
			"POST":   true,
			"PUT":    true,
			"DELETE": true,
		}
		if !allowedMethods[c.Request.Method] {
			c.JSON(http.StatusMethodNotAllowed, gin.H{
				"error": "Method Not Allowed",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
