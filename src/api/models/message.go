package models

import (
	"../../internals/database"
	"../../internals/model"
)

type Message struct {
	model.CommonModel
	ChatID  uint   `json:"chat_id"`
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
	ReceiverID uint `json:"receiver_id"`
}

func (m *Message) CreateMessage() {
	db := database.GetDatabase()

	db.DB.Create(&m)
}
