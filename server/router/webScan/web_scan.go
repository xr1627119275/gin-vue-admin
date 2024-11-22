package webScan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WebScanRouter struct{}

// InitWebScanRouter 初始化 web扫描 路由信息
func (s *WebScanRouter) InitWebScanRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	w_scanRouter := Router.Group("w_scan").Use(middleware.OperationRecord())
	w_scanRouterWithoutRecord := Router.Group("w_scan")
	w_scanRouterWithoutAuth := PublicRouter.Group("w_scan")
	{
		w_scanRouter.POST("createWebScan", w_scanApi.CreateWebScan)             // 新建web扫描
		w_scanRouter.DELETE("deleteWebScan", w_scanApi.DeleteWebScan)           // 删除web扫描
		w_scanRouter.DELETE("deleteWebScanByIds", w_scanApi.DeleteWebScanByIds) // 批量删除web扫描
		w_scanRouter.PUT("updateWebScan", w_scanApi.UpdateWebScan)              // 更新web扫描
	}
	{
		w_scanRouterWithoutRecord.GET("findWebScan", w_scanApi.FindWebScan)       // 根据ID获取web扫描
		w_scanRouterWithoutRecord.GET("getWebScanList", w_scanApi.GetWebScanList) // 获取web扫描列表
		w_scanRouterWithoutRecord.POST("parseScan", w_scanApi.ParseScan)          // 解析web 资产扫描
	}
	{
		w_scanRouterWithoutAuth.GET("getWebScanPublic", w_scanApi.GetWebScanPublic) // 获取web扫描列表
	}
}
