package service

import (
	"chess/cons"
	"chess/model"
)

// PawnRules 兵
func PawnRules(row, column int, role model.Roles, board *model.ChessBoard) (paths [][]int) {
	x, y := 0, 0
	if role == cons.A {
		// 下
		x, y = row+1, column
		if x < cons.ChessBoardRowNum {
			if board[x][y] == nil {
				paths = append(paths, []int{x, y})
			}
		}
		// 如果还在初始位置
		if row == 1 {
			x, y = row+2, column
			if board[x][y] == nil {
				paths = append(paths, []int{x, y})
			}
		}
		// 右下
		x, y = row+1, column+1
		if x < cons.ChessBoardRowNum && y < cons.ChessBoardColNum {
			if board[x][y] != nil && board[x][y].Owner.GameInfo.Role == cons.B {
				paths = append(paths, []int{x, y})
			}
		}
		// 左下
		x, y = row+1, column-1
		if x < cons.ChessBoardRowNum && y >= 0 {
			if board[x][y] != nil && board[x][y].Owner.GameInfo.Role == cons.B {
				paths = append(paths, []int{row + 1, column - 1})
			}
		}
	} else {
		// 上
		x, y = row-1, column
		if x >= 0 {
			if board[x][y] == nil {
				paths = append(paths, []int{x, y})
			}
		}
		// 如果还在初始位置
		if row == 6 {
			x, y = row-2, column
			if board[x][y] == nil {
				paths = append(paths, []int{x, y})
			}
		}
		// 右上
		x, y = row-1, column+1
		if x >= 0 && y < cons.ChessBoardColNum {
			if board[x][y] != nil && board[x][y].Owner.GameInfo.Role == cons.A {
				paths = append(paths, []int{x, y})
			}
		}
		// 左上
		x, y = row-1, column-1
		if x >= 0 && y >= 0 {
			if board[x][y] != nil && board[x][y].Owner.GameInfo.Role == cons.A {
				paths = append(paths, []int{x, y})
			}
		}
	}

	return
}

// RookRules 车
func RookRules(row, column int, role model.Roles, board *model.ChessBoard) (paths [][]int) {
	var fight model.Roles = cons.A
	if role == cons.A {
		fight = cons.B
	}

	i, j := 0, 0
	// 上
	for i = row - 1; i >= 0; i-- {
		if board[i][column] == nil || board[i][column].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{i, column})
		}
	}

	// 下
	for i = row + 1; i < cons.ChessBoardRowNum && board[i][column] == nil; i++ {
		if board[i][column] == nil || board[i][column].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{i, column})
		}
	}

	// 左
	for j = column - 1; j >= 0 && board[row][j] == nil; j-- {
		if board[row][j] == nil || board[row][j].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{row, j})
		}
	}

	// 右
	for j = column + 1; j < cons.ChessBoardColNum && board[row][j] == nil; j++ {
		if board[row][j] == nil || board[row][j].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{row, j})
		}
	}

	return
}

// BishopRules 象
func BishopRules(row, column int, role model.Roles, board *model.ChessBoard) (paths [][]int) {
	var fight model.Roles = cons.A
	if role == cons.A {
		fight = cons.B
	}

	i, j := 0, 0
	// 右上
	for i, j = row-1, column+1; i >= 0 && j < cons.ChessBoardColNum; i, j = i-1, j+1 {
		if board[i][j] == nil || board[i][j].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{i, j})
		}
	}

	// 右下
	for i, j = row+1, column+1; i < cons.ChessBoardRowNum && j < cons.ChessBoardColNum; i, j = i+1, j+1 {
		if board[i][j] == nil || board[i][j].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{i, j})
		}
	}

	// 左上
	for i, j = row-1, column-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == nil || board[i][j].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{i, j})
		}
	}

	// 左下
	for i, j = row+1, column-1; i < cons.ChessBoardRowNum && j >= 0; i, j = i+1, j-1 {
		if board[i][j] == nil || board[i][j].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{i, j})
		}
	}

	return
}

// KingRules 王
func KingRules(row, column int, role model.Roles, board *model.ChessBoard) (paths [][]int) {
	var fight model.Roles = cons.A
	if role == cons.A {
		fight = cons.B
	}

	// 八个方位
	i, j := 0, 0
	directions := [][]int{
		{row - 1, column},     // 上
		{row + 1, column},     // 下
		{row, column - 1},     // 左
		{row, column + 1},     // 右
		{row - 1, column - 1}, // 左上
		{row + 1, column - 1}, // 左下
		{row - 1, column + 1}, // 右上
		{row + 1, column + 1}, // 右下
	}
	for _, arr := range directions {
		i, j = arr[0], arr[1]
		if i < 0 || i > cons.ChessBoardRowNum-1 || j < 0 || j > cons.ChessBoardColNum {
			continue
		}
		if board[i][j] == nil || board[i][j].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{i, j})
		}
	}

	return
}

// QueenRules 后
func QueenRules(row, column int, role model.Roles, board *model.ChessBoard) (paths [][]int) {
	var fight model.Roles = cons.A
	if role == cons.A {
		fight = cons.B
	}

	i, j := 0, 0
	// 上
	for i = row - 1; i >= 0; i-- {
		if board[i][column] == nil || board[i][column].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{i, column})
		}
	}

	// 下
	for i = row + 1; i < 8; i++ {
		if board[i][column] == nil || board[i][column].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{i, column})
		}
	}

	// 左
	for j = column - 1; i >= 0; j-- {
		if board[row][j] == nil || board[row][j].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{row, j})
		}
	}

	// 右
	for j = column + 1; i < 8; j++ {
		if board[row][j] == nil || board[row][j].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{row, j})
		}
	}

	// 右上
	for i, j = row-1, column+1; i >= 0 && j < 8; i, j = i-1, j+1 {
		if board[i][j] == nil || board[i][j].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{i, j})
		}
	}

	// 右下
	for i, j = row+1, column+1; i < 8 && j < 8; i, j = i+1, j+1 {
		if board[i][j] == nil || board[i][j].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{i, j})
		}
	}

	// 左上
	for i, j = row-1, column-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == nil || board[i][j].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{i, j})
		}
	}

	// 左下
	for i, j = row+1, column-1; i < 8 && j >= 0; i, j = i+1, j-1 {
		if board[i][j] == nil || board[i][j].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{i, j})
		}
	}

	return
}

// KnightRules 马
func KnightRules(row, column int, role model.Roles, board *model.ChessBoard) (paths [][]int) {
	var fight model.Roles = cons.A
	if role == cons.A {
		fight = cons.B
	}

	// 四个方位
	i, j := 0, 0
	directions := [][]int{
		{row - 2, column - 1}, // 左上1
		{row + 2, column - 1}, // 左下1
		{row - 2, column + 1}, // 右上1
		{row + 2, column + 1}, // 右下1
		{row - 1, column - 2}, // 左上2
		{row + 1, column - 2}, // 左下2
		{row - 1, column + 2}, // 右上2
		{row + 1, column + 2}, // 右下2
	}
	for _, arr := range directions {
		i, j = arr[0], arr[1]
		if i < 0 || i > cons.ChessBoardRowNum-1 || j < 0 || j > cons.ChessBoardColNum {
			continue
		}
		if board[i][j] == nil || board[i][j].Owner.GameInfo.Role == fight {
			paths = append(paths, []int{i, j})
		}
	}

	return
}
