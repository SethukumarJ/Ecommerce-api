package interfaces

import (
	"github.com/golang-jwt/jwt"
	domain "ecommerce/pkg/domain"
)

type JWTUsecase interface {
	GenerateAccessToken(userid string, userName string, role string) (string, error)
	VerifyToken(token string) (bool, *domain.SignedDetails)
	GetTokenFromString(signedToken string, claims *domain.SignedDetails) (*jwt.Token, error)
	GenerateRefreshToken(userid string, userName string, role string) (string, error)

}