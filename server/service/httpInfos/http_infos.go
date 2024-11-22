package httpInfos

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/httpInfos"
	httpInfosReq "github.com/flipped-aurora/gin-vue-admin/server/model/httpInfos/request"
)

type HttpInfosService struct{}

// CreateHttpInfos 创建httpInfos表记录
// Author [piexlmax](https://github.com/piexlmax)
func (httpInfoService *HttpInfosService) CreateHttpInfos(httpInfo *httpInfos.HttpInfos) (err error) {
	err = global.GVA_DB.Create(httpInfo).Error
	return err
}

// DeleteHttpInfos 删除httpInfos表记录
// Author [piexlmax](https://github.com/piexlmax)
func (httpInfoService *HttpInfosService) DeleteHttpInfos(id string) (err error) {
	err = global.GVA_DB.Delete(&httpInfos.HttpInfos{}, "id = ?", id).Error
	return err
}

// DeleteHttpInfosByIds 批量删除httpInfos表记录
// Author [piexlmax](https://github.com/piexlmax)
func (httpInfoService *HttpInfosService) DeleteHttpInfosByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]httpInfos.HttpInfos{}, "id in ?", ids).Error
	return err
}

// UpdateHttpInfos 更新httpInfos表记录
// Author [piexlmax](https://github.com/piexlmax)
func (httpInfoService *HttpInfosService) UpdateHttpInfos(httpInfo httpInfos.HttpInfos) (err error) {
	err = global.GVA_DB.Model(&httpInfos.HttpInfos{}).Where("id = ?", httpInfo.Id).Updates(&httpInfo).Error
	return err
}

// GetHttpInfos 根据id获取httpInfos表记录
// Author [piexlmax](https://github.com/piexlmax)
func (httpInfoService *HttpInfosService) GetHttpInfos(id string) (httpInfo httpInfos.HttpInfos, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&httpInfo).Error
	return
}

// GetHttpInfosInfoList 分页获取httpInfos表记录
// Author [piexlmax](https://github.com/piexlmax)
func (httpInfoService *HttpInfosService) GetHttpInfosInfoList(info httpInfosReq.HttpInfosSearch) (list []httpInfos.HttpInfos, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&httpInfos.HttpInfos{})
	var httpInfos []httpInfos.HttpInfos
	// 如果有条件搜索 下方会自动创建搜索语句

	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("created_at desc")
	}

	err = db.Find(&httpInfos).Error
	return httpInfos, total, err
}
