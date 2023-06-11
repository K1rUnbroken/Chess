package model

import (
	"github.com/gorilla/websocket"
)

type Hub struct {
	Rooms map[int]*Room
	Srv   chan *Message
}

type Client struct {
	Conn     *websocket.Conn
	Username string
	Receive  chan []byte
	RoomId   int

	// 用户处于什么状态："大厅"-消息发送到系统chan；"房间"-消息发送到room chan
	Status int
}

type GameClient struct {
	Conn     *websocket.Conn
	Username string
	Receive  chan []byte
	Role     Roles
}

type Message struct {
	Data []byte
	Cli  *Client
}

type GameMessage struct {
	Data []byte
	Cli  *GameClient
}

// ChessSource 棋局资源
type ChessSource struct {
	Move []byte

	// 棋盘
	Board *ChessBoard

	Clients []*Client
}

type Room struct {
	Clients []*Client
	Srv     chan *Message
}

type GameSource struct {
	Turn    Roles
	Clients []*GameClient
	Data    chan *GameMessage

	// 当前被选中的棋子
	Selected *ChessPieceObj

	// 当前被选中的棋子的可到达的路径
	ValidPaths [][]int
}

// ChessPieceObj 棋子
type ChessPieceObj struct {
	Id       string
	Type     ChessPieceType
	Owner    *GameClient
	Location []int
}

// ChessBoard 8x8方格
type ChessBoard [8][8]*ChessPieceObj

type ChessPieceType int

// Roles A方为0，B方为1
type Roles int
