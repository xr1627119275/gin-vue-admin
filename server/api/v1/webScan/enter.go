package webScan

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ WebScanApi }

var w_scanService = service.ServiceGroupApp.WebScanServiceGroup.WebScanService
