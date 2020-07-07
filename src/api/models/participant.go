package models

import (
	"../../internals/database"
	"../../internals/model"
)

type Participant struct {
	model.CommonModel
	ChatID uint
	UserID uint
}

func (p *Participant) CreateChatParticipant() {

	db := database.GetDatabase()

	db.DB.Create(&p)
}
