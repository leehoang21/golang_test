package handler

import (
	"base/internal/handler/filter"
	"base/internal/service/token"
	"base/internal/service/user"
	"base/internal/utils/web"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService  user.Service
	tokenService token.Service
	*web.JsonRender
	contextWith web.ContextWith
}

func NewUserHandler(
	userService user.Service,
	tokenService token.Service,
	contextWith web.ContextWith,
) UserHandler {
	return UserHandler{
		userService:  userService,
		tokenService: tokenService,
		contextWith:  contextWith,
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
	var input user.UserInput
	web.AssertNil(ctx.BindJSON(&input))
	var us, err = u.userService.Create(ctx.Request.Context(), input)
	web.AssertNil(err)
	u.SendData(ctx, us)
}

// UpdateHandler
// @Tags         User
// @Summary      Create user
// @Description  Create user
// @ID           user-update
// @Accept       json
// @Produce      json
// @Param        x-user-id  header    string       false  "user id"
// @Param        data       body      models.User  true   "user data"
// @Success      200        {object}  models.User
// @Router       /api/v1/users [put]
func (u UserHandler) UpdateHandler(ctx *gin.Context) {
	var input user.UserInput
	web.AssertNil(ctx.BindJSON(&input))
	userID, e := u.contextWith.GetUserID(ctx)
	web.AssertNil(e)
	var us, err = u.userService.Update(ctx.Request.Context(), userID, input)
	web.AssertNil(err)
	u.SendData(ctx, us)
}

// DeleteHandler
// @Tags         User
// @Summary      delete user
// @Description  delete user
// @ID           user-delete
// @Accept       json
// @Produce      json
// @Param        id  		in:query    string       false  "user id"
// @Success      200        {object}  null
// @Router       /api/v1/users/:id [delete]
func (u UserHandler) DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := u.contextWith.GetUserID(ctx)
	if id == userID || err != nil {
		web.AssertNil(web.Forbidden("access denied"))
	}
	err = u.userService.Delete(ctx.Request.Context(), id)
	web.AssertNil(err)
	u.SendData(ctx, nil)
}

// GetListHandler
// @Tags         User
// @Summary      get list user
// @Description  get list user
// @ID           user-get-list
// @Accept       json
// @Produce      json
// @Param        x-user-id  header    string       false  "user id"
// @Param        data       body      models.User  true   "user data"
// @Success      200        {object}  []models.User
// @Router       /api/v1/users [get]
func (u UserHandler) GetListHandler(ctx *gin.Context) {
	var f = filter.NewUserListPrams()
	web.AssertNil(ctx.BindQuery(f))
	var us, total, err = u.userService.List(ctx.Request.Context(), f)
	web.AssertNil(err)
	u.SendData(ctx, responseList{
		Data:  us,
		Total: total,
	})
}

// GetHandler
// @Tags         User
// @Summary      get  user
// @Description  get  user
// @ID           user-get
// @Accept       json
// @Produce      json
// @Param        x-user-id  header    string       false  "user get"
// @Param        id       in:query    string  true   "user id"
// @Success      200        {object}  models.User
// @Router       /api/v1/users [get]
func (u UserHandler) GetHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var us, err = u.userService.GetByID(ctx.Request.Context(), id)
	web.AssertNil(err)
	u.SendData(ctx, us)
}

// ResetPassHandler
// @Tags         User
// @Summary      Reset pass user
// @Description  Reset pass user
// @ID           user-reset-pass
// @Accept       json
// @Produce      json
// @Param        x-user-id  header    string       false  "user id"
// @Param        id  	    in:query  string       false  "user id"
// @Param        data       body      user.UserResetPassInput  true   "user data"
// @Success      200        {object}  models.User
// @Router       /api/v1/users/reset-pass [put]
func (u UserHandler) ResetPassHandler(ctx *gin.Context) {
	var f user.UserResetPassInput
	web.AssertNil(ctx.BindJSON(&f))
	userID, e := u.contextWith.GetUserID(ctx)
	web.AssertNil(e)
	var us, err = u.userService.ResetPass(ctx.Request.Context(), userID, f)
	web.AssertNil(err)
	err = u.tokenService.RevokeAllByUserID(ctx, userID)
	web.AssertNil(err)
	u.SendData(ctx, us)
}
