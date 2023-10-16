package users

type Service interface {
	EnterChatroom(enterChatroomSpec EnterChatroomSpec) error
	LeaveChatroom(leaveChatroomSpec LeaveChatroomSpec) error
}

type Repository interface {
	EnterChatroom(email string, chatroom string) error
	LeaveChatroom(email string) error
}