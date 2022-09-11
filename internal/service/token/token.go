package token

import (
	"base/internal/models"
	"base/internal/repository"
	"context"

	"base/internal/utils/validator"
)

type Service interface {
	Create(ctx context.Context, token TokenCreateInput) (*models.Token, error)
	RevokeAllByUserID(ctx context.Context, userID string) error
	RevokeByID(ctx context.Context, id string) error
}

type tokenService struct {
	tokenRepo repository.Token
	validator validator.Validator
}

func NewTokenService(tokenRepo repository.Token,
	validator validator.Validator) Service {
	return &tokenService{
		tokenRepo: tokenRepo,
		validator: validator,
	}
}
