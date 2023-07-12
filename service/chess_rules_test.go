package service

import (
	"chess/cons"
	"chess/model"
	"testing"
)

func TestPawnRules(t *testing.T) {
	var (
		in1             = 1
		in2             = 0
		in3 model.Roles = cons.A
	)
	res := PawnRules(in1, in2, in3, board)
	expected := [][]int{
		{2, 0},
		{3, 0},
	}

	if len(res) != len(expected) {
		t.Error("PawnRules wrong")
	}
}

func TestRookRules(t *testing.T) {
	var (
		in1             = 2
		in2             = 0
		in3 model.Roles = cons.A
	)
	_ = RookRules(in1, in2, in3, board)
}

func TestBishopRules(t *testing.T) {
	var (
		in1             = 2
		in2             = 4
		in3 model.Roles = cons.A
	)
	_ = BishopRules(in1, in2, in3, board)
}

func TestKingRules(t *testing.T) {
	var (
		in1             = 2
		in2             = 3
		in3 model.Roles = cons.A
	)
	_ = KingRules(in1, in2, in3, board)
}

func TestQueenRules(t *testing.T) {
	var (
		in1             = 2
		in2             = 6
		in3 model.Roles = cons.A
	)
	_ = QueenRules(in1, in2, in3, board)
}

func TestKnightRules(t *testing.T) {
	var (
		in1             = 3
		in2             = 2
		in3 model.Roles = cons.A
	)
	_ = KnightRules(in1, in2, in3, board)
}
