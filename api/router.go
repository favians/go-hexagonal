package api

import (
	"chat-hex/api/v1/auth"
	"chat-hex/api/v1/chatrooms"
	"chat-hex/api/v1/messages"
	"chat-hex/api/v1/users"
	"chat-hex/middlewares"

	echo "github.com/labstack/echo/v4"
)

//RegisterPaths Register all V1 API with routing path
func RegisterPaths(e *echo.Echo, authController *auth.Controller, usersController *users.Controller, chatroomsController *chatrooms.Controller, messagesController *messages.Controller) {
	if authController == nil {
		panic("auth controller cannot be nil")
	}

	if usersController == nil {
		panic("users controller cannot be nil")
	}

	if messagesController == nil {
		panic("messages controller cannot be nil")
	}

	//auth
	authV1 := e.Group("v1/auth")
	authV1.POST("/token", authController.GenerateToken)

	//users
	usersV1 := e.Group("v1/users")
	usersV1.PATCH("/enter-chatroom", middlewares.Auth(usersController.EnterChatroom, authController.AuthError))
	usersV1.PATCH("/leave-chatroom", middlewares.Auth(usersController.LeaveChatroom, authController.AuthError))

	//chatrooms
	chatroomsV1 := e.Group("v1/chatrooms")
	chatroomsV1.GET("", middlewares.Auth(chatroomsController.GetChatrooms, authController.AuthError))

	//messages
	messagesV1 := e.Group("v1/messages")
	messagesV1.POST("", middlewares.Auth(messagesController.InsertMessage, authController.AuthError))
	messagesV1.GET("/:chatroom", middlewares.Auth(messagesController.GetMessagesByChatroom, authController.AuthError))

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
