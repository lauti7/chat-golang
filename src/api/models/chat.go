package models

import (
	"../../internals/database"
	"../../internals/model"
)

type Chat struct {
	model.CommonModel
	Type     string    `json:"type"`
	Messages []Message `json:"messages,omitempty" gorm:"foreignkey:chat_id"`
	Users    []User    `json:"users,omitempty"`
}

func (c *Chat) CreateChat() {

	db := database.GetDatabase()

	db.DB.Create(&c)
}

//Check if exists a chat between users
func (c *Chat) Between(auth uint, other uint) {

	db := database.GetDatabase()

	//Get ID from existed chat
	db.DB.Raw("select c.* from chats as c where exists (select * from users as u inner join participants as p on p.user_id = u.id where c.id = p.chat_id and p.user_id = ?) and exists (select * from users as u inner join participants as p on p.user_id = u.id where c.id = p.chat_id and p.user_id = ?) limit 1", auth, other).Scan(&c)

	if c.ID > 0 {
		c.GetChatParticipant(auth)
		c.GetChattMessages()
	} else {
		c.CreateChat()

		createdParticipants := [2]Participant{}
		createdParticipants[0] = Participant{UserID: auth, ChatID: c.ID}
		createdParticipants[1] = Participant{UserID: other, ChatID: c.ID}

		for _, p := range createdParticipants {
			p.CreateChatParticipant()
		}

		c.GetChatParticipant(auth)

	}

}

func FindChat(id uint) Chat {

	var chat Chat

	db := database.GetDatabase()

	db.DB.Where("id=?", id).Preload("Messages").Find(&chat)

	return chat
}

func (c *Chat) GetChatUsers() {

	db := database.GetDatabase()

	db.DB.Raw("select u.* from participants as p inner join users as u on p.user_id = u.id where p.chat_id = ?", c.ID).Scan(&c.Users)

}

func (c *Chat) GetChatParticipant(auth uint) {
	db := database.GetDatabase()

	db.DB.Raw("select u.* from participants as p inner join users as u on p.user_id = u.id where p.chat_id = ? and p.user_id != ?", c.ID, auth).Scan(&c.Users)
}

func (c *Chat) GetChattMessages() {
	db := database.GetDatabase()

	db.DB.Model(&c).Related(&c.Messages)
}
