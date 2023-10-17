package users

type Service interface {
	EnterChatroom(enterChatroomSpec EnterChatroomSpec) error
	LeaveChatroom(leaveChatroomSpec LeaveChatroomSpec) error
	HashPassword(password string) (*string, error)
	FindUserByEmail(email string) (*User, error)
	CheckPassword(hashedPassword string, clearPassword string) error
}

type Repository interface {
	FindUserByEmail(email string) (*User, error)
	EnterChatroom(email string, chatroom string) error
	LeaveChatroom(email string) error
}