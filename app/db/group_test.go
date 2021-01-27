package db

import (
	"testing"

	"github.com/2n1/proxysubscriber/app/defs/entity"
)

func TestGroupAdd(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	id, err := AddGroup("Test1","")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("group id:", id)
}

func TestFindGroup(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	group, err := FindGroup(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(group.ID, group.Name)
}
func TestCountGroup(t *testing.T) {
	defer Close()
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	count, err := CountGroup("id=?", 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("count of #1 group:", count)
}
func TestCountGroupByName(t *testing.T) {
	defer Close()
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	count, err := CountGroupByName("Test")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("count of Test group:", count)
}
func TestFindGroups(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	groups, err := FindGroups(0, 15, "name LIKE ?", "%Test%")
	if err != nil {
		t.Fatal(err)
	}
	for _, group := range groups.Data.([]*entity.Group) {
		t.Log(group.ID, group.Name)
	}
}
func TestFindAllGroups(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	groups, err := FindAllGroups(0, 15)
	if err != nil {
		t.Fatal(err)
	}
	for _, group := range groups.Data.([]*entity.Group) {
		t.Log(group.ID, group.Name)
	}
}
func TestEditGroup(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	aff, err := EditGroup(1, "Test 222")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("edited:", aff)
}
func TestDeleteGroup(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	aff, err := DeleteGroup(2)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Deleted:", aff)
}
