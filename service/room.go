package service

import (
	"chess/model"
	"github.com/gorilla/websocket"
	"sync"
)

var (
	Lock   sync.Mutex
	RoomId = 0
)

const (
	King   = 1
	Queen  = 2
	Bishop = 3
	Rook   = 4
	Knight = 5
	Pawn   = 6
)

// Read 读出用户的命令，交给Hub处理
func Read(cli *model.Client) {
	conn := cli.Conn
	defer func() {

	}()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}

		msg := &model.Message{
			Data: data,
			Cli:  cli,
		}
		switch cli.Status {
		// 大厅
		case 0:
			HubListener.Srv <- msg
		// 房间
		case 1:
		// 游戏中
		case 2:

		}

	}
}

func Write(cli *model.Client) {
	conn := cli.Conn
	defer func() {

	}()

	for {
		msg := <-cli.Receive
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			panic(err)
		}
	}
}

// Listener 监听用户的基础命令
/*
	READY  			准备
 	FINDROOM		查找所有可以加入的房间
 	QUIT  			退出大厅（断开连接）
 	ADDROOM XX		加入房间XX
 	CREATROOM		创建房间
*/
