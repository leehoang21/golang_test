package router

import (
	//"base/internal/handler"
	//"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HandlerFuncs struct {
	UserListHandler gin.HandlerFunc
}

// Create handlers
func Create(handlerFuncs HandlerFuncs) *gin.Engine {
	root := gin.Default()
	root.Use(cors.Default())

	api := root.Group("/api")

	userGroup := api.Group("/users")
	userGroup.GET("/list", handlerFuncs.UserListHandler)
	return root
}
