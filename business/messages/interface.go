package messages

type Service interface {
	InsertMessage(insertMessageSpec InsertMessageSpec) error
}

type Repository interface {
	InsertMessage(message Message) error
}