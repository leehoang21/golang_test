package handler

import (
	"base/internal/models"
	"base/internal/service/auth"
	"base/internal/service/token"
	"base/internal/service/user"
	"base/internal/utils/web"

	"github.com/gin-gonic/gin"
)

// LoginHandler
// @Tags         User
// @Summary      Login user
// @Description  Login user
// @ID           auth-login
// @Accept       json
// @Produce      json
// @Param        x-user-id  header    string       false  "user id"
// @Param        data       body      user.LoginInput  true   "user data"
// @Success      200        {object}  models.User
// @Router       /api/v1/auth/login [post]
func (u UserHandler) LoginHandler(ctx *gin.Context) {
	var f = user.LoginInput{}
	web.AssertNil(ctx.BindJSON(&f))
	var us, err = u.userService.Login(ctx.Request.Context(), f)
	web.AssertNil(err)
	var tkInput = token.TokenCreateInput{
		UserID:   us.ID,
		Platform: f.Platform,
	}
	token, err := u.tokenService.Create(ctx, tkInput)
	web.AssertNil(err)
	u.SendData(ctx, auth.ResponseLogin{
		AccessToken: token.ID,
		User:        us,
	})
}

// LogoutHandler
// @Tags         User
// @Summary      Logout user
// @Description  Logout user
// @ID           user-logout
// @Accept       json
// @Produce      json
// @Param        x-user-id  header    string       false  "user id"
// @Param        data       body      models.User  true   "user data"
// @Success      200        {object}  models.User
// @Router       /api/v1/users [put]
func (u UserHandler) LogoutHandler(ctx *gin.Context) {
	var f models.User
	web.AssertNil(ctx.BindJSON(&f))
	userID, err := u.contextWith.GetUserID(ctx)
	web.AssertNil(err)
	err = u.tokenService.RevokeAllByUserID(ctx, userID)
	web.AssertNil(err)
	u.SendData(ctx, nil)
}
