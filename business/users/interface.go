package users

type Service interface {
	LeaveChatroom(leaveChatroomSpec LeaveChatroomSpec) error
}

type Repository interface {
	LeaveChatroom(email string) error
}