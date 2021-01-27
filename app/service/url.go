package service

import (
	"github.com/2n1/proxysubscriber/app/db"
	"github.com/2n1/proxysubscriber/app/defs/entity"
	"github.com/2n1/proxysubscriber/app/errs"
)

func GroupURL(groupID int64) (string, *errs.Err) {
	if err := db.Open(); err != nil {
		return "", errs.DBOpenFailed(err)
	}
	defer db.Close()
	urlID, err := db.GenURL(groupID)
	if err != nil {
		return "", errs.DBOperationWithMsgFailed("生成分组URL失败", err)
	}
	return urlID, nil
}

func FindGroupURLIfExists(groupID int64) (*entity.URL, *errs.Err) {
	if err := db.Open(); err != nil {
		return nil, errs.DBOpenFailed(err)
	}
	defer db.Close()
	url, err := db.FindURLIfExists(groupID)
	if err != nil {
		return nil, errs.DBOperationWithMsgFailed("获取分组URL失败", err)
	}
	return url, nil
}
func GetGroupIDFromURL(urlID string) (int64, *errs.Err) {
	if err := db.Open(); err != nil {
		return 0, errs.DBOpenFailed(err)
	}
	defer db.Close()
	groupID, err := db.GetGroupIDFromURL(urlID)
	if err != nil {
		return 0, errs.DBOperationWithMsgFailed("获取分组ID失败", err)
	}
	return groupID, nil
}
