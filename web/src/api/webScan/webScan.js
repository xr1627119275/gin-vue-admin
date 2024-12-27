import service from '@/utils/request'

// @Tags WebScan
// @Summary 创建web扫描
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WebScan true "创建web扫描"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /w_scan/createWebScan [post]
export const createWebScan = (data) => {
  return service({
    url: '/w_scan/createWebScan',
    method: 'post',
    data
  })
}

export const createAFrogWebScan = (data) => {
  return service({
    url: '/w_scan/createAFrogWebScan',
    method: 'post',
    data
  })
}

// @Tags WebScan
// @Summary 删除web扫描
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WebScan true "删除web扫描"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /w_scan/deleteWebScan [delete]
export const deleteWebScan = (params) => {
  return service({
    url: '/w_scan/deleteWebScan',
    method: 'delete',
    params
  })
}

// @Tags WebScan
// @Summary 批量删除web扫描
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除web扫描"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /w_scan/deleteWebScan [delete]
export const deleteWebScanByIds = (params) => {
  return service({
    url: '/w_scan/deleteWebScanByIds',
    method: 'delete',
    params
  })
}

// @Tags WebScan
// @Summary 更新web扫描
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WebScan true "更新web扫描"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /w_scan/updateWebScan [put]
export const updateWebScan = (data) => {
  return service({
    url: '/w_scan/updateWebScan',
    method: 'put',
    data
  })
}

// @Tags WebScan
// @Summary 用id查询web扫描
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WebScan true "用id查询web扫描"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /w_scan/findWebScan [get]
export const findWebScan = (params) => {
  return service({
    url: '/w_scan/findWebScan',
    method: 'get',
    params
  })
}

// @Tags ParseScan
// @Summary 解析web扫描
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /w_scan/parseScan [post]
export const parseScan = (data) => {
  return service({
    url: '/w_scan/parseScan',
    method: 'post',
    data
  })
}

// @Tags WebScan
// @Summary 分页获取web扫描列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取web扫描列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /w_scan/getWebScanList [get]
export const getWebScanList = (params) => {
  return service({
    url: '/w_scan/getWebScanList',
    method: 'get',
    params
  })
}
