package service

import (
	"testing"

	"github.com/2n1/proxysubscriber/app/cfg"
)

func TestAddGroup(t *testing.T) {
	if err := cfg.InitFrom("../../config.json");err!=nil{
		t.Fatal(err)
	}
	id, err := AddGroup("你好")
	if err!=nil{
		t.Fatal(err.String())
	}
	t.Log("group id:",id)
}