package service

import (
	"github.com/2n1/proxysubscriber/app/cfg"
	"github.com/2n1/proxysubscriber/app/db"
	"github.com/2n1/proxysubscriber/app/defs/entity"
	"github.com/2n1/proxysubscriber/app/errs"
)

func AddNode(data db.InputData) (int64, *errs.Err) {
	if err := db.Open(); err != nil {
		return 0, errs.DBOpenFailed(err)
	}
	defer db.Close()
	if _, ok := data["name"]; !ok {
		return 0, errs.InvalidParameterWithMsgError("请输入节点名称", nil)
	}
	if _, ok := data["group_id"]; !ok {
		return 0, errs.InvalidParameterWithMsgError("请选择所属分组 ", nil)
	}
	count, err := db.CountNode("name=? AND group_id=?", data["name"], data["group_id"])
	if err != nil {
		return 0, errs.DBOperationFailed(err)
	}
	if count > 0 {
		return 0, errs.ExistsError("分组已存在")
	}

	id, err := db.AddNode(data)
	if err != nil {
		return 0, errs.DBOperationFailed(err)
	}
	return id, nil
}

func EditNode(id int64, data db.InputData) (int64, *errs.Err) {
	if err := db.Open(); err != nil {
		return 0, errs.DBOpenFailed(err)
	}
	defer db.Close()
	if _, ok := data["name"]; !ok {
		return 0, errs.InvalidParameterWithMsgError("请输入节点名称", nil)
	}
	if _, ok := data["group_id"]; !ok {
		return 0, errs.InvalidParameterWithMsgError("请选择所属分组 ", nil)
	}
	count, err := db.CountNode("name=? AND group_id=? AND id<>?", data["name"], data["group_id"], id)
	if err != nil {
		return 0, errs.DBOperationFailed(err)
	}
	if count > 0 {
		return 0, errs.ExistsError("节点已存在")
	}
	aff, err := db.EditNode(id, data)
	if err != nil {
		return 0, errs.DBOperationFailed(err)
	}
	return aff, nil
}
func DeleteNode(id int64) (int64, *errs.Err) {
	if err := db.Open(); err != nil {
		return 0, errs.DBOpenFailed(err)
	}
	defer db.Close()
	aff, err := db.DeleteNode(id)
	if err != nil {
		return 0, errs.DBOperationFailed(err)
	}
	return aff, nil
}
func FindNodeByID(id int64) (*entity.Node, *errs.Err) {
	if err := db.Open(); err != nil {
		return nil, errs.DBOpenFailed(err)
	}
	defer db.Close()
	node, err := db.FindNode(id)
	if err != nil {
		return nil, errs.DBOperationFailed(err)
	}
	return node, nil
}
func FindNodes(page int, condition string, params ...interface{}) (*db.Paginate, *errs.Err) {
	if err := db.Open(); err != nil {
		return nil, errs.DBOpenFailed(err)
	}
	defer db.Close()
	paginate, err := db.FindNodes(page, cfg.Cfg.PageSize, condition, params...)
	if err != nil {
		return nil, errs.DBOperationFailed(err)
	}
	return paginate, nil
}
func FindAllNodes(page int,groupID int64) (*db.Paginate, *errs.Err) {
	if err := db.Open(); err != nil {
		return nil, errs.DBOpenFailed(err)
	}
	defer db.Close()
	paginate, err := db.FindAllNodes(page, cfg.Cfg.PageSize,groupID)
	if err != nil {
		return nil, errs.DBOperationFailed(err)
	}
	return paginate, nil
}
