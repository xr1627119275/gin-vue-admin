package highPort

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	HighRiskPortConfigApi
	PasswordRuleApi
}

var (
	HRPCService = service.ServiceGroupApp.HighPortServiceGroup.HighRiskPortConfigService
	PRService   = service.ServiceGroupApp.HighPortServiceGroup.PasswordRuleService
)
