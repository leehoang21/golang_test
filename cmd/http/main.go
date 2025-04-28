package main

import (
	"base/config"
	"base/internal/handler"
	_ "base/internal/init"
	"base/internal/repository/querymgo"
	"base/internal/router"
	"base/internal/router/middleware"
	"base/internal/service/feature"
	"base/internal/service/group_role"
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
	var userRepo = querymgo.NewUserRepo(db, "users", "usr")
	var tokenRepo = querymgo.NewTokenRepo(db, "tokens", "tk")
	var groupRoleRepo = querymgo.NewRoleRepo(db, "group_roles", "gr")
	var featureRepo = querymgo.NewFeatureRepo(db, "features", "ft")

	var validatorService = validator.NewValidator()
	var userService = user.NewUserService(userRepo, validatorService)
	var tokenService = token.NewTokenService(tokenRepo, validatorService)
	var gRoleService = group_role.NewGroupRoleService(groupRoleRepo, validatorService)
	var featureService = feature.NewFeatureService(featureRepo, validatorService)

	contextWith := web.NewContextWith()
	mid := middleware.NewMiddleware(tokenRepo, userRepo, groupRoleRepo, featureRepo, contextWith)

	handlerUser := handler.NewUserHandler(userService, tokenService, contextWith)
	handlerGroupRole := handler.NewGroupRoleHandler(gRoleService, tokenService, contextWith)
	handlerFeature := handler.NewFeatureHandler(featureService, contextWith)
	var rounterFuncs = router.HandlerFuncs{
		UserCreateHandler:    handlerUser.CreateHandler,
		UserUpdateHandler:    handlerUser.UpdateHandler,
		UserResetPassHandler: handlerUser.ResetPassHandler,
		UserListHandler:      handlerUser.GetListHandler,
		UserDeleteHandler:    handlerUser.DeleteHandler,
		UserGetHandler:       handlerUser.GetHandler,

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
	}

	routerApi := rounterFuncs.Create(mid)

	err := routerApi.Run()
	if err != nil {
		return
	}
}
