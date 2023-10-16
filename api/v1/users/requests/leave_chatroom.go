package requests

import "chat-hex/business/users"

type LeaveChatroomRequest struct {
	Email     string `json:"email"`
	Chatroom string `json:"chatroom"`
}

func (req *LeaveChatroomRequest) ToLeaveChatroomSpec() *users.LeaveChatroomSpec {
	var leaveChatroomSpec users.LeaveChatroomSpec

	leaveChatroomSpec.Email = req.Email
	leaveChatroomSpec.Chatroom = req.Chatroom

	return &leaveChatroomSpec
}