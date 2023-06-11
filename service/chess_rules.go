package service

import "chess/model"

// PawnRules 兵
func PawnRules(row, column int, role model.Roles, board *model.ChessBoard) (paths [][]int) {
	if role == A {
		if row == 1 {
			paths = append(paths, []int{row + 2, column})
		}
		if row+1 < 8 && board[row+1][column] == nil {
			paths = append(paths, []int{row + 1, column})
		}
		// 判断其右下方是否有棋子
		if row+1 < 8 && column+1 < 8 && board[row+1][column+1] != nil {
			paths = append(paths, []int{row + 1, column + 1})
		}
		// 判断其左下方是否有棋子
		if row+1 < 8 && column-1 > 0 && board[row+1][column-1] != nil {
			paths = append(paths, []int{row + 1, column - 1})
		}
	} else {
		if row == 7 {
			paths = append(paths, []int{row, column - 2})
		}
		if row-1 > 0 && board[row-1][column] == nil {
			paths = append(paths, []int{row, column - 1})
		}
		// 判断其右上方是否有棋子
		if row-1 > 0 && column+1 < 8 && board[row-1][column+1] != nil {
			paths = append(paths, []int{row - 1, column + 1})
		}
		// 判断其左上方是否有棋子
		if row-1 > 0 && column-1 > 0 && board[row-1][column-1] != nil {
			paths = append(paths, []int{row - 1, column - 1})
		}
	}

	return
}

// RookRules 车
func RookRules(row, column int, role model.Roles, board *model.ChessBoard) (paths [][]int) {
	i, j := 0, 0
	// 空位置
	// 上
	for i = row - 1; i > 0 && board[i][column] == nil; i-- {
		paths = append(paths, []int{i, column})
	}
	top := []int{i, column}

	// 下
	for i = row + 1; i < 8 && board[i][column] == nil; i++ {
		paths = append(paths, []int{i, column})
	}
	button := []int{i, column}

	// 左
	for j = column - 1; j > 0 && board[row][j] == nil; j-- {
		paths = append(paths, []int{row, j})
	}
	left := []int{row, j}

	// 右
	for j = column + 1; j < 8 && board[row][j] == nil; j++ {
		paths = append(paths, []int{row, j})
	}
	right := []int{row, j}

	// 非空位置
	if role == A {
		if board[top[0]][top[1]].Owner.Role != A {
			paths = append(paths, top)
		}
		if board[button[0]][button[1]].Owner.Role != A {
			paths = append(paths, button)
		}
		if board[left[0]][left[1]].Owner.Role != A {
			paths = append(paths, left)
		}
		if board[right[0]][right[1]].Owner.Role != A {
			paths = append(paths, right)
		}
	} else {
		if board[top[0]][top[1]].Owner.Role != B {
			paths = append(paths, top)
		}
		if board[button[0]][button[1]].Owner.Role != B {
			paths = append(paths, button)
		}
		if board[left[0]][left[1]].Owner.Role != B {
			paths = append(paths, left)
		}
		if board[right[0]][right[1]].Owner.Role != B {
			paths = append(paths, right)
		}
	}

	return
}

// BishopRules 象
func BishopRules(row, column int, role model.Roles, board *model.ChessBoard) (paths [][]int) {
	i, j := 0, 0
	// 空位置
	// 右上
	for i, j = row-1, column+1; i < 8 && j > 0 && board[i][j] == nil; i, j = i-1, j+1 {
		paths = append(paths, []int{i, j})
	}
	rt := []int{i, j}

	// 右下
	for i, j = row+1, column+1; i < 8 && j > 0 && board[i][j] == nil; i, j = i+1, j+1 {
		paths = append(paths, []int{i, j})
	}
	rb := []int{i, j}

	// 左上
	for i, j = row-1, column-1; i < 8 && j > 0 && board[i][j] == nil; i, j = i-1, j-1 {
		paths = append(paths, []int{i, j})
	}
	lt := []int{i, j}

	// 左下
	for i, j = row+1, column-1; i < 8 && j > 0 && board[i][j] == nil; i, j = i+1, j-1 {
		paths = append(paths, []int{i, j})
	}
	lb := []int{i, j}

	// 非空位置
	if role == A {
		if board[rt[0]][rt[1]].Owner.Role != A {
			paths = append(paths, rt)
		}
		if board[rb[0]][rb[1]].Owner.Role != A {
			paths = append(paths, rb)
		}
		if board[lt[0]][lt[1]].Owner.Role != A {
			paths = append(paths, lt)
		}
		if board[lb[0]][lb[1]].Owner.Role != A {
			paths = append(paths, lb)
		}
	} else {
		if board[rt[0]][rt[1]].Owner.Role != B {
			paths = append(paths, rt)
		}
		if board[rb[0]][rb[1]].Owner.Role != B {
			paths = append(paths, rb)
		}
		if board[lt[0]][lt[1]].Owner.Role != B {
			paths = append(paths, lt)
		}
		if board[lb[0]][lb[1]].Owner.Role != B {
			paths = append(paths, lb)
		}
	}

	return
}

