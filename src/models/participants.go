package models

type Participants struct {
	CommonModel
	User   User
	Chat   Chat
	ChatID uint
	UserID uint
}
