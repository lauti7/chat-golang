package main

import (
	chatController "./api/chat"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	// messageController "../../api/message"
	// participantController "../../api/participant"
	userController "./api/user"
	"./internals/database"
)

var upgrader = websocket.Upgrader{}

func main() {

	_ = database.GetDatabase()

	// clientManager := ClientManager{
	// 	Clients:    make(map[Client]bool),
	// 	Broadcast:  make(chan Message),
	// 	Typing:     make(chan Message),
	// 	Register:   make(chan Client),
	// 	Unregister: make(chan Client),
	// }

	//Channels start waiting for receiving
	// go clientManager.run()

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
	}

	// server.GET("/ws", clientManager.handleConnections)

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
