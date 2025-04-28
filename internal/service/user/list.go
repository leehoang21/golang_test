package user

import (
	"base/internal/handler/filter"
	"base/internal/models"
	"context"
)

func (s userService) List(ctx context.Context, input *filter.UserListParams) ([]models.User, int64, error) {
	var res []models.User
	var total, err = s.userRepo.R_SearchAndCount(ctx, input, &res)
	return res, total, err
}
