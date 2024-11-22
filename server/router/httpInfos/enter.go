package httpInfos

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ HttpInfosRouter }

var httpInfoApi = api.ApiGroupApp.HttpInfosApiGroup.HttpInfosApi
