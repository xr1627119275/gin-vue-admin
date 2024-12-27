package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webScan"
	"time"
)

type WebScanSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	ID             string
	HOST           string `json:"host" form:"host" `
	request.PageInfo
}

type RunWebScanParam struct {
	RunType string `json:"type"`
	Target  string `json:"target"`
	webScan.WebScan
}

type ParseWebScanSearch struct {
	Type    string `json:"type" `
	TaskID  string `json:"task_id" `
	Content string `json:"content" `
}
