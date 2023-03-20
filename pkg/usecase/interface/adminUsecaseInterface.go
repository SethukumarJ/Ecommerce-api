package interfaces

import (
	"context"

	domain "ecommerce/pkg/domain"
)

type AdminUsecase interface {
	
	FindAdminById(ctx context.Context, id string) (domain.AdminResponse, error)
	FindAdminByName(ctx context.Context, name string) (domain.AdminResponse, error)
	CreateAdmin(ctx context.Context, user domain.Admins) (domain.AdminResponse, error)
	DeleteAdmin(ctx context.Context, userid string) error
}
