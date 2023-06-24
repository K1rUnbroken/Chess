package api

import (
	"chess/cons"
	"chess/model"
	"chess/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Connect 连接大厅
func Connect(c *gin.Context) {
	fmt.Println("here")
	// 协议升级为websocket
	upgrader := service.GenUpgrader()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	// 创建客户端对象
	username := c.Query("username")
	if username == "" {
		panic("username not set")
	}
	cli := model.Client{
		Conn:     conn,
		Username: fmt.Sprintf("%s", username),
		Receive:  make(chan []byte, 1024),
		Status:   cons.InHall,
	}

	// 开启读写
	go service.Read(&cli)
	go service.Write(&cli)
}
