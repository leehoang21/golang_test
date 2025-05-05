package router

import (
	"base/internal/router/middleware"

	"github.com/gin-gonic/gin"
)

func (handlerFuncs HandlerFuncs) createTaskGroup(g *gin.RouterGroup,
	mid middleware.Middleware, groupName string) {
	TaskGroup := g.Group(groupName).Use(mid.MidBasicType(groupName))
	TaskGroup.POST("", handlerFuncs.TaskCreateHandler)
	TaskGroup.PUT("/:id", handlerFuncs.TaskUpdateHandler)
	TaskGroup.GET("", handlerFuncs.TaskListHandler)
	TaskGroup.GET("/:id", handlerFuncs.TaskGetHandler)
	TaskGroup.DELETE("/:id", handlerFuncs.TaskDeleteHandler)
}
