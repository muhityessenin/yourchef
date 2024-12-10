package routes

import (
	"github.com/gin-gonic/gin"
	"user_service/internal/api/handler"
)

func InitRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler) {
	router.POST("/register", userHandler.CreateUser)
	router.POST("/login", userHandler.Login)
	router.GET("/users/", userHandler.GetUsers)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)
	router.GET("/users/:id", userHandler.GetUserById)
	router.GET("/users/search", userHandler.SearchUser)
}
