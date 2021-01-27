package db

import (
	"database/sql"
	"io/ioutil"

	"github.com/2n1/proxysubscriber/app/cfg"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type InputData map[string]interface{}

func (d InputData) Get(key string) interface{} {
	if v, ok := d[key]; ok {
		return v
	}
	return nil
}

func (d InputData) GetString(key string) string {
	if v := d.Get(key); v != nil {
		if vv, ok := v.(string); ok {
			return vv
		}
	}
	return ""
}
func (d InputData) GetInt(key string) int {
	if v := d.Get(key); v != nil {
		if vv, ok := v.(int); ok {
			return vv
		}
	}
	return 0
}
func (d InputData) GetInt64(key string) int64 {
	if v := d.Get(key); v != nil {
		if vv, ok := v.(int64); ok {
			return vv
		}
	}
	return 0
}

func OpenFile(dbfile string) (err error) {
	db, err = sql.Open("sqlite3", dbfile)
	if err != nil {
		return err
	}
	return nil
}
func Open() error {
	return OpenFile(cfg.Cfg.DbFile)
}

func Close() error {
	return db.Close()
}

func CreateTables(sqlfile string) error {
	buf, err := ioutil.ReadFile(sqlfile)
	if err != nil {
		return err
	}
	if _, err := db.Exec(string(buf)); err != nil {
		return err
	}
	return nil
}
