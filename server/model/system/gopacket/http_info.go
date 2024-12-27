package gopacket

import "github.com/flipped-aurora/gin-vue-admin/server/model/httpInfos"

type HttpInfo struct {
	httpInfos.HttpInfos
	Headers   map[string]string `json:"headers" gorm:"serializer:json;type:text;"`
	ResHeader map[string]string `json:"resHeader" gorm:"serializer:json;type:text;"`
}

//func (HttpInfo) TableName() string {
//	return "http_infos"
//}
