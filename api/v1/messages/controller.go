package messages

import (
	"go-hexagonal/api/common"
	"go-hexagonal/api/v1/messages/requests"
	"go-hexagonal/api/v1/messages/responses"
	"go-hexagonal/business/messages"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service messages.Service
}

func NewController(service messages.Service) *Controller {
	return &Controller{
		service,
	}
}

func (controller *Controller) InsertMessage(c echo.Context) error {
	insertMessageRequest := new(requests.InsertMessageRequest)

	err := c.Bind(insertMessageRequest)
	if err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err = controller.service.InsertMessage(*insertMessageRequest.ToInsertMessageSpec())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func (controller *Controller) GetMessagesByChatroom(c echo.Context) error {
	chatroom := c.Param("chatroom")

	messages, err := controller.service.GetMessagesByChatroom(chatroom)
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := responses.NewGetMessagesByChatroomResponse(messages)

	return c.JSON(common.NewSuccessResponse(response))
}