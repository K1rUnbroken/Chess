package service

import (
	"chess/cons"
	"chess/model"
	"fmt"
	"strconv"
	"time"
)

// InitGame 初始化游戏资源
func InitGame(clients []*model.Client) *model.Game {
	clients[0].GameInfo = &model.GameInfo{
		Role: cons.A,
	}
	clients[1].GameInfo = &model.GameInfo{
		Role: cons.B,
	}

	// 初始化棋盘
	board := InitChessBoard(clients)

	game := &model.Game{
		Id:         model.GameId(time.Now().Unix()),
		Turn:       cons.A,
		Clients:    clients,
		Ch:         make(chan *model.Message, 20),
		Selected:   nil,
		ValidPaths: make([][]int, 0),
		Board:      board,
	}

	clients[0].GameInfo.GameId = game.Id
	clients[1].GameInfo.GameId = game.Id

	return game
}

// InitChessBoard 初始化棋盘
func InitChessBoard(clients []*model.Client) *model.ChessBoard {
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

	return board[xx][yy]
}

// CheckLocationValid 检查要移动到的位置是否在给出的可以到达的路径范围内
func CheckLocationValid(destLoc []int, validPaths [][]int) bool {
	x, y := destLoc[0], destLoc[1]
	for _, value := range validPaths {
		if value[0] == x && value[1] == y {
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
		paths = PawnRules(row, column, obj.Owner.GameInfo.Role, board)
	case cons.Rook:
		paths = RookRules(row, column, obj.Owner.GameInfo.Role, board)
	case cons.Bishop:
		paths = BishopRules(row, column, obj.Owner.GameInfo.Role, board)
	case cons.King:
		paths = KingRules(row, column, obj.Owner.GameInfo.Role, board)
	case cons.Queen:
		paths = QueenRules(row, column, obj.Owner.GameInfo.Role, board)
	case cons.Knight:
		paths = KnightRules(row, column, obj.Owner.GameInfo.Role, board)
	}

	return paths
}

func UpdateChessBoard(src *model.ChessPieceObj, destLoc []int, board *model.ChessBoard) {
	board[destLoc[0]][destLoc[1]] = src
	board[src.Location.X][src.Location.Y] = nil
	src.Location.X, src.Location.Y = destLoc[0], destLoc[1]
}

func PrintChessBoard(cli *model.Client, board *model.ChessBoard) string {
	// 一个格子宽2，高1
	width := 2
	// 空格
	space := ""
	for i := 0; i < width; i++ {
		space += " "
	}
	// 换行
	newLine := "\n"

	// 顶部的行号
	str := fmt.Sprintf("%2s", " ")
	for i := 0; i < cons.ChessBoardColNum; i++ {
		str += fmt.Sprintf("%2d", i)
	}
	str += newLine

	// 上边框
	str += " A|"
	for i := 0; i < width*cons.ChessBoardColNum; i++ {
		str += "-"
	}
	str += "|" + newLine

	// 左侧的列号+左边框+棋子+右边框
	for i := 0; i < 8; i++ {
		str += fmt.Sprintf("%2d", i) + "|"
		for j := 0; j < 8; j++ {
			if board[i][j] != nil {
				value := board[i][j].Type
				str += fmt.Sprintf("%s", value)
			} else {
				str += space
			}
		}
		str += "|" + newLine
	}

	// 底部的边框
	str += " B|"
	for i := 0; i < width*cons.ChessBoardColNum; i++ {
		str += "-"
	}
	str += "|" + newLine

	// 提示
	info := ""
	game := Hub.Games[cli.GameInfo.GameId]
	if cli.GameInfo.Role == game.Turn {
		info = "现在是您的回合..."
	} else {
		info = "现在是对方的回合..."
	}
	str += info + "\n"

	return str
}
