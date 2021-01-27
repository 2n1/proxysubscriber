package service

import (
	"github.com/2n1/proxysubscriber/app/db"
	"github.com/2n1/proxysubscriber/app/errs"
)

func Total() (groupCount, nodeCount int, err *errs.Err) {
	if err := db.Open(); err != nil {
		return 0, 0, errs.DBOpenFailed(err)
	}
	defer db.Close()
	groupCount, e := db.CountGroup("")
	if e != nil {
		return 0, 0, errs.DBOperationWithMsgFailed("统计分组失败", e)
	}
	nodeCount, e = db.CountNode("")
	if e != nil {
		return 0, 0, errs.DBOperationWithMsgFailed("统计节点失败", e)
	}
	return groupCount, nodeCount, nil
}
