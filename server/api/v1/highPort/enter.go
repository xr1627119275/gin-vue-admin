package highPort

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ HighRiskPortConfigApi }

var HRPCService = service.ServiceGroupApp.HighPortServiceGroup.HighRiskPortConfigService
