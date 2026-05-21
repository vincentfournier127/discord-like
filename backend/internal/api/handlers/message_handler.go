package handlers

import "discord-like/internal/repository"

type MessageHandler struct {
	mr *repository.MessageRepository
}

func NewMessageHandler(mr *repository.MessageRepository) *MessageHandler {
	return &MessageHandler{
		mr: mr,
	}
}

func (mh *MessageHandler) SendMessage() {

}

func (mh *MessageHandler) FetchChatroomMessages() {

}
