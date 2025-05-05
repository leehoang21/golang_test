package router

import (
	"github.com/gin-gonic/gin"
)

func (handlerFuncs HandlerFuncs) createReportGroup(g *gin.RouterGroup,
	groupName string) {
	ReportGroup := g.Group(groupName)
	ReportGroup.GET("/send-telegram-report", handlerFuncs.SendTelegramReportHandler)
}
