package messages

import (
	"chat-hex/api/common"
	"chat-hex/api/v1/messages/requests"
	"chat-hex/api/v1/messages/responses"
	"chat-hex/business/commands"
	"chat-hex/business/messages"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service messages.Service
	commandsService commands.Service
}

func NewController(service messages.Service, commandsService commands.Service) *Controller {
	return &Controller{
		service,
		commandsService,
	}
}

func (controller *Controller) InsertMessage(c echo.Context) error {
	insertMessageRequest := new(requests.InsertMessageRequest)

	err := c.Bind(insertMessageRequest)
	if err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	messageHasCommandStructure := controller.service.MessageHasCommandStructure(*insertMessageRequest.ToInsertMessageSpec())
	if messageHasCommandStructure {
		err = controller.commandsService.ProcessCommand(*insertMessageRequest.ToCommandSpec())
		if err != nil {
			return c.JSON(common.NewErrorBusinessResponse(err))
		}
	} else {
		err = controller.service.InsertMessage(*insertMessageRequest.ToInsertMessageSpec())
		if err != nil {
			return c.JSON(common.NewErrorBusinessResponse(err))
		}
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