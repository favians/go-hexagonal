package chatrooms

import (
	"chat-hex/api/common"
	"chat-hex/api/v1/chatrooms/responses"
	"chat-hex/business/chatrooms"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service chatrooms.Service
}

func NewController(service chatrooms.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) GetChatrooms(c echo.Context) error {
	chatrooms, err := controller.service.GetChatrooms()
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := responses.NewGetChatroomsResponse(chatrooms)

	return c.JSON(common.NewSuccessResponse(response))
}