package message

import (
	messageModel "../models"
	"github.com/gin-gonic/gin"
)

func NewMessage(c *gin.Context) {

  var msg messageModel.Message

  err := c.ShouldBindJSON(&msg)

  if err != nil {
    c.JSON(400, gin.H{
      "err": "Bad Request."
    })
  }

  msg.CreateMessage()

	//Si el usario que recibe esta conectado enviar websocket, si no devuelve 200 status:true

}
