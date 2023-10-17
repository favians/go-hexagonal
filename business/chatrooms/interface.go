package chatrooms

type Service interface {
	FindChatroomByCode(code string) (*Chatroom, error)
	GetChatrooms() ([]Chatroom, error)
}

type Repository interface {
	FindChatroomByCode(code string) (*Chatroom, error)
	GetChatrooms() ([]Chatroom, error)
}