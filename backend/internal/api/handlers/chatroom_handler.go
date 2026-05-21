package handlers

import (
	"discord-like/internal/repository"
)

type ChatroomHandler struct {
	cr *repository.ChatroomRepository
}

func NewChatroomHandler(cr *repository.ChatroomRepository) *ChatroomHandler {
	return &ChatroomHandler{
		cr: cr,
	}
}

func (ch *ChatroomHandler) CreateChatroom() {

}

func (ch *ChatroomHandler) GetUserChatrooms() {

}

func (ch *ChatroomHandler) AddUserToChatroom() {

}
