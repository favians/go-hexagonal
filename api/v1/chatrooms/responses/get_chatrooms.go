package responses

import "chat-hex/business/chatrooms"

type chatroomResponse struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Code string `json:"code"`
}

type getChatroomsResponse struct {
	Chatrooms []chatroomResponse `json:"chatrooms"`
}

func NewGetChatroomsResponse(chatrooms []chatrooms.Chatroom) getChatroomsResponse {
	getChatroomsResponse := getChatroomsResponse{}

	for _, value := range chatrooms {
		var chatroomResponse chatroomResponse
		
		chatroomResponse.Name = value.Name
		chatroomResponse.Desc = value.Desc
		chatroomResponse.Code = value.Code

		getChatroomsResponse.Chatrooms = append(getChatroomsResponse.Chatrooms, chatroomResponse)
	}

	if getChatroomsResponse.Chatrooms == nil {
		getChatroomsResponse.Chatrooms = []chatroomResponse{}
	}

	return getChatroomsResponse
}