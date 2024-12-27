package webScan

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webScan"
	webScanReq "github.com/flipped-aurora/gin-vue-admin/server/model/webScan/request"
	webScanService "github.com/flipped-aurora/gin-vue-admin/server/service/webScan"
	"github.com/flipped-aurora/gin-vue-admin/server/task/portscan"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
)

type WebScanApi struct{}

// CreateWebScan 创建web扫描
// @Tags WebScan
// @Summary 创建web扫描
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webScan.WebScan true "创建web扫描"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /w_scan/createWebScan [post]
func (w_scanApi *WebScanApi) CreateWebScan(c *gin.Context) {
	var params webScanReq.RunWebScanParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var RunType = params.RunType
	err = w_scanService.CreateWebScan(&params.WebScan)
	go webScanService.WebScanServiceApp.RunScan(params, RunType)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithData(params.WebScan.GVA_MODEL_UUID.ID, c)
}

func (w_scanApi *WebScanApi) CreateAFrogWebScan(c *gin.Context) {
	var params webScanReq.RunWebScanParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, webscan := webScanService.WebScanServiceApp.RunAFrogScan(params)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithData(webscan, c)
}

// DeleteWebScan 删除web扫描
// @Tags WebScan
// @Summary 删除web扫描
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webScan.WebScan true "删除web扫描"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /w_scan/deleteWebScan [delete]
func (w_scanApi *WebScanApi) DeleteWebScan(c *gin.Context) {
	ID := c.Query("ID")
	err := w_scanService.DeleteWebScan(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWebScanByIds 批量删除web扫描
// @Tags WebScan
// @Summary 批量删除web扫描
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /w_scan/deleteWebScanByIds [delete]
func (w_scanApi *WebScanApi) DeleteWebScanByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := w_scanService.DeleteWebScanByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWebScan 更新web扫描
// @Tags WebScan
// @Summary 更新web扫描
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body webScan.WebScan true "更新web扫描"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /w_scan/updateWebScan [put]
func (w_scanApi *WebScanApi) UpdateWebScan(c *gin.Context) {
	var w_scan webScan.WebScan
	err := c.ShouldBindJSON(&w_scan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = w_scanService.UpdateWebScan(w_scan)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWebScan 用id查询web扫描
// @Tags WebScan
// @Summary 用id查询web扫描
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webScan.WebScan true "用id查询web扫描"
// @Success 200 {object} response.Response{data=webScan.WebScan,msg=string} "查询成功"
// @Router /w_scan/findWebScan [get]
func (w_scanApi *WebScanApi) FindWebScan(c *gin.Context) {
	ID := c.Query("ID")
	rew_scan, err := w_scanService.GetWebScan(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(rew_scan, c)
}

// GetWebScanList 分页获取web扫描列表
// @Tags WebScan
// @Summary 分页获取web扫描列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query webScanReq.WebScanSearch true "分页获取web扫描列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /w_scan/getWebScanList [get]
func (w_scanApi *WebScanApi) GetWebScanList(c *gin.Context) {
	var pageInfo webScanReq.WebScanSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := w_scanService.GetWebScanInfoList(pageInfo)
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

// GetWebScanPublic 不需要鉴权的web扫描接口
// @Tags WebScan
// @Summary 不需要鉴权的web扫描接口
// @accept application/json
// @Produce application/json
// @Param data query webScanReq.WebScanSearch true "分页获取web扫描列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /w_scan/getWebScanPublic [get]
func (w_scanApi *WebScanApi) GetWebScanPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的web扫描接口信息",
	}, "获取成功", c)
}

// ParseScan 解析web资产扫描
// @Tags WebScan
// @Summary 解析web资产扫描
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /w_scan/parseScan [post]
func (w_scanApi *WebScanApi) ParseScan(c *gin.Context) {

	var parseParam webScanReq.ParseWebScanSearch
	err := c.ShouldBindJSON(&parseParam)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	i := portscan.NewImportOfflineResult(parseParam.Type)

	if parseParam.TaskID != "" {
		scanInfo, err := w_scanService.GetWebScan(parseParam.TaskID)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		i.Parse([]byte(scanInfo.Result))
	}
	if parseParam.Content != "" {
		i.Parse([]byte(parseParam.Content))
	}
	for ip, ipa := range i.IpResult.IPResult {
		log.Printf("%s %s", ip, ipa)
		for port, pa := range ipa.Ports {
			log.Printf("%s %s", port, pa)
		}
	}

	response.OkWithData(i, c)
	//response.OkWithMessage("运行成功", c)

}
