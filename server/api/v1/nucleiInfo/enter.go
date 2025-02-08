package nucleiInfo

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ NucleiApi }

var nucleiService = service.ServiceGroupApp.NucleiInfoServiceGroup.NucleiService
