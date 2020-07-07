package migrations

import (
	"../model"
)

type User struct {
	model.CommonModel
	Username string `json:"user_name"`
}
