package nucleiInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/nucleiInfo"
	nucleiInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/nucleiInfo/request"
)

type NucleiService struct{}

// CreateNuclei 创建nucleiInfo记录
// Author [yourname](https://github.com/yourname)
func (nucleiService *NucleiService) CreateNuclei(nuclei *nucleiInfo.Nuclei) (err error) {
	err = global.GVA_DB.Create(nuclei).Error
	return err
}

// DeleteNuclei 删除nucleiInfo记录
// Author [yourname](https://github.com/yourname)
func (nucleiService *NucleiService) DeleteNuclei(ID string) (err error) {
	err = global.GVA_DB.Delete(&nucleiInfo.Nuclei{}, "id = ?", ID).Error
	return err
}

// DeleteNucleiByIds 批量删除nucleiInfo记录
// Author [yourname](https://github.com/yourname)
func (nucleiService *NucleiService) DeleteNucleiByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]nucleiInfo.Nuclei{}, "id in ?", IDs).Error
	return err
}

// UpdateNuclei 更新nucleiInfo记录
// Author [yourname](https://github.com/yourname)
func (nucleiService *NucleiService) UpdateNuclei(nuclei nucleiInfo.Nuclei) (err error) {
	err = global.GVA_DB.Model(&nucleiInfo.Nuclei{}).Where("id = ?", nuclei.ID).Updates(&nuclei).Error
	return err
}

// GetNuclei 根据ID获取nucleiInfo记录
// Author [yourname](https://github.com/yourname)
func (nucleiService *NucleiService) GetNuclei(ID string) (nuclei nucleiInfo.Nuclei, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&nuclei).Error
	return
}

// GetNucleiInfoList 分页获取nucleiInfo记录
// Author [yourname](https://github.com/yourname)
func (nucleiService *NucleiService) GetNucleiInfoList(info nucleiInfoReq.NucleiSearch) (list []nucleiInfo.Nuclei, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&nucleiInfo.Nuclei{})
	var nucleis []nucleiInfo.Nuclei
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

	err = db.Find(&nucleis).Error
	return nucleis, total, err
}
func (nucleiService *NucleiService) GetNucleiPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
