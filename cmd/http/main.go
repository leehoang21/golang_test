package main

import (
	"base/config"
	"base/internal/handler"
	_ "base/internal/init"
	"base/internal/notification"
	"base/internal/repository/querymgo"
	"base/internal/router"
	"base/internal/router/middleware"
	"base/internal/service/estimate"
	"base/internal/service/feature"
	"base/internal/service/group_role"
	"base/internal/service/report"
	"base/internal/service/task"
	"base/internal/service/token"
	"base/internal/service/user"
	"base/internal/utils/validator"
	"base/internal/utils/web"
	"context"

	_ "base/docs"
	"base/internal/connection"
)

// @title           DJM API
// @version         1.0
// @description     This is DJM api docs.
// @termsOfService  http://swagger.io/terms/

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath
func main() {
	var ctx = context.Background()
	var cf = config.LoadEnv()
	var db = connection.ConnectDB(ctx, cf.DB)
	var sender = notification.NewSender(cf.TelegramConfig)
	var userRepo = querymgo.NewUserRepo(db, "users", "usr")
	var tokenRepo = querymgo.NewTokenRepo(db, "tokens", "tk")
	var groupRoleRepo = querymgo.NewRoleRepo(db, "group_roles", "gr")
	var featureRepo = querymgo.NewFeatureRepo(db, "features", "ft")
	var taskRepo = querymgo.NewTaskRepo(db, "tasks", "t")
	var estimateRepo = querymgo.NewEstimateRepo(db, "estimates", "es")

	var validatorService = validator.NewValidator()
	var userService = user.NewUserService(userRepo, validatorService)
	var tokenService = token.NewTokenService(tokenRepo, validatorService)
	var gRoleService = group_role.NewGroupRoleService(groupRoleRepo, validatorService)
	var featureService = feature.NewFeatureService(featureRepo, validatorService)
	var taskService = task.NewtaskService(taskRepo, validatorService)
	var estimateService = estimate.NewestimateService(estimateRepo, validatorService)
	var reportService = report.NewReportService(taskRepo, validatorService)

	contextWith := web.NewContextWith()
	mid := middleware.NewMiddleware(tokenRepo, userRepo, groupRoleRepo, featureRepo, contextWith)

	handlerUser := handler.NewUserHandler(userService, tokenService, contextWith)
	handlerGroupRole := handler.NewGroupRoleHandler(gRoleService, contextWith)
	handlerFeature := handler.NewFeatureHandler(featureService, contextWith)
	handlerTask := handler.NewTaskHandler(taskService, contextWith)
	handlerEstimate := handler.NewEstimateHandler(estimateService, contextWith)
	handlerReport := handler.NewReportHandler(reportService, contextWith, sender)

	var rounterFuncs = router.HandlerFuncs{
		UserCreateHandler:    handlerUser.CreateHandler,
		UserUpdateHandler:    handlerUser.UpdateHandler,
		UserResetPassHandler: handlerUser.ResetPassHandler,
		UserListHandler:      handlerUser.GetListHandler,
		UserDeleteHandler:    handlerUser.DeleteHandler,
		UserGetHandler:       handlerUser.GetHandler,
		UserProfileHandler:   handlerUser.ProfileHandler,

		LogoutHandler: handlerUser.LogoutHandler,
		LoginHandler:  handlerUser.LoginHandler,

		GroupRoleCreateHandler: handlerGroupRole.CreateHandler,
		GroupRoleUpdateHandler: handlerGroupRole.UpdateHandler,
		GroupRoleListHandler:   handlerGroupRole.GetListHandler,
		GroupRoleDeleteHandler: handlerGroupRole.DeleteHandler,
		GroupRoleGetHandler:    handlerGroupRole.GetHandler,

		FeatureCreateHandler: handlerFeature.CreateHandler,
		FeatureUpdateHandler: handlerFeature.UpdateHandler,
		FeatureListHandler:   handlerFeature.GetListHandler,
		FeatureDeleteHandler: handlerFeature.DeleteHandler,
		FeatureGetHandler:    handlerFeature.GetHandler,

		TaskCreateHandler: handlerTask.CreateHandler,
		TaskUpdateHandler: handlerTask.UpdateHandler,
		TaskListHandler:   handlerTask.GetListHandler,
		TaskDeleteHandler: handlerTask.DeleteHandler,
		TaskGetHandler:    handlerTask.GetHandler,

		EstimateCreateHandler: handlerEstimate.CreateHandler,
		EstimateUpdateHandler: handlerEstimate.UpdateHandler,
		EstimateDeleteHandler: handlerEstimate.DeleteHandler,
		EstimateGetHandler:    handlerEstimate.GetHandler,

		SendTelegramReportHandler: handlerReport.SendTelegramReportHandler,
	}

	routerApi := rounterFuncs.Create(mid)

	err := routerApi.Run()
	if err != nil {
		return
	}
}
