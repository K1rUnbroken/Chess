package service

//----------------------------------------------cmd usage-----------------------

var (
	QuitUsageErr = `
Usage:
	quit

Example: 
	quit
`
	FindRoomUsageErr = `
Usage:
	findRoom

Example: 
	findRoom
`
	AddRoomUsageErr = `
Usage:
	addRoom <RoomId>

Example: 
	addRoom 123456
`
	CreateRoomUsageErr = `
Usage:
	createRoom

Example: 
	createRoom
`
	ReadyUsageErr = `
Usage:
	ready

Example: 
	ready
`
	SelectUsageErr = `
Usage:
	select <RowNum> <ColumnNum>

Example: 
	select 1 0
`
	MoveUsageErr = `
Usage:
	move <RowNum> <ColumnNum>

Example: 
	move 2 0
`
	HallCmdNotExistsErr = `
Error:
	Command not exists.

Supported commands:
	quit				退出大厅，断开连接
	findRoom			查找所有可加入的房间
	addRoom <roomId>		加入房间
	createRoom			创建房间
`
	RoomCmdNotExistsErr = `
Error:
	Command not exists.

Supported commands:
	quit				退出房间
	ready				准备
`
	GameCmdNotExistsErr = `
Error:
	Command not exists.

Supported commands:
	select <RowNum> <ColumnNum>		选择棋子进行操作
	move <RowNum> <ColumnNum>		移动棋子
`
)

//--------------------------------------------room----------------------------------

var (
	RoomNotExistsErr = "Error: Room not exists."
)

//--------------------------------------------game---------------------------------

var (
	NotMyTurnErr     = "Error: Not Your Turn, please wait for the opponent."
	ObjNotExistsErr  = "Error: The target chess piece does not exist. Please reselect."
	ObjNotOwnErr     = "Error: You can't pick the opponent's pieces. Please reselect."
	SelectedFirstErr = "Error: Please select a chess piece first."
	UnableMoveToErr  = "Error: Target position not reachable"
)
