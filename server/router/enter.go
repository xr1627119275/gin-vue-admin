package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Port"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/highPort"
	"github.com/flipped-aurora/gin-vue-admin/server/router/httpInfos"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/webScan"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	HttpInfos httpInfos.RouterGroup
	WebScan   webScan.RouterGroup
	Port      Port.RouterGroup
	HighPort  highPort.RouterGroup
}
