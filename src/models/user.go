package models

type User struct {
	CommonModel
	Username string `json:"user_name"`
}
