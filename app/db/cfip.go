package db

import (
	"errors"
	"strings"

	"github.com/2n1/proxysubscriber/app/defs/entity"
)

func addCfips(data InputData) error {
	//defer Close()
	if data == nil || len(data) == 0 {
		return errors.New("请给定要更新的数据")
	}
	var params []interface{}
	sqlStr := `INSERT INTO cfips (cu_ip,cu_label,ct_ip,ct_label,cm_ip,cm_label)
		VALUES(?,?,?,?,?,?)`

	params = append(params, data.GetString("cu_ip"))
	params = append(params, data.GetString("cu_label"))
	params = append(params, data.GetString("ct_ip"))
	params = append(params, data.GetString("ct_label"))
	params = append(params, data.GetString("cm_ip"))
	params = append(params, data.GetString("cm_label"))

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(params...)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
func editCfips(data InputData) error {
	//defer Close()
	if data == nil || len(data) == 0 {
		return errors.New("请给定要更新的数据")
	}
	var params []interface{}
	sqlStrBuf := strings.Builder{}
	sqlStrBuf.WriteString("UPDATE cfips SET ")
	if v, ok := data["cu_ip"]; ok {
		sqlStrBuf.WriteString("cu_ip=?,")
		params = append(params, v)
	} else {
		sqlStrBuf.WriteString("cu_ip=cu_ip,")
	}
	if v, ok := data["cu_label"]; ok {
		sqlStrBuf.WriteString("cu_label=?,")
		params = append(params, v)
	} else {
		sqlStrBuf.WriteString("cu_label=cu_label,")
	}
	if v, ok := data["ct_ip"]; ok {
		sqlStrBuf.WriteString("ct_ip=?,")
		params = append(params, v)
	} else {
		sqlStrBuf.WriteString("ct_ip=ct_ip,")
	}
	if v, ok := data["ct_label"]; ok {
		sqlStrBuf.WriteString("ct_label=?,")
		params = append(params, v)
	} else {
		sqlStrBuf.WriteString("ct_label=ct_label,")
	}
	if v, ok := data["cm_ip"]; ok {
		sqlStrBuf.WriteString("cm_ip=?,")
		params = append(params, v)
	} else {
		sqlStrBuf.WriteString("cm_ip=cm_ip,")
	}
	if v, ok := data["cm_label"]; ok {
		sqlStrBuf.WriteString("cm_label=?")
		params = append(params, v)
	} else {
		sqlStrBuf.WriteString("cm_label=cm_label")
	}
	stmt, err := db.Prepare(sqlStrBuf.String())
	if err != nil {
		return err
	}
	_, err = stmt.Exec(params...)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func UpdateCfips(data InputData) error {
	//defer Close()
	stmt, err := db.Prepare("SELECT count(*) FROM cfips")
	if err != nil {
		return err
	}
	defer stmt.Close()
	var c int64
	row := stmt.QueryRow()
	if err := row.Scan(&c); err != nil {
		return err
	}
	if c == 0 {
		return addCfips(data)
	}
	return editCfips(data)
}
func GetCfips() (*entity.CloudFlareIP, error) {
	//defer Close()
	stmt, err := db.Prepare("SELECT cu_ip,cu_label,ct_ip,ct_label,cm_ip,cm_label FROM cfips LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var c = new(entity.CloudFlareIP)
	row := stmt.QueryRow()
	if err := row.Scan(&c.ChinaUnicomIP, &c.ChinaUnicomLable, &c.ChinaTelecomIP, &c.ChinaTelecomLable, &c.ChinaMobileIP, &c.ChinaMobileLable); err != nil {
		return nil, err
	}
	return c, nil
}
