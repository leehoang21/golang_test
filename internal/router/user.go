package router

import (
	"base/internal/router/middleware"

	"github.com/gin-gonic/gin"
)

func (handlerFuncs HandlerFuncs) createUserGroup(g *gin.RouterGroup,
	mid middleware.Middleware, groupName string) {
	userGroup := g.Group(groupName) //.Use(mid.MidBasicType(""))
	userGroup.POST("", handlerFuncs.UserCreateHandler)
	userGroup.PUT("/:id", handlerFuncs.UserUpdateHandler)
	userGroup.PUT("/:id/reset-pass", handlerFuncs.UserResetPassHandler)
	userGroup.GET("", handlerFuncs.UserListHandler)
	userGroup.GET("/:id", handlerFuncs.UserGetHandler)
	userGroup.DELETE("/:id", handlerFuncs.UserDeleteHandler)
}
