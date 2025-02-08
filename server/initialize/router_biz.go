package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		httpInfosRouter := router.RouterGroupApp.HttpInfos
		httpInfosRouter.InitHttpInfosRouter(privateGroup, publicGroup)
	}
	{
		webScanRouter := router.RouterGroupApp.WebScan
		webScanRouter.InitWebScanRouter(privateGroup, publicGroup)
	}
	{
		highPortRouter := router.RouterGroupApp.HighPort
		highPortRouter.InitHighRiskPortConfigRouter(privateGroup, publicGroup)
		highPortRouter.InitPasswordRuleRouter(privateGroup, publicGroup)
	}
	{
		nucleiInfoRouter := router.RouterGroupApp.NucleiInfo
		nucleiInfoRouter.InitNucleiRouter(privateGroup, publicGroup)
	}
}
