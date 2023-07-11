package cons

// 表示用户状态，用于Client.Status属性
const (
	InHall  = 1 // 在大厅
	InRoom  = 2 // 在房间
	BeReady = 3 // 已准备
	InGame  = 4 // 游戏中
)

// 棋子的类型
const (
	King   = "王" // 王
	Queen  = "后" // 后
	Bishop = "象" // 象
	Rook   = "车" // 车
	Knight = "马" // 马
	Pawn   = "兵" // 兵
)

// cmd
const (
	QUIT       = "quit"
	FINDROOM   = "findRoom"
	ADDROOM    = "addRoom"
	CREATEROOM = "createRoom"
	READY      = "ready"
	SELECT     = "select"
	MOVE       = "move"
)

// 游戏对局中的A方或B方
const (
	A = "A"
	B = "B"
)

// 8x8棋盘
const (
	ChessBoardRowNum = 8
	ChessBoardColNum = 8
)
