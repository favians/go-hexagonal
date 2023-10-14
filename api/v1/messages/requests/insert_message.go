package requests

import (
	"go-hexagonal/business/messages"
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