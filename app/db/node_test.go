package db

import (
	"testing"

	"github.com/2n1/proxysubscriber/app/defs/entity"
)

func TestAddNode(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	data := map[string]interface{}{
		"name":   "vemss1",
		"server": "bar.foo",
		"port":   443,
		"passwd": "ca8d9a0a-57fa-4969-ae30-609936080c78",
		"cipher": "auto",
		"node_type":"vmess",
		"group_id":1,
		"sni":"bar.foo",
		"alter_id":64,
		"ws_path":"/foobar",
		"ws_host":"bar.foo",
		"cf_ip":1,
	}
	id, err := AddNode(data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("node id:", id)
}
func TestFindNode(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	node, err := FindNode(6)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v",*node)
}
func TestEditNode(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	data := map[string]interface{}{
		"name":   "vemss222",
		"server": "bar.foodd",
		"port":   1443,
		"passwd": "Ea8d9a0a-57fa-4969-ae30-609936080c78",
		"cipher": "Dauto",
		"node_type":"Fvmess",
		"group_id":22,
		"sni":"AEbar.foo",
		"alter_id":12,
		"ws_path":"/foobarDD",
		"ws_host":"FFDbar.foo",
		"cf_ip":10,
	}
	aff, err := EditNode(6,data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("edit node aff:", aff)
}

func TestFindNodes(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	nodes, err := FindNodes(0, 15, "node_type=?","ss")
	if err != nil {
		t.Fatal(err)
	}
	for _,node:=range nodes.Data.([]*entity.Node){
		t.Logf("%#v",*node)
	}
}
func TestFindAllNodes(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	nodes, err := FindAllNodes(0, 4,1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(nodes.Page,nodes.TotalRecord,nodes.TotalPage)
	for _,node:=range nodes.Data.([]*entity.Node){
		t.Logf("%#v",*node)
	}
}
func TestDeleteNode(t *testing.T) {

}