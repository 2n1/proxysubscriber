package db

import (
	"log"
	"strconv"
	"strings"

	"github.com/2n1/proxysubscriber/app/defs/entity"
)

const nodeFields = "name,group_id,node_type,server,port,passwd,cipher,sni,alter_id,ws_path,ws_host,cf_ip"

func AddNode(data InputData) (int64, error) {
	//defer Close()
	sqlStr := "INSERT INTO nodes (name,group_id,node_type,server,port,passwd,cipher,sni,alter_id,ws_path,ws_host,cf_ip) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(data.GetString("name"), data.GetInt("group_id"), data.GetString("node_type"), data.GetString("server"), data.GetInt("port"), data.GetString("passwd"), data.GetString("cipher"), data.GetString("sni"), data.GetInt("alter_id"), data.GetString("ws_path"), data.GetString("ws_host"), data.GetInt("cf_ip"))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	return result.LastInsertId()
}
func CountNode(condition string, params ...interface{}) (int, error) {
	//defer Close()
	sqlStrBuf := strings.Builder{}
	sqlStrBuf.WriteString("SELECT COUNT(*) FROM nodes WHERE (1=1)")
	if len(condition) > 0 {
		sqlStrBuf.WriteString(" AND (")
		sqlStrBuf.WriteString(condition)
		sqlStrBuf.WriteString(")")
	}
	stmt, err := db.Prepare(sqlStrBuf.String())
	if err != nil {
		return 0, err
	}
	var c int
	row := stmt.QueryRow(params...)
	if err := row.Scan(&c); err != nil {
		return 0, err
	}
	defer stmt.Close()
	return c, nil
}

func CountNodeByName(name string) (int, error) {
	return CountGroup("name=?", name)
}
func EditNode(id int64, data InputData) (int64, error) {
	//defer Close()
	var params []interface{}
	sqlStrBuf := strings.Builder{}
	sqlStrBuf.WriteString("UPDATE nodes SET ")
	fields := []string{"name", "group_id", "node_type", "server", "port", "passwd", "cipher", "sni", "alter_id", "ws_path", "ws_host", "cf_ip"}
	lastFieldsIdx := len(fields) - 1
	for idx, fn := range fields {
		sqlStrBuf.WriteString(fn)
		sqlStrBuf.WriteString("=")
		if v, ok := data[fn]; ok {
			sqlStrBuf.WriteString("?")
			params = append(params, v)
		} else {
			sqlStrBuf.WriteString(fn)
		}
		if idx < lastFieldsIdx {
			sqlStrBuf.WriteString(",")
		}
	}
	sqlStrBuf.WriteString(" WHERE id=?")
	log.Println(sqlStrBuf.String())
	stmt, err := db.Prepare(sqlStrBuf.String())
	if err != nil {
		return 0, err
	}
	params = append(params, id)
	result, err := stmt.Exec(params...)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	return result.RowsAffected()
}
func DeleteNode(id int64) (int64, error) {
	//defer Close()
	sqlStr := "DELETE FROM nodes WHERE id=?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	return result.RowsAffected()
}
func FindNode(id int64) (*entity.Node, error) {
	//defer Close()
	sqlStr := "SELECT id,name,group_id,node_type,server,port,passwd,cipher,sni,alter_id,ws_path,ws_host,cf_ip FROM nodes WHERE id=?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id)
	r := new(entity.Node)
	if err := row.Scan(&r.ID, &r.Name, &r.GroupID, &r.NodeType, &r.Server, &r.Port, &r.Password, &r.Cipher, &r.SNI, &r.AlterID, &r.WSPath, &r.WSHost, &r.CFIP); err != nil {
		return nil, err
	}
	defer stmt.Close()
	return r, nil
}
func FindNodes(page, pageSize int, condition string, params ...interface{}) (*Paginate, error) {
	//defer Close()
	count, err := CountNode(condition, params...)
	if err != nil {
		return nil, err
	}
	sqlStrBuf := strings.Builder{}
	sqlStrBuf.WriteString(`SELECT n.id,n.name,group_id,node_type,server,port,passwd,cipher,sni,alter_id,ws_path,ws_host,cf_ip 
							,g.name as group_name
							FROM nodes as n 
							INNER JOIN groups as g
							on n.group_id=g.id
							WHERE (1=1)
						`)
	if len(condition) > 0 {
		sqlStrBuf.WriteString(" AND (")
		sqlStrBuf.WriteString(condition)
		sqlStrBuf.WriteString(")")
	}
	sqlStrBuf.WriteString(" ORDER BY n.id DESC")
	if page >= 0 {
		sqlStrBuf.WriteString(" LIMIT ")
		sqlStrBuf.WriteString(strconv.Itoa(page * pageSize))
		sqlStrBuf.WriteString(",")
		sqlStrBuf.WriteString(strconv.Itoa(pageSize))
	}
	stmt, err := db.Prepare(sqlStrBuf.String())
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(params...)
	if err != nil {
		return nil, err
	}
	var rr []*entity.Node
	for rows.Next() {
		r := new(entity.Node)
		if err := rows.Scan(&r.ID, &r.Name, &r.GroupID, &r.NodeType, &r.Server, &r.Port, &r.Password, &r.Cipher, &r.SNI, &r.AlterID, &r.WSPath, &r.WSHost, &r.CFIP, &r.GroupName); err != nil {
			continue
		}
		rr = append(rr, r)
	}
	defer stmt.Close()
	p := NewPaginate(count, page, pageSize, rr)
	return p, nil
}

func FindAllNodes(page, pageSize int, groupID int64) (*Paginate, error) {
	condition := ""
	var params []interface{}
	if groupID > 0 {
		condition = "group_id=?"
		params = []interface{}{
			groupID,
		}
	}
	return FindNodes(page, pageSize, condition, params...)
}
