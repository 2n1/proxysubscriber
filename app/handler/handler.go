package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/2n1/proxysubscriber/app/cfg"
	"github.com/2n1/proxysubscriber/app/defs/entity"
	"github.com/2n1/proxysubscriber/app/defs/forms"
	"github.com/2n1/proxysubscriber/app/errs"
	"github.com/2n1/proxysubscriber/app/service"
	"github.com/2n1/proxysubscriber/app/util"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) *errs.Err {
	groupCount, nodeCount, err := service.Total()
	if err!=nil{
		return err
	}
	data := gin.H{
		"groupCount":groupCount,
		"nodeCount":nodeCount,
	}
	htmlResponse(c, data, "index.html")
	return nil
}

func Group(c *gin.Context) *errs.Err {
	msg := getFlash(c)
	page := getPage(c)
	data := gin.H{
		"title": "分组列表",
		"msg":   msg,
		"page":  page,
	}
	paginate, err := service.FindAllGroups(page)
	if err != nil {
		return err
	}
	data["paginate"] = paginate

	htmlResponse(c, data, "group.html")
	return nil
}

func GroupInput(c *gin.Context) *errs.Err {
	editIdStr := c.Param("id")
	var editId int64
	var err error
	isEdit := strings.Index(c.Request.RequestURI, "/group/edit") > -1
	if editId, err = strconv.ParseInt(editIdStr, 10, 64); err != nil {
		if isEdit {
			return errs.InvalidParameterError(err)
		}
	}
	title := "添加分组"
	action := "add"
	if isEdit {
		title = "修改分组"
		action = fmt.Sprintf("edit/%d", editId)
	}

	data := gin.H{
		"title":  title,
		"action": action,
	}
	if c.Request.Method == http.MethodGet {
		if isEdit {
			if entity, err := service.FindGroupByID(editId); err != nil {
				return errs.DBOperationWithMsgFailed("获取分组失败，请检查是否存在", err)
			} else {
				data["entity"] = entity
			}

		}
		htmlResponse(c, data, "group-input.html")
		return nil
	}
	if c.Request.Method == http.MethodPost {
		var form entity.Group
		if err := c.ShouldBind(&form); err != nil {
			return errs.InvalidParameterError(err)
		}
		if isEdit {
			if _, err := service.EditGroup(editId, form.Name); err != nil {
				return err
			}
			redirectWithFlash(c, "/man/group", "分组修改成功")
			return nil

		}
		if _, err := service.AddGroup(form.Name); err != nil {
			return err
		}
		redirectWithFlash(c, "/man/group", "分组添加成功")
		return nil
	}
	return nil
}

