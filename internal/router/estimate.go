package router

import (
	"base/internal/router/middleware"

	"github.com/gin-gonic/gin"
)

func (handlerFuncs HandlerFuncs) createEstimateGroup(g *gin.RouterGroup,
	mid middleware.Middleware, groupName string) {
	EstimateGroup := g.Group(groupName).Use(mid.MidBasicType(groupName))
	EstimateGroup.POST("", handlerFuncs.EstimateCreateHandler)
	EstimateGroup.PUT("/:id", handlerFuncs.EstimateUpdateHandler)
	EstimateGroup.GET("/:id", handlerFuncs.EstimateGetHandler)
	EstimateGroup.DELETE("/:id", handlerFuncs.EstimateDeleteHandler)
}
