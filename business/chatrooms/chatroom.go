package chatrooms

type Chatroom struct {
	Id   string
	Name string
	Desc string
	Code string
}

func NewChatroom(
	id string,
	name string,
	desc string,
	code string) Chatroom {

	return Chatroom{
		Id:   id,
		Name: name,
		Desc: desc,
		Code: code,
	}
}