package x_nuclei

import (
	"context"
	nuclei "github.com/projectdiscovery/nuclei/v3/lib"
	"github.com/projectdiscovery/nuclei/v3/pkg/templates"
)

var nucleiEngine *nuclei.NucleiEngine

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func init() {
	ctx := context.Background()
	//ne, err := nuclei.NewNucleiEngineCtx(ctx, nuclei.WithTemplateFilters(nuclei.TemplateFilters{IDs: []string{"CVE-2021-3129"}}))
	nucleiEngine, _ = nuclei.NewNucleiEngineCtx(ctx, nuclei.WithTemplateFilters(nuclei.TemplateFilters{Severity: "critical"}))
	err := nucleiEngine.LoadAllTemplates()
	if err != nil {
		panic(err)
		return
	}
}

func GetNucleiTemplates() []*templates.Template {
	return nucleiEngine.Store().Templates()
}
func main() {

	//ne.ParseTemplate("")
	//ne.Store().LoadTemplatesWithTags([]string{
	//	"C:\\Users\\16271\\GolandProjects\\test-nuclei",
	//}, []string{"high", "critical"})
	//templates := ne.Store().Templates()
	//fmt.Println(templates)
	//if err != nil {
	//	panic(err)
	//}
	//// setup waitgroup to handle concurrency
	////err = ne.LoadAllTemplates()
	//// scan 1 = run dns templates on scanme.sh
	//
	//ne.LoadTargets([]string{"127.0.0.1:8085"}, true)
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//
	//err = ne.ExecuteWithCallback(func(event *output.ResultEvent) {
	//	if len(event.ExtractedResults) > 0 {
	//		log.Println("[success]", event.ExtractedResults)
	//	}
	//})
	// wait for all scans to finish
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
