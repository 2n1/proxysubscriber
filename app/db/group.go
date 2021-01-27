package db

import (
	"strconv"
	"strings"

	"github.com/2n1/proxysubscriber/app/defs/entity"
)

func AddGroup(name,url string) (int64, error) {
	//defer Close()
	sqlStr := "INSERT INTO groups (name,url) VALUES (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(name,url)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	return result.LastInsertId()
}
func CountGroup(condition string, params ...interface{}) (int, error) {
	//defer Close()
	sqlStrBuf := strings.Builder{}
	sqlStrBuf.WriteString("SELECT COUNT(*) FROM groups WHERE (1=1)")
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

func CountGroupByName(name string) (int, error) {
	return CountGroup("name=?", name)
}
func EditGroup(id int64, name string) (int64, error) {
	//defer Close()
	sqlStr := "UPDATE groups SET name =? WHERE id=?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(name, id)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	return result.RowsAffected()
}
func RefreshGroupURL(id int64, url string) (int64, error) {
	//defer Close()
	sqlStr := "UPDATE groups SET url =? WHERE id=?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(url, id)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	return result.RowsAffected()
}
func DeleteGroup(id int64) (int64, error) {
	//defer Close()
	sqlStr := "DELETE FROM groups WHERE id=?"
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
func DeleteGroupWithNodes(id int64) (int64, int64, error) {
	//defer Close()
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	sqlStr := "DELETE FROM groups WHERE id=?"
	stmt, err := tx.Prepare(sqlStr)
	if err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	result, err := stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	defer stmt.Close()

	sqlStr = "DELETE FROM nodes WHERE group_id=?"
	stmtNode, err := tx.Prepare(sqlStr)
	if err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	resultNode, err := stmtNode.Exec(id)
	if err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	defer stmtNode.Close()
	var affGroup, affNode int64
	if affGroup, err = result.RowsAffected(); err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	if affNode, err = resultNode.RowsAffected(); err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	tx.Commit()
	return affGroup, affNode, nil
}
func FindGroup(id int64) (*entity.Group, error) {
	//defer Close()
	sqlStr := "SELECT id,name,url FROM groups WHERE id=?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id)
	r := new(entity.Group)
	if err := row.Scan(&r.ID, &r.Name,&r.Url); err != nil {
		return nil, err
	}
	defer stmt.Close()
	return r, nil
}
func FindGroups(page, pageSize int, condition string, params ...interface{}) (*Paginate, error) {
	//defer Close()
	count, err := CountGroup(condition, params...)
	if err != nil {
		return nil, err
	}
	sqlStrBuf := strings.Builder{}
	sqlStrBuf.WriteString("SELECT id,name,url FROM groups WHERE (1=1)")
	if len(condition) > 0 {
		sqlStrBuf.WriteString(" AND (")
		sqlStrBuf.WriteString(condition)
		sqlStrBuf.WriteString(")")
	}
	sqlStrBuf.WriteString(" ORDER BY id DESC")
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
	var rr []*entity.Group
	for rows.Next() {
		r := new(entity.Group)
		if err := rows.Scan(&r.ID, &r.Name,&r.Url); err != nil {
			continue
		}
		rr = append(rr, r)
	}
	defer stmt.Close()
	return NewPaginate(count, page, pageSize, rr), nil
}

func FindAllGroups(page, pageSize int) (*Paginate, error) {
	return FindGroups(page, pageSize, "")
}
