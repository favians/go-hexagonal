package users

import (
	"chat-hex/business"
	"chat-hex/business/chatrooms"
	"chat-hex/util/validator"

	"golang.org/x/crypto/bcrypt"
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

func (s *service) HashPassword(password string) (*string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}

	preResult := string(bytes)
	result := &preResult
	return result, nil
}

func (s *service) FindUserByEmail(email string) (*User, error){
	user, err := s.repository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) CheckPassword(passwordToCheck string, clearPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordToCheck), []byte(clearPassword))
	if err != nil {
		return err
	}
	
	return nil
}