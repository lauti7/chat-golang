package migrations

import (
	"../model"
)

type Chat struct {
	model.CommonModel
	Type string `json:"type"`
}
