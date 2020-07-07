package models

import (
	"../../internals/model"
)

type Message struct {
	model.CommonModel
	ChatID  uint   `json:"chat_id"`
	UserID  uint   `json:"user_id"`
	Content string `json:"content"`
}
