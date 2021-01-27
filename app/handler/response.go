package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/2n1/proxysubscriber/app/cfg"
	"github.com/2n1/proxysubscriber/app/errs"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const flashSessionName = "flashMsg"
const authSessionName = "auth"

func response(c *gin.Context, data map[string]interface{}, template string, t int, err *errs.Err) {
	d := map[string]interface{}{
		"year":      time.Now().Year(),
		"site_name": cfg.Cfg.SiteName,
		"domain":    cfg.Cfg.BaseURL,
		"is_demo":   cfg.Cfg.IsDemo,
	}
	if err != nil {
		d["err"] = err
		switch t {
		case TypeHTML:
			c.HTML(http.StatusOK, template, d)
		case TypeJSON:
			c.JSON(http.StatusOK, d)
		}
		return
	}
	if data != nil && len(data) > 0 {
		for k, v := range data {
			d[k] = v
		}
	}
	switch t {
	case TypeHTML:
		c.HTML(http.StatusOK, template, d)
	case TypeJSON:
		c.JSON(http.StatusOK, d)
	}
}

func errResponse(c *gin.Context, t int, err *errs.Err) {
	response(c, nil, "err.html", t, err)
}
func htmlResponse(c *gin.Context, data map[string]interface{}, template string) {
	response(c, data, template, TypeHTML, nil)
}
func jsonResponse(c *gin.Context, data map[string]interface{}) {
	response(c, data, "", TypeJSON, nil)
}
func redirect(c *gin.Context, url string) {
	c.Redirect(http.StatusMovedPermanently, url)
}
func redirectWithFlash(c *gin.Context, url, msg string) {
	flash(c, msg)
	redirect(c, url)
}
func flash(c *gin.Context, msg string) {
	ses := sessions.Default(c)
	ses.Set(flashSessionName, msg)
	if err := ses.Save(); err != nil {
		log.Println("save session failed:", err)
	}
}
func getFlash(c *gin.Context) string {
	ses := sessions.Default(c)
	msg := ses.Get(flashSessionName)
	if nil == msg {
		return ""
	}
	ses.Set(flashSessionName, nil)
	ses.Save()
	ses.Delete(flashSessionName)
	if m, ok := msg.(string); ok {
		return m
	}
	return ""
}
func saveAuth(c *gin.Context, email string) {
	ses := sessions.Default(c)
	ses.Set(authSessionName, email)
	if err := ses.Save(); err != nil {
		log.Println("save session failed:", err)
	}
}
func getAuth(c *gin.Context) string {
	ses := sessions.Default(c)
	msg := ses.Get(authSessionName)
	if nil == msg {
		return ""
	}
	if v, ok := msg.(string); ok {
		return v
	}
	return ""
}
func clearAuth(c *gin.Context) {
	ses := sessions.Default(c)
	ses.Set(authSessionName, nil)
	ses.Save()
	ses.Delete(authSessionName)
	ses.Save()
}
func getPage(c *gin.Context) int {
	pageStr := c.Query("page")
	if page, err := strconv.Atoi(pageStr); err == nil {
		return page - 1
	}
	return 0
}
func getIDWithName(c *gin.Context, key string) int64 {
	s := c.Param(key)
	if v, err := strconv.ParseInt(s, 10, 64); err == nil {
		return v
	}
	return 0
}
func getID(c *gin.Context) int64 {
	return getIDWithName(c, "id")
}
