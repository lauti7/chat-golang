package user

import (
	userModel "../models"
	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Username string `json:"user_name" binding:"required"`
}

type GetUserInput struct {
	ID uint `json:id`
}

func GetUsers(c *gin.Context) {

	authId := c.Request.Header["Authorization"][0]

	users := userModel.GetAllUsers(authId)

	c.JSON(200, gin.H{
		"users": users,
	})

}

func CreateUser(c *gin.Context) {

	var userInput UserInput

	err := c.ShouldBindJSON(&userInput)

	if err != nil {
		c.JSON(400, gin.H{
			"error":        err,
			"cleanMessage": "Check if data that you send is correct",
		})
	}

	newUser := userModel.CreateUser(userInput.Username)

	c.JSON(200, gin.H{
		"user": newUser,
	})
}

func Login(c *gin.Context) {
	var user userModel.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error":        err,
			"cleanMessage": "Check if data that you send is correct",
		})
	}

	user.FindByUsername()

	c.JSON(200, gin.H{
		"user": user,
	})

}

func GetUserChats(c *gin.Context) {
	var user userModel.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error":        err,
			"cleanMessage": "Check if data that you send is correct",
		})
	}

	user.GetChats()

	c.JSON(200, gin.H{
		"chats": user.Chats,
	})

}
