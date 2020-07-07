package database

import (
	"../migrations"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

type MySql struct {
	DB *gorm.DB
}

var chatDatabase = MySql{}
var once sync.Once

func GetDatabase() MySql {
	once.Do(func() {
		fmt.Println("Callind Do Once...")
		chatDatabase.ConnectDatabase()
	})

	return chatDatabase
}

func (c *MySql) ConnectDatabase() {
	db, err := gorm.Open("mysql", "root@(127.0.0.1:3306)/chatgolang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to DB")
	}

	c.DB = db

	c.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&migrations.Chat{}, &migrations.Message{}, &migrations.User{}, &migrations.Participant{})
	c.DB.Model(&migrations.Message{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	c.DB.Model(&migrations.Message{}).AddForeignKey("chat_id", "chats(id)", "CASCADE", "CASCADE")
	c.DB.Model(&migrations.Participant{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	c.DB.Model(&migrations.Participant{}).AddForeignKey("chat_id", "chats(id)", "CASCADE", "CASCADE")

}
