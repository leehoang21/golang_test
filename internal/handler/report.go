package handler

import (
	"base/internal/notification"
	"base/internal/service/report"
	"base/internal/utils/web"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	reportService report.Service
	sender        *notification.Sender
	*web.JsonRender
	contextWith web.ContextWith
}

func NewReportHandler(
	reportService report.Service,
	contextWith web.ContextWith,
	sender *notification.Sender,
) ReportHandler {
	return ReportHandler{
		reportService: reportService,
		contextWith:   contextWith,
		sender:        sender,
	}
}

// SendTelegramReportHandler
// @Tags         SendTelegramReport
// @Summary      SendTelegramReport
// @Description  SendTelegramReport
// @ID           Telegram-Report-send
// @Accept       json
// @Produce      json
// @Router       /api/v1/send-telegram-report [get]
func (u ReportHandler) SendTelegramReportHandler(ctx *gin.Context) {
	var mess, err = u.reportService.GenerateMorningReport(ctx.Request.Context())
	web.AssertNil(err)
	err = u.sender.SendTelegramMessage(mess)
	web.AssertNil(err)
	u.SendData(ctx, nil)
}
