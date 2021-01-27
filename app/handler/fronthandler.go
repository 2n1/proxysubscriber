package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/2n1/proxysubscriber/app/defs/entity"
	"github.com/2n1/proxysubscriber/app/defs/forms"
	"github.com/2n1/proxysubscriber/app/defs/subscribe"
	"github.com/2n1/proxysubscriber/app/errs"
	"github.com/2n1/proxysubscriber/app/service"
	"github.com/2n1/proxysubscriber/app/util"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

const installLockFile = "./install.lock"

func FrontendIndex(c *gin.Context) *errs.Err {
	u := "/man"
	if !util.IsExists(installLockFile) {
		u = "/install"
	}
	redirect(c, u)
	return nil
}
func Install(c *gin.Context) *errs.Err {
	if util.IsExists(installLockFile) {
		return errs.InstalledError()
	}
	if c.Request.Method == http.MethodGet {
		htmlResponse(c, nil, "install.html")
		return nil
	}
	if c.Request.Method == http.MethodPost {
		var form forms.InstallForm
		if err:=c.ShouldBind(&form);err!=nil{
			return errs.InvalidParameterError(err)
		}
		if err:=service.Install(installLockFile,&form);err!=nil {
			return err
		}
		redirect(c,"/login")
		return nil
	}
	return nil
}
func LoginHandler(c *gin.Context) *errs.Err {
	if c.Request.Method == http.MethodGet {
		htmlResponse(c, nil, "login.html")
		return nil
	}
	if c.Request.Method == http.MethodPost {
		var form forms.LoginForm
		if err := c.ShouldBind(&form); err != nil {
			return errs.InvalidParameterWithMsgError("输入有误", err)
		}
		auth, err := service.FindAuth()
		if err != nil {
			return err
		}
		if auth.Email != form.Email || !util.VerifyPassword(form.Password, auth.Password) {
			return errs.InvalidAuthError()
		}
		saveAuth(c, auth.Email)
		redirect(c, "/man")
		return nil
	}
	return nil
}

