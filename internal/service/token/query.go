package token

import (
	"context"
	"eclectric/internal/models"
)

func (s *tokenService) Create(ctx context.Context, tkInput *models.Token) (*models.Token, error) {
	err := s.tokenRepo.Create(ctx, tkInput)
	if err != nil {
		return nil, err
	}
	return tkInput, nil
}

func (s *tokenService) RevokeAllByUserID(ctx context.Context, userID string) error {
	return s.tokenRepo.RevokeAllByUserID(ctx, userID)
}

func (s *tokenService) RevokeByID(ctx context.Context, id string) error {
	return s.tokenRepo.RevokeID(ctx, id)
}
