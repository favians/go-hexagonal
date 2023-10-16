package requests

import "chat-hex/business/users"

type EnterChatroomRequest struct {
	Email    string `json:"email"`
	Chatroom string `json:"chatroom"`
}

func (req *EnterChatroomRequest) ToEnterChatroomSpec() *users.EnterChatroomSpec {
	var enterChatroomSpec users.EnterChatroomSpec

	enterChatroomSpec.Email = req.Email
	enterChatroomSpec.Chatroom = req.Chatroom

	return &enterChatroomSpec
}