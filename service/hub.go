package service

import (
	"chess/model"
	"fmt"
	"strings"
)

const (
	A = 1
	B = 2
)

func RunHub() {
	for {
		msg := <-HubListener.Srv
		dataStr := string(msg.Data)
		parts := strings.Split(dataStr, " ")
		switch parts[0] {
		case "FINDROOM":
			// 找出所有可以加入的房间号，推送给该用户
		case "ADDROOM":

		case "CREATEROOM":
			// 创建一个RoomListener管理该房间
			roomListener := &model.Room{
				Clients: make([]*model.Client, 0),
				Srv:     make(chan *model.Message, 20),
			}
			go RunRoom(roomListener)
		}
	}
}

// RunRoom 管理每个房间
func RunRoom(roomListener *model.Room) {
	for {
		msg := <-roomListener.Srv
		dataStr := string(msg.Data)
		switch dataStr {
		case "QUIT":
			// 消息通知
			info := msg.Cli.Username + "退出了房间"
			for _, value := range roomListener.Clients {
				if value != msg.Cli {
					value.Receive <- []byte(info)
				}
			}

		case "READY":
			// 消息通知
			info := msg.Cli.Username + "已准备"
			for _, value := range roomListener.Clients {
				if value != msg.Cli {
					value.Receive <- []byte(info)
				}
			}

			// 判断是否都准备了
			if roomListener.Clients[0].Status == 1 && roomListener.Clients[1].Status == 1 {
				// 创建一个GameSource进行对局
				gameClients := []*model.GameClient{
					// A方
					{
						Conn:     roomListener.Clients[0].Conn,
						Username: roomListener.Clients[0].Username,
						Receive:  make(chan []byte),
						Role:     A,
					},
					// B方
					{
						Conn:     roomListener.Clients[1].Conn,
						Username: roomListener.Clients[1].Username,
						Receive:  make(chan []byte),
						Role:     B,
					},
				}
				gameSource := &model.GameSource{
					Turn:    0,
					Data:    make(chan *model.Message, 20),
					Clients: gameClients,
				}
				go RunGame(gameSource)
			}
		}
	}
}

func RunGame(gameSource *model.GameSource) {
	// 初始化棋局资源
	board := InitChessBoard(gameSource.Clients)

	for {
		msg := <-gameSource.Data
		parts := strings.Split(string(msg.Data), " ")
		switch parts[0] {
		// SELECT XX: 选中一个棋子，获取可以到达的路径-------------------------------------------------
		case "SELECT":
			// 不是该用户的回合
			if msg.Cli != gameSource.Clients[gameSource.Turn] {
				// 消息通知
				info := "等待对手操作。。。"
				msg.Cli.Receive <- []byte(info)
				continue
			}
			// 取得待操作的棋子对象
			obj := GetChessPieceObj(parts[1], board)
			// 找出所有可以到达的路径
			paths := GetAllPaths(obj, board)
			gameSource.Selected = obj
			gameSource.ValidPaths = paths
			// 消息通知
			info := "your chooses:"
			for _, value := range paths {
				info = fmt.Sprintf(" [%d,%d] ", value[0], value[1])
			}
			msg.Cli.Receive <- []byte(info)
		// MOVE XX: 将上一步选中的棋子移动到某个位置----------------------------------------------------
		case "MOVE":
			// 不是该用户的回合
			if msg.Cli != gameSource.Clients[gameSource.Turn] {
				// 消息通知
				info := "等待对手操作。。。"
				msg.Cli.Receive <- []byte(info)
				continue
			}
			// 取得目标位置的棋子对象
			obj := GetChessPieceObj(parts[1], board)
			// 判断该操作是否有效
			if ok := CheckValid(obj, gameSource.ValidPaths); !ok {
				info := "invalid operation"
				msg.Cli.Receive <- []byte(info)
			}

		}
	}

}
