import service from '@/utils/request'
// @Tags Nuclei
// @Summary 创建nucleiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Nuclei true "创建nucleiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /nuclei/createNuclei [post]
export const createNuclei = (data) => {
  return service({
    url: '/nuclei/createNuclei',
    method: 'post',
    data
  })
}

// @Tags Nuclei
// @Summary 删除nucleiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Nuclei true "删除nucleiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /nuclei/deleteNuclei [delete]
export const deleteNuclei = (params) => {
  return service({
    url: '/nuclei/deleteNuclei',
    method: 'delete',
    params
  })
}

// @Tags Nuclei
// @Summary 批量删除nucleiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除nucleiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /nuclei/deleteNuclei [delete]
export const deleteNucleiByIds = (params) => {
  return service({
    url: '/nuclei/deleteNucleiByIds',
    method: 'delete',
    params
  })
}

// @Tags Nuclei
// @Summary 更新nucleiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Nuclei true "更新nucleiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /nuclei/updateNuclei [put]
export const updateNuclei = (data) => {
  return service({
    url: '/nuclei/updateNuclei',
    method: 'put',
    data
  })
}

// @Tags Nuclei
// @Summary 用id查询nucleiInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Nuclei true "用id查询nucleiInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /nuclei/findNuclei [get]
export const findNuclei = (params) => {
  return service({
    url: '/nuclei/findNuclei',
    method: 'get',
    params
  })
}

// @Tags Nuclei
// @Summary 分页获取nucleiInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取nucleiInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /nuclei/getNucleiList [get]
export const getNucleiList = (params) => {
  return service({
    url: '/nuclei/getNucleiList',
    method: 'get',
    params
  })
}

// @Tags Nuclei
// @Summary 分页获取nuclei 模板列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "模板列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /nuclei/getNucleiTemplateList [post]
export const getNucleiTemplateList = (data) => {
  return service({
    url: '/nuclei/getNucleiTemplateList',
    method: 'post',
    data
  })
}

export const createNucleiScan = (data) => {
  return service({
    url: '/nuclei/createScan',
    method: 'post',
    data
  })
}


export const getNucleiPocData = (params) => {
  return service({
    url: '/nuclei/getNucleiPocData',
    method: 'get',
    params
  })
}

// @Tags Nuclei
// @Summary 不需要鉴权的nucleiInfo接口
// @accept application/json
// @Produce application/json
// @Param data query nucleiInfoReq.NucleiSearch true "分页获取nucleiInfo列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /nuclei/getNucleiPublic [get]
export const getNucleiPublic = () => {
  return service({
    url: '/nuclei/getNucleiPublic',
    method: 'get',
  })
}
