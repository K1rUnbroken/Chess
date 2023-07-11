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
			break
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
		conn.Close()
	}()

	for {
		select {
		case msg := <-cli.Receive:
			err := conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				break
			}
		}
	}
}
