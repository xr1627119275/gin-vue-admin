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
