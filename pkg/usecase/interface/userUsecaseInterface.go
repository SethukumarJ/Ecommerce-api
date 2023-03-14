package interfaces

import (
	"context"

	domain "ecommerce/pkg/domain"
)

type UserUseCase interface {
	AllUsers(ctx context.Context) ([]domain.UserResponse, error)
	FindUser(ctx context.Context, id string) (domain.UserResponse, error)
	CreateUser(ctx context.Context, user domain.Users) (domain.UserResponse, error)
	Delete(ctx context.Context, userid string) error
}
