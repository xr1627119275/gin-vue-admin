package gopacket

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type HttpInfo struct {
	global.GVA_MODEL
	Host       string              `json:"host" gorm:"comment:域名"`                       // 域名
	ResContent string              `json:"res_content" gorm:"type:text;comment:响应内容字符串"` // 响应内容
	ResData    []byte              `json:"res_data" gorm:"comment:响应数据"`
	ReqContent string              `json:"req_content" gorm:"type:text;comment:请求内容字符串"` // 响应内容
	ReqData    []byte              `json:"req_data" gorm:"comment:请求数据"`
	Headers    map[string][]string `json:"headers" gorm:"serializer:json;type:text;"`
	ResHeader  map[string][]string `json:"res_header" gorm:"serializer:json;type:text;"`
	StatusCode int                 `json:"status_code" gorm:"comment:响应码"`
	RemoteAddr string              `json:"remote_addr" gorm:"comment:原始IP"`
	Path       string              `json:"path" gorm:"type:text;comment:路径"` // PATH路径
	Method     string              `json:"method" gorm:"type:text;comment:method"`
	Proto      string              `json:"proto" gorm:"type:text;comment:proto"`
}

func (HttpInfo) TableName() string {
	return "http_infos"
}
