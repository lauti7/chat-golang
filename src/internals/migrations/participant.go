package migrations

import (
	"../model"
)

type Participant struct {
	model.CommonModel
	ChatID uint
	UserID uint
}
