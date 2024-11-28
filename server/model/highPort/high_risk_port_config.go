// 自动生成模板HighRiskPortConfig
package highPort

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 高危端口 结构体  HighRiskPortConfig
type HighRiskPortConfig struct {
	global.GVA_MODEL
	PortNumber      *int    `json:"portNumber" form:"portNumber" gorm:"column:port_number;comment:端口号;" binding:"required"`      //端口号
	PortName        *string `json:"portName" form:"portName" gorm:"column:port_name;comment:端口名称;" binding:"required"`           //端口名称
	RiskLevel       *string `json:"riskLevel" form:"riskLevel" gorm:"column:risk_level;comment:风险等级;" binding:"required"`        //风险等级
	PortDescription *string `json:"portDescription" form:"portDescription" gorm:"column:port_description;comment:描述;type:text;"` //描述
}

// TableName 高危端口 HighRiskPortConfig自定义表名 high_risk_port_config
func (HighRiskPortConfig) TableName() string {
	return "high_risk_port_config"
}
