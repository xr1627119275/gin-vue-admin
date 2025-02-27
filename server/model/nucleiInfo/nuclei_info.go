// 自动生成模板Nuclei
package nucleiInfo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// nucleiInfo 结构体  Nuclei
type Nuclei struct {
	global.GVA_MODEL
	Template_id *string `json:"template_id" form:"template_id" gorm:"column:template_id;comment:;"` //模板ID
}

// TableName nucleiInfo Nuclei自定义表名 nuclei
func (Nuclei) TableName() string {
	return "nuclei"
}

// nucleiInfo 结构体  Nuclei
type NucleiScan struct {
	Target      *string  `json:"target" form:"target" gorm:"column:target;comment:;"`
	TemplateIds []string `json:"template_ids" form:"template_ids" gorm:"column:template_ids;comment:;"`
	global.GVA_MODEL
}

// TableName nucleiInfo Nuclei自定义表名 nuclei
func (NucleiScan) TableName() string {
	return "nucleiScan"
}
