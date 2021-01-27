package db

import (
	"errors"

	"github.com/2n1/proxysubscriber/app/defs/entity"
	"github.com/2n1/proxysubscriber/app/util"
)

func GenURL(groupID int64) (string, error) {
	tx, err := db.Begin()
	if err != nil {
		return "", err
	}
	stmtDel, err := tx.Prepare("DELETE FROM urls WHERE group_id=?")
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer stmtDel.Close()
	if _, err := stmtDel.Exec(groupID); err != nil {
		tx.Rollback()
		return "", err
	}
	stmt, err := tx.Prepare("INSERT INTO urls(id,group_id) VALUES(?,?)")
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer stmt.Close()
	id := util.GenID()
	if _, err := stmt.Exec(id, groupID); err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()
	return id, nil
}

func FindURLIfExists(groupID int64) (*entity.URL, error) {
	stmt, err := db.Prepare("SELECT id FROM urls WHERE group_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(groupID)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("获取数据失败")
	}
	for rows.Next() {
		var u entity.URL
		if err := rows.Scan(&u.ID); err != nil {
			return nil, err
		}
		return &u, nil
	}

	return nil, nil
}
func GetGroupIDFromURL(urlID string) (int64, error) {
	stmt, err := db.Prepare("SELECT id FROM groups WHERE url=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(urlID)
	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
