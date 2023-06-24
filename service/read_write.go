package service

import (
	"chess/model"
	"github.com/gorilla/websocket"
)

func Read(cli *model.Client) {
	conn := cli.Conn
	defer func() {
		conn.Close()
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

		Hub.Ch <- msg
	}
}

func Write(cli *model.Client) {
	conn := cli.Conn
	defer func() {

	}()

	for {
		select {
		case msg := <-cli.Receive:
			err := conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				panic(err)
			}
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
