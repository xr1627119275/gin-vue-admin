package highPort

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/highPort"
	highPortReq "github.com/flipped-aurora/gin-vue-admin/server/model/highPort/request"
	"gorm.io/gorm/clause"
)

type HighRiskPortConfigService struct{}

// CreateHighRiskPortConfig 创建高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) CreateHighRiskPortConfig(HRPC *highPort.HighRiskPortConfig) (err error) {
	err = global.GVA_DB.Create(HRPC).Error
	return err
}

// DeleteHighRiskPortConfig 删除高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) DeleteHighRiskPortConfig(ID string) (err error) {
	err = global.GVA_DB.Delete(&highPort.HighRiskPortConfig{}, "id = ?", ID).Error
	return err
}

// DeleteHighRiskPortConfigByIds 批量删除高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) DeleteHighRiskPortConfigByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]highPort.HighRiskPortConfig{}, "id in ?", IDs).Error
	return err
}

// UpdateHighRiskPortConfig 更新高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) UpdateHighRiskPortConfig(HRPC highPort.HighRiskPortConfig) (err error) {
	err = global.GVA_DB.Model(&highPort.HighRiskPortConfig{}).Where("id = ?", HRPC.ID).Updates(&HRPC).Error
	return err
}

// GetHighRiskPortConfig 根据ID获取高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) GetHighRiskPortConfig(ID string) (HRPC highPort.HighRiskPortConfig, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&HRPC).Error
	return
}

// GetHighRiskPortLogsInfoList 分页获取高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) GetHighRiskPortLogsInfoList(info highPortReq.HighRiskPortConfigSearch) (list []highPort.HighRiskPortLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&highPort.HighRiskPortLog{}).Preload("PortConfig")
	var HRPCs []highPort.HighRiskPortLog
	if info.PortNumber != nil && *info.PortNumber != 0 {
		db = db.Joins("PortConfig")
		db.Clauses(clause.Eq{
			Column: "PortConfig.port_number",
			Value:  info.PortNumber,
		})
	}

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

	err = db.Find(&HRPCs).Error
	return HRPCs, total, err
}

// GetHighRiskPortConfigInfoList 分页获取高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) GetHighRiskPortConfigInfoList(info highPortReq.HighRiskPortConfigSearch) (list []highPort.HighRiskPortConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&highPort.HighRiskPortConfig{})
	var HRPCs []highPort.HighRiskPortConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["port_number"] = true
	orderMap["risk_level"] = true
	orderMap["order"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("`order` desc")
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&HRPCs).Error
	return HRPCs, total, err
}
func (HRPCService *HighRiskPortConfigService) GetHighRiskPortConfigPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
