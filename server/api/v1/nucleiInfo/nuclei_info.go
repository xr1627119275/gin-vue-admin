package nucleiInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/nucleiInfo"
	nucleiInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/nucleiInfo/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/x_nuclei"
	"github.com/gin-gonic/gin"
	"github.com/projectdiscovery/nuclei/v3/pkg/templates"
	"go.uber.org/zap"
	"strings"
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
// @Router /nuclei/getNucleiTemplateList [post]
func (nucleiApi *NucleiApi) GetNucleiTemplateList(c *gin.Context) {

	var pageInfo nucleiInfoReq.NucleiSearch
	err := c.BindJSON(&pageInfo)

	datas := x_nuclei.GetNucleiTemplates(pageInfo.TemplateFilters)

	if pageInfo.Input != "" {
		temp := make([]*templates.Template, 0)
		for _, template := range datas {
			if strings.Contains(template.ID+template.Info.Name+template.Info.Description, pageInfo.Input) {
				temp = append(temp, template)
			}
		}
		datas = temp
	}
	total := len(datas)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Page := pageInfo.Page
	pageInfo.Page = Page - 1
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}

	var nextIndex = (pageInfo.Page + 1) * pageInfo.PageSize
	if nextIndex > total {
		nextIndex = total
	}
	datas = datas[pageInfo.Page*pageInfo.PageSize : nextIndex]
	// 清理空的
	for i := 0; i < len(datas); i++ {
		if datas[i] == nil {
			datas = append(datas[:i], datas[i+1:]...)
			i--
		}
	}

	response.OkWithDetailed(response.PageResult{
		List:     datas,
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

// CreateScan 创建扫描
// @Tags Nuclei
// @Summary 创建扫描
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body nucleiInfo.NucleiScan true "创建扫描"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /nuclei/createScan [post]
func (nucleiApi *NucleiApi) CreateScan(c *gin.Context) {
	var nucleiScanQuery nucleiInfoReq.NucleiScanQuery
	err := c.ShouldBindJSON(&nucleiScanQuery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	logId, err := nucleiService.CreateNucleiScan(&nucleiScanQuery)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{
		"id": logId,
	}, "创建成功", c)
}

// GetNucleiPocData 获取pocData
// @Tags Nuclei
// @Summary 获取pocData
// @accept application/json
// @Produce application/json
// @Param data query nucleiInfoReq.NucleiSearch true "获取pocData"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /nuclei/getNucleiPocData [get]
func (nucleiApi *NucleiApi) GetNucleiPocData(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	id := c.Query("id")
	data, _ := nucleiService.GetNucleiPocData(id)
	response.OkWithDetailed(gin.H{
		"data": string(data),
	}, "获取成功", c)
}
