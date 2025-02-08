package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/highPort"
	"github.com/flipped-aurora/gin-vue-admin/server/model/httpInfos"
	"github.com/flipped-aurora/gin-vue-admin/server/model/nucleiInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webScan"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(httpInfos.HttpInfos{}, httpInfos.HttpWeakPassword{}, webScan.WebScan{}, highPort.HighRiskPortConfig{}, highPort.HighRiskPortLog{}, highPort.HighRiskPortScan{}, highPort.PasswordRule{}, nucleiInfo.Nuclei{})
	if err != nil {
		return err
	}
	return nil
}
