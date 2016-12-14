package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func DrgDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDrgs,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"drgs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DrgResource(),
			},
		},
	}
}

func readDrgs(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DrgDatasourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}
