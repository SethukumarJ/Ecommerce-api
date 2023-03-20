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
// @summary Get admin
// @description Get one users
// @tags admins
// @id FindAdminyId
// @produce json
// @Param    adminId   query      string  true  "admin Id : "
// @Router /admin/id [get]
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
func (cr *AdminHandler) FindAdminyId(c *gin.Context) {
	adminId := c.Query("adminId")
	fmt.Println(adminId)

	user, err := cr.adminUsecase.FindAdminById(c,adminId)
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
// @id CeateAdmin
// @param RegisterAdmin body domain.Admins{} true "admin signup with username, phonenumber, email ,password"
// @produce json
// @Router /admin [Post]
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
func (cr *AdminHandler) CreateAdmin(c *gin.Context) {
	var newUser domain.Admins

	if err := c.Bind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	admin,err := cr.adminUsecase.CreateAdmin(c,newUser)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)
		utils.ResponseJSON(*c, response)
		return
	}

	newUser.Password = ""
	response := response.SuccessResponse(true, "SUCCESS", admin)
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary list all active users for admin
// @ID list all active users
// @Tags Admin-User Profile
// @Produce json
// @Security BearerAuth
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
	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)

}