func UrlHandler(c *gin.Context) *errs.Err {
	urlID := c.Param("id")
	client := c.Query("c")
	cfStr := c.Query("cf")
	cfOnlyStr := c.Query("co")
	cfOnly := cfOnlyStr == "1"
	cf := strings.Split(cfStr, ",")
	if client == "" {
		return errs.InvalidParameterError(nil)
	}
	if len(cf) == 0 && cfOnly {
		return errs.InvalidParameterError(nil)
	}

	groupID, err := service.GetGroupIDFromURL(urlID)
	if err != nil {
		return err
	}

	paginate, err := service.FindNodes(-1, "group_id=?", groupID)
	if err != nil {
		return err
	}
	var cfips *entity.CloudFlareIP
	if len(cf) > 0 {
		var e *errs.Err
		cfips, e = service.GetCfips()
		if e != nil {
			return e
		}
	}
	var nodes []*entity.Node
	if v, ok := paginate.Data.([]*entity.Node); ok {
		nodes = v
	}
	var sb strings.Builder
	if client == "v2ray" {
		for _, node := range nodes {
			switch node.NodeType {
			case "vmess":
				sb.WriteString(vmessURL(node, cf, cfips, cfOnly))
			case "trojan":
				sb.WriteString(trojanURL(node))
			case "ss":
				sb.WriteString(ssURL(node))
			default:
				continue
			}
			sb.WriteString("\n")
		}
		c.String(200, util.Base64Encode(sb.String()))
		return nil

	}
	if client == "clash" {
		sb.WriteString(clashYaml(nodes, cf, cfips, cfOnly))
	}
	c.String(200, sb.String())
	return nil
}
func vmessURL(node *entity.Node, cf []string, cfips *entity.CloudFlareIP, cfOnly bool) string {
	type server struct {
		servername string
		label      string
	}
	var servers []*server
	for _, c := range cf {
		if strings.TrimSpace(c) == "" {
			continue
		}
		s := new(server)
		switch c {
		case "cu":
			if cfips.ChinaUnicomIP == "" {
				continue
			}
			s.servername = cfips.ChinaUnicomIP
			s.label = cfips.ChinaUnicomLable
		case "ct":
			if cfips.ChinaTelecomIP == "" {
				continue
			}
			s.servername = cfips.ChinaTelecomIP
			s.label = cfips.ChinaTelecomLable
		case "cm":
			if cfips.ChinaMobileIP == "" {
				continue
			}
			s.servername = cfips.ChinaMobileIP
			s.label = cfips.ChinaMobileLable
		default:
			continue
		}
		servers = append(servers, s)
	}
	var sb []string
	for _, srv := range servers {
		sni := node.SNI
		if sni == "" {
			sni = node.Server
		}
		data := map[string]interface{}{
			"add":  srv.servername,
			"aid":  node.AlterID,
			"host": sni,
			"id":   node.Password,
			"net":  "ws",
			"path": node.WSPath,
			"port": node.Port,
			"tls":  "tls",
			"type": "none",
			"v":    2,
			"ps":   node.Name + "_" + srv.label,
		}

		bytes, _ := json.Marshal(data)
		sb = append(sb, "vmess://"+util.Base64Encode(string(bytes)))
	}
	if !cfOnly || len(servers) < 1 {
		data := map[string]interface{}{
			"add":  node.Server,
			"aid":  node.AlterID,
			"host": node.SNI,
			"id":   node.Password,
			"net":  "ws",
			"path": node.WSPath,
			"port": node.Port,
			"tls":  "tls",
			"type": "none",
			"v":    2,
			"ps":   node.Name,
		}
		bytes, _ := json.Marshal(data)
		sb = append(sb, "vmess://"+util.Base64Encode(string(bytes)))
	}

	return strings.Join(sb, "\n")
}
func trojanURL(node *entity.Node) string {
	return fmt.Sprintf("trojan://%s@%s:%d?allowInsecure=0&allowInsecureHostname=0&allowInsecureCertificate=0&sessionTicket=0&tfo=0#%s",
		node.Password,
		node.Server,
		node.Port,
		url.QueryEscape(node.Name))
}
func ssURL(node *entity.Node) string {
	return fmt.Sprintf("ss://%s@%s:%d/#%s",
		util.Base64Encode(node.Cipher+":"+node.Password),
		node.Server,
		node.Port,
		url.QueryEscape(node.Name),
	)
}
func clashYaml(nodes []*entity.Node, cf []string, cfips *entity.CloudFlareIP, cfOnly bool) string {
	ns := &subscribe.ClashSubscribe{
		Proxies: []subscribe.NodeInfo{},
		ProxyGroups: []*subscribe.ClashSubscribeProxyItem{
			{
				Name:    "Proxy",
				Type:    "select",
				Proxies: nil,
			},
		},
		Rules: []string{
			"MATCH,Proxy",
		},
	}
	type server struct {
		servername string
		label      string
	}
	var servers []*server
	for _, c := range cf {
		if strings.TrimSpace(c) == "" {
			continue
		}
		s := new(server)
		switch c {
		case "cu":
			if cfips.ChinaUnicomIP == "" {
				continue
			}
			s.servername = cfips.ChinaUnicomIP
			s.label = cfips.ChinaUnicomLable
		case "ct":
			if cfips.ChinaTelecomIP == "" {
				continue
			}
			s.servername = cfips.ChinaTelecomIP
			s.label = cfips.ChinaTelecomLable
		case "cm":
			if cfips.ChinaMobileIP == "" {
				continue
			}
			s.servername = cfips.ChinaMobileIP
			s.label = cfips.ChinaMobileLable
		default:
			continue
		}
		servers = append(servers, s)
	}
	for _, node := range nodes {
		sni := node.SNI
		if sni == "" {
			sni = node.Server
		}
		wsHost := node.WSHost
		if wsHost == "" {
			wsHost = node.Server
		}
		switch node.NodeType {
		case "vmess":
			if node.CFIP == 1 {
				for _, srv := range servers {
					n := subscribe.NewVmessNodeWithSNIAndWSHost(node.Name+"_"+srv.label, srv.servername, node.Password, node.WSPath, node.Port, node.AlterID, sni, wsHost)
					ns.Proxies = append(ns.Proxies, n)
				}
			}
			if !cfOnly || len(servers) < 1 {
				n := subscribe.NewVmessNodeWithSNIAndWSHost(node.Name, node.Server, node.Password, node.WSPath, node.Port, node.AlterID, sni, wsHost)
				ns.Proxies = append(ns.Proxies, n)
			}
		case "trojan":
			n := subscribe.NewTrojanNodeWithSNI(node.Name, node.Server, node.Password, sni, node.Port)
			ns.Proxies = append(ns.Proxies, n)
		case "ss":
			n := subscribe.NewSSNode(node.Name, node.Server, node.Cipher, node.Password, node.Port)
			ns.Proxies = append(ns.Proxies, n)
		}
	}
	ns.Update1sProxyGroupsItemProxies()
	bytes, _ := yaml.Marshal(ns)
	return string(bytes)
}
