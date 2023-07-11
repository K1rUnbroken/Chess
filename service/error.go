package service

//----------------------------------------------cmd usage-----------------------

func QuitUsageErr() string {
	msg := `
Usage:
	quit

Example: 
	quit`

	return msg
}

func FindRoomUsageErr() string {
	msg := `
Usage:
	findRoom

Example: 
	findRoom`

	return msg
}

func AddRoomUsageErr() string {
	msg := `
Usage:
	addRoom <RoomId>

Example: 
	addRoom 123456`

	return msg
}

func CreateRoomUsageErr() string {
	msg := `
Usage:
	createRoom

Example: 
	createRoom`

	return msg
}

func ReadyUsageErr() string {
	msg := `
Usage:
	ready

Example: 
	ready`

	return msg
}

func SelectUsageErr() string {
	msg := `
Usage:
	select <RowNum> <ColumnNum>

Example: 
	select 1 0`

	return msg
}

func MoveUsageErr() string {
	msg := `
Usage:
	move <RowNum> <ColumnNum>

Example: 
	move 2 0`

	return msg
}

func HallCmdNotExistsErr() string {
	msg := `
Error:
	Command not exists.

Supported commands:
	quit				退出大厅，断开连接
	findRoom			查找所有可加入的房间
	addRoom <roomId>		加入房间
	createRoom			创建房间
`
	return msg
}

func RoomCmdNotExistsErr() string {
	msg := `
Error:
	Command not exists.

Supported commands:
	quit				退出房间
	ready				准备
`
	return msg
}

func GameCmdNotExistsErr() string {
	msg := `
Error:
	Command not exists.

Supported commands:
	select <RowNum> <ColumnNum>		选择棋子进行操作
	move <RowNum> <ColumnNum>		移动棋子
`
	return msg
}

//--------------------------------------------room----------------------------------

func RoomNotExistsErr() string {
	return "Error: Room not exists."
}

//--------------------------------------------game---------------------------------

func NotMyTurnErr() string {
	return "Error: Not Your Turn, please wait for the opponent."
}

func ObjNotExistsErr() string {
	return "Error: The target chess piece does not exist. Please reselect."
}

func ObjNotOwn() string {
	return "Error: You can't pick the opponent's pieces. Please reselect."
}

func SelectedFirstErr() string {
	return "Error: Please select a chess piece first."
}

func UnableMoveToErr() string {
	return "Error: Target position not reachable"
}
