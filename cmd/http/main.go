package main

import (
	//1
	"base/config"
	"base/internal/handler"
	_ "base/internal/init"
	"base/internal/repository/querymgo"
	"base/internal/router"
	"base/internal/service/user"
	"base/internal/utils/validator"
	"context"

	//2
	"base/internal/connection"
)

func main() {
	var ctx = context.Background()
	var db = connection.ConnectDB(ctx, config.DB{})
	var userRepo = querymgo.NewUserRepo(db, "users", "usr")

	var validatorService = validator.NewValidator()
	var userService = user.NewUserService(userRepo, validatorService)
	var handlerFuncs = router.HandlerFuncs{
		UserListHandler: handler.NewUserHandler(userService).ListHandler,
	}
	routerApi := router.Create(handlerFuncs)
	routerApi.Run(":8080")
}
