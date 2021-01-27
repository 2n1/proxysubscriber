package subscribe

type NodeInfo interface {
	NodeName() string
}

type ClashSubscribeProxyItem struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Proxies []string `yaml:"proxies"`
}

type ClashSubscribe struct {
	Proxies     []NodeInfo `yaml:"proxies"`
	ProxyGroups []*ClashSubscribeProxyItem `yaml:"proxy-groups"`
	Rules       []string `yaml:"rules"`
}

func (s *ClashSubscribe) Update1sProxyGroupsItemProxies() {
	s.ProxyGroups[0].Proxies = []string{}
	for _, p := range s.Proxies {
		s.ProxyGroups[0].Proxies = append(s.ProxyGroups[0].Proxies, p.NodeName())
	}
}

type VmessNode struct {
	Name           string `json:"name" yaml:"name"`
	Server         string `json:"server" yaml:"server"`
	Port           int    `json:"port" yaml:"port"`
	Type           string `json:"type" yaml:"type"`
	UUID           string `json:"uuid" yaml:"uuid"`
	AlterID        int    `json:"alter_id" yaml:"alterId"`
	Cipher         string `json:"cipher" yaml:"cipher"`
	TLS            bool   `json:"tls" yaml:"tls"`
	SkipCertVerify bool   `json:"skip_cert_verify" yaml:"skip-cert-verify"`
	Servername     string `json:"servername" yaml:"servername"`
	Network        string `json:"network" yaml:"network"`
	WSPath         string `json:"ws_path" yaml:"ws-path"`
	WSHeaders      struct {
		Host string `json:"host" yaml:"Host"`
	} `json:"ws_headers" yaml:"ws-headers"`
}

func (n *VmessNode) NodeName() string {
	return n.Name
}

func NewVmessNode(name, server, uuid, wsPath string, port, alter int) *VmessNode {
	return NewVmessNodeWithSNIAndWSHost(name, server, uuid, wsPath, port, alter, "", "")

}
func NewVmessNodeWithSNIAndWSHost(name, server, uuid, wsPath string, port, alter int, sni, wsServer string) *VmessNode {
	if sni == "" {
		sni = server
	}
	if wsServer == "" {
		wsServer = server
	}
	return &VmessNode{
		Name:           name,
		Server:         server,
		Port:           port,
		Type:           "vmess",
		UUID:           uuid,
		AlterID:        alter,
		Cipher:         "auto",
		TLS:            true,
		SkipCertVerify: false,
		Servername:     sni,
		Network:        "ws",
		WSPath:         wsPath,
		WSHeaders: struct {
			Host string `json:"host" yaml:"Host"`
		}{Host: wsServer},
	}
}

type SSNode struct {
	Name     string `json:"name" yaml:"name"`
	Server   string `json:"server" yaml:"server"`
	Port     int    `json:"port" yaml:"port"`
	Type     string `json:"type" yaml:"type"`
	Cipher   string `json:"cipher" yaml:"cipher"`
	Password string `json:"password" yaml:"password"`
}

func (n *SSNode) NodeName() string {
	return n.Name
}

func NewSSNode(name, server, cipher, password string, port int) *SSNode {
	return &SSNode{
		Name:     name,
		Server:   server,
		Port:     port,
		Type:     "ss",
		Cipher:   cipher,
		Password: password,
	}
}

type TrojanNode struct {
	Name           string   `json:"name" yaml:"name"`
	Server         string   `json:"server" yaml:"server"`
	Port           int      `json:"port" yaml:"port"`
	Type           string   `json:"type" yaml:"type"`
	Password       string   `json:"password" yaml:"password"`
	SNI            string   `json:"sni" yaml:"sni"`
	Alpn           []string `json:"alpn" yaml:"alpn"`
	SkipCertVerify bool     `json:"skip_cert_verify" yaml:"skip-cert-verify"`
}

func (n *TrojanNode) NodeName() string {
	return n.Name
}

func NewTrojanNode(name, server, password string, port int) *TrojanNode {
	return NewTrojanNodeWithSNI(name, server, password, "", port)
}
func NewTrojanNodeWithSNI(name, server, password, sni string, port int) *TrojanNode {
	if sni == "" {
		sni = server
	}
	return &TrojanNode{
		Name:     name,
		Server:   server,
		Port:     port,
		Type:     "trojan",
		Password: password,
		SNI:      sni,
		Alpn: []string{
			"h2",
			"http/1.1",
		},
		SkipCertVerify: false,
	}
}
