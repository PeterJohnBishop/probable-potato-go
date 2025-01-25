package server

import (
	"probable-potato-go/main.go/routes"

	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	r := gin.Default()

	// Register item routes
	itemRoutes := r.Group("/items")
	userRoutes := r.Group("/users")
	routes.RegisterUserRoutes(userRoutes)
	routes.RegisterItemRoutes(itemRoutes)

	return r
}
