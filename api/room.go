package api

import (
	"chess/model"
	"chess/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Connect 连接大厅
func Connect(c *gin.Context) {
	// 协议升级为websocket
	upgrader := service.GenUpgrader()
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}

	// 创建客户端对象
	username, ok := c.Get("username")
	if !ok {
		panic("username not set")
	}
	cli := model.Client{
		Conn:     conn,
		Username: fmt.Sprintf("%s", username),
	}

	// 开启读写
	go service.Read(&cli)
	go service.Write(&cli)
	// 监听用户命令
	go service.Listener(&cli)
}
