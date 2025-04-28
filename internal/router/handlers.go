package router

import "github.com/gin-gonic/gin"

type HandlerFuncs struct {
	// user
	UserCreateHandler    gin.HandlerFunc
	UserUpdateHandler    gin.HandlerFunc
	UserResetPassHandler gin.HandlerFunc
	UserDeleteHandler    gin.HandlerFunc
	UserListHandler      gin.HandlerFunc
	UserGetHandler       gin.HandlerFunc

	// auth
	LoginHandler  gin.HandlerFunc
	LogoutHandler gin.HandlerFunc

	//group role
	GroupRoleCreateHandler        gin.HandlerFunc
	GroupRoleUpdateHandler        gin.HandlerFunc
	GroupRoleDeleteHandler        gin.HandlerFunc
	GroupRoleListHandler          gin.HandlerFunc
	GroupRoleGetHandler           gin.HandlerFunc
	GroupRoleAddMembersHandler    gin.HandlerFunc
	GroupRoleUpdateMembersHandler gin.HandlerFunc
	GroupRoleDeleteMembersHandler gin.HandlerFunc

	//feature
	FeatureCreateHandler           gin.HandlerFunc
	FeatureUpdateHandler           gin.HandlerFunc
	FeatureDeleteHandler           gin.HandlerFunc
	FeatureListHandler             gin.HandlerFunc
	FeatureGetHandler              gin.HandlerFunc
	FeatureAddPermissionHandler    gin.HandlerFunc
	FeatureUpdatePermissionHandler gin.HandlerFunc
	FeatureDeletePermissionHandler gin.HandlerFunc

	// push api
}
