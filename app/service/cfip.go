package service

import (
	"github.com/2n1/proxysubscriber/app/db"
	"github.com/2n1/proxysubscriber/app/defs/entity"
	"github.com/2n1/proxysubscriber/app/errs"
)

func GetCfips() (*entity.CloudFlareIP, *errs.Err) {
	if err := db.Open(); err != nil {
		return nil, errs.DBOpenFailed(err)
	}
	defer db.Close()
	cfips, err := db.GetCfips()
	if err != nil {
		return nil, errs.DBOpenFailed(err)
	}
	return cfips, nil
}

func UpdateCfips(data db.InputData) *errs.Err {
	if err := db.Open(); err != nil {
		return errs.DBOpenFailed(err)
	}
	defer db.Close()
	if err := db.UpdateCfips(data); err != nil {
		return errs.DBOperationWithMsgFailed("更新优选IP失败", err)
	}
	return nil
}
