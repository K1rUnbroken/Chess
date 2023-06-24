package service

import (
	"chess/cons"
	"chess/model"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	Lock = sync.Mutex{}
)

// RunHub 管理全局所有的房间，处理用户的命令
func RunHub() {
	for {
		select {
		case msg := <-Hub.Ch:
			dataStr := string(msg.Data)
			parts := strings.Split(dataStr, " ")
			cli := msg.Cli

			switch cli.Status {
			case cons.InHall:
				switch parts[0] {
				case cons.QUIT:
					cli.Conn.Close()
				case cons.FINDROOM:
					info := "roomID:"
					for k, _ := range Hub.Rooms {
						info += fmt.Sprintf("%d\n", k)
					}
					cli.Receive <- []byte(info)
				case cons.ADDROOM:
					if len(parts) != 2 {
						// error: 参数错误
					}
					roomId, _ := strconv.Atoi(parts[1])
					room, ok := Hub.Rooms[model.RoomId(roomId)]
					if !ok {
						// error: 房间不存在
					}
					Lock.Lock()
					room.Clients = append(room.Clients, cli)
					Lock.Unlock()

					cli.Status = cons.InRoom
					cli.RoomId = model.RoomId(roomId)

					// 消息通知自己
					info := "您加入了房间" + fmt.Sprintf("%d", roomId)
					cli.Receive <- []byte(info)
					//// 消息通知房间内的其他人
					//info = cli.Username + "加入了房间"
					//for _, value := range Hub.Rooms[cli.RoomId].Clients {
					//	if value != cli {
					//
					//	}
					//}
				case cons.CREATEROOM:
					room := &model.Room{
						Id:      model.RoomId(time.Now().Unix()),
						Clients: make([]*model.Client, 0),
						Ch:      make(chan *model.Message, 20),
					}
					Lock.Lock()
					Hub.Rooms[room.Id] = room
					room.Clients = append(room.Clients, cli)
					Lock.Unlock()

					cli.Status = cons.InRoom
					cli.RoomId = room.Id

					info := "您创建了房间" + fmt.Sprintf("%d", room.Id)
					cli.Receive <- []byte(info)
				default:
					// error:命令不存在->Usage()
				}
			case cons.InRoom:
				room := Hub.Rooms[cli.RoomId]
				switch dataStr {
				case cons.QUIT:
					clients := room.Clients
					cliNum := len(clients)
					// 在Room.Clients中删除该client
					Lock.Lock()
					if clients[0] == cli {
						room.Clients = append(room.Clients[cliNum:])
					} else {
						room.Clients = append(room.Clients[:cliNum])
					}
					Lock.Unlock()
					// 如果房间为空，在Hub.Rooms中删除该room
					Lock.Lock()
					if cliNum == 0 {
						delete(Hub.Rooms, cli.RoomId)
					}
					Lock.Unlock()
					// 消息通知自己
					info := "您退出了房间" + fmt.Sprintf("%d", cli.RoomId)
					cli.Receive <- []byte(info)
					// 消息通知房间中的其他人
					if cliNum > 0 {
						info = cli.Username + "退出了房间"
						for _, value := range room.Clients {
							if value != cli {
								value.Receive <- []byte(info)
							}
						}
					}

					cli.Conn.Close()
				case cons.READY:
					// 消息通知自己
					info := "您已准备"
					cli.Receive <- []byte(info)
					// 消息通知房间中的其他人
					info = cli.Username + "已准备"
					for _, value := range room.Clients {
						if value != cli {
							value.Receive <- []byte(info)
						}
					}
					cli.Status = cons.BeReady

					// 房间内的玩家都准备后就开始游戏
					if len(room.Clients) == 2 && room.Clients[0].Status == cons.BeReady && room.Clients[1].Status == cons.BeReady {
						game := InitGame(cli.RoomId)
						GameHub.Games[game.Id] = game

						info = "游戏开始"
						for _, value := range room.Clients {
							value.Receive <- []byte(info)
						}

						for _, value := range game.Clients {
							info = "test"
							value.Receive <- []byte(info)
							info = PrintChessBoard(value.Role, game.Board)
							fmt.Println(info)
							value.Receive <- []byte(info)
						}
					}
				default:
					fmt.Println("error")
					// error:命令不存在->Usage()
				}
			case cons.InGame:

			}
		}
	}
}

// RunGameHub 管理全局所有的游戏，处理用户在游戏中的操作
func RunGameHub() {
	for {
		select {
		case msg := <-GameHub.Ch:
			parts := strings.Split(string(msg.Data), " ")
			cli := msg.Cli
			game := GameHub.Games[cli.GameId]
			switch parts[0] {
			// ------------------------------SELECT <ROW> <COLUMN> :选中一个棋子，获取可以到达的路径---------------------------------
			case cons.SELECT:
				// 参数检查
				if len(parts) != 2 {
					info := "命令错误"
					cli.Receive <- []byte(info)
					continue
				}
				// 不是该用户的回合
				if cli != game.Clients[game.Turn] {
					info := "请等待对手操作"
					cli.Receive <- []byte(info)
					continue
				}
				// 取得待操作的棋子对象
				obj := GetChessPieceObj(parts[1], parts[2], game.Board)
				// 找出所有可以到达的路径
				paths := GetAllPaths(obj, GameHub.Games[cli.GameId].Board)
				GameHub.Games[cli.GameId].Selected = obj
				GameHub.Games[cli.GameId].ValidPaths = paths
				// 消息通知
				info := "可以选择的路径：\n"
				for _, value := range paths {
					info += fmt.Sprintf(" [%d,%d] ", value[0], value[1])
				}
				cli.Receive <- []byte(info)
			// -------------------------------MOVE <ROW> <COLUMN>: 将选中的棋子移动到指定位置------------------------------
			case cons.MOVE:
				// 棋子未选中
				if GameHub.Games[cli.GameId].Selected == nil {
					info := "请先使用SELECT选中棋子"
					cli.Receive <- []byte(info)
					continue
				}
				GameHub.Games[cli.GameId].Selected = nil
				// 取得目标位置的棋子对象
				obj := GetChessPieceObj(parts[1], parts[2], game.Board)
				// 判断目标位置是否可以到达
				if ok := CheckLocationValid(obj, game.ValidPaths); !ok {
					info := "无法移动到指定位置"
					cli.Receive <- []byte(info)
				}
				// 更新棋盘
				UpdateChessBoard(game.Selected, obj, game.Board)
				// 将新的棋盘发送给玩家
				info := PrintChessBoard(cli.Role, game.Board)
				cli.Receive <- []byte(info)
			}
		}
	}
}
