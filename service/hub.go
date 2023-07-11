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

// RunHub 管理全局，处理用户的命令
func RunHub() {
	for {
		select {
		case msg := <-Hub.Ch:
			parts := strings.Split(string(msg.Data), " ")
			cli := msg.Cli

			switch cli.Status {
			//-----------------------------------------Status: InHall---------------------------------------------------
			case cons.InHall:
				switch parts[0] {
				case cons.QUIT:
					// cmd语法检查
					if info, ok := CheckCmdFormat(cons.QUIT, parts); !ok {
						cli.Receive <- []byte(info)
						continue
					}

					cli.Conn.Close()
				case cons.FINDROOM:
					// cmd语法检查
					if info, ok := CheckCmdFormat(cons.FINDROOM, parts); !ok {
						cli.Receive <- []byte(info)
						continue
					}

					// 消息通知
					info := "roomID:"
					for k := range Hub.Rooms {
						info += fmt.Sprintf("%d\n", k)
					}
					MsgToClient(info, "", cli)
				case cons.ADDROOM:
					// cmd语法检查
					if info, ok := CheckCmdFormat(cons.ADDROOM, parts); !ok {
						cli.Receive <- []byte(info)
						continue
					}

					// 房间是否存在
					roomId, _ := strconv.Atoi(parts[1])
					room, ok := Hub.Rooms[model.RoomId(roomId)]
					if !ok {
						MsgToClient(RoomNotExistsErr(), "", cli)
						continue
					}

					Lock.Lock()
					room.Clients = append(room.Clients, cli)
					Lock.Unlock()

					cli.Status = cons.InRoom
					cli.RoomId = model.RoomId(roomId)

					// 消息通知
					info1 := "您加入了房间" + fmt.Sprintf("%d", roomId)
					info2 := cli.Username + "加入了房间"
					MsgToClient(info1, info2, cli)
				case cons.CREATEROOM:
					// cmd语法检查
					if info, ok := CheckCmdFormat(cons.CREATEROOM, parts); !ok {
						cli.Receive <- []byte(info)
						continue
					}

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

					// 消息通知
					info := "您创建了房间" + fmt.Sprintf("%d", room.Id)
					MsgToClient(info, "", cli)
				default:
					MsgToClient(HallCmdNotExistsErr(), "", cli)
				}
			//-------------------------------------------Status: InRoom-------------------------------------------------
			case cons.InRoom:
				room := Hub.Rooms[cli.RoomId]
				switch parts[0] {
				case cons.QUIT:
					// cmd语法检查
					if info, ok := CheckCmdFormat(cons.QUIT, parts); !ok {
						cli.Receive <- []byte(info)
						continue
					}

					// 在Room.Clients中删除该client
					clients := room.Clients
					cliNum := len(clients)
					Lock.Lock()
					if clients[0] == cli {
						clients = append(clients[cliNum:])
					} else {
						clients = append(clients[:cliNum])
					}
					cliNum = len(clients)
					Lock.Unlock()

					// 状态更新
					cli.Status = cons.InHall

					// 消息通知
					info1 := "您退出了房间" + fmt.Sprintf("%d", cli.RoomId)
					info2 := cli.Username + "退出了房间"
					MsgToClient(info1, info2, cli)

					// 如果房间为空，在Hub.Rooms中删除该room
					Lock.Lock()
					if cliNum == 0 {
						delete(Hub.Rooms, cli.RoomId)
					}
					Lock.Unlock()
				case cons.READY:
					// cmd语法检查
					if info, ok := CheckCmdFormat(cons.READY, parts); !ok {
						cli.Receive <- []byte(info)
						continue
					}

					// 消息通知
					info1 := "您已准备"
					info2 := cli.Username + "已准备"
					MsgToClient(info1, info2, cli)

					cli.Status = cons.BeReady

					// 房间内的玩家都准备后就开始游戏
					if len(room.Clients) == 2 && room.Clients[0].Status == cons.BeReady && room.Clients[1].Status == cons.BeReady {
						game := InitGame(room.Clients)
						Hub.Games[game.Id] = game

						for _, value := range room.Clients {
							value.Status = cons.InGame
						}

						// 消息通知
						for _, value := range game.Clients {
							info1 = "游戏开始，您是" + string(value.GameInfo.Role) + "方\n"
							info1 += PrintChessBoard(value, game.Board)
							MsgToClient(info1, "", value)

							// 测试，输出到终端
							if game.Turn == value.GameInfo.Role {
								fmt.Println(info1)
							}
						}
					}
				default:
					MsgToClient(RoomCmdNotExistsErr(), "", cli)
				}
			//----------------------------------------------------Status: InGame----------------------------------------
			case cons.InGame:
				game := Hub.Games[cli.GameInfo.GameId]
				switch parts[0] {
				case cons.SELECT:
					// cmd语法检查
					if info, ok := CheckCmdFormat(cons.SELECT, parts); !ok {
						cli.Receive <- []byte(info)
						continue
					}

					// 是否是该用户的回合
					if cli.GameInfo.Role != game.Turn {
						MsgToClient(NotMyTurnErr(), "", cli)
						continue
					}

					// 找出选中棋子的所有可到达的路径
					obj := GetChessPieceObj(parts[1], parts[2], game.Board)
					if obj == nil {
						MsgToClient(ObjNotExistsErr(), "", cli)
						continue
					}
					if obj.Owner != cli {
						MsgToClient(ObjNotOwn(), "", cli)
						continue
					}
					paths := GetAllPaths(obj, game.Board)

					// 更新对局信息
					game.Selected = obj
					game.ValidPaths = paths

					// 消息通知
					info := "可以选择的路径：\n"
					for _, value := range paths {
						info += fmt.Sprintf(" [%d,%d] ", value[0], value[1])
					}
					MsgToClient(info, "", cli)
				case cons.MOVE:
					// cmd语法检查
					if info, ok := CheckCmdFormat(cons.MOVE, parts); !ok {
						cli.Receive <- []byte(info)
						continue
					}

					// 必须先选中要操作的棋子
					if game.Selected == nil {
						MsgToClient(SelectedFirstErr(), "", cli)
						continue
					}

					// 目标位置是否可以到达
					x, _ := strconv.Atoi(parts[1])
					y, _ := strconv.Atoi(parts[2])
					destLoc := []int{x, y}
					if ok := CheckLocationValid(destLoc, game.ValidPaths); !ok {
						MsgToClient(UnableMoveToErr(), "", cli)
						continue
					}

					// 更新棋盘
					UpdateChessBoard(game.Selected, destLoc, game.Board)

					// 更新对局信息
					if game.Turn == cons.A {
						game.Turn = cons.B
					} else {
						game.Turn = cons.A
					}
					game.Selected = nil

					// 消息通知
					opponent := game.Clients[0]
					if game.Clients[0] == cli {
						opponent = game.Clients[1]
					}
					info1 := PrintChessBoard(cli, game.Board)
					info2 := PrintChessBoard(opponent, game.Board)
					MsgToClient(info1, info2, cli)

					// 游戏结束
					if game.Board[destLoc[0]][destLoc[1]].Type == cons.King {
						info1 = "Game over. You win!!!"
						info2 = "Game over. You lost!!!"
						MsgToClient(info1, info2, cli)

					}

					// 测试，打印棋盘
					fmt.Println(info1)

				default:
					MsgToClient(GameCmdNotExistsErr(), "", cli)
				}
			}
		}
	}
}
