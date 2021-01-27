package db

import "testing"

func TestUpdateCfips(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	data:=map[string]interface{} {
		"cu_ip":"104.22.65.143",
		"cu_label":"CU",
		"ct_ip":"",
		"ct_label":"CT",
		"cm_ip":"104.19.103.16",
		"cm_label":"CM",
	}
	if err:=UpdateCfips(data);err!=nil{
		t.Fatal(err)
	}
}

func TestGetCfips(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	cfips, err := GetCfips()
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(cfips.ChinaUnicomIP, cfips.ChinaUnicomLable, cfips.ChinaTelecomIP,cfips.ChinaTelecomLable,
		cfips.ChinaMobileIP,cfips.ChinaMobileLable)
}
