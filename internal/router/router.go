package router

import (
	"eclectric/internal/router/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

// Create handlers
func (handlerFuncs HandlerFuncs) Create(mid middleware.Middleware) *gin.Engine {
	root := gin.Default()
	root.Use(cors.Default())
	root.Use(mid.Recovery())
	api := root.Group("/api/v1")
	api.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	handlerFuncs.createAuthGroup(api, mid, "/auth")
	handlerFuncs.createUserGroup(api, mid, "/users")
	return root
}
