package messages

import (
	"chat-hex/business"
	primitiveIDGenerator "chat-hex/util/primiviteIDGenerator"
	"chat-hex/util/validator"
	"time"
)

type InsertMessageSpec struct {
	Content  string `validate:"required"`
	Sender   string `validate:"required"`
	Chatroom string `validate:"required"`
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) InsertMessage(insertMessageSpec InsertMessageSpec) error {
	err := validator.GetValidator().Struct(insertMessageSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	id := primitiveIDGenerator.GenerateID()
	message := NewMessage(
		id,
		insertMessageSpec.Content,
		insertMessageSpec.Sender,
		time.Now(),
		insertMessageSpec.Chatroom,
	)

	err = s.repository.InsertMessage(message)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetMessagesByChatroom(chatroom string) ([]Message, error) {
	messages, err := s.repository.GetMessagesByChatroom(chatroom)
	if err != nil {
		return []Message{}, err
	}

	return messages, err
}

func (s *service) MessageHasCommandStructure(insertMessageSpec InsertMessageSpec) bool {
	err := validator.GetValidator().Struct(insertMessageSpec)
	if err != nil {
		return false
	}

	firstCharacter:= insertMessageSpec.Content[0:1]
	return firstCharacter == "/"
}