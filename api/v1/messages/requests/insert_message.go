package requests

import (
	"chat-hex/business/commands"
	"chat-hex/business/messages"
)

type InsertMessageRequest struct {
	Content  string `json:"content"`
	Sender   string `json:"sender"`
	Chatroom string `json:"chatroom"`
}

func (req *InsertMessageRequest) ToInsertMessageSpec() *messages.InsertMessageSpec {
	var insertMessageSpec messages.InsertMessageSpec

	insertMessageSpec.Content = req.Content
	insertMessageSpec.Sender = req.Sender
	insertMessageSpec.Chatroom = req.Chatroom

	return &insertMessageSpec
}

func (req *InsertMessageRequest) ToCommandSpec() *commands.CommandSpec {
	var commandSpec commands.CommandSpec

	commandSpec.Content = req.Content
	commandSpec.Sender = req.Sender
	commandSpec.Chatroom = req.Chatroom

	return &commandSpec
}