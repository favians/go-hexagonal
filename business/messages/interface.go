package messages

type Service interface {
	InsertMessage(insertMessageSpec InsertMessageSpec) error
	GetMessagesByChatroom(chatroom string) ([]Message, error)
	MessageHasCommandStructure(insertMessageSpec InsertMessageSpec) bool
}

type Repository interface {
	InsertMessage(message Message) error
	GetMessagesByChatroom(chatroom string) ([]Message, error)
}