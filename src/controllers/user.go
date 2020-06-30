package controllers

import (
  "github.com/gin-gonic/gin"
  "../models"
  "../database"
)

type CreateUserInput struct {
  Username string `json:"user_name" binding:"required"`
}

func GetUsers(c *gin.Context) {

  db := database.GetDatabase()

  var users []models.User

  db.DB.Find(&users)

  c.JSON(200, gin.H{
    "users": users,
  })

}

func CreateUser(c *gin.Context) {
  db := database.GetDatabase()

  var userInput CreateUserInput

  err := c.ShouldBindJSON(&userInput)

  if err != nil {
    c.JSON(500, gin.H{
      "error": err,
      "clearMessage": "Check if data that you send is correct",
    })
  }

  user := models.User{Username: userInput.Username}

  db.DB.Create(&user)

  c.JSON(200, gin.H{
    "user": user,
  })

}
