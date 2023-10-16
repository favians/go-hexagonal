package api

import (
	"chat-hex/api/v1/messages"
	"chat-hex/api/v1/users"

	echo "github.com/labstack/echo/v4"
)

//RegisterPaths Register all V1 API with routing path
func RegisterPaths(e *echo.Echo, usersController *users.Controller, messagesController *messages.Controller) {
	if usersController == nil {
		panic("users controller cannot be nil")
	}

	if messagesController == nil {
		panic("messages controller cannot be nil")
	}

	//users
	usersV1 := e.Group("v1/users")
	usersV1.PATCH("/enter-chatroom", usersController.EnterChatroom)
	usersV1.PATCH("/leave-chatroom", usersController.LeaveChatroom)

	//messages
	messagesV1 := e.Group("v1/messages")
	messagesV1.POST("", messagesController.InsertMessage)
	messagesV1.GET("/:chatroom", messagesController.GetMessagesByChatroom)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
