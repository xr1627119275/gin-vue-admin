// 自动生成模板PasswordRule
package highPort

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 密码规则 结构体  PasswordRule
type PasswordRule struct {
	global.GVA_MODEL
	RulePattern *string `json:"rulePattern" form:"rulePattern" gorm:"column:rule_pattern;comment:;" binding:"required"` //规则模式
	Description *string `json:"description" form:"description" gorm:"column:description;comment:;" binding:"required"`  //描述
	Status      *bool   `json:"status" form:"status" gorm:"column:status;comment:;" binding:"required"`                 //状态
	RiskLevel   *int    `json:"riskLevel" form:"riskLevel" gorm:"column:risk_level;comment:;" binding:"required"`       //风险等级
}

// TableName 密码规则 PasswordRule自定义表名 PasswordRules
func (PasswordRule) TableName() string {
	return "PasswordRules"
}
