package nucleiInfo

import (
	"context"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	nucleiInfoReq "github.com/flipped-aurora/gin-vue-admin/server/model/nucleiInfo/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/x_nuclei"
	"github.com/google/uuid"
	nuclei "github.com/projectdiscovery/nuclei/v3/lib"
	"github.com/projectdiscovery/nuclei/v3/pkg/output"
	"go.uber.org/zap"
	"io"
	"log"
	"os"
)

// CreateNucleiScan 创建扫描记录
func (nucleiService *NucleiService) CreateNucleiScan(nucleiScanQuery *nucleiInfoReq.NucleiScanQuery) (logId string, err error) {

	engine, err := nuclei.NewNucleiEngineCtx(context.Background(), //nuclei.WithTemplateFilters(nuclei.TemplateFilters{Tags: []string{"php", "cve2024", "iconv"}}),
		nuclei.WithTemplateFilters(nucleiScanQuery.PocFilter),
	)
	if err != nil {
		global.GVA_LOG.Error("NewNucleiEngineCtx 失败", zap.Error(err))
		return
	}
	logId = uuid.NewString()

	engine.LoadTargets(nucleiScanQuery.Targets, true)
	go func() {
		err = engine.ExecuteWithCallback(func(event *output.ResultEvent) {
			if len(event.ExtractedResults) > 0 {
				log.Println("[success]", event.ExtractedResults)
				data, _ := json.Marshal(event)
				utils.HighLogMessage[logId] <- string(data)
			}
		})

		utils.HighLogMessage[logId] <- "{{over_end}}"
		if err != nil {
			global.GVA_LOG.Error("nuclei 执行失败", zap.Error(err))
		}
	}()

	return
}

func (nucleiService *NucleiService) GetNucleiPocData(id string) ([]byte, error) {
	pocs := x_nuclei.GetNucleiAllTemplates()

	for _, item := range pocs {
		if item.ID == id {
			path := item.Path
			//out = make(chan string)
			//go func() {
			f, err := os.Open(path)
			if err != nil {
				return nil, err
			}
			defer f.Close()

			return io.ReadAll(f)
			//	scanner := bufio.NewScanner(f)
			//	for scanner.Scan() {
			//		out <- scanner.Text()
			//	}
			//}()
			//return <-out, nil
		}
	}
	return nil, nil

}
