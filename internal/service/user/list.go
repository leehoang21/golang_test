package user

import (
	"base/internal/handler/filter"
	"base/internal/model"
	"context"
)

func (s userService) List(ctx context.Context, f *filter.UserListParams) ([]model.User, error) {
	var res = []model.User{}
	return res, s.userRepo.Search(ctx, f, &res)
}
