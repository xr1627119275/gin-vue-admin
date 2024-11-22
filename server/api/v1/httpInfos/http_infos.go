package httpInfos

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/httpInfos"
	httpInfosReq "github.com/flipped-aurora/gin-vue-admin/server/model/httpInfos/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HttpInfosApi struct{}

// CreateHttpInfos 创建httpInfos表
// @Tags HttpInfos
// @Summary 创建httpInfos表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body httpInfos.HttpInfos true "创建httpInfos表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /httpInfo/createHttpInfos [post]
func (httpInfoApi *HttpInfosApi) CreateHttpInfos(c *gin.Context) {
	var httpInfo httpInfos.HttpInfos
	err := c.ShouldBindJSON(&httpInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = httpInfoService.CreateHttpInfos(&httpInfo)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteHttpInfos 删除httpInfos表
// @Tags HttpInfos
// @Summary 删除httpInfos表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body httpInfos.HttpInfos true "删除httpInfos表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /httpInfo/deleteHttpInfos [delete]
func (httpInfoApi *HttpInfosApi) DeleteHttpInfos(c *gin.Context) {
	id := c.Query("id")
	err := httpInfoService.DeleteHttpInfos(id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteHttpInfosByIds 批量删除httpInfos表
// @Tags HttpInfos
// @Summary 批量删除httpInfos表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /httpInfo/deleteHttpInfosByIds [delete]
func (httpInfoApi *HttpInfosApi) DeleteHttpInfosByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := httpInfoService.DeleteHttpInfosByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateHttpInfos 更新httpInfos表
// @Tags HttpInfos
// @Summary 更新httpInfos表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body httpInfos.HttpInfos true "更新httpInfos表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /httpInfo/updateHttpInfos [put]
func (httpInfoApi *HttpInfosApi) UpdateHttpInfos(c *gin.Context) {
	var httpInfo httpInfos.HttpInfos
	err := c.ShouldBindJSON(&httpInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = httpInfoService.UpdateHttpInfos(httpInfo)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindHttpInfos 用id查询httpInfos表
// @Tags HttpInfos
// @Summary 用id查询httpInfos表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query httpInfos.HttpInfos true "用id查询httpInfos表"
// @Success 200 {object} response.Response{data=httpInfos.HttpInfos,msg=string} "查询成功"
// @Router /httpInfo/findHttpInfos [get]
func (httpInfoApi *HttpInfosApi) FindHttpInfos(c *gin.Context) {
	id := c.Query("id")
	rehttpInfo, err := httpInfoService.GetHttpInfos(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(rehttpInfo, c)
}

// GetHttpInfosList 分页获取httpInfos表列表
// @Tags HttpInfos
// @Summary 分页获取httpInfos表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query httpInfosReq.HttpInfosSearch true "分页获取httpInfos表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /httpInfo/getHttpInfosList [get]
func (httpInfoApi *HttpInfosApi) GetHttpInfosList(c *gin.Context) {
	var pageInfo httpInfosReq.HttpInfosSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := httpInfoService.GetHttpInfosInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetHttpInfosPublic 不需要鉴权的httpInfos表接口
// @Tags HttpInfos
// @Summary 不需要鉴权的httpInfos表接口
// @accept application/json
// @Produce application/json
// @Param data query httpInfosReq.HttpInfosSearch true "分页获取httpInfos表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /httpInfo/getHttpInfosPublic [get]
func (httpInfoApi *HttpInfosApi) GetHttpInfosPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的httpInfos表接口信息",
	}, "获取成功", c)
}
