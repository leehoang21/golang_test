package user

import (
	"base/internal/models"
	hashpassword "base/internal/utils/hash_password"
	"base/internal/utils/web"
	"context"
)

func (s *userService) Update(ctx context.Context, id string, user *models.User) (*models.User, error) {
	var uExist *models.User
	if err := s.userRepo.R_SelectByID(ctx, id, &uExist); err != nil {
		return nil, err
	}
	if uEmail, _ := s.userRepo.GetByEmail(ctx, user.Email); uEmail != nil && id != uEmail.ID {
		return nil, web.BadRequest("Email đã tồn tại")
	}
	if user.Password.String() != "" {
		pass, err := user.Password.GererateHashedPassword()
		if err != nil {
			return nil, err
		}
		user.Password = hashpassword.NewPassword(pass)
	}
	user.ID = id
	err := s.userRepo.R_Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) ResetPass(ctx context.Context, id string, user UserResetPassInput) (*models.User, error) {
	var uExist *models.User
	if err := s.userRepo.R_SelectByID(ctx, id, &uExist); err != nil {
		return nil, err
	}
	err := uExist.Password.ComparePassword(hashpassword.Password(user.PasswordOld))
	if err != nil {
		return nil, err
	}
	pass, err := hashpassword.NewPassword(user.PasswordNew).GererateHashedPassword()
	if err != nil {
		return nil, err
	}
	uExist.Password = hashpassword.NewPassword(pass)
	err = s.userRepo.R_Update(ctx, uExist)
	if err != nil {
		return nil, err
	}
	return uExist, nil
}
