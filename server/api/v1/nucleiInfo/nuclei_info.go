package nucleiInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/nucleiInfo"
	nucleiInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/nucleiInfo/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/x_nuclei"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NucleiApi struct{}

// CreateNuclei 创建nucleiInfo
// @Tags Nuclei
// @Summary 创建nucleiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body nucleiInfo.Nuclei true "创建nucleiInfo"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /nuclei/createNuclei [post]
func (nucleiApi *NucleiApi) CreateNuclei(c *gin.Context) {
	var nuclei nucleiInfo.Nuclei
	err := c.ShouldBindJSON(&nuclei)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = nucleiService.CreateNuclei(&nuclei)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteNuclei 删除nucleiInfo
// @Tags Nuclei
// @Summary 删除nucleiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body nucleiInfo.Nuclei true "删除nucleiInfo"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /nuclei/deleteNuclei [delete]
func (nucleiApi *NucleiApi) DeleteNuclei(c *gin.Context) {
	ID := c.Query("ID")
	err := nucleiService.DeleteNuclei(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteNucleiByIds 批量删除nucleiInfo
// @Tags Nuclei
// @Summary 批量删除nucleiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /nuclei/deleteNucleiByIds [delete]
func (nucleiApi *NucleiApi) DeleteNucleiByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := nucleiService.DeleteNucleiByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateNuclei 更新nucleiInfo
// @Tags Nuclei
// @Summary 更新nucleiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body nucleiInfo.Nuclei true "更新nucleiInfo"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /nuclei/updateNuclei [put]
func (nucleiApi *NucleiApi) UpdateNuclei(c *gin.Context) {
	var nuclei nucleiInfo.Nuclei
	err := c.ShouldBindJSON(&nuclei)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = nucleiService.UpdateNuclei(nuclei)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindNuclei 用id查询nucleiInfo
// @Tags Nuclei
// @Summary 用id查询nucleiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query nucleiInfo.Nuclei true "用id查询nucleiInfo"
// @Success 200 {object} response.Response{data=nucleiInfo.Nuclei,msg=string} "查询成功"
// @Router /nuclei/findNuclei [get]
func (nucleiApi *NucleiApi) FindNuclei(c *gin.Context) {
	ID := c.Query("ID")
	renuclei, err := nucleiService.GetNuclei(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(renuclei, c)
}

// GetNucleiList 分页获取nucleiInfo列表
// @Tags Nuclei
// @Summary 分页获取nucleiInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query nucleiInfoReq.NucleiSearch true "分页获取nucleiInfo列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /nuclei/getNucleiList [get]
func (nucleiApi *NucleiApi) GetNucleiList(c *gin.Context) {
	var pageInfo nucleiInfoReq.NucleiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := nucleiService.GetNucleiInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetNucleiTemplateList 分页获取nuclei 模板列表
// @Tags Nuclei
// @Summary 分页获取nuclei模板列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query nucleiInfoReq.NucleiSearch true "分页获取nucleiInfo列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /nuclei/getNucleiTemplateList [get]
func (nucleiApi *NucleiApi) GetNucleiTemplateList(c *gin.Context) {
	templates := x_nuclei.GetNucleiTemplates()
	total := len(templates)
	// templates 分页获取
	var pageInfo nucleiInfoReq.NucleiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Page := pageInfo.Page
	pageInfo.Page = Page - 1
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}
	templates = templates[pageInfo.Page*pageInfo.PageSize : (pageInfo.Page+1)*pageInfo.PageSize]
	// 清理空的
	for i := 0; i < len(templates); i++ {
		if templates[i] == nil {
			templates = append(templates[:i], templates[i+1:]...)
			i--
		}
	}
	response.OkWithDetailed(response.PageResult{
		List:     templates,
		Total:    int64(total),
		Page:     Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetNucleiPublic 不需要鉴权的nucleiInfo接口
// @Tags Nuclei
// @Summary 不需要鉴权的nucleiInfo接口
// @accept application/json
// @Produce application/json
// @Param data query nucleiInfoReq.NucleiSearch true "分页获取nucleiInfo列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /nuclei/getNucleiPublic [get]
func (nucleiApi *NucleiApi) GetNucleiPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	nucleiService.GetNucleiPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的nucleiInfo接口信息",
	}, "获取成功", c)
}
