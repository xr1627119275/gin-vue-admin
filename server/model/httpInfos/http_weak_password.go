package httpInfos

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// HttpWeakPassword表 结构体  HttpInfos
type HttpWeakPassword struct {
	global.GVA_MODEL
	HttpInfosId int       `json:"http_infos_id"`
	HttpInfos   HttpInfos `json:"http_infos"`
	Password    string    `json:"password" gorm:"column:password"`
}

// TableName httpInfos表 HttpInfos自定义表名 http_infos
func (HttpWeakPassword) TableName() string {
	return "http_weak_password_records"
}
