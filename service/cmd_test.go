package service

import (
	"chess/cons"
	"testing"
)

func TestCheckCmdFormat(t *testing.T) {
	var (
		inArg1 cmd = cons.ADDROOM
		inArg2     = []string{"findRoom", "123"}
	)
	want1, want2 := "", true

	res1, res2 := CheckCmdFormat(inArg1, inArg2)
	if !(res1 == want1 && res2 == want2) {
		t.Error("wrong")
	}
}
