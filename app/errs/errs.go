package errs

import (
	"fmt"

	"github.com/2n1/proxysubscriber/app/errs/code"
)

type Err struct {
	Code  int
	Msg   string
	Cause error
}

func (e *Err) Error() string {
	return e.Msg
}
func (e *Err) String() string {
	if e.Cause == nil {
		return e.Msg
	}
	return fmt.Sprintf("%s: %s", e.Msg, e.Cause.Error())
}

func From(code int, msg string, err error) *Err {
	return &Err{
		Code:  code,
		Msg:   msg,
		Cause: err,
	}
}

func New(code int, msg string) *Err {
	return From(code, msg, nil)
}

func DBOpenFailed(err error) *Err {
	return From(code.DBOpenFailed, "数据库连接失败", err)
}
func ExistsError(msg string) *Err {
	return New(code.Exists, msg)
}
func DBOperationFailed(err error) *Err {
	return DBOperationWithMsgFailed("数据库操作失败", err)
}
func DBOperationWithMsgFailed(msg string, err error) *Err {
	return From(code.DBOperationFailed, msg, err)
}
func InvalidParameterWithMsgError(msg string, err error) *Err {
	return From(code.InvalidParameter, msg, err)
}
func InvalidParameterError(err error) *Err {
	return InvalidParameterWithMsgError("参数错误", err)
}
func IsDemoError() *Err {
	return New(code.IsDemo, "演示站点不允许该操作")
}
func InvalidAuthError() *Err {
	return New(code.InvalidAuth, "Email或密码错误")
}
func UnAuthError() *Err {
	return New(code.UnAuth, "请登录")
}
func InstalledError() *Err {
	return New(code.Installed, "请不要重复安装")
}
func IoFailedError(err error) *Err {
	return IoFailedWithMsgError("io操作失败",err)
}
func IoFailedWithMsgError(msg string,err error) *Err {
	return From(code.IoOperationFailed,msg, err)
}