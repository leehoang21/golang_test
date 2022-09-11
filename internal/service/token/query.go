package token

import (
	"base/internal/models"
	"context"
)

func (s *tokenService) Create(ctx context.Context, tkInput TokenCreateInput) (*models.Token, error) {
	tk := tkInput.ToModel()
	err := s.tokenRepo.R_Create(ctx, tk)
	if err != nil {
		return nil, err
	}
	return tk, nil
}

func (s *tokenService) RevokeAllByUserID(ctx context.Context, userID string) error {
	return s.tokenRepo.RevokeAllByUserID(ctx, userID)
}

func (s *tokenService) RevokeByID(ctx context.Context, id string) error {
	return s.tokenRepo.RevokeID(ctx, id)
}
