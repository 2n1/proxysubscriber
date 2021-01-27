package handler

import (
	"log"

	"github.com/2n1/proxysubscriber/app/errs"
	"github.com/gin-gonic/gin"
)

const (
	TypeJSON = iota
	TypeHTML
	TypeText
)

func w(f func(*gin.Context) *errs.Err, t int) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := f(c); err != nil {
			if err.Cause!=nil {
				log.Println(err.String())
			} else {
				log.Println(err.Error())
			}
			//if err.Code == code.UnAuth {
			//	redirect(c,"/login")
			//} else {
				errResponse(c, t, err)
			//}
			c.Abort()
		}
	}
}
func J(f func(*gin.Context) *errs.Err) func(*gin.Context) {
	return w(f, TypeJSON)
}
func H(f func(*gin.Context) *errs.Err) func(*gin.Context) {
	return w(f, TypeHTML)
}
func S(f func(*gin.Context) *errs.Err) func(*gin.Context) {
	return w(f, TypeText)
}
