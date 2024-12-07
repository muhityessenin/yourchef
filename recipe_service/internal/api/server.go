package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"recipe_service/internal/api/handler"
	"recipe_service/internal/api/routes"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(recipeHandler *handler.RecipeHandler) *Server {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(MethodNotAllowedMiddleware())

	routes.InitRoutes(router.Group("/recipe"), recipeHandler)

	return &Server{router}
}

func (s *Server) Run(infoLog *log.Logger, errorLog *log.Logger) {
	infoLog.Printf("starting server on: 8000")
	err := s.engine.Run(":8000")
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
