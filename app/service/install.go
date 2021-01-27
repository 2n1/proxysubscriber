package service

import (
	"os"

	"github.com/2n1/proxysubscriber/app/cfg"
	"github.com/2n1/proxysubscriber/app/db"
	"github.com/2n1/proxysubscriber/app/defs/forms"
	"github.com/2n1/proxysubscriber/app/errs"
	"github.com/2n1/proxysubscriber/app/util"
)

func Install(lock string, form *forms.InstallForm) *errs.Err {
	 os.Remove(cfg.Cfg.DbFile)
	file, err := os.Create(lock)
	if err != nil {
		return errs.IoFailedWithMsgError("创建锁定文件失败", err)
	}
	defer file.Close()
	if err := db.Open(); err != nil {
		return errs.DBOpenFailed(err)
	}
	defer db.Close()
	if err := db.CreateTables(cfg.Cfg.SQLFile); err != nil {
		return errs.DBOperationWithMsgFailed("创建数据表失败", err)
	}
	pwd, e := util.Password(form.Password)
	if e != nil {
		return errs.InvalidParameterError(e)
	}
	if _, err := db.AddAuth(form.Email, pwd); err != nil {
		return errs.DBOperationWithMsgFailed("创建用户失败", err)
	}
	data := map[string]interface{}{
		"cu_ip":    "104.22.65.143",
		"cu_label": "CU",
		"ct_ip":    "104.18.202.108",
		"ct_label": "CT",
		"cm_ip":    "104.19.103.16",
		"cm_label": "CM",
	}
	if err := db.UpdateCfips(data); err != nil {
		return errs.DBOperationWithMsgFailed("设置优选IP失败", err)
	}
	return nil
}
