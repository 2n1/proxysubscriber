package db

import "github.com/2n1/proxysubscriber/app/defs/entity"

func AddAuth(email, password string) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO auths (email,passwd) VALUES (?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(email, password)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
func EditAuth(email, password string) (int64, error) {
	stmt, err := db.Prepare("UPDATE auths SET email=?,passwd=?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(email, password)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
func FindAuth() (*entity.Authorization, error) {
	stmt, err := db.Prepare("SELECT email,passwd FROM auths LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow()
	var auth entity.Authorization
	if err := row.Scan(&auth.Email, &auth.Password); err != nil {
		return nil, err
	}
	return &auth, nil
}
func UpdateAuth(email, password string) (int64, error) {
	return 0, nil
}
