package web

import (
	"base/internal/models"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type ContextWith interface {
	SetUser(ctx *gin.Context, u models.User)
	GetUser(ctx *gin.Context) (*models.User, error)
	GetToken(r *http.Request) string
	GetTokenSocket(r http.Header, requets url.URL) string
	GetTokenPublic(r *http.Request) string
}

func NewContextWith() ContextWith {
	return &ClientCache{}
}

type ClientCache struct {
	Token string
	User  models.User
}

const xCacheClient = "x-cache-client"
const xUser = "x-user"

func (c *ClientCache) GetUser(ctx *gin.Context) (*models.User, error) {
	u, isExit := ctx.Get(xUser)
	if !isExit {
		return nil, BadRequest("user not found")
	}
	//check type
	user, ok := u.(models.User)
	if !ok {
		return nil, BadRequest("user not found")
	}
	return &user, nil
}

func (c *ClientCache) SetUser(ctx *gin.Context, u models.User) {
	ctx.Set(xUser, u)
}

func (c *ClientCache) ContextWithUser(ctx *gin.Context, u *ClientCache) {
	ctx.Set(xCacheClient, u)
}

func (c *ClientCache) GetTokenFromContext(ctx *gin.Context) *ClientCache {
	if token, ok := ctx.Get(xCacheClient); ok {
		return token.(*ClientCache)
	}
	return nil
}

const bearerHeader = "Bearer "
const accessToken = "access_token"

func (c *ClientCache) GetToken(r *http.Request) string {
	var authHeader = r.Header.Get("Authorization")
	if strings.HasPrefix(authHeader, bearerHeader) {
		return strings.TrimPrefix(authHeader, bearerHeader)
	}
	return r.URL.Query().Get(accessToken)
}

func (c *ClientCache) GetTokenSocket(r http.Header, requets url.URL) string {
	var authHeader = r.Get("Authorization")
	if strings.HasPrefix(authHeader, bearerHeader) {
		return strings.TrimPrefix(authHeader, bearerHeader)
	}
	return requets.Query().Get(accessToken)
}

func (c *ClientCache) GetTokenPublic(r *http.Request) string {
	var authHeader = r.Header.Get("public")
	if strings.HasPrefix(authHeader, bearerHeader) {
		return strings.TrimPrefix(authHeader, bearerHeader)
	}
	return ""
}
