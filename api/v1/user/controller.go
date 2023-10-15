package user

import (
	"chat-hex/api/common"
	"chat-hex/api/v1/user/request"
	"chat-hex/api/v1/user/response"
	"chat-hex/business/user"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

//Controller Get item API controller
type Controller struct {
	service user.Service
}

//NewController Construct item API controller
func NewController(service user.Service) *Controller {
	return &Controller{
		service,
	}
}

//GetItemByID Get item by ID echo handler
func (controller *Controller) FindUserByID(c echo.Context) error {
	id := c.Param("id")

	user, err := controller.service.FindUserByID(id)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetUserResponse(*user)

	return c.JSON(common.NewSuccessResponse(response))
}

//FindAllUserWithPagination Find All User with pagination handler
func (controller *Controller) FindAllUserWithPagination(c echo.Context) error {

	pageQueryParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageQueryParam)
	if err != nil || page <= 0 {
		page = 1
	}

	rowPerPageQueryParam := c.QueryParam("row_per_page")
	rowPerPage, err := strconv.Atoi(rowPerPageQueryParam)
	if err != nil || rowPerPage <= 0 {
		rowPerPage = 10
	}

	skip := (page * rowPerPage) - rowPerPage

	users, err := controller.service.FindAllUserWithPagination(skip, rowPerPage)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetAllUserResponse(users, page, rowPerPage)

	return c.JSON(common.NewSuccessResponse(response))
}

// InsertUser Create new user echo handler
func (controller *Controller) InsertUser(c echo.Context) error {
	insertUserRequest := new(request.InsertUserRequest)

	if err := c.Bind(insertUserRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.InsertUser(*insertUserRequest.ToUpsertUserSpec(), "creator")
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

// UpdateUser update existing user
func (controller *Controller) UpdateUser(c echo.Context) error {
	id := c.Param("id")

	updateUserRequest := new(request.UpdateUserRequest)

	if err := c.Bind(updateUserRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.UpdateUser(id, updateUserRequest.Name, "modifier", updateUserRequest.Version)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
