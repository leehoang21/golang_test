package middleware

import (
	"eclectric/internal/base/mlog"
	"eclectric/internal/repository"
	"eclectric/internal/utils/web"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	MidBasicType(role ...string) gin.HandlerFunc
	Recovery() gin.HandlerFunc
}

func NewMiddleware(tokenRepo repository.Token,
	userRepo repository.User,
	contextWith web.ContextWith,
) Middleware {
	return &mid{
		ContextWith: contextWith,
		tokenRepo:   tokenRepo,
		userRepo:    userRepo,
	}
}

type mid struct {
	web.ContextWith
	web.JsonRender
	tokenRepo repository.Token
	userRepo  repository.User
}

func (m mid) MidBasicType(role ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("vo day")
		var tokenID = m.GetToken(ctx.Request)
		fmt.Println("vo da id")
		var tok, err = m.tokenRepo.GetByID(ctx, tokenID)
		if err != nil || tok == nil {
			err = web.Unauthorized("access token not found")
			m.SendErrorForce(ctx, err, http.StatusUnauthorized)
			ctx.Abort()
			return
		}
		m.SetUserID(ctx, tok.UserID)
		ctx.Next()
	}
}

var logMiddle = mlog.NewTagLog("middle")

func (m mid) Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logMiddle.Error(err)

				if httpError, ok := err.(web.IHttpError); ok {
					m.SendErrorForce(c, err.(error), httpError.StatusCode())
				} else {
					fmt.Println(string(debug.Stack()))
					m.SendErrorForce(c, err.(error), http.StatusInternalServerError)
				}
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
