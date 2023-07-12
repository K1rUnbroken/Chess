package service

import (
	"chess/cons"
	"chess/model"
	"fmt"
	"github.com/gorilla/websocket"
	"testing"
)

var (
	clients = []*model.Client{
		{
			Conn: &websocket.Conn{},
			GameInfo: &model.GameInfo{
				GameId: 123,
				Role:   cons.A,
			},
		},
		{
			Conn: &websocket.Conn{},
			GameInfo: &model.GameInfo{
				Role: cons.B,
			},
		},
	}
	testChessPieceObj = &model.ChessPieceObj{
		Type:  cons.Pawn,
		Owner: clients[0],
		Location: struct {
			X int
			Y int
		}{X: 1, Y: 0},
	}
	board = InitChessBoard(clients)
)

func TestInitGame(t *testing.T) {
	game := InitGame(clients)
	Hub.Games[game.Id] = game
}

func TestInitChessBoard(t *testing.T) {
	board = InitChessBoard(clients)
	fmt.Println("board:\n", PrintChessBoard(clients[0], board))
}

func TestGetChessPieceObj(t *testing.T) {
	x, y := "1", "0"
	var expected model.ChessPieceType = cons.Pawn

	if res := GetChessPieceObj(x, y, board); res.Type != expected {
		t.Error("GetChessPieceObj wrong")
	}
}

func TestCheckLocationValid(t *testing.T) {
	in1, in2 := []int{1, 0}, [][]int{{1, 0}, {2, 0}}
	expected := true

	if res := CheckLocationValid(in1, in2); res != expected {
		t.Error("wrong")
	}
}

func TestGetAllPaths(t *testing.T) {
	expected := [][]int{{2, 0}, {3, 0}}
	res := GetAllPaths(testChessPieceObj, board)

	if len(expected) != len(res) {
		t.Error("wrong")
	}
}

func TestUpdateChessBoard(t *testing.T) {
	in := []int{2, 0}

	UpdateChessBoard(testChessPieceObj, in, board)
	if board[1][0] != nil || board[2][0] != testChessPieceObj {
		t.Error("wrong")
	}

	board = InitChessBoard(clients)
}

func TestPrintChessBoard(t *testing.T) {
	PrintChessBoard(clients[1], board)
}
