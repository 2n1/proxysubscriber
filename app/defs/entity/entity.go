package entity

type Group struct {
	ID int64 `json:"id"`
	Name string `json:"name" form:"name" binding:"required"`
	Url string
}

type Node struct {
	ID int64 `json:"id"`
	Name     string `json:"name" form:"name" binding:"required`
	GroupID  int    `json:"group_id" form:"group_id" binding:"required`
	NodeType string `json:"node_type" form:"type" binding:"required`
	Server   string `json:"server" form:"server" binding:"required`
	Port     int    `json:"port" form:"port" binding:"required`
	Password string `json:"password" form:"password" binding:"required`
	Cipher   string `json:"cipher" form:"cipher"`
	SNI      string `json:"sni" form:"sni"`
	AlterID  int `json:"alter_id" form:"alter_id"`
	WSPath   string `json:"ws_path" form:"ws_path"`
	WSHost   string `json:"ws_host" form:"ws_host"`
	CFIP int `json:"cf_ip" form:"cf_ip"`
	GroupName string `json:"group_name"`
}

type CloudFlareIP struct {
	ChinaUnicomIP     string `json:"cu_ip" form:"cu_ip"`
	ChinaUnicomLable  string `json:"cu_label" form:"cu_label"`
	ChinaTelecomIP    string `json:"ct_ip" form:"ct_ip"`
	ChinaTelecomLable string `json:"ct_label" form:"ct_label"`
	ChinaMobileIP     string `json:"cm_ip" form:"cm_ip"`
	ChinaMobileLable  string `json:"cm_label" form:"cm_label"`
}

type Authorization struct {
	ID int64 `json:"id"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type URL struct {
	ID string
	GroupID int64
}