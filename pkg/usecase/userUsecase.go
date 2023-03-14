package usecase

import (
	"context"

	domain "ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	services "ecommerce/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

// AllUsers implements interfaces.UserUseCase
func (c *userUseCase) AllUsers(ctx context.Context) ([]domain.UserResponse, error) {
	users, err := c.userRepo.AllUsers(ctx)
	return users, err
}

// Delete implements interfaces.UserUseCase
func (c *userUseCase) Delete(ctx context.Context, userId string) error {
	return c.userRepo.Delete(ctx, userId)
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}



func (c *userUseCase) FindUser(ctx context.Context, id string) (domain.UserResponse, error) {
	user, err := c.userRepo.FindUser(ctx, id)
	return user, err
}

func (c *userUseCase) CreateUser(ctx context.Context, user domain.Users) (domain.UserResponse, error) {
	userres, err := c.userRepo.CreateUser(ctx, user)

	return userres, err
}

