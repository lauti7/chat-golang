package chat

import (
	chatModel "../models"
	"github.com/gin-gonic/gin"
)

type CreateChatInput struct {
	Type         string              `json:"type" binding:"required"`
	Participants [2]ParticipantInput `json:"participants"`
}

type ParticipantInput struct {
	UserID uint `json:"user_id"`
}

type GetChatInput struct {
	Participants [2]ParticipantInput `json:"participants" binding:"required"`
}

func CreateChat(c *gin.Context) {
	var chatInput CreateChatInput

	err := c.ShouldBindJSON(&chatInput)

	if err != nil {
		c.JSON(400, gin.H{
			"clearMessage": "your request is missing something",
		})
	}

	var chat chatModel.Chat

	//Check if exists, if not, create it
	chat.Between(chatInput.Participants[0].UserID, chatInput.Participants[1].UserID)
	c.JSON(200, gin.H{
		"chat": chat,
	})
}
