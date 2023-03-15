package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"ecommerce/pkg/config"
	domain "ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interface"
	services "ecommerce/pkg/usecase/interface"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo   interfaces.UserRepository
	adminRepo  interfaces.AdminRepository
	mailConfig config.MailConfig
	config     config.Config
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

// CreateUser implements interfaces.UserUseCase
func (c *userUseCase) CreateUser(ctx context.Context, user domain.Users) (domain.UserResponse, error) {
	fmt.Println("create user from service")
	userres, err := c.userRepo.FindUser(ctx, user.Email)
	fmt.Println("found user", err)

	if err == nil {
		return domain.UserResponse{}, errors.New("username already exists")
	}

	if err != nil && err != sql.ErrNoRows {
		return userres, err
	}

	//hashing password
	user.Password = HashPassword(user.Password)
	fmt.Println("password", user.Password)
	_, err = c.userRepo.CreateUser(ctx, user)
	if err != nil {
		return domain.UserResponse{}, err
	}
	return userres, nil
}

// FindUser implements interfaces.UserUseCase
func (c *userUseCase) FindUser(ctx context.Context, id string) (domain.UserResponse, error) {
	userres, err := c.userRepo.FindUser(ctx, id)

	if err != nil {
		return domain.UserResponse{}, err
	}

	return userres, nil
}

func (c *userUseCase) SendVerificationEmail(email string) error {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	subject := "Account Verification"
	msg := []byte(
		"From: Events Radar <eventsRadarversion1@gmail.com>\r\n" +
			"To: " + email + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
			"<html>" +
			"  <head>" +
			"    <style>" +
			"      .blue-button {" +
			"        background-color: blue;" +
			"        color: white;" +
			"        padding: 10px 20px;" +
			"        border-radius: 5px;" +
			"        text-decoration: none;" +
			"        font-size: 16px;" +
			"      }" +
			"    </style>" +
			"  </head>" +
			"  <body>" +
			"    <p>Click the button on verify your accout:</p>" +
			"    <a class=\"blue-button\" href=\"https://eventsradar.online/user/verify-account?token=" + tokenString + "\" target=\"_blank\">Access Credentials</a>" +
			"  </body>" +
			"</html>")

	// send email with HTML message
	if err := c.mailConfig.SendMail(c.config, email, msg); err != nil {
		return err
	}

	err = c.userRepo.StoreVerificationDetails(email, tokenString)
	if err != nil {
		return err
	}

	return nil
}

// HashPassword hashes the password
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func NewUserUseCase(
	userRepo interfaces.UserRepository,
	adminRepo interfaces.AdminRepository,
	mailConfig config.MailConfig,
	config config.Config) services.UserUseCase {
	return &userUseCase{
		userRepo:   userRepo,
		adminRepo:  adminRepo,
		mailConfig: mailConfig,
		config:     config,
	}
}
