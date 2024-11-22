package webScan

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ WebScanRouter }

var w_scanApi = api.ApiGroupApp.WebScanApiGroup.WebScanApi
