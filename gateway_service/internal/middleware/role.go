package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Role not found"})
			return
		}

		userRole := role.(string)
		authorized := false
		for _, r := range allowedRoles {
			if r == userRole {
				authorized = true
				break
			}
		}

		if !authorized {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			return
		}

		c.Next()
	}
}
