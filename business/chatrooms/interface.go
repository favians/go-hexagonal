package chatrooms

type Service interface {
	FindChatroomByCode(code string) (*Chatroom, error)
}

type Repository interface {
	FindChatroomByCode(code string) (*Chatroom, error)
}