package responses

import (
	"go-hexagonal/business/messages"
	"time"
)

type getMessageResponse struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	Sender    string `json:"sender"`
	Timestamp time.Time `json:"timestamp"`
}

type getMessagesByChatroomResponse struct {
	Messages []getMessageResponse `json:"messages"`
}

func NewGetMessagesByChatroomResponse(messages []messages.Message) getMessagesByChatroomResponse {
	getMessagesByChatroomResponse := getMessagesByChatroomResponse{}

	for _, value := range messages {
		var getMessageResponse getMessageResponse

		getMessageResponse.Id = value.Id
		getMessageResponse.Content = value.Content
		getMessageResponse.Sender = value.Sender
		getMessageResponse.Timestamp = value.Timestamp

		getMessagesByChatroomResponse.Messages = append(getMessagesByChatroomResponse.Messages, getMessageResponse)
	}

	if getMessagesByChatroomResponse.Messages == nil {
		getMessagesByChatroomResponse.Messages = []getMessageResponse{}
	}

	return getMessagesByChatroomResponse
}