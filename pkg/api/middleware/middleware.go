package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"ecommerce/pkg/response"
	"ecommerce/pkg/utils"
	usecase "ecommerce/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	AuthorizeJwtUser() gin.HandlerFunc
	AuthorizeJwtAdmin() gin.HandlerFunc
}

type middleware struct {
	jwtUsecase  usecase.JWTUsecase
	userUsecase usecase.UserUseCase
}

func NewMiddlewareUser(jwtUserUsecase usecase.JWTUsecase, userUsecase usecase.UserUseCase) Middleware {
	return &middleware{
		jwtUsecase:  jwtUserUsecase,
		userUsecase: userUsecase,
	}

}

func (cr *middleware) AuthorizeJwtUser() gin.HandlerFunc {
	return (func(c *gin.Context) {

		//getting from header
		autheader := c.Request.Header["Authorization"]
		auth := strings.Join(autheader, " ")
		bearerToken := strings.Split(auth, " ")
		fmt.Printf("\n\ntocen : %v\n\n", autheader)

		if len(bearerToken) != 2 {
			err := errors.New("request does not contain an access token")
			response := response.ErrorResponse("Failed to autheticate jwt", err.Error(), nil)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()

			return
		}

		authtoken := bearerToken[1]
		ok, claims := cr.jwtUsecase.VerifyToken(authtoken)
		source := fmt.Sprint(claims.Source)
		fmt.Println("///////////////token role", claims.Role)
		if claims.Role != "user" {
			err := errors.New("your role of the token is not valid")
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		}

		if !ok && source == "accesstoken" {
			err := errors.New("your access token is not valid")
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		} else if !ok && source == "refreshtoken" {
			err := errors.New("your refresh token is not valid")
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		} else if ok && source == "refreshtoken" {

			c.Writer.Header().Set("Authorization", authtoken)
			c.Next()
		} else {

			userid := fmt.Sprint(claims.UserId)
			c.Writer.Header().Set("userid", userid)
			c.Next()

		}

	})
}

func (cr *middleware) AuthorizeJwtAdmin() gin.HandlerFunc {
	return (func(c *gin.Context) {

		//getting from header
		autheader := c.Request.Header["Authorization"]
		auth := strings.Join(autheader, " ")
		bearerToken := strings.Split(auth, " ")
		fmt.Printf("\n\ntocen : %v\n\n", autheader)

		if len(bearerToken) != 2 {
			err := errors.New("request does not contain an access token")
			response := response.ErrorResponse("Failed to autheticate jwt", err.Error(), nil)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()

			return
		}

		authtoken := bearerToken[1]
		ok, claims := cr.jwtUsecase.VerifyToken(authtoken)
		source := fmt.Sprint(claims.Source)
		fmt.Println("///////////////token role", claims.Role)
		if claims.Role != "admin" {
			err := errors.New("your role of the token is not valid")
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		}

		if !ok && source == "accesstoken" {
			err := errors.New("your access token is not valid")
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		} else if !ok && source == "refreshtoken" {
			err := errors.New("your refresh token is not valid")
			response := response.ErrorResponse("Error", err.Error(), source)
			c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			utils.ResponseJSON(*c, response)
			c.Abort()
			return
		} else if ok && source == "refreshtoken" {

			c.Writer.Header().Set("Authorization", authtoken)
			c.Next()
		} else {

			adminid := fmt.Sprint(claims.UserId)
			c.Writer.Header().Set("adminid", adminid)
			c.Next()

		}

	})
}
