package users

import (
	"chat-hex/api/common"
	"chat-hex/api/v1/users/requests"
	"chat-hex/business/users"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service users.Service
}

func NewController(service users.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) EnterChatroom(c echo.Context) error {
	enterChatroomRequest := new(requests.EnterChatroomRequest)

	err := c.Bind(enterChatroomRequest)
	if err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err = controller.service.EnterChatroom(*enterChatroomRequest.ToEnterChatroomSpec())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return nil
}

func (controller *Controller) LeaveChatroom(c echo.Context) error {
	leaveChatroomRequest := new(requests.LeaveChatroomRequest)

	err := c.Bind(leaveChatroomRequest)
	if err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err = controller.service.LeaveChatroom(*leaveChatroomRequest.ToLeaveChatroomSpec())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return nil
}
