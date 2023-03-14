package interfaces

import "context"

// AuthUseCase is the interface for authentication usecase
type AuthUsecase interface {
	VerifyUser(ctx context.Context, string, password string) error
	VerifyAdmin(ctx context.Context,email string, password string) error
	VerifyAccount(ctx context.Context,email string, code string) error
}
