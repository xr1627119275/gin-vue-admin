import service from '@/utils/request'
// @Tags HighRiskPortConfig
// @Summary 创建高危端口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HighRiskPortConfig true "创建高危端口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /HRPC/createHighRiskPortConfig [post]
export const createHighRiskPortConfig = (data) => {
  return service({
    url: '/HRPC/createHighRiskPortConfig',
    method: 'post',
    data
  })
}

// @Tags HighRiskPortConfig
// @Summary 删除高危端口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HighRiskPortConfig true "删除高危端口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /HRPC/deleteHighRiskPortConfig [delete]
export const deleteHighRiskPortConfig = (params) => {
  return service({
    url: '/HRPC/deleteHighRiskPortConfig',
    method: 'delete',
    params
  })
}

// @Tags HighRiskPortConfig
// @Summary 批量删除高危端口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除高危端口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /HRPC/deleteHighRiskPortConfig [delete]
export const deleteHighRiskPortConfigByIds = (params) => {
  return service({
    url: '/HRPC/deleteHighRiskPortConfigByIds',
    method: 'delete',
    params
  })
}

// @Tags HighRiskPortConfig
// @Summary 更新高危端口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HighRiskPortConfig true "更新高危端口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /HRPC/updateHighRiskPortConfig [put]
export const updateHighRiskPortConfig = (data) => {
  return service({
    url: '/HRPC/updateHighRiskPortConfig',
    method: 'put',
    data
  })
}

// @Tags HighRiskPortConfig
// @Summary 用id查询高危端口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HighRiskPortConfig true "用id查询高危端口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /HRPC/findHighRiskPortConfig [get]
export const findHighRiskPortConfig = (params) => {
  return service({
    url: '/HRPC/findHighRiskPortConfig',
    method: 'get',
    params
  })
}

// @Tags HighRiskPortConfig
// @Summary 分页获取高危端口列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取高危端口列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /HRPC/getHighRiskPortConfigList [get]
export const getHighRiskPortConfigList = (params) => {
  return service({
    url: '/HRPC/getHighRiskPortConfigList',
    method: 'get',
    params
  })
}

export const getHighRiskPortLogsList = (params) => {
  return service({
    url: '/HRPC/getHighRiskPortLogsList',
    method: 'get',
    params
  })
}

// @Tags HighRiskPortConfig
// @Summary 不需要鉴权的高危端口接口
// @accept application/json
// @Produce application/json
// @Param data query highPortReq.HighRiskPortConfigSearch true "分页获取高危端口列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /HRPC/getHighRiskPortConfigPublic [get]
export const getHighRiskPortConfigPublic = () => {
  return service({
    url: '/HRPC/getHighRiskPortConfigPublic',
    method: 'get'
  })
}
