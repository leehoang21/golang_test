package user

import (
	"base/internal/models"
	hashpassword "base/internal/utils/hash_password"
	"base/internal/utils/web"
	"context"
)

func (s *userService) Update(ctx context.Context, id string, input Input) (*models.User, error) {
	var uExist *models.User
	if err := s.validator.ValidateStruct(input); err != nil {
		return nil, err
	}

	if err := s.userRepo.R_SelectByID(ctx, id, &uExist); err != nil {
		return nil, err
	}

	if uEmail, _ := s.userRepo.GetByEmail(ctx, input.Email); uEmail != nil && id != uEmail.ID {
		return nil, web.BadRequest("Email đã tồn tại")
	}

	baseModel := uExist.BaseModel
	uExist = input.ToModel()
	if uExist.Password.String() != "" {
		pass, err := uExist.Password.GenerateHashedPassword()
		if err != nil {
			return nil, err
		}
		uExist.Password = hashpassword.NewPassword(pass)
	}

	uExist.BaseModel = baseModel
	uExist.ID = id
	err := s.userRepo.R_Update(ctx, uExist)
	return uExist, err
}

func (s *userService) ResetPass(ctx context.Context, id string, user ResetPassInput) (*models.User, error) {
	var uExist *models.User
	if err := s.userRepo.R_SelectByID(ctx, id, &uExist); err != nil {
		return nil, err
	}
	err := uExist.Password.ComparePassword(hashpassword.Password(user.PasswordOld))
	if err != nil {
		return nil, err
	}
	pass, err := hashpassword.NewPassword(user.PasswordNew).GenerateHashedPassword()
	if err != nil {
		return nil, err
	}
	uExist.Password = hashpassword.NewPassword(pass)
	err = s.userRepo.R_Update(ctx, uExist)
	return uExist, err
}
