package auth

import (
	"chat-hex/api/common"
	"chat-hex/api/v1/auth/requests"
	"chat-hex/business/auth"
	"chat-hex/business/users"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service auth.Service
	usersService users.Service
}

func NewController(service auth.Service, usersService users.Service) *Controller {
	return &Controller{
		service,
		usersService,
	}
}

func (controller *Controller) GenerateToken(c echo.Context) error {
	generateTokenRequest := new(requests.GenerateTokenRequest)

	err := c.Bind(generateTokenRequest)
	if err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	user, err := controller.usersService.FindUserByEmail(generateTokenRequest.Email)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	err = controller.usersService.CheckPassword(user.Password, generateTokenRequest.Password)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response, err := controller.service.GenerateJWT(*generateTokenRequest.ToGenerateTokenSpec())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponse(response))
}

func (controller *Controller) AuthError(c echo.Context) error {
	return c.JSON(common.NewForbiddenResponse())
}