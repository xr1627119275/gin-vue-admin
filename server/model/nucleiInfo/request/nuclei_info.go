package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	nuclei "github.com/projectdiscovery/nuclei/v3/lib"
	"time"
)

type NucleiSearch struct {
	StartCreatedAt         *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt           *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Input                  string     `json:"input" form:"input"`
	nuclei.TemplateFilters `json:"poc_filter" form:"poc_filter"`
	request.PageInfo
}
