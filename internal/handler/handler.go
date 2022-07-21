package handler

import (
	"base/internal/handler/filter"
	"base/internal/service/user"
	"base/internal/utils/web"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService user.UserService
	web.JsonRender
}

func NewUserHandler(
	userService user.UserService,
) UserHandler {
	return UserHandler{
		UserService: userService,
	}
}

func (u UserHandler) ListHandler(ctx *gin.Context) {
	var f = filter.NewUserListPrams()
	web.WebArsernil(ctx.BindQuery(&f))
	var us, err = u.UserService.List(ctx.Request.Context(), f)
	web.WebArsernil(err)
	u.SendData(ctx, us)
}
