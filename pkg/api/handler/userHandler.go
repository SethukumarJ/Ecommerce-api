package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	domain "ecommerce/pkg/domain"
	"ecommerce/pkg/response"
	services "ecommerce/pkg/usecase/interface"
	utils "ecommerce/pkg/utils"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}


func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}




// FindOne godoc
// @summary Get one users
// @description Get one users
// @tags users
// @id FindOne
// @produce json
// @Param        userId   query      string  true  "User Id : "
// @Router /users [get]
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
func (cr *UserHandler) FindByID(c *gin.Context) {
	userId := c.Query("userId")
	fmt.Println(userId)

	user, err := cr.userUseCase.FindUser(c.Request.Context(), userId)
	fmt.Printf("\n\nuser  : %v\n\nerr  %v\n\n", user, err)

	if err != nil {
		response := response.ErrorResponse("FAILL", err.Error(), nil)
		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", user)
	utils.ResponseJSON(*c, response)

}

// FindAll godoc
// @summary Get all users
// @description Save user
// @tags users
// @id Save
// @param RegisterAdmin body domain.Users{} true "admin signup with username, phonenumber email ,password"
// @produce json
// @Router /api/users [Post]
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
func (cr *UserHandler) Save(c *gin.Context) {
	var newUser domain.Users

	if err := c.Bind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := cr.userUseCase.CreateUser(c.Request.Context(), newUser)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	user.Password = ""
	response := response.SuccessResponse(true, "SUCCESS", user)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// DeleteOne godoc
// @summary Delete one users
// @description Delete one users
// @tags users
// @id DeleteOne
// @produce json
// @Param        userId   query      string  true  "User Id : "
// @Router /users [delete]
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
func (cr *UserHandler) Delete(c *gin.Context) {
	userId := c.Query("userId")

	ctx := c.Request.Context()
	user, err := cr.userUseCase.FindUser(ctx, userId)

	if err != nil {
		response := response.ErrorResponse("FAILL", err.Error(), nil)
		utils.ResponseJSON(*c, response)
		return
	}

	if user == (domain.UserResponse{}) {
		response := response.ErrorResponse("FAILL", "There is no users with your id check id", nil)
		utils.ResponseJSON(*c, response)
		return
	}

	err = cr.userUseCase.Delete(ctx, userId)
	if err != nil {
		response := response.ErrorResponse("FAILL", err.Error(), nil)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", nil)
	utils.ResponseJSON(*c, response)
}
