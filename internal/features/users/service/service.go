package users_service

import (
	"context"
	"github.com/Daty26/todo-app/internal/core/domain"
)

type UsersService struct {
	userRepository UserRepository
}

func NewUsersService(userRepository UserRepository) *UsersService {
	return &UsersService{userRepository: userRepository}
}

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	PatchUser(ctx context.Context, id int, user domain.User) (domain.User, error)
}
