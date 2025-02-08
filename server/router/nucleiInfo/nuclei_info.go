package nucleiInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NucleiRouter struct{}

// InitNucleiRouter 初始化 nucleiInfo 路由信息
func (s *NucleiRouter) InitNucleiRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	nucleiRouter := Router.Group("nuclei").Use(middleware.OperationRecord())
	nucleiRouterWithoutRecord := Router.Group("nuclei")
	nucleiRouterWithoutAuth := PublicRouter.Group("nuclei")
	{
		nucleiRouter.POST("createNuclei", nucleiApi.CreateNuclei)             // 新建nucleiInfo
		nucleiRouter.DELETE("deleteNuclei", nucleiApi.DeleteNuclei)           // 删除nucleiInfo
		nucleiRouter.DELETE("deleteNucleiByIds", nucleiApi.DeleteNucleiByIds) // 批量删除nucleiInfo
		nucleiRouter.PUT("updateNuclei", nucleiApi.UpdateNuclei)              // 更新nucleiInfo
	}
	{
		nucleiRouterWithoutRecord.GET("findNuclei", nucleiApi.FindNuclei)                       // 根据ID获取nucleiInfo
		nucleiRouterWithoutRecord.GET("getNucleiList", nucleiApi.GetNucleiList)                 // 获取nucleiInfo列表
		nucleiRouterWithoutRecord.GET("getNucleiTemplateList", nucleiApi.GetNucleiTemplateList) // 获取nucleiInfo列表
	}
	{
		nucleiRouterWithoutAuth.GET("getNucleiPublic", nucleiApi.GetNucleiPublic) // nucleiInfo开放接口
	}
}
