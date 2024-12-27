import service from '@/utils/request'
// @Tags PasswordRule
// @Summary 创建密码规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasswordRule true "创建密码规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PR/createPasswordRule [post]
export const createPasswordRule = (data) => {
  return service({
    url: '/PR/createPasswordRule',
    method: 'post',
    data
  })
}

// @Tags PasswordRule
// @Summary 删除密码规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasswordRule true "删除密码规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PR/deletePasswordRule [delete]
export const deletePasswordRule = (params) => {
  return service({
    url: '/PR/deletePasswordRule',
    method: 'delete',
    params
  })
}

// @Tags PasswordRule
// @Summary 批量删除密码规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除密码规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PR/deletePasswordRule [delete]
export const deletePasswordRuleByIds = (params) => {
  return service({
    url: '/PR/deletePasswordRuleByIds',
    method: 'delete',
    params
  })
}

// @Tags PasswordRule
// @Summary 更新密码规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PasswordRule true "更新密码规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PR/updatePasswordRule [put]
export const updatePasswordRule = (data) => {
  return service({
    url: '/PR/updatePasswordRule',
    method: 'put',
    data
  })
}

// @Tags PasswordRule
// @Summary 用id查询密码规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PasswordRule true "用id查询密码规则"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PR/findPasswordRule [get]
export const findPasswordRule = (params) => {
  return service({
    url: '/PR/findPasswordRule',
    method: 'get',
    params
  })
}

// @Tags PasswordRule
// @Summary 分页获取密码规则列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取密码规则列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PR/getPasswordRuleList [get]
export const getPasswordRuleList = (params) => {
  return service({
    url: '/PR/getPasswordRuleList',
    method: 'get',
    params
  })
}

// @Tags PasswordRule
// @Summary 不需要鉴权的密码规则接口
// @accept application/json
// @Produce application/json
// @Param data query highPortReq.PasswordRuleSearch true "分页获取密码规则列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PR/getPasswordRulePublic [get]
export const getPasswordRulePublic = () => {
  return service({
    url: '/PR/getPasswordRulePublic',
    method: 'get'
  })
}
// DictConfig 密码字典配置
// @Tags PasswordRule
// @Summary 密码字典配置
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /PR/dict [POST]
export const setPassDict = (data) => {
  return service({
    url: '/PR/dict',
    method: 'POST',
    data
  })
}
// GetDict 获取弱密码字典
// @Tags PasswordRule
// @Summary 获取弱密码字典
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /PR/dict [GET]
export const getPassDict = () => {
  return service({
    url: '/PR/dict',
    method: 'GET'
  })
}
