package service

import (
	"chess/model"
	"testing"
)

func TestGetToken(t *testing.T) {
	claims := model.MyClaims{Username: "kyr"}
	res, _ := GetToken(&claims)

	if res == "" {
		t.Error("GetToken wrong")
	}
}
