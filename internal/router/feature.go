package router

import (
	"base/internal/router/middleware"

	"github.com/gin-gonic/gin"
)

func (handlerFuncs HandlerFuncs) createFeatureGroup(g *gin.RouterGroup,
	mid middleware.Middleware, groupName string) {
	groupRoleGroup := g.Group(groupName).Use(mid.MidBasicType(groupName))
	groupRoleGroup.POST("", handlerFuncs.FeatureCreateHandler)
	groupRoleGroup.PUT("/:id", handlerFuncs.FeatureUpdateHandler)
	groupRoleGroup.GET("", handlerFuncs.FeatureListHandler)
	groupRoleGroup.GET("/:id", handlerFuncs.FeatureGetHandler)
	groupRoleGroup.DELETE("/:id", handlerFuncs.FeatureDeleteHandler)
}
