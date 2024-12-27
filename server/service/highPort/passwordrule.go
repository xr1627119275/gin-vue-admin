package highPort

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/highPort"
	highPortReq "github.com/flipped-aurora/gin-vue-admin/server/model/highPort/request"
)

type PasswordRuleService struct{}

// CreatePasswordRule 创建密码规则记录
// Author [yourname](https://github.com/yourname)
func (PRService *PasswordRuleService) CreatePasswordRule(PR *highPort.PasswordRule) (err error) {
	err = global.GVA_DB.Create(PR).Error
	return err
}

// DeletePasswordRule 删除密码规则记录
// Author [yourname](https://github.com/yourname)
func (PRService *PasswordRuleService) DeletePasswordRule(ID string) (err error) {
	err = global.GVA_DB.Delete(&highPort.PasswordRule{}, "id = ?", ID).Error
	return err
}

// DeletePasswordRuleByIds 批量删除密码规则记录
// Author [yourname](https://github.com/yourname)
func (PRService *PasswordRuleService) DeletePasswordRuleByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]highPort.PasswordRule{}, "id in ?", IDs).Error
	return err
}

// UpdatePasswordRule 更新密码规则记录
// Author [yourname](https://github.com/yourname)
func (PRService *PasswordRuleService) UpdatePasswordRule(PR highPort.PasswordRule) (err error) {
	err = global.GVA_DB.Model(&highPort.PasswordRule{}).Where("id = ?", PR.ID).Updates(&PR).Error
	return err
}

// GetPasswordRule 根据ID获取密码规则记录
// Author [yourname](https://github.com/yourname)
func (PRService *PasswordRuleService) GetPasswordRule(ID string) (PR highPort.PasswordRule, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&PR).Error
	return
}

// GetPasswordRuleInfoList 分页获取密码规则记录
// Author [yourname](https://github.com/yourname)
func (PRService *PasswordRuleService) GetPasswordRuleInfoList(info highPortReq.PasswordRuleSearch) (list []highPort.PasswordRule, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&highPort.PasswordRule{})
	var PRs []highPort.PasswordRule
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&PRs).Error
	return PRs, total, err
}
func (PRService *PasswordRuleService) GetPasswordRulePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// DictConfig 密码字典配置
// Author [yourname](https://github.com/yourname)
func (PRService *PasswordRuleService) DictConfig() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&highPort.PasswordRule{})
	return db.Error
}

// GetDict 获取弱密码字典
// Author [yourname](https://github.com/yourname)
func (PRService *PasswordRuleService) GetDict() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&highPort.PasswordRule{})
	return db.Error
}
