package highPort

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/highPort"
	highPortReq "github.com/flipped-aurora/gin-vue-admin/server/model/highPort/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/fileio"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"path/filepath"
)

type PasswordRuleApi struct{}

// CreatePasswordRule 创建密码规则
// @Tags PasswordRule
// @Summary 创建密码规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body highPort.PasswordRule true "创建密码规则"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /PR/createPasswordRule [post]
func (PRApi *PasswordRuleApi) CreatePasswordRule(c *gin.Context) {
	var PR highPort.PasswordRule
	err := c.ShouldBindJSON(&PR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = PRService.CreatePasswordRule(&PR)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePasswordRule 删除密码规则
// @Tags PasswordRule
// @Summary 删除密码规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body highPort.PasswordRule true "删除密码规则"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /PR/deletePasswordRule [delete]
func (PRApi *PasswordRuleApi) DeletePasswordRule(c *gin.Context) {
	ID := c.Query("ID")
	err := PRService.DeletePasswordRule(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePasswordRuleByIds 批量删除密码规则
// @Tags PasswordRule
// @Summary 批量删除密码规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /PR/deletePasswordRuleByIds [delete]
func (PRApi *PasswordRuleApi) DeletePasswordRuleByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := PRService.DeletePasswordRuleByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePasswordRule 更新密码规则
// @Tags PasswordRule
// @Summary 更新密码规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body highPort.PasswordRule true "更新密码规则"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /PR/updatePasswordRule [put]
func (PRApi *PasswordRuleApi) UpdatePasswordRule(c *gin.Context) {
	var PR highPort.PasswordRule
	err := c.ShouldBindJSON(&PR)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = PRService.UpdatePasswordRule(PR)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPasswordRule 用id查询密码规则
// @Tags PasswordRule
// @Summary 用id查询密码规则
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query highPort.PasswordRule true "用id查询密码规则"
// @Success 200 {object} response.Response{data=highPort.PasswordRule,msg=string} "查询成功"
// @Router /PR/findPasswordRule [get]
func (PRApi *PasswordRuleApi) FindPasswordRule(c *gin.Context) {
	ID := c.Query("ID")
	rePR, err := PRService.GetPasswordRule(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rePR, c)
}

// GetPasswordRuleList 分页获取密码规则列表
// @Tags PasswordRule
// @Summary 分页获取密码规则列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query highPortReq.PasswordRuleSearch true "分页获取密码规则列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /PR/getPasswordRuleList [get]
func (PRApi *PasswordRuleApi) GetPasswordRuleList(c *gin.Context) {
	var pageInfo highPortReq.PasswordRuleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := PRService.GetPasswordRuleInfoList(pageInfo)
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

// GetPasswordRulePublic 不需要鉴权的密码规则接口
// @Tags PasswordRule
// @Summary 不需要鉴权的密码规则接口
// @accept application/json
// @Produce application/json
// @Param data query highPortReq.PasswordRuleSearch true "分页获取密码规则列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PR/getPasswordRulePublic [get]
func (PRApi *PasswordRuleApi) GetPasswordRulePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	PRService.GetPasswordRulePublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的密码规则接口信息",
	}, "获取成功", c)
}

// DictConfig 密码字典配置
// @Tags PasswordRule
// @Summary 密码字典配置
// @accept application/json
// @Produce application/json
// @Param data query highPortReq.PasswordRuleSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /PR/dict [POST]
func (PRApi *PasswordRuleApi) DictConfig(c *gin.Context) {
	var passDict highPortReq.PassDictReq

	var err error
	for {
		err = c.ShouldBindJSON(&passDict)
		if err != nil {
			break
		}
		err = fileio.WriteString(filepath.Join("pass.txt"), passDict.Content)
		break
	}
	// 请添加自己的业务逻辑
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	go func() {
		global.WeekPassList, _ = fileio.ReadLines(filepath.Join("pass.txt"))
	}()

	response.OkWithMessage("修改成功", c)
}

// GetDict 获取弱密码字典
// @Tags PasswordRule
// @Summary 获取弱密码字典
// @accept application/json
// @Produce application/json
// @Param data query highPortReq.PasswordRuleSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /PR/dict [GET]
func (PRApi *PasswordRuleApi) GetDict(c *gin.Context) {
	// 请添加自己的业务逻辑
	data, err := fileio.ReadLines(filepath.Join("pass.txt"))
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData(data, c)
}
