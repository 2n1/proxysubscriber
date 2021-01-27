package cfg

import "testing"

func TestInitFrom(t *testing.T) {
	if err:=InitFrom("../../config.json");err!=nil{
		t.Fatal(err)
	}
	t.Log(Cfg.DbFile,Cfg.PageSize,Cfg.SQLFile)
}
