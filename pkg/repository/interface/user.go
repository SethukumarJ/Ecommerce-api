package interfaces

import (
	"context"

	"ecommerce/pkg/domain"
)

type UserRepository interface {
	AllUsers(ctx context.Context) ([]domain.UserResponse, error)
	FindUserByID(ctx context.Context, id string) (domain.UserResponse, error)
	FindUserByName(ctx context.Context, id string) (domain.UserResponse, error)
	CreateUser(ctx context.Context, user domain.Users) (domain.UserResponse, error)
	DeleteUser(ctx context.Context, userid string) error
	VerifyAccount(ct context.Context,email string, code string) error
	StoreVerificationDetails(email string, code string) error
}
