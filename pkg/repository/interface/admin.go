package interfaces

import (
	"context"

	"ecommerce/pkg/domain"
)

type AdminRepository interface {
	
	FindAdmin(ctx context.Context, id string) (domain.AdminResponse, error)
	CreateAdmin(ctx context.Context, user domain.Admins) (domain.AdminResponse, error)
	
	
}
