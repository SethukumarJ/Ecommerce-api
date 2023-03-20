package interfaces

import (
	"context"

	domain "ecommerce/pkg/domain"
)

type UserUseCase interface {
	AllUsers(ctx context.Context) ([]domain.UserResponse, error)
	FindUserByID(ctx context.Context, id string) (domain.UserResponse, error)
	FindUserByName(ctx context.Context, name string) (domain.UserResponse, error)
	CreateUser(ctx context.Context, user domain.Users) (domain.UserResponse, error)
	DeleteUser(ctx context.Context, userid string) error
}
