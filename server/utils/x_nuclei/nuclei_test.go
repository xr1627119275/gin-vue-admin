package x_nuclei

import (
	"context"
	"fmt"
	nuclei "github.com/projectdiscovery/nuclei/v3/lib"
	"github.com/projectdiscovery/nuclei/v3/pkg/output"
	"log"
	"testing"
)

func Test_nucleiTemplates(t *testing.T) {
	//ctx := context.Background()
	//ne, err := nuclei.NewNucleiEngineCtx(ctx, nuclei.WithTemplateFilters(nuclei.TemplateFilters{IDs: []string{"CVE-2021-3129"}}))
	//ne, err := nuclei.NewNucleiEngineCtx(ctx)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//err = ne.LoadAllTemplates()
	//if err != nil {
	//	return
	//}
	ne, err := nuclei.NewNucleiEngineCtx(context.Background()) //nuclei.WithTemplateFilters(nuclei.TemplateFilters{Tags: []string{"php", "cve2024", "iconv"}}),
	//nuclei.WithTemplateFilters(nuclei.TemplateFilters{Severity: "critical"}),
	//nuclei.WithTemplateFilters(nuclei.TemplateFilters{ExcludeIDs: }),

	if err != nil {
		panic(err)
	}
	// load targets and optionally probe non http/https targets

	//ne.ParseTemplate("")
	//ne.Store().LoadTemplates([]string{
	//	"C:\\Users\\xrdev\\GolandProjects\\test-nuclei",
	//}, []string{"high", "critical"})

	//templates := ne.Store().Templates()
	//for _, item := range templates {
	//	fmt.Print(item.ID,"")
	//}

	// setup waitgroup to handle concurrency
	//err = ne.LoadAllTemplates()
	// scan 1 = run dns templates on scanme.sh

	ne.LoadTargets([]string{"192.168.220.129:8080"}, true)
	//ne.LoadTargets([]string{"192.168.220.129:8080"}, true)
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//
	err = ne.ExecuteWithCallback(func(event *output.ResultEvent) {
		if len(event.ExtractedResults) > 0 {
			log.Println("[success]", event.ExtractedResults)
		}
	})
	fmt.Println(err)
	// wait for all scans to finish
}

func Test_nucleiTemplatesGet(t *testing.T) {
	ne, _ := nuclei.NewNucleiEngineCtx(context.Background(), //nuclei.WithTemplateFilters(nuclei.TemplateFilters{Tags: []string{"php", "cve2024", "iconv"}}),
		//nuclei.WithTemplateFilters(nuclei.TemplateFilters{Severity: "critical"}),
		nuclei.WithTemplateFilters(nuclei.TemplateFilters{Tags: []string{"cve2021"}}),
	)
	temps := ne.GetTemplates()
	fmt.Println(len(temps))
	//loader.New(loader.NewConfig(&types.Options{
	//	ExcludeTags: ne.GetTemplates(),
	//} , ))
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
