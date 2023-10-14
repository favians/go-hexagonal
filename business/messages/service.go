package messages

import (
	"go-hexagonal/business"
	primitiveIDGenerator "go-hexagonal/util/primiviteIDGenerator"
	"go-hexagonal/util/validator"
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