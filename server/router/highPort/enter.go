package highPort

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	HighRiskPortConfigRouter
	PasswordRuleRouter
}

var (
	HRPCApi = api.ApiGroupApp.HighPortApiGroup.HighRiskPortConfigApi
	PRApi   = api.ApiGroupApp.HighPortApiGroup.PasswordRuleApi
)
