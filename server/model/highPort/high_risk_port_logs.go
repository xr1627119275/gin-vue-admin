// 自动生成模板HighRiskPort
package highPort

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 高危端口 结构体  HighRiskPortLog
type HighRiskPortLog struct {
	global.GVA_MODEL
	Info         string `json:"info" gorm:"column:info"`
	PortConfigId uint16
	PortConfig   *HighRiskPortConfig
	Ip           string
	FromIP       string
	Mac          string
	FromMac      string
}

// TableName 高危端口记录 HighRiskPort自定义表名 high_risk_port_logs
func (HighRiskPortLog) TableName() string {
	return "high_risk_port_logs"
}
