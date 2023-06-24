package model

import (
	"github.com/gorilla/websocket"
)

// Hub
/*
	功能：程序运行时创建一个HallHub，用于管理所有房间和处理用户命令
	属性：
 		Rooms	存储所有的房间
 		Ch		通道，传输用户命令交给HallHub处理
	支持的用户命令如下：
		（1）quit				退出程序
		（2）findRoom			查找所有可以加入的房间，返回房间号
		（3）addRoom	<RoomId>	加入指定房间号
		（4）createRoom			创建一个新房间
*/
type Hub struct {
	Rooms map[RoomId]*Room
	Ch    chan *Message
}

// Room
/*
	功能：每创建一个新的房间，HallHub会创建一个RoomHub来管理该房间以及处理该房间内用户的命令
	属性：
		Clients		存储房间内的所有客户端
		Ch			通道，传输用户命令交给RoomHub处理
*/
type Room struct {
	Id      RoomId
	Clients []*Client
	Ch      chan *Message
}

// Client
/*
	功能：代表一个客户端实例
	属性：
		Conn 		ws连接
		Username	用户姓名
		Receive		通道，用于接受RoomHub广播给自己的消息
		RoomId		所属房间的id
		Status 		用户的状态
*/
type Client struct {
	Conn     *websocket.Conn
	Username string
	Receive  chan []byte
	RoomId   RoomId
	Status   ClientStatus
}

type GameHub struct {
	Games map[GameId]*Game
	Ch    chan *GameMessage
}

// Game
/*
	功能：表示一局游戏资源，每开始一局游戏，RoomHub就会创建一个GameHub来管理该局游戏以及处理玩家的命令
	属性：
		Turn		表示当前轮到谁的回合，属性值为A或B
		Clients		进行该局游戏的两个客户端对象（这里是GameClient而不是Client）
		Ch			通道，将用户的命令操作传输给GameHub处理
		Selected	表示当前被选中的棋子（用户每次移动棋子之前需要先选中它）
		ValidPaths	表示被选中的棋子的所有可以到达的路径
		Board		表示棋盘
*/
type Game struct {
	Id         GameId
	Turn       Roles
	Clients    []*GameClient
	Ch         chan *GameMessage
	Selected   *ChessPieceObj
	ValidPaths [][]int
	Board      *ChessBoard
}

// GameClient
/*
	功能：表示游戏内的客户端对象
	属性：
		Conn		即Client.Conn，与Client.Conn指向的连接时同一个
		Username	值与Client.Username相同
		Receive		即Client.Receive，与Client.Receive指向的chan是同一个
		Role		表示该玩家的角色，要么是A方，要么是B方
*/
type GameClient struct {
	GameId   GameId
	Conn     *websocket.Conn
	Username string
	Receive  chan []byte
	Role     Roles
}

// Message
/*
	功能：用于把客户端发送的消息传输给HallHub或RoomHub处理
*/
type Message struct {
	Data []byte
	Cli  *Client
}

// GameMessage
/*
	功能：在游戏中，将用户的消息打包传输给GameHub处理
*/
type GameMessage struct {
	Data []byte
	Cli  *GameClient
}

// ChessPieceObj
/*
	功能：表示一个棋子对象
	属性：
		Id			string类型，其实是该棋子的名称，但是每个棋子的名称不同
		Type		表示棋子的类型
		Owner		棋子所属的玩家
		Location	棋子的坐标
*/
type ChessPieceObj struct {
	Type     ChessPieceType
	Owner    *GameClient
	Location struct {
		X int
		Y int
	}
}

// ChessBoard 8x8方格的棋盘
type ChessBoard [8][8]*ChessPieceObj

type ChessPieceType string

// Roles A方为0，B方为1
type Roles int

type ClientStatus int

type RoomId int
type GameId int
type Cmd string
