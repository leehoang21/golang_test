package router

import (
	"base/internal/router/middleware"

	"github.com/gin-gonic/gin"
)

func (handlerFuncs HandlerFuncs) createAuthGroup(g *gin.RouterGroup,
	mid middleware.Middleware, groupName string) {
	authGroup := g.Group(groupName)
	authGroup.POST("/login", handlerFuncs.LoginHandler)
	authGroup.Use(mid.MidBasicType("")).POST("/logout", handlerFuncs.LogoutHandler)
}
