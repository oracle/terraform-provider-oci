package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func InstanceConsoleHistoriesResource() *schema.Resource {
	return &schema.Resource{
		Create: createInstanceConsoleHistories,
		Read:   readInstanceConsoleHistories,
		Delete: deleteInstanceConsoleHistories,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createInstanceConsoleHistories(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	ichCrud := &InstanceConsoleHistoriesResourceCrud{D: d, Client: client}
	return crud.CreateResource(d, ichCrud)
}

func readInstanceConsoleHistories(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	ichCrud := &InstanceConsoleHistoriesResourceCrud{D: d, Client: client}
	return crud.ReadResource(ichCrud)
}

func deleteInstanceConsoleHistories(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	ichCrud := &InstanceConsoleHistoriesResourceCrud{D: d, Client: client}
	return crud.DeleteResource(ichCrud)
}
