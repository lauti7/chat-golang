package main

import (
	"github.com/gin-gonic/gin"
	chatController "./api/chat"
	// messageController "../../api/message"
	// participantController "../../api/participant"
	userController "./api/user"
	"./internals/database"
)

func main() {

	_ = database.GetDatabase()

	usersManager := UsersManager{
		OnlineUsers:    make(map[uint]OnlineUser),
		RegisterChannel:   make(chan OnlineUser),
		UnregisterChannel: make(chan OnlineUser),
	}

	//Channels start waiting for receiving
	go usersManager.registration()

	server := gin.Default()
	server.Use(CORSMiddleware())

	api := server.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "API IS ALIVE",
			})
		})

		api.GET("/users", userController.GetUsers)
		api.POST("/users", userController.CreateUser)
		api.POST("/users/login", userController.Login)
		api.POST("/chat", chatController.CreateChat)
		api.GET("/ws", checkAndFindUser() ,func(c *gin.Context) {
			user, _ := c.Keys["user"]
			usersManager.handleWS(c.Writer, c.Request, user)
		})
	}


	server.Run()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
