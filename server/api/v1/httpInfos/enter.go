package httpInfos

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ HttpInfosApi }

var httpInfoService = service.ServiceGroupApp.HttpInfosServiceGroup.HttpInfosService
