package router

import (
	"base/internal/router/middleware"

	"github.com/gin-gonic/gin"
)

func (handlerFuncs HandlerFuncs) createUserGroup(g *gin.RouterGroup,
	mid middleware.Middleware, groupName string) {
	authGroup := g.Group(groupName).Use(mid.MidBasicType(""))
	authGroup.POST("/create", handlerFuncs.LoginHandler)
	authGroup.POST("/reset-pass", handlerFuncs.LoginHandler)
	authGroup.GET("/profile", handlerFuncs.LoginHandler)
}
