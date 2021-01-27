package cfg

import (
	"encoding/json"
	"io/ioutil"
)

type sessionConfig struct {
	SecretKey string `json:"secret_key"`
	Name      string `json:"name"`
}
type config struct {
	Addr string `json:"addr"`
	PageSize int    `json:"page_size"`
	DbFile   string `json:"db_file"`
	SQLFile  string `json:"sql_file"`
	SiteName string `json:"site_name"`
	BaseURL  string `json:"base_url"`
	IsDemo   bool   `json:"is_demo"`
	Mode string `json:"mode"`
	Session  *sessionConfig
}

var Cfg *config

func InitFrom(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	Cfg = new(config)
	if err := json.Unmarshal(buf, Cfg); err != nil {
		return err
	}
	return nil
}
func Init() error {
	return InitFrom("./config.json")
}
