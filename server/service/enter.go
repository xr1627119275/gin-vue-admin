package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/Port"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/highPort"
	"github.com/flipped-aurora/gin-vue-admin/server/service/httpInfos"
	"github.com/flipped-aurora/gin-vue-admin/server/service/nucleiInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/webScan"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	HttpInfosServiceGroup  httpInfos.ServiceGroup
	WebScanServiceGroup    webScan.ServiceGroup
	PortServiceGroup       Port.ServiceGroup
	HighPortServiceGroup   highPort.ServiceGroup
	NucleiInfoServiceGroup nucleiInfo.ServiceGroup
}
