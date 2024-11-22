package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/httpInfos"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webScan"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(httpInfos.HttpInfos{}, webScan.WebScan{})
	if err != nil {
		return err
	}
	return nil
}
