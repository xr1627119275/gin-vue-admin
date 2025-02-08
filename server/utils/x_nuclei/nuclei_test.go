package x_nuclei

import "testing"

func Test_nucleiTemplates(t *testing.T) {
	temps := GetNucleiTemplates()
	t.Log(temps)
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
