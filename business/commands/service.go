package commands

import (
	"chat-hex/business"
	"fmt"
	"strings"
)

type CommandSpec struct {
	Content  string `validate:"required"`
	Sender   string `validate:"required"`
	Chatroom string `validate:"required"`
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) ProcessCommand(commandSpec CommandSpec) error {
	pieces := strings.Split(commandSpec.Content, "=")

	possibleCommand := strings.ToLower(pieces[0])
	if possibleCommand == CommandStock {
		stockCode := strings.ToLower(pieces[1])
		if len(stockCode) <= 0 {
			return business.ErrInvalidCommand
		}

		err := s.AsyncStockCommand(stockCode, commandSpec.Chatroom)
		if err != nil {
			return business.ErrInvalidCommand
		}

		return nil
	}

	return business.ErrInvalidCommand
}

func (s *service) AsyncStockCommand(stockCode string, chatroom string) error {
	fmt.Println("STOCK CODE IS", stockCode, "CHATROOM IS", chatroom)
	return nil
}