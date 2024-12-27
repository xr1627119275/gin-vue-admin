package httpInfos

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HttpInfosRouter struct{}

// InitHttpInfosRouter 初始化 httpInfos表 路由信息
func (s *HttpInfosRouter) InitHttpInfosRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	httpInfoRouter := Router.Group("httpInfo").Use(middleware.OperationRecord())
	httpInfoRouterWithoutRecord := Router.Group("httpInfo")
	httpInfoRouterWithoutAuth := PublicRouter.Group("httpInfo")
	{
		httpInfoRouter.POST("createHttpInfos", httpInfoApi.CreateHttpInfos)             // 新建httpInfos表
		httpInfoRouter.DELETE("deleteHttpInfos", httpInfoApi.DeleteHttpInfos)           // 删除httpInfos表
		httpInfoRouter.DELETE("deleteHttpInfosByIds", httpInfoApi.DeleteHttpInfosByIds) // 批量删除httpInfos表
		httpInfoRouter.PUT("updateHttpInfos", httpInfoApi.UpdateHttpInfos)              // 更新httpInfos表
	}
	{
		httpInfoRouterWithoutRecord.GET("findHttpInfos", httpInfoApi.FindHttpInfos)       // 根据ID获取httpInfos表
		httpInfoRouterWithoutRecord.GET("getHttpInfosList", httpInfoApi.GetHttpInfosList) // 获取httpInfos表列表

		httpInfoRouterWithoutRecord.GET("getHttpWeakPassWordInfosList", httpInfoApi.GetHttpWeakPassWordInfosList) // 获取httpInfos表列表

	}
	{
		httpInfoRouterWithoutAuth.GET("getHttpInfosPublic", httpInfoApi.GetHttpInfosPublic) // 获取httpInfos表列表
	}
}
