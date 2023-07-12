package service

import "testing"

func TestGenUpgrader(t *testing.T) {
	res := GenUpgrader()
	if res == nil {
		t.Error("GenUpgrader wrong")
	}
}
