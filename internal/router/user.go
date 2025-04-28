package router

import (
	"base/internal/router/middleware"

	"github.com/gin-gonic/gin"
)

func (handlerFuncs HandlerFuncs) createUserGroup(g *gin.RouterGroup,
	mid middleware.Middleware, groupName string) {
	userGroup := g.Group(groupName).Use(mid.MidBasicType(groupName))
	userGroup.POST("", handlerFuncs.UserCreateHandler)
	userGroup.GET("", handlerFuncs.UserListHandler)
	userGroup.GET("/:id", handlerFuncs.UserGetHandler)
	userGroup.DELETE("/:id", handlerFuncs.UserDeleteHandler)

	g.Group(groupName).Use(mid.MidBasicType("")).PUT("", handlerFuncs.UserUpdateHandler)
	g.Group(groupName).Use(mid.MidBasicType("")).PUT("/reset-pass", handlerFuncs.UserResetPassHandler)
}
