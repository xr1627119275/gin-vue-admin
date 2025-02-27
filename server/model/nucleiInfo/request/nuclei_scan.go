package request

import nuclei "github.com/projectdiscovery/nuclei/v3/lib"

type NucleiScanQuery struct {
	Targets   []string               `json:"targets" form:"targets"`
	PocFilter nuclei.TemplateFilters `json:"poc_filter" form:"poc_filter"`
	PocIds    []string               `json:"poc_ids" form:"poc_ids"`
}
