package usecase

import (
	"context"

	domain "ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	services "ecommerce/pkg/usecase/interface"
)

type adminUseCase struct {
	adminRepo interfaces.AdminRepository
}

// FindAdminById implements interfaces.AdminUsecase
func (c *adminUseCase) FindAdminById(ctx context.Context, id string) (domain.AdminResponse, error) {
	user, err := c.adminRepo.FindAdmin(ctx, id)
	
	return user, err
}

// FindAdminByName implements interfaces.AdminUsecase
func (c *adminUseCase) FindAdminByName(ctx context.Context, name string) (domain.AdminResponse, error) {
	user, err := c.adminRepo.FindAdmin(ctx, name)
	
	return user, err
}

// CreateAdmin implements interfaces.AdminUsecase
func (c *adminUseCase) CreateAdmin(ctx context.Context, user domain.Admins) (domain.AdminResponse, error) {
	userres, err := c.adminRepo.CreateAdmin(ctx, user)

	return userres, err
}

// Delete implements interfaces.AdminUsecase
func (c *adminUseCase) DeleteAdmin(ctx context.Context, userid string) error {
	 err := c.adminRepo.DeleteAdmin(ctx, userid)

	return err
}

// FindAdmin implements interfaces.AdminUsecase
func (c *adminUseCase) FindAdmin(ctx context.Context, id string) (domain.AdminResponse, error) {
	user, err := c.adminRepo.FindAdmin(ctx, id)
	
	return user, err
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUsecase {
	return &adminUseCase{
		adminRepo: repo,
	}
}
