package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Port"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/highPort"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/httpInfos"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/nucleiInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/webScan"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup     system.ApiGroup
	ExampleApiGroup    example.ApiGroup
	HttpInfosApiGroup  httpInfos.ApiGroup
	WebScanApiGroup    webScan.ApiGroup
	PortApiGroup       Port.ApiGroup
	HighPortApiGroup   highPort.ApiGroup
	NucleiInfoApiGroup nucleiInfo.ApiGroup
}
