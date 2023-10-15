package api

import (
	"go-hexagonal/api/v1/messages"
	"go-hexagonal/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

//RegisterPaths Register all V1 API with routing path
func RegisterPaths(e *echo.Echo, messagesController *messages.Controller, userController *user.Controller) {
	if messagesController == nil {
		panic("messages controller cannot be nil")
	}

	if userController == nil {
		panic("user controller cannot be nil")
	}

	//messages
	messagesV1 := e.Group("v1/messages")
	messagesV1.POST("", messagesController.InsertMessage)
	messagesV1.GET("/:chatroom", messagesController.GetMessagesByChatroom)

	//user
	userV1 := e.Group("v1/users")
	userV1.GET("/:id", userController.FindUserByID)
	userV1.GET("", userController.FindAllUserWithPagination)
	userV1.POST("", userController.InsertUser)
	userV1.PUT("/:id", userController.UpdateUser)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
