package main

import (
	"base/config"
	"base/internal/connection"
	"base/internal/data/enums"
	"base/internal/repository/querymgo"
	"base/internal/service/feature"
	"base/internal/service/group_role"
	"base/internal/service/user"
	"base/internal/utils/validator"
	"context"
	"log"
	"os"
)

// superuser
func superUser() user.Input {
	return user.Input{
		Email:    os.Getenv("EV_ADMIN_EMAIL"),
		Phone:    "",
		Password: os.Getenv("EV_ADMIN_PASSWORD"),
		Fullname: "Super Admin",
		Birthday: "",
	}
}

// role
func role(id string) group_role.GroupRoleInput {
	return group_role.GroupRoleInput{
		Name:        "admin",
		Description: "admin",
		Status:      enums.StatustypeEnabled.String(),
		MemberIds:   []string{id},
	}
}

// feature
func featureUser() feature.Input {
	return feature.Input{
		Name:      "admin",
		KeyMenu:   "/users",
		URLView:   "/users",
		Api:       "/users",
		Action:    []string{"/users [post]", "/users [get]", "/users/:id [get]", "/users/:id [delete]"},
		Status:    enums.StatustypeEnabled.String(),
		RoleNames: []string{"admin"},
	}
}

func featureGroupRole() feature.Input {
	return feature.Input{
		Name:      "group_role",
		KeyMenu:   "/group-roles",
		URLView:   "/group-roles",
		Api:       "/group-roles",
		Action:    []string{"/group-roles [post]", "/group-roles [put]", "/group-roles [get]", "/group-roles/:id [get]", "/group-roles/:id [delete]"},
		Status:    enums.StatustypeEnabled.String(),
		RoleNames: []string{"admin"},
	}
}

func featureFeature() feature.Input {
	return feature.Input{
		Name:      "feature",
		KeyMenu:   "/features",
		URLView:   "/features",
		Api:       "/features",
		Action:    []string{"/features [post]", "/features [get]", "/features/:id [get]", "/features/:id [delete]"},
		Status:    enums.StatustypeEnabled.String(),
		RoleNames: []string{"admin"},
	}

}

func main() {
	var ctx = context.Background()
	var cf = config.LoadEnv()
	var db = connection.ConnectDB(ctx, cf.DB)
	var userRepo = querymgo.NewUserRepo(db, "users", "usr")
	var groupRoleRepo = querymgo.NewRoleRepo(db, "group_roles", "gr")
	var featureRepo = querymgo.NewFeatureRepo(db, "features", "ft")

	var validatorService = validator.NewValidator()
	var userService = user.NewUserService(userRepo, validatorService)
	var gRoleService = group_role.NewGroupRoleService(groupRoleRepo, validatorService)
	var featureService = feature.NewFeatureService(featureRepo, validatorService)

	u, err := userService.Create(ctx, superUser())
	if err != nil {
		log.Fatal(err)
	}
	//role
	_, err = gRoleService.Create(ctx, role(u.ID))
	if err != nil {
		log.Fatal(err)
	}
	//feature
	_, err = featureService.Create(ctx, featureUser())
	if err != nil {
		log.Fatal(err)

	}
	_, err = featureService.Create(ctx, featureGroupRole())
	if err != nil {
		log.Fatal(err)
	}
	_, err = featureService.Create(ctx, featureFeature())
	if err != nil {
		log.Fatal(err)
	}

}
