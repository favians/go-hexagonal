package users

import (
	"chat-hex/business"
	"chat-hex/business/chatrooms"
	"chat-hex/util/validator"
)

type EnterChatroomSpec struct {
	Email     string `validate:"required"`
	Chatroom string `validate:"required"`
}

type LeaveChatroomSpec struct {
	Email     string `validate:"required"`
	Chatroom string `validate:"required"`
}

type service struct {
	repository Repository
	chatroomsService	chatrooms.Service
}

func NewService(repository Repository, chatroomsService chatrooms.Service) Service {
	return &service{repository, chatroomsService}
}

func (s *service) EnterChatroom(enterChatroomSpec EnterChatroomSpec) error {
	err := validator.GetValidator().Struct(enterChatroomSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	_, err = s.chatroomsService.FindChatroomByCode(enterChatroomSpec.Chatroom)
	if err != nil {
		return business.ErrInvalidSpec
	}

	err = s.repository.EnterChatroom(enterChatroomSpec.Email, enterChatroomSpec.Chatroom)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) LeaveChatroom(leaveChatroomSpec LeaveChatroomSpec) error {
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