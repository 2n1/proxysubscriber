package service

import (
	"github.com/2n1/proxysubscriber/app/cfg"
	"github.com/2n1/proxysubscriber/app/db"
	"github.com/2n1/proxysubscriber/app/defs/entity"
	"github.com/2n1/proxysubscriber/app/errs"
	"github.com/2n1/proxysubscriber/app/util"
)

func AddGroup(name string) (int64, *errs.Err) {
	if err := db.Open(); err != nil {
		return 0, errs.DBOpenFailed(err)
	}
	defer db.Close()
	countGroupByName, err := db.CountGroupByName(name)
	if err != nil {
		return 0, errs.DBOperationFailed(err)
	}
	if countGroupByName > 0 {
		return 0, errs.ExistsError("分组已存在")
	}
	url:=util.GenID()
	id, err := db.AddGroup(name,url)
	if err != nil {
		return 0, errs.DBOperationFailed(err)
	}
	return id, nil
}

func EditGroup(id int64, name string) (int64, *errs.Err) {
	if err := db.Open(); err != nil {
		return 0, errs.DBOpenFailed(err)
	}
	defer db.Close()
	countGroup, err := db.CountGroup("name=? AND id<>?", name, id)
	if err != nil {
		return 0, errs.DBOperationFailed(err)
	}
	if countGroup > 0 {
		return 0, errs.ExistsError("分组已存在")
	}
	aff, err := db.EditGroup(id, name)
	if err != nil {
		return 0, errs.DBOperationFailed(err)
	}
	return aff, nil
}
func RefreshGroupURL(id int64, ) (string, *errs.Err) {
	if err := db.Open(); err != nil {
		return "", errs.DBOpenFailed(err)
	}
	defer db.Close()
	u:=util.GenID()
	_, err := db.RefreshGroupURL(id, u)
	if err != nil {
		return "", errs.DBOperationFailed(err)
	}
	return u, nil
}
func DeleteGroup(id int64) (int64, *errs.Err) {
	if err := db.Open(); err != nil {
		return 0, errs.DBOpenFailed(err)
	}
	defer db.Close()
	aff, err := db.DeleteGroup(id)
	if err != nil {
		return 0, errs.DBOperationFailed(err)
	}
	return aff, nil
}
func DeleteGroupWithNode(id int64) (int64, int64, *errs.Err) {
	if err := db.Open(); err != nil {
		return 0, 0, errs.DBOpenFailed(err)
	}
	defer db.Close()
	gAff, nAff, err := db.DeleteGroupWithNodes(id)
	if err != nil {
		return 0, 0, errs.DBOperationFailed(err)
	}
	return gAff, nAff, nil
}
func FindGroupByID(id int64) (*entity.Group, *errs.Err) {
	if err := db.Open(); err != nil {
		return nil, errs.DBOpenFailed(err)
	}
	defer db.Close()
	group, err := db.FindGroup(id)
	if err != nil {
		return nil, errs.DBOperationFailed(err)
	}
	return group, nil
}
func FindGroups(page int, condition string, params ...interface{}) (*db.Paginate, *errs.Err) {
	if err := db.Open(); err != nil {
		return nil, errs.DBOpenFailed(err)
	}
	defer db.Close()
	paginate, err := db.FindGroups(page, cfg.Cfg.PageSize, condition, params)
	if err != nil {
		return nil, errs.DBOperationFailed(err)
	}
	return paginate, nil
}
func FindAllGroups(page int) (*db.Paginate, *errs.Err) {
	if err := db.Open(); err != nil {
		return nil, errs.DBOpenFailed(err)
	}
	defer db.Close()
	paginate, err := db.FindAllGroups(page, cfg.Cfg.PageSize)
	if err != nil {
		return nil, errs.DBOperationFailed(err)
	}
	return paginate, nil
}
