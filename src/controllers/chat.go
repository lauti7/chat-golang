package controllers

import (
  "github.com/gin-gonic/gin"
  "../models"
  "../database"
)

type CreateChatInput struct {
  Type string `json:"type" binding:"required"`
}

func CreateChat(c *gin.Context) {

  

}
