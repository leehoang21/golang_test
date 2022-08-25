package handler

import (
	"eclectric/internal/models"
	"eclectric/internal/service/token"
	"eclectric/internal/service/user"
	"eclectric/internal/utils/web"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService  user.Service
	TokenService token.Service
	*web.JsonRender
	contextWith web.ContextWith
}

func NewUserHandler(
	userService user.Service,
	contextWith web.ContextWith,
) UserHandler {
	return UserHandler{
		UserService: userService,
		contextWith: contextWith,
	}
}

// CreateHandler
// @Tags         User
// @Summary      Create user
// @Description  Create user
// @ID           user-create
// @Accept       json
// @Produce      json
// @Param        x-user-id  header    string       false  "user id"
// @Param        data       body      models.User  true   "user data"
// @Success      200        {object}  models.User
// @Router       /api/v1/users [post]
func (u UserHandler) CreateHandler(ctx *gin.Context) {
	var f models.User
	web.AssertNil(ctx.BindJSON(&f))
	var us, err = u.UserService.Create(ctx.Request.Context(), &f)
	web.AssertNil(err)
	u.SendData(ctx, us)
}

// UpdateHandler
// @Tags         User
// @Summary      Create user
// @Description  Create user
// @ID           user-create
// @Accept       json
// @Produce      json
// @Param        x-user-id  header    string       false  "user id"
// @Param        data       body      models.User  true   "user data"
// @Success      200        {object}  models.User
// @Router       /api/v1/users [put]
func (u UserHandler) UpdateHandler(ctx *gin.Context) {
	var f models.User
	web.AssertNil(ctx.BindJSON(&f))
	uID := ctx.GetString("x-user-id")
	var us, err = u.UserService.Update(ctx.Request.Context(), uID, &f)
	web.AssertNil(err)
	u.SendData(ctx, us)
}
