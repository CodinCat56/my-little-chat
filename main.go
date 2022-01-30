package main

import (
	"net/http" // For serving our index.html file

	"github.com/gin-gonic/gin"  // For creating the server
	"gopkg.in/olahol/melody.v1" // For sending and receiving messages (a websocket framework)
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.Run(":8080")
}
