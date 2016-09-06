package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func ConsoleHistoryDataDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readConsoleHistoryData,
		Schema: map[string]*schema.Schema{
			"console_history_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"data": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"length": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"offset": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func readConsoleHistoryData(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &ConsoleHistoryDataDatasourceCrud{D: d, Client: client}

	return crud.ReadResource(reader)
}
