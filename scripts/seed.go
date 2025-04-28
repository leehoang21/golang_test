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
)

func permissionUser(role string) feature.Input {
	return feature.Input{
		RoleNames: []string{role},
		Name:      "User",
		KeyMenu:   "/users",
		URLView:   "/users",
		Status:    enums.StatustypeEnabled.String(),
		Api:       "/users",
		Action:    []string{"/users [post]", "/users [get]", "/users/:id [get]", "/users/:id [delete]"},
	}
}

func permissionGroupRole(role string) feature.Input {
	return feature.Input{
		RoleNames: []string{role},
		Name:      "GroupRole",
		KeyMenu:   "/group-roles",
		URLView:   "/group-roles",
		Status:    enums.StatustypeEnabled.String(),
		Api:       "/group-roles",
		Action:    []string{"/group-roles [post]", "/group-roles [get]", "/group-roles/:id [get]", "/group-roles/:id [put]", "/group-roles/:id [delete]"},
	}
}

func permissionFeature(role string) feature.Input {
	return feature.Input{
		RoleNames: []string{role},
		Name:      "Feature",
		KeyMenu:   "/features",
		URLView:   "/features",
		Status:    enums.StatustypeEnabled.String(),
		Api:       "/features",
		Action:    []string{"/features [post]", "/features [get]", "/features/:id [get]", "/features/:id [put]", "/features/:id [delete]"},
	}
}

func superAdmin() user.Input {
	return user.Input{
		Email:    "admin@admin.com",
		Phone:    "",
		Password: "Admin123@",
		Fullname: "super admin",
		Birthday: "",
	}
}

func roleAdmin(userId string) group_role.GroupRoleInput {
	return group_role.GroupRoleInput{
		Name:        "admin",
		Description: "super admin",
		Status:      enums.StatustypeEnabled.String(),
		Members:     []string{userId},
	}
}

func seedSuperAdmin(ctx context.Context, userService user.Service, groupRoleService group_role.Service, featureService feature.Service) {
	sAdmin, err := userService.Create(ctx, superAdmin())
	if err != nil {
		log.Fatal("Create super admin error:", err)
	}

	sAdminGroupRole, err := groupRoleService.Create(ctx, roleAdmin(sAdmin.ID))
	if err != nil {
		log.Fatal("Create group role error:", err)
	}

	p := permissionUser(sAdminGroupRole.Name)
	_, err = featureService.Create(ctx, p)
	if err != nil {
		log.Fatal("Create feature error:", err)
	}

	p = permissionGroupRole(sAdminGroupRole.Name)
	_, err = featureService.Create(ctx, p)
	if err != nil {
		log.Fatal("Create feature error:", err)
	}

	p = permissionFeature(sAdminGroupRole.Name)
	_, err = featureService.Create(ctx, p)
	if err != nil {
		log.Fatal("Create feature error:", err)
	}

	log.Println("✅ Tạo admin thành công.")
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
	var groupRoleService = group_role.NewGroupRoleService(groupRoleRepo, validatorService)
	var featureService = feature.NewFeatureService(featureRepo, validatorService)

	seedSuperAdmin(ctx, userService, groupRoleService, featureService)
}
