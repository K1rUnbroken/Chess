package service

import (
	"chess/model"
)

func InitChessBoard(clients []*model.GameClient) *model.ChessBoard {
	board := new(model.ChessBoard)

	var idType = [][]interface{}{
		{"rook1", Rook},
		{"knight", Knight},
		{"bishop1", Bishop},
		{"king", King},
		{"queen", Queen},
		{"bishop2", Bishop},
		{"knight2", Knight},
		{"rook2", Rook},
	}
	for i, value := range idType {
		// ASide
		board[0][i] = &model.ChessPieceObj{
			Id:       value[0].(string),
			Type:     value[1].(model.ChessPieceType),
			Owner:    clients[0],
			Location: []int{0, i},
		}
		// BSide
		board[7][i] = &model.ChessPieceObj{
			Id:       value[0].(string),
			Type:     value[1].(model.ChessPieceType),
			Owner:    clients[1],
			Location: []int{7, i},
		}
	}
	for i := 0; i < 8; i++ {
		// ASide
		board[1][i] = &model.ChessPieceObj{
			Id:       "pawn" + string(rune(i+1)),
			Type:     Pawn,
			Owner:    clients[0],
			Location: []int{1, i},
		}
		// BSide
		board[6][i] = &model.ChessPieceObj{
			Id:       "pawn" + string(rune(i+1)),
			Type:     Pawn,
			Owner:    clients[1],
			Location: []int{6, i},
		}
	}

	return board
}

func GetChessPieceObj(dest string, board *model.ChessBoard) *model.ChessPieceObj {
	for _, value := range board {
		for _, value2 := range value {
			if value2.Id == dest {
				return value2
			}
		}
	}

	return nil
}

func CheckValid(destObj *model.ChessPieceObj, validPaths [][]int) bool {
	if destObj == nil {
		return false
	}

	i, j := destObj.Location[0], destObj.Location[1]
	
	return true
}

// GetAllPaths 选中一个棋子，得到它的所有路径
func GetAllPaths(obj *model.ChessPieceObj, board *model.ChessBoard) (paths [][]int) {
	row, column := obj.Location[0], obj.Location[1]
	switch obj.Type {
	case Pawn:
		paths = PawnRules(row, column, obj.Owner.Role, board)
	case Rook:
		paths = RookRules(row, column, obj.Owner.Role, board)
	case Bishop:
		paths = BishopRules(row, column, obj.Owner.Role, board)
	case King:
		paths = KingRules(row, column, obj.Owner.Role, board)
	case Queen:
		paths = QueenRules(row, column, obj.Owner.Role, board)
	}

	return paths
}
