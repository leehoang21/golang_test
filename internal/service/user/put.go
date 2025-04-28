package user

import (
	"base/internal/models"
	hashpassword "base/internal/utils/hash_password"
	"base/internal/utils/web"
	"context"
	"log"
)

func (s *userService) Update(ctx context.Context, id string, input Input) (*models.User, error) {
	var uExist *models.User
	if err := s.validator.ValidateStruct(input); err != nil {
		return nil, err
	}

	if err := s.userRepo.R_SelectByID(ctx, id, &uExist); err != nil {
		return nil, err
	}

	uExist = input.ToModel()
	uEmail, _ := s.userRepo.GetByEmail(ctx, uExist.Email)
	log.Println(uEmail)
	if uEmail, _ := s.userRepo.GetByEmail(ctx, uExist.Email); uEmail != nil && id != uEmail.ID {
		return nil, web.BadRequest("Email đã tồn tại")
	}
	if uExist.Password.String() != "" {
		pass, err := uExist.Password.GenerateHashedPassword()
		if err != nil {
			return nil, err
		}
		uExist.Password = hashpassword.NewPassword(pass)
	}
	uExist.ID = id
	err := s.userRepo.R_Update(ctx, uExist)
	if err != nil {
		return nil, err
	}
	return uExist, nil
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
	if err != nil {
		return nil, err
	}
	return uExist, nil
}
