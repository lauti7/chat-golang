package models

type Message struct {
	CommonModel
	Chat    Chat
	ChatID  uint
	User    User
	UserID  uint
	Content string
}
