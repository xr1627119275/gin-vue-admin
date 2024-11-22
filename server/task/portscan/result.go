package portscan

import (
	"fmt"
	"sync"
)

const (
	IpOpenedPortFilterNumber = 50 // IP开放端口数量，超过该数量则认为存在安全设备过滤
)

// Config 端口扫描的参数配置
type Config struct {
	Target           string `json:"target"`
	ExcludeTarget    string `json:"executeTarget"`
	Port             string `json:"port"`
	OrgId            *int   `json:"orgId"`
	Rate             int    `json:"rate"`
	IsPing           bool   `json:"ping"`
	Tech             string `json:"tech"`
	IsIpLocation     bool   `json:"ipLocation"`
	IsHttpx          bool   `json:"httpx"`
	IsScreenshot     bool   `json:"screenshot"`
	IsFingerprintHub bool   `json:"fingerprinthub"`
	IsIconHash       bool   `json:"iconhash"`
	IsFingerprintx   bool   `json:"fingerprintx"`
	CmdBin           string `json:"cmdBin"`
	IsLoadOpenedPort bool   `json:"loadOpenedPort"`
	IsPortscan       bool   `json:"isPortscan"`
	WorkspaceId      int    `json:"workspaceId"`
	IsProxy          bool   `json:"proxy"`
}

// PortAttrResult 端口属性结果
type PortAttrResult struct {
	RelatedId int
	Source    string
	Tag       string
	Content   string
}

type HttpResult struct {
	RelatedId int
	Source    string
	Tag       string
	Content   string
}

// PortResult 端口结果
type PortResult struct {
	Status    string
	PortAttrs []PortAttrResult
	HttpInfo  []HttpResult
}

// IPResult IP结果
type IPResult struct {
	OrgId    *int
	Location string
	Status   string
	Ports    map[int]*PortResult `json:"ports"`
}

// Result 端口扫描结果
type Result struct {
	sync.RWMutex
	IPResult map[string]*IPResult `json:"ipResult"`
}

type OfflineResult interface {
	ParseContentResult(content []byte) (ipResult Result)
}

type ImportOfflineResult struct {
	resultType       string
	offlineInterface OfflineResult
	IpResult         Result
}

func NewImportOfflineResult(resultType string) *ImportOfflineResult {
	i := &ImportOfflineResult{resultType: resultType}
	switch resultType {
	//case "nmap":
	//	i.offlineInterface = new(Nmap)
	//case "masscan":
	//	i.offlineInterface = new(Masscan)
	case "fscan":
		i.offlineInterface = new(FScan)
		//case "gogo":
		//	i.offlineInterface = new(Gogo)
		//case "goby":
		//	i.offlineInterface = new(Goby)
	}
	return i
}

func NewImportOfflineResultWithInterface(resultType string, resultInterface OfflineResult) *ImportOfflineResult {
	i := &ImportOfflineResult{resultType: resultType}
	i.offlineInterface = resultInterface

	return i
}

func (i *ImportOfflineResult) Parse(content []byte) {
	if i.offlineInterface == nil {
		fmt.Errorf("invalid offline result:%s", i.resultType)
		return
	}
	i.IpResult = i.offlineInterface.ParseContentResult(content)
}

func (r *Result) HasIP(ip string) bool {
	r.RLock()
	defer r.RUnlock()

	_, ok := r.IPResult[ip]
	return ok
}

func (r *Result) SetIP(ip string) {
	r.Lock()
	defer r.Unlock()

	r.IPResult[ip] = &IPResult{Ports: make(map[int]*PortResult)}
}

func (r *Result) HasPort(ip string, port int) bool {
	r.RLock()
	defer r.RUnlock()

	_, ok := r.IPResult[ip].Ports[port]
	return ok
}

func (r *Result) SetPort(ip string, port int) {
	r.Lock()
	defer r.Unlock()

	r.IPResult[ip].Ports[port] = &PortResult{PortAttrs: []PortAttrResult{}}
}

func (r *Result) SetPortAttr(ip string, port int, par PortAttrResult) {
	r.Lock()
	defer r.Unlock()

	r.IPResult[ip].Ports[port].PortAttrs = append(r.IPResult[ip].Ports[port].PortAttrs, par)
}

func (r *Result) SetPortHttpInfo(ip string, port int, result HttpResult) {
	r.Lock()
	defer r.Unlock()

	r.IPResult[ip].Ports[port].HttpInfo = append(r.IPResult[ip].Ports[port].HttpInfo, result)
}
