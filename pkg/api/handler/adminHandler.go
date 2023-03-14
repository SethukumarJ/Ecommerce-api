package handler

import (
	"fmt"
	"net/http"

	domain "ecommerce/pkg/domain"
	"ecommerce/pkg/response"
	usecase "ecommerce/pkg/usecase/interface"
	"ecommerce/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUsecase usecase.AdminUsecase
	userUsecase  usecase.UserUseCase
}

func NewAdminHandler(
	adminUsecase usecase.AdminUsecase,
	userUsecase usecase.UserUseCase,
) AdminHandler {
	return AdminHandler{
		adminUsecase: adminUsecase,
		userUsecase:  userUsecase,
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
func (cr *AdminHandler) FindAdmin(c *gin.Context) {
	userId := c.Query("userId")
	fmt.Println(userId)

	user, err := cr.adminUsecase.FindAdmin(c,userId)
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
func (cr *AdminHandler) CreateAdmin(c *gin.Context) {
	var newUser domain.Admins

	if err := c.Bind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	_,err := cr.adminUsecase.CreateAdmin(c,newUser)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	newUser.Password = ""
	response := response.SuccessResponse(true, "SUCCESS", newUser)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary list all active users for admin
// @ID list all active users
// @Tags Admin-User Profile
// @Produce json
// @Security BearerAuth
// @Param  page   query  string  true  "Page number: "
// @Param  pagesize   query  string  true  "Page capacity : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/list-users [get]
func (cr *UserHandler) ViewAllUsers(c *gin.Context) {

	users, err := cr.userUseCase.AllUsers(c.Request.Context())

	if err != nil {
		response := response.ErrorResponse("error while getting users from database", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "Listed All Users", users)
	utils.ResponseJSON(*c, response)

}