func GroupDel(c *gin.Context) *errs.Err {
	idStr := c.Param("id")
	var id int64
	var err error
	if id, err = strconv.ParseInt(idStr, 10, 64); err != nil {
		return errs.InvalidParameterWithMsgError("参数错误", err)
	}
	if _, _, err := service.DeleteGroupWithNode(id); err != nil {
		return errs.DBOperationWithMsgFailed("删除分组失败", err)
	}
	redirectWithFlash(c, "/man/group", "分组删除成功")
	return nil
}
func Node(c *gin.Context) *errs.Err {
	page := getPage(c)
	msg := getFlash(c)
	data := gin.H{
		"title": "节点列表",
		"msg":   msg,
		"page":  page,
	}
	groupIDStr, ok := c.GetQuery("group")
	var groupID int64
	if ok {
		if v, err := strconv.ParseInt(groupIDStr, 10, 64); err == nil {
			groupID = v
		}
	}
	data["group"] = groupID
	paginate, err := service.FindAllNodes(page, groupID)
	if err != nil {
		return err
	}
	data["paginate"] = paginate
	htmlResponse(c, data, "node.html")
	return nil
}
func NodeDel(c *gin.Context)*errs.Err{
	id:=getID(c)
	if _, err := service.DeleteNode(id);err!=nil{
		return err
	}
	redirectWithFlash(c,"/man/node", "节点删除成功")
	return nil
}
func NodeInput(c *gin.Context) *errs.Err {
	groups, e := service.FindAllGroups(-1)
	if e != nil {
		return e
	}
	editIdStr := c.Param("id")
	var editId int64
	var err error
	isEdit := strings.Index(c.Request.RequestURI, "/node/edit") > -1
	if editId, err = strconv.ParseInt(editIdStr, 10, 64); err != nil {
		if isEdit {
			return errs.InvalidParameterError(err)
		}
	}
	title := "添加节点"
	action := "add"
	if isEdit {
		title = "修改节点"
		action = fmt.Sprintf("edit/%d", editId)
	}

	data := gin.H{
		"title":  title,
		"action": action,
		"groups": groups.Data,
		"isEdit": isEdit,
	}
	if c.Request.Method == http.MethodGet {
		if isEdit {
			if entity, err := service.FindNodeByID(editId); err != nil {
				return errs.DBOperationWithMsgFailed("获取节点失败，请检查是否存在", err)
			} else {
				data["entity"] = entity
			}
		}
		htmlResponse(c, data, "node-input.html")
		return nil
	}
	if c.Request.Method == http.MethodPost {
		var form entity.Node
		if err := c.ShouldBind(&form); err != nil {
			return errs.InvalidParameterError(err)
		}
		data := map[string]interface{}{
			"name":      form.Name,
			"group_id":  form.GroupID,
			"node_type": form.NodeType,
			"server":    form.Server,
			"port":      form.Port,
			"passwd":    form.Password,
			"cipher":    form.Cipher,
			"sni":       form.SNI,
			"alter_id":  form.AlterID,
			"ws_path":   form.WSPath,
			"ws_host":   form.WSHost,
			"cf_ip":     form.CFIP,
		}
		if isEdit {
			if _, err := service.EditNode(editId, data); err != nil {
				return err
			}
			redirectWithFlash(c, "/man/node", "节点修改成功")
			return nil
		}
		if _, err := service.AddNode(data); err != nil {
			return err
		}
		redirectWithFlash(c, "/man/node", "节点添加成功")
		return nil
	}
	return nil
}
func GroupUrl(c *gin.Context) *errs.Err {
	groupID := getID(c)
	if c.Request.Method == http.MethodGet {
		g, err := service.FindGroupByID(groupID)
		if err != nil {
			return err
		}
		data := map[string]interface{}{
			"url":     g.Url,
			"groupID": groupID,
		}
		htmlResponse(c, data, "url.html")
		return nil
	}

	return nil
}
func RefreshGroupUrl(c *gin.Context) *errs.Err {
	groupID := getID(c)
	if c.Request.Method == http.MethodPost {
		uid, err := service.RefreshGroupURL(groupID)
		if err != nil {
			return err
		}
		jsonResponse(c, gin.H{"url": uid})
		return nil
	}
	return nil
}
func Cfips(c *gin.Context) *errs.Err {
	if c.Request.Method == http.MethodGet {
		cfips, err := service.GetCfips()
		if err != nil {
			return err
		}
		data := gin.H{
			"cf":  cfips,
			"msg": getFlash(c),
		}
		htmlResponse(c, data, "cfip.html")
		return nil
	}
	if c.Request.Method == http.MethodPost {
		var form entity.CloudFlareIP
		if err := c.ShouldBind(&form); err != nil {
			return errs.InvalidParameterWithMsgError("参数错误", err)
		}
		data := map[string]interface{}{
			"cu_ip":    form.ChinaUnicomIP,
			"cu_label": form.ChinaUnicomLable,
			"ct_ip":    form.ChinaTelecomIP,
			"ct_label": form.ChinaTelecomLable,
			"cm_ip":    form.ChinaMobileIP,
			"cm_label": form.ChinaMobileLable,
		}
		if err := service.UpdateCfips(data); err != nil {
			return err
		}
		redirectWithFlash(c, c.Request.RequestURI, "优选IP更新成功")
	}
	return nil
}
func UpdateAuth(c *gin.Context) *errs.Err {
	if c.Request.Method == http.MethodGet {
		auth, err := service.FindAuth()
		if err != nil {
			return err
		}
		data := gin.H{
			"email": auth.Email,
			"msg":   getFlash(c),
		}
		htmlResponse(c, data, "update-auth.html")
		return nil
	}
	if c.Request.Method == http.MethodPost {
		if cfg.Cfg.IsDemo {
			return errs.IsDemoError()
		}
		var form forms.UpdateAuthForm
		if err := c.ShouldBind(&form); err != nil {
			return errs.InvalidParameterWithMsgError("输入错误", err)
		}
		if form.NPassword != form.RPassword {
			return errs.InvalidParameterWithMsgError("两次输入的密码不一致", nil)
		}
		auth, err := service.FindAuth()
		if err != nil {
			return err
		}
		if !util.VerifyPassword(form.Password, auth.Password) {
			return errs.InvalidParameterWithMsgError("现用密码错误", nil)
		}
		if _, err := service.EditAuth(form.Email, form.NPassword); err != nil {
			return err
		}
		redirectWithFlash(c, "/man/auth", "修改成功")
		return nil
	}
	return nil
}
func Logout(c *gin.Context) *errs.Err {
	clearAuth(c)
	redirect(c,"/login")
	return nil
}