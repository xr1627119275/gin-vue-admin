import service from '@/utils/request'

// @Tags HttpInfos
// @Summary 创建httpInfos表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HttpInfos true "创建httpInfos表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /httpInfo/createHttpInfos [post]
export const createHttpInfos = (data) => {
  return service({
    url: '/httpInfo/createHttpInfos',
    method: 'post',
    data
  })
}

// @Tags HttpInfos
// @Summary 删除httpInfos表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HttpInfos true "删除httpInfos表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /httpInfo/deleteHttpInfos [delete]
export const deleteHttpInfos = (params) => {
  return service({
    url: '/httpInfo/deleteHttpInfos',
    method: 'delete',
    params
  })
}

// @Tags HttpInfos
// @Summary 批量删除httpInfos表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除httpInfos表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /httpInfo/deleteHttpInfos [delete]
export const deleteHttpInfosByIds = (params) => {
  return service({
    url: '/httpInfo/deleteHttpInfosByIds',
    method: 'delete',
    params
  })
}

// @Tags HttpInfos
// @Summary 更新httpInfos表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.HttpInfos true "更新httpInfos表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /httpInfo/updateHttpInfos [put]
export const updateHttpInfos = (data) => {
  return service({
    url: '/httpInfo/updateHttpInfos',
    method: 'put',
    data
  })
}

// @Tags HttpInfos
// @Summary 用id查询httpInfos表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.HttpInfos true "用id查询httpInfos表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /httpInfo/findHttpInfos [get]
export const findHttpInfos = (params) => {
  return service({
    url: '/httpInfo/findHttpInfos',
    method: 'get',
    params
  })
}

// @Tags HttpInfos
// @Summary 分页获取httpInfos表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取httpInfos表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /httpInfo/getHttpInfosList [get]
export const getHttpInfosList = (params) => {
  return service({
    url: '/httpInfo/getHttpInfosList',
    method: 'get',
    params
  })
}

export const getHttpWeakPassWordInfosList = (params) => {
  return service({
    url: '/httpInfo/getHttpWeakPassWordInfosList',
    method: 'get',
    params
  })
}
