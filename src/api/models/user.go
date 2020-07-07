package models

import (
	"../../internals/database"
	"../../internals/model"
)

type User struct {
	model.CommonModel
	Username string `json:"user_name"`
	Chats    []Chat `json:"chats,omitempty"`
}

func GetAllUsers(authId string) []User {

	var users []User

	db := database.GetDatabase()

	db.DB.Raw("select * from users where id != ?", authId).Scan(&users)

	return users

}

func CreateUser(username string) User {

	user := User{
		Username: username,
	}

	db := database.GetDatabase()

	db.DB.Create(&user)

	return user
}

func FindUser(id uint) User {
	user := User{}

	db := database.GetDatabase()

	db.DB.Find(&user, id)

	return user
}

func (u *User) FindByUsername() {
	db := database.GetDatabase()

	db.DB.Where("username=?", u.Username).Find(&u)
}

func (u *User) GetChats() {
	db := database.GetDatabase()

	db.DB.Model(&u).Related(&u.Chats)

	chats := &u.Chats

	for _, chat := range *chats {
		chat.GetChatUsers()
	}
}
