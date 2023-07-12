package service

import (
	"chess/cons"
	"chess/model"
)

type cmd string

func CheckCmdFormat(cmd cmd, parts []string) (string, bool) {
	switch cmd {
	case cons.QUIT:
		if len(parts) != 1 {
			return QuitUsageErr, false
		}
	case cons.FINDROOM:
		if len(parts) != 1 {
			return FindRoomUsageErr, false
		}
	case cons.ADDROOM:
		if len(parts) != 2 {
			return AddRoomUsageErr, false
		}
	case cons.CREATEROOM:
		if len(parts) != 1 {
			return CreateRoomUsageErr, false
		}
	case cons.READY:
		if len(parts) != 1 {
			return ReadyUsageErr, false
		}
	case cons.SELECT:
		if len(parts) != 3 {
			return SelectUsageErr, false
		}
	case cons.MOVE:
		if len(parts) != 3 {
			return MoveUsageErr, false
		}
	}

	return "", true
}

func MsgToClient(toSelf, toOpponent string, cli *model.Client) {
	// 通知自己
	if toSelf != "" {
		cli.Receive <- []byte(toSelf)
	}

	// 通知对手
	if toOpponent != "" {
		for _, value := range Hub.Rooms[cli.RoomId].Clients {
			if value != cli {
				value.Receive <- []byte(toOpponent)
			}
		}
	}
}
