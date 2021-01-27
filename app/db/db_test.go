package db

import "testing"

func TestInit(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	defer Close()
}

func TestCreateTables(t *testing.T) {
	if err := OpenFile("../../ps.db"); err != nil {
		t.Fatal(err)
	}
	defer Close()
	if err := CreateTables("../../ps.sql"); err != nil {
		t.Fatal(err)
	}
}
