package service

import (
	"github.com/2n1/proxysubscriber/app/db"
	"github.com/2n1/proxysubscriber/app/defs/entity"
	"github.com/2n1/proxysubscriber/app/errs"
	"github.com/2n1/proxysubscriber/app/util"
)

func AddAuth(email, password string) (int64, *errs.Err) {
	if err := db.Open(); err != nil {
		return 0, errs.DBOpenFailed(err)
	}
	defer db.Close()
	hashedPassword, err := util.Password(password)
	if err != nil {
		return 0, errs.InvalidParameterWithMsgError("计算密码Hash失败", err)
	}
	id, err := db.AddAuth(email, hashedPassword)
	if err != nil {
		return 0, errs.DBOperationWithMsgFailed("添加授权失败", err)
	}
	return id, nil
}
func EditAuth(email, password string) (int64, *errs.Err) {
	if err := db.Open(); err != nil {
		return 0, errs.DBOpenFailed(err)
	}
	defer db.Close()
	hashedPassword, err := util.Password(password)
	if err != nil {
		return 0, errs.InvalidParameterWithMsgError("计算密码Hash失败", err)
	}
	aff, err := db.EditAuth(email, hashedPassword)
	if err != nil {
		return 0, errs.DBOperationWithMsgFailed("更新授权失败", err)
	}
	return aff, nil
}
func FindAuth() (*entity.Authorization, *errs.Err) {
	if err := db.Open(); err != nil {
		return nil, errs.DBOpenFailed(err)
	}
	defer db.Close()
	auth, err := db.FindAuth()
	if err != nil {
		return nil, errs.DBOperationWithMsgFailed("获取授权失败", err)
	}
	return auth, nil
}
