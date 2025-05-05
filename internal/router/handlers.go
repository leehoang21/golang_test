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
	UserProfileHandler   gin.HandlerFunc

	// auth
	LoginHandler  gin.HandlerFunc
	LogoutHandler gin.HandlerFunc

	//group role
	GroupRoleCreateHandler gin.HandlerFunc
	GroupRoleUpdateHandler gin.HandlerFunc
	GroupRoleDeleteHandler gin.HandlerFunc
	GroupRoleListHandler   gin.HandlerFunc
	GroupRoleGetHandler    gin.HandlerFunc

	//feature
	FeatureCreateHandler gin.HandlerFunc
	FeatureUpdateHandler gin.HandlerFunc
	FeatureDeleteHandler gin.HandlerFunc
	FeatureListHandler   gin.HandlerFunc
	FeatureGetHandler    gin.HandlerFunc

	// task
	TaskCreateHandler gin.HandlerFunc
	TaskUpdateHandler gin.HandlerFunc
	TaskDeleteHandler gin.HandlerFunc
	TaskListHandler   gin.HandlerFunc
	TaskGetHandler    gin.HandlerFunc
	//estimate
	EstimateCreateHandler gin.HandlerFunc
	EstimateUpdateHandler gin.HandlerFunc
	EstimateDeleteHandler gin.HandlerFunc
	EstimateGetHandler    gin.HandlerFunc

	//report
	SendTelegramReportHandler gin.HandlerFunc
}
