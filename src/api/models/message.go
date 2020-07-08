package models

import (
	"../../internals/model"
	"../../internals/database"
)

type Message struct {
	model.CommonModel
	ChatID  uint   `json:"chat_id"`
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}

func (m *Message) CreateMessage() {
	db := database.GetDatabase()

	db.DB.Create(&m)
}
