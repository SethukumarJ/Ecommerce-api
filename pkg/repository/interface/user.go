package interfaces

import (
	"context"

	"ecommerce/pkg/domain"
)

type UserRepository interface {
	AllUsers(ctx context.Context) ([]domain.UserResponse, error)
	FindUser(ctx context.Context, id string) (domain.UserResponse, error)
	CreateUser(ctx context.Context, user domain.Users) (domain.UserResponse, error)
	Delete(ctx context.Context, userid string) error
	VerifyAccount(ct context.Context,email string, code string) error
	StoreVerificationDetails(email string, code string) error
}
