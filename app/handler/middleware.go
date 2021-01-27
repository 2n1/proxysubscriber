package handler

import (
	"github.com/2n1/proxysubscriber/app/errs"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(c *gin.Context) *errs.Err {
	auth := getAuth(c)
	if auth == "" {
		return errs.UnAuthError()
	}
	return nil
}
