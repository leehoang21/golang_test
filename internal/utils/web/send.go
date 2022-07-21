package web

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type JsonRender struct {
}

func (r *JsonRender) SendData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data":   data,
		"status": "success",
		"code":   200,
	})
}

func (r *JsonRender) SendString(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func (r *JsonRender) SendDataNotFound(ctx *gin.Context, data interface{}, isNotFound bool) {
	var status = "success"
	if isNotFound {
		status = "error"
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data":   data,
		"status": status,
		"code":   200,
	})
}

func (r *JsonRender) SendError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": err.Error(),
		"status":  "error",
		"code":    200,
	})
}

func (r *JsonRender) Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"data":   nil,
		"status": "success",
		"code":   200,
	})
}

func Recover() {
	if r := recover(); r != nil {
		glog.Error(r, string(debug.Stack()))
	}
}
