package router

import (
	"base/internal/router/middleware"

	"github.com/gin-gonic/gin"
)

func (handlerFuncs HandlerFuncs) createGroupRoleGroup(g *gin.RouterGroup,
	mid middleware.Middleware, groupName string) {
	groupRoleGroup := g.Group(groupName).Use(mid.MidBasicType(groupName))
	groupRoleGroup.POST("", handlerFuncs.GroupRoleCreateHandler)
	groupRoleGroup.PUT("/:id", handlerFuncs.GroupRoleUpdateHandler)
	groupRoleGroup.GET("", handlerFuncs.GroupRoleListHandler)
	groupRoleGroup.GET("/:id", handlerFuncs.GroupRoleGetHandler)
	groupRoleGroup.DELETE("/:id", handlerFuncs.GroupRoleDeleteHandler)
	groupRoleGroup.POST("/members", handlerFuncs.GroupRoleAddMembersHandler)
	groupRoleGroup.PUT("/members/:id", handlerFuncs.GroupRoleUpdateMembersHandler)
	groupRoleGroup.DELETE("/members/:id", handlerFuncs.GroupRoleDeleteMembersHandler)
}
