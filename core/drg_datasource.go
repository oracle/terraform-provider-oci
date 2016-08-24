package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func DatasourceCoreDrgs() *schema.Resource {
	return &schema.Resource{
		Read: readDrgs,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"drgs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceCoreDrg(),
			},
		},
	}
}

func readDrgs(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DrgsSync{D: d, Client: client}
	return crud.ReadResource(sync)
}
