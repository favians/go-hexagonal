package users

import (
	"chat-hex/business"
	"chat-hex/util/validator"
	"fmt"
)

type LeaveChatroomSpec struct {
	Email     string `validate:"required"`
	Chatroom string `validate:"required"`
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) LeaveChatroom(leaveChatroomSpec LeaveChatroomSpec) error {
	fmt.Println("HOLA HOLA", leaveChatroomSpec.Email, leaveChatroomSpec.Chatroom)
	err := validator.GetValidator().Struct(leaveChatroomSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	err = s.repository.LeaveChatroom(leaveChatroomSpec.Email)
	if err != nil {
		return err
	}

	return nil
}