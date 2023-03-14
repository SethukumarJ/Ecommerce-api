package interfaces

import (
	"context"

	domain "ecommerce/pkg/domain"
)

type AdminUsecase interface {
	
	FindAdmin(ctx context.Context, id string) (domain.AdminResponse, error)
	CreateAdmin(ctx context.Context, user domain.Admins) (domain.AdminResponse, error)
	Delete(ctx context.Context, userid string) error
}