// KingRules 王
func KingRules(row, column int, role model.Roles, board *model.ChessBoard) (paths [][]int) {
	var fight model.Roles = A
	if role == A {
		fight = B
	}

	// 上
	if row-1 > 0 {
		if board[row-1][column] == nil || board[row-1][column].Owner.Role == fight {
			paths = append(paths, []int{row - 1, column})
		}
	}

	// 下
	if row+1 > 0 {
		if board[row+1][column] == nil || board[row+1][column].Owner.Role == fight {
			paths = append(paths, []int{row + 1, column})
		}
	}

	// 左
	if column-1 > 0 {
		if board[row][column-1] == nil || board[row][column-1].Owner.Role == fight {
			paths = append(paths, []int{row, column - 1})
		}
	}

	// 右
	if column+1 > 0 {
		if board[row][column+1] == nil || board[row][column+1].Owner.Role == fight {
			paths = append(paths, []int{row, column + 1})
		}
	}

	// 右上
	if row-1 > 0 && column+1 < 8 {
		if board[row-1][column+1] == nil || board[row-1][column+1].Owner.Role == fight {
			paths = append(paths, []int{row - 1, column + 1})
		}
	}

	// 右下
	if row+1 < 0 && column+1 < 8 {
		if board[row+1][column+1] == nil || board[row+1][column+1].Owner.Role == fight {
			paths = append(paths, []int{row + 1, column + 1})
		}
	}

	// 左上
	if row-1 > 0 && column-1 < 8 {
		if board[row-1][column-1] == nil || board[row-1][column-1].Owner.Role == fight {
			paths = append(paths, []int{row - 1, column - 1})
		}
	}

	// 左下
	if row+1 < 0 && column-1 < 8 {
		if board[row+1][column-1] == nil || board[row+1][column-1].Owner.Role == fight {
			paths = append(paths, []int{row + 1, column - 1})
		}
	}

	return
}

// QueenRules 后
func QueenRules(row, column int, role model.Roles, board *model.ChessBoard) (paths [][]int) {
	var fight model.Roles = A
	if role == A {
		fight = B
	}

	i, j := 0, 0
	// 上
	for i = row - 1; i > 0; i-- {
		if board[i][column] == nil || board[i][column].Owner.Role == fight {
			paths = append(paths, []int{i, column})
		}
	}

	// 下
	for i = row + 1; i < 8; i++ {
		if board[i][column] == nil || board[i][column].Owner.Role == fight {
			paths = append(paths, []int{i, column})
		}
	}

	// 左
	for j = column - 1; i > 0; j-- {
		if board[row][j] == nil || board[row][j].Owner.Role == fight {
			paths = append(paths, []int{row, j})
		}
	}

	// 右
	for j = column + 1; i < 8; j++ {
		if board[row][j] == nil || board[row][j].Owner.Role == fight {
			paths = append(paths, []int{row, j})
		}
	}

	// 右上
	for i, j = row-1, column+1; i > 0 && j < 8; i, j = i-1, j+1 {
		if board[i][j] == nil || board[i][j].Owner.Role == fight {
			paths = append(paths, []int{i, j})
		}
	}

	// 右下
	for i, j = row+1, column+1; i < 8 && j < 8; i, j = i+1, j+1 {
		if board[i][j] == nil || board[i][j].Owner.Role == fight {
			paths = append(paths, []int{i, j})
		}
	}

	// 左上
	for i, j = row-1, column-1; i > 0 && j > 0; i, j = i-1, j-1 {
		if board[i][j] == nil || board[i][j].Owner.Role == fight {
			paths = append(paths, []int{i, j})
		}
	}

	// 左下
	for i, j = row+1, column-1; i < 8 && j > 0; i, j = i+1, j-1 {
		if board[i][j] == nil || board[i][j].Owner.Role == fight {
			paths = append(paths, []int{i, j})
		}
	}

	return
}
