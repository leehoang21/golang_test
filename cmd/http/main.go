package main

import (
	"base/config"
	"base/internal/handler"
	_ "base/internal/init"
	"base/internal/repository/querymgo"
	"base/internal/router"
	"base/internal/router/middleware"
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
	var db = connection.ConnectDB(ctx, config.LoadEnv().DB)
	var userRepo = querymgo.NewUserRepo(db, "users", "usr")
	var tokenRepo = querymgo.NewTokenRepo(db, "tokens", "tk")

	var validatorService = validator.NewValidator()
	var userService = user.NewUserService(userRepo, validatorService)
	//var tokenService = token.NewTokenService(tokenRepo, validatorService)
	contextWith := web.NewContextWith()
	mid := middleware.NewMiddleware(tokenRepo, userRepo, contextWith)

	handlerUser := handler.NewUserHandler(userService, contextWith)
	var rounterFuncs = router.HandlerFuncs{
		UserCreateHandler: handlerUser.CreateHandler,
		LogoutHandler:     handlerUser.LogoutHandler,
		LoginHandler:      handlerUser.LoginHandler,
		UserUpdateHandler: handlerUser.UpdateHandler,
	}

	routerApi := rounterFuncs.Create(mid)
	routerApi.Run(":8080")
}
