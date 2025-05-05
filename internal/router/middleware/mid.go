package middleware

import (
	"base/internal/base/mlog"
	"base/internal/data/enums"
	"base/internal/models"
	"base/internal/repository"
	"base/internal/utils/web"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	MidBasicType(groupName string) gin.HandlerFunc
	Recovery() gin.HandlerFunc
}

func NewMiddleware(tokenRepo repository.Token,
	userRepo repository.User,
	roleRepo repository.GroupRole,
	featureRepo repository.Feature,
	contextWith web.ContextWith,

) Middleware {
	return &mid{
		ContextWith: contextWith,
		tokenRepo:   tokenRepo,
		userRepo:    userRepo,
		roleRepo:    roleRepo,
		featureRepo: featureRepo,
	}
}

type mid struct {
	web.ContextWith
	web.JsonRender
	tokenRepo   repository.Token
	userRepo    repository.User
	roleRepo    repository.GroupRole
	featureRepo repository.Feature
}

func (m mid) MidBasicType(groupName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenID = m.GetToken(ctx.Request)
		var tok, err = m.tokenRepo.GetByID(ctx, tokenID)
		if err != nil || tok == nil {
			err = web.Unauthorized("access token not found")
			m.SendErrorForce(ctx, err, http.StatusUnauthorized)
			ctx.Abort()
			return
		}
		var us models.User
		err = m.userRepo.R_SelectByID(ctx, tok.UserID, &us)
		if err != nil {
			err = web.Unauthorized("user not found")
			m.SendErrorForce(ctx, err, http.StatusUnauthorized)
			ctx.Abort()
			return
		}
		m.SetUser(ctx, us)
		//check permission
		permission(tok.UserID, m, ctx, groupName)

	}
}

func permission(userId string, m mid, ctx *gin.Context, groupName string) {
	if groupName == "" {
		ctx.Next()
		return
	}
	f, err := m.featureRepo.GetByApi(ctx, groupName)
	if err != nil || len(f.RoleNames) == 0 || f.Status != enums.StatustypeEnabled.String() {
		err = web.Forbidden("access denied")
		m.SendErrorForce(ctx, err, http.StatusForbidden)
		ctx.Abort()
		return
	}

	gRoles, err := m.roleRepo.GetByUserID(ctx, userId)
	for _, p := range f.RoleNames {
		for _, g := range gRoles {
			if g.Name == p {
				ctx.Next()
				return
			}
		}

	}
	err = web.Forbidden("access denied")
	m.SendErrorForce(ctx, err, http.StatusForbidden)
	ctx.Abort()
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
