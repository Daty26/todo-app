package users_service

import (
	"context"
	"fmt"
	"github.com/Daty26/todo-app/internal/core/domain"
	core_errors "github.com/Daty26/todo-app/internal/core/errors"
)

func (s *UsersService) GetUsers(
	ctx context.Context,
	limit *int, offset *int,
) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("limit must be non-negative: %w", core_errors.ErrInvalidArgument)
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset must be non-negastive: %w", core_errors.ErrInvalidArgument)
	}
	getUsers, err := s.userRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get users from repo: %w", err)
	}

	return getUsers, nil
}
