package highPort

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HighRiskPortConfigRouter struct{}

// InitHighRiskPortConfigRouter 初始化 高危端口 路由信息
func (s *HighRiskPortConfigRouter) InitHighRiskPortConfigRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	HRPCRouter := Router.Group("HRPC").Use(middleware.OperationRecord())
	HRPCRouterWithoutRecord := Router.Group("HRPC")
	HRPCRouterWithoutAuth := PublicRouter.Group("HRPC")
	{
		HRPCRouter.POST("createHighRiskPortConfig", HRPCApi.CreateHighRiskPortConfig)             // 新建高危端口
		HRPCRouter.DELETE("deleteHighRiskPortConfig", HRPCApi.DeleteHighRiskPortConfig)           // 删除高危端口
		HRPCRouter.DELETE("deleteHighRiskPortConfigByIds", HRPCApi.DeleteHighRiskPortConfigByIds) // 批量删除高危端口
		HRPCRouter.PUT("updateHighRiskPortConfig", HRPCApi.UpdateHighRiskPortConfig)              // 更新高危端口
	}
	{
		HRPCRouterWithoutRecord.POST("portScan", HRPCApi.PortScan)      // 获取扫描结果
		HRPCRouterWithoutRecord.GET("getPortScan", HRPCApi.GetPortScan) // 获取高危端口日志列表

		HRPCRouterWithoutRecord.GET("findHighRiskPortConfig", HRPCApi.FindHighRiskPortConfig)       // 根据ID获取高危端口
		HRPCRouterWithoutRecord.GET("getHighRiskPortConfigList", HRPCApi.GetHighRiskPortConfigList) // 获取高危端口列表
		HRPCRouterWithoutRecord.GET("getHighRiskPortLogsList", HRPCApi.GetHighRiskPortLogsList)     // 获取高危端口日志列表

	}
	{
		HRPCRouterWithoutAuth.GET("getHighRiskPortConfigPublic", HRPCApi.GetHighRiskPortConfigPublic) // 高危端口开放接口
	}
}
