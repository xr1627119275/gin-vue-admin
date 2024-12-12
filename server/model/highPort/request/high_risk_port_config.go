package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type HighRiskPortConfigSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
	Sort       string `json:"sort" form:"sort"`
	Order      string `json:"order" form:"order"`
	Status     *int   `json:"status" form:"status"`
	PortNumber *int   `json:"portNumber" form:"portNumber"`
}

type PortScanSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
	Target   string `json:"target" form:"target"`
	Port     string `json:"port" form:"port"`
	TaskName string `json:"taskName" form:"taskName"`
}
