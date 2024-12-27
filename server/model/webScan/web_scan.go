// 自动生成模板WebScan
package webScan

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/customType"
)

type Tag []string

func (t *Tag) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Tag) Value() (driver.Value, error) {
	return json.Marshal(t)
}

// web扫描 结构体  WebScan
type WebScan struct {
	global.GVA_MODEL_UUID
	HOST       string             `json:"host" form:"host" gorm:"column:host;comment:;"`                 //HOST
	PORT       string             `json:"port" form:"port" gorm:"column:port;comment:;"`                 //HOST
	Target     string             `json:"target" form:"target" gorm:"column:target;comment:;"`           //Target
	Result     string             `json:"result" form:"result" gorm:"column:result;comment:;type:text;"` //扫描结果
	Args       customType.JsonMap `json:"args" form:"args" gorm:"column:args;comment:;type:text;"`       //参数
	ScanOver   bool               `json:"scan_over" gorm:"column:scan_over;comment:;default:false;"`     //扫描结束
	JsonResult string             `json:"json_result" gorm:"column:json_result;comment:;type:longtext;"` //json结果
}

// TableName web扫描 WebScan自定义表名 web_scan
func (WebScan) TableName() string {
	return "web_scan"
}
