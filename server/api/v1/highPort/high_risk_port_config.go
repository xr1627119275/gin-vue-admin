package highPort

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/highPort"
	highPortReq "github.com/flipped-aurora/gin-vue-admin/server/model/highPort/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HighRiskPortConfigApi struct{}

// CreateHighRiskPortConfig 创建高危端口
// @Tags HighRiskPortConfig
// @Summary 创建高危端口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body highPort.HighRiskPortConfig true "创建高危端口"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /HRPC/createHighRiskPortConfig [post]
func (HRPCApi *HighRiskPortConfigApi) CreateHighRiskPortConfig(c *gin.Context) {
	var HRPC highPort.HighRiskPortConfig
	err := c.ShouldBindJSON(&HRPC)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = HRPCService.CreateHighRiskPortConfig(&HRPC)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteHighRiskPortConfig 删除高危端口
// @Tags HighRiskPortConfig
// @Summary 删除高危端口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body highPort.HighRiskPortConfig true "删除高危端口"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /HRPC/deleteHighRiskPortConfig [delete]
func (HRPCApi *HighRiskPortConfigApi) DeleteHighRiskPortConfig(c *gin.Context) {
	ID := c.Query("ID")
	err := HRPCService.DeleteHighRiskPortConfig(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteHighRiskPortConfigByIds 批量删除高危端口
// @Tags HighRiskPortConfig
// @Summary 批量删除高危端口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /HRPC/deleteHighRiskPortConfigByIds [delete]
func (HRPCApi *HighRiskPortConfigApi) DeleteHighRiskPortConfigByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := HRPCService.DeleteHighRiskPortConfigByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateHighRiskPortConfig 更新高危端口
// @Tags HighRiskPortConfig
// @Summary 更新高危端口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body highPort.HighRiskPortConfig true "更新高危端口"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /HRPC/updateHighRiskPortConfig [put]
func (HRPCApi *HighRiskPortConfigApi) UpdateHighRiskPortConfig(c *gin.Context) {
	var HRPC highPort.HighRiskPortConfig
	err := c.ShouldBindJSON(&HRPC)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = HRPCService.UpdateHighRiskPortConfig(HRPC)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindHighRiskPortConfig 用id查询高危端口
// @Tags HighRiskPortConfig
// @Summary 用id查询高危端口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query highPort.HighRiskPortConfig true "用id查询高危端口"
// @Success 200 {object} response.Response{data=highPort.HighRiskPortConfig,msg=string} "查询成功"
// @Router /HRPC/findHighRiskPortConfig [get]
func (HRPCApi *HighRiskPortConfigApi) FindHighRiskPortConfig(c *gin.Context) {
	ID := c.Query("ID")
	reHRPC, err := HRPCService.GetHighRiskPortConfig(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reHRPC, c)
}

// GetHighRiskPortConfigList 分页获取高危端口列表
// @Tags HighRiskPortConfig
// @Summary 分页获取高危端口列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query highPortReq.HighRiskPortConfigSearch true "分页获取高危端口列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /HRPC/getHighRiskPortConfigList [get]
func (HRPCApi *HighRiskPortConfigApi) GetHighRiskPortConfigList(c *gin.Context) {
	var pageInfo highPortReq.HighRiskPortConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := HRPCService.GetHighRiskPortConfigInfoList(pageInfo)
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

// GetHighRiskPortLogsList 分页获取高危端口日志
// @Tags HighRiskPortConfig
// @Summary 分页获取高危日志列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query highPortReq.HighRiskPortConfigSearch true "分页获取高危端口列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /HRPC/getHighRiskPortLogsList [get]
func (HRPCApi *HighRiskPortConfigApi) GetHighRiskPortLogsList(c *gin.Context) {
	var pageInfo highPortReq.HighRiskPortConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := HRPCService.GetHighRiskPortLogsInfoList(pageInfo)
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

// GetHighRiskPortConfigPublic 不需要鉴权的高危端口接口
// @Tags HighRiskPortConfig
// @Summary 不需要鉴权的高危端口接口
// @accept application/json
// @Produce application/json
// @Param data query highPortReq.HighRiskPortConfigSearch true "分页获取高危端口列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /HRPC/getHighRiskPortConfigPublic [get]
func (HRPCApi *HighRiskPortConfigApi) GetHighRiskPortConfigPublic(c *gin.Context) {
	HRPCService.GetHighRiskPortConfigPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的高危端口接口信息",
	}, "获取成功", c)
}

func (HRPCApi *HighRiskPortConfigApi) PortScan(c *gin.Context) {
	var portSearch highPortReq.PortScanSearch
	err := c.ShouldBindJSON(&portSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, scanInfo := HRPCService.PortScan(portSearch)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(scanInfo, "获取成功", c)
}

func (HRPCApi *HighRiskPortConfigApi) GetPortScan(c *gin.Context) {
	id := c.Query("id")
	data, err := HRPCService.GetPortScan(id)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}
