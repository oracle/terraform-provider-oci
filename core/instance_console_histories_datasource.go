package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func InstanceConsoleHistoriesDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readInstanceConsoleHistoriesDatasource,
		Schema: map[string]*schema.Schema{
			"instance_console_history_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"console_history": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1000000, // 1 MB
			},
		},
	}
}

func readInstanceConsoleHistoriesDatasource(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &InstanceConsoleHistoriesDatasourceCrud{
		D:      d,
		Client: client,
	}

	return crud.ReadResource(reader)
}
