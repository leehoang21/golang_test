package web

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type ContextWith interface {
	SetUserID(ctx *gin.Context, uID string)
	GetUserID(ctx *gin.Context) (string, error)
	GetToken(r *http.Request) string
	GetTokenSocket(r http.Header, requets url.URL) string
	GetTokenPublic(r *http.Request) string
}

func NewContextWith() ContextWith {
	return &ClientCache{}
}

type ClientCache struct {
	Token  string
	UserID string
}

const xCacheClient = "x-cache-client"
const xUserID = "x-user-id"

func (c *ClientCache) GetUserID(ctx *gin.Context) (string, error) {
	uID := ctx.GetString(xUserID)
	if uID == "" {
		return uID, BadRequest("user-id not found")
	}
	return uID, nil
}

func (c *ClientCache) SetUserID(ctx *gin.Context, uID string) {
	ctx.Set(xUserID, uID)
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
