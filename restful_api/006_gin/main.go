package main

import (
	domain "github.com/JairoGLoz/senior-go-projects/senior-go-projects/restful_api/000_domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(authMiddleware())
	router.GET("/users/:id", handleGetUser)
	router.POST("/users/", handleCreateUser)

	router.Run(":8080")
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Pass the request to the next handler if the Authorization header is valid
		c.Next()
	}
}

func handleGetUser(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "Requested user with id: %s", id)
}

func handleCreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
		return
	}

	if user.Age < 18 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User must be 18 years or older"})
		return
	}

	c.String(http.StatusOK, "User created successfully!")
}
