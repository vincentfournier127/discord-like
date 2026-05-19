package model

import "time"

type Message struct {
	ID          int       `json:"id"`
	Chatroom_ID int       `json:"chatroom_id"`
	Sender_ID   int       `json:"sender_id"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
}
