package service

import (
	"chess/cons"
	"chess/model"
	"fmt"
	"strconv"
	"time"
)

// InitGame 初始化游戏资源
func InitGame(roomId model.RoomId) *model.Game {
	// 创建游戏对局内的Client对象
	gameClients := []*model.GameClient{
		// A方
		{
			Conn:     Hub.Rooms[roomId].Clients[0].Conn,
			Username: Hub.Rooms[roomId].Clients[0].Username,
			Receive:  Hub.Rooms[roomId].Clients[0].Receive,
			Role:     cons.A,
		},
		// B方
		{
			Conn:     Hub.Rooms[roomId].Clients[1].Conn,
			Username: Hub.Rooms[roomId].Clients[1].Username,
			Receive:  Hub.Rooms[roomId].Clients[1].Receive,
			Role:     cons.B,
		},
	}
	// 初始化棋盘
	board := InitChessBoard(gameClients)

	game := &model.Game{
		Id:         model.GameId(time.Now().Unix()),
		Turn:       cons.A,
		Clients:    gameClients,
		Ch:         make(chan *model.GameMessage, 20),
		Selected:   nil,
		ValidPaths: make([][]int, 0),
		Board:      board,
	}

	game.Clients[0].GameId = game.Id
	game.Clients[1].GameId = game.Id

	return game
}

// InitChessBoard 初始化棋盘
func InitChessBoard(clients []*model.GameClient) *model.ChessBoard {
	board := new(model.ChessBoard)

	// 除了Pawn之外的其他棋子
	arr := []model.ChessPieceType{cons.Rook, cons.Knight, cons.Bishop, cons.King, cons.Queen, cons.Bishop, cons.Knight, cons.Rook}
	for i, value := range arr {
		// ASide
		board[0][i] = &model.ChessPieceObj{
			Type:  value,
			Owner: clients[0],
			Location: struct {
				X int
				Y int
			}{
				X: 0,
				Y: i,
			},
		}
		// BSide
		board[7][i] = &model.ChessPieceObj{
			Type:  value,
			Owner: clients[1],
			Location: struct {
				X int
				Y int
			}{
				X: 7,
				Y: i,
			},
		}
	}

	// Pawn
	for i := 0; i < cons.ChessBoardColNum; i++ {
		// ASide
		board[1][i] = &model.ChessPieceObj{
			Type:  cons.Pawn,
			Owner: clients[0],
			Location: struct {
				X int
				Y int
			}{
				X: 1,
				Y: i,
			},
		}
		// BSide
		board[6][i] = &model.ChessPieceObj{
			Type:  cons.Pawn,
			Owner: clients[1],
			Location: struct {
				X int
				Y int
			}{
				X: 6,
				Y: i,
			},
		}
	}

	return board
}

// GetChessPieceObj 根据x坐标和y坐标返回该棋子对象
func GetChessPieceObj(x, y string, board *model.ChessBoard) *model.ChessPieceObj {
	xx, _ := strconv.Atoi(x)
	yy, _ := strconv.Atoi(y)

	// 超出棋盘范围
	if xx < 0 || xx > cons.ChessBoardColNum || yy < 0 || yy > cons.ChessBoardRowNum {
		return nil
	}

	for _, value := range board {
		for _, value2 := range value {
			if value2.Location.X == xx && value2.Location.Y == yy {
				return value2
			}
		}
	}

	return nil
}

// CheckLocationValid 检查要移动到的位置是否在可以移动的路径范围内
func CheckLocationValid(destObj *model.ChessPieceObj, validPaths [][]int) bool {
	if destObj == nil {
		return false
	}

	i, j := destObj.Location.X, destObj.Location.Y
	for _, value := range validPaths {
		if value[0] == i && value[1] == j {
			return true
		}
	}

	return false
}

// GetAllPaths 选中一个棋子，得到它的所有路径
func GetAllPaths(obj *model.ChessPieceObj, board *model.ChessBoard) (paths [][]int) {
	row, column := obj.Location.X, obj.Location.Y
	switch obj.Type {
	case cons.Pawn:
		paths = PawnRules(row, column, obj.Owner.Role, board)
	case cons.Rook:
		paths = RookRules(row, column, obj.Owner.Role, board)
	case cons.Bishop:
		paths = BishopRules(row, column, obj.Owner.Role, board)
	case cons.King:
		paths = KingRules(row, column, obj.Owner.Role, board)
	case cons.Queen:
		paths = QueenRules(row, column, obj.Owner.Role, board)
	}

	return paths
}

func UpdateChessBoard(src, dst *model.ChessPieceObj, board *model.ChessBoard) {
	board[src.Location.X][src.Location.Y] = nil
	board[dst.Location.X][dst.Location.Y] = dst
}

func PrintChessBoard(role model.Roles, board *model.ChessBoard) string {
	// 一个格子宽2，高1
	width := 2
	// 空格
	space := " "
	for i := 0; i < width; i++ {
		space += "&nbsp"
	}
	// 换行
	newLine := "\n"

	//str := "&nbsp"
	//for i := 0; i < cons.ChessBoardColNum; i++ {
	//	str += fmt.Sprintf("%02d", i)
	//}
	str := "|"
	for i := 0; i < width*cons.ChessBoardColNum; i++ {
		str += "-"
	}
	str += "|" + newLine + "|"

	if role == cons.A {
		for i := 7; i >= 0; i-- {
			for j := 0; j < 8; j++ {
				if board[i][j] != nil {
					value := board[i][j].Type
					str += fmt.Sprintf("%s", value)
				} else {
					str += space
				}
			}
			str += "|" + newLine + "|"
		}
	} else {
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				if board[i][j] != nil {
					value := board[i][j].Type
					str += fmt.Sprintf("%s", value)
				} else {
					str += space
				}
			}
			str += "|" + newLine + "|"
		}
	}

	for i := 0; i < width*cons.ChessBoardColNum; i++ {
		str += "-"
	}
	str += "|" + newLine

	return str
}
