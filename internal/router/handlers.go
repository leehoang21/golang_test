package router

import "github.com/gin-gonic/gin"

type HandlerFuncs struct {
	// user
	UserCreateHandler gin.HandlerFunc
	UserUpdateHandler gin.HandlerFunc
	UserDeleteHandler gin.HandlerFunc
	// auth
	LoginHandler  gin.HandlerFunc
	LogoutHandler gin.HandlerFunc
	// push api
}
