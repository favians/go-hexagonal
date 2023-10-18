package messages

import "time"

type Message struct {
	Id	string
	Content	string
	Sender	string
	Timestamp	time.Time
	ChatRoom	string
}

func NewMessage(
	id string,
	content string,
	sender string,
	timestamp time.Time,
	chatroom string) Message {
	
	return Message{
		Id:	id,
		Content: content,
		Sender: sender,
		Timestamp: timestamp,
		ChatRoom: chatroom,
	}
}