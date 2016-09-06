package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func ConsoleHistoryResource() *schema.Resource {
	return &schema.Resource{
		Create: createConsoleHistory,
		Read:   readConsoleHistory,
		Delete: deleteConsoleHistory,
		Schema: map[string]*schema.Schema{
			"availability_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createConsoleHistory(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	ichCrud := &ConsoleHistoryResourceCrud{D: d, Client: client}
	return crud.CreateResource(d, ichCrud)
}

func readConsoleHistory(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	ichCrud := &ConsoleHistoryResourceCrud{D: d, Client: client}
	return crud.ReadResource(ichCrud)
}

func deleteConsoleHistory(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	ichCrud := &ConsoleHistoryResourceCrud{D: d, Client: client}
	return crud.DeleteResource(ichCrud)
}
