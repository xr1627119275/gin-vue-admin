package highPort

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PasswordRuleRouter struct{}

func (s *PasswordRuleRouter) InitPasswordRuleRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	PRRouter := Router.Group("PR").Use(middleware.OperationRecord())
	PRRouterWithoutRecord := Router.Group("PR")
	PRRouterWithoutAuth := PublicRouter.Group("PR")
	{
		PRRouter.POST("createPasswordRule", PRApi.CreatePasswordRule)
		PRRouter.DELETE("deletePasswordRule", PRApi.DeletePasswordRule)
		PRRouter.DELETE("deletePasswordRuleByIds", PRApi.DeletePasswordRuleByIds)
		PRRouter.PUT("updatePasswordRule", PRApi.UpdatePasswordRule)
		PRRouter.POST("dict", PRApi.DictConfig)
		PRRouter.GET("dict", PRApi.GetDict)

	}
	{
		PRRouterWithoutRecord.GET("findPasswordRule", PRApi.FindPasswordRule)
		PRRouterWithoutRecord.GET("getPasswordRuleList", PRApi.GetPasswordRuleList)
	}
	{
		PRRouterWithoutAuth.GET("getPasswordRulePublic", PRApi.GetPasswordRulePublic)
	}
}
