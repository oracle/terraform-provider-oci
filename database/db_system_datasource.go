package database

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func DBSystemDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDBSystems,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_systems": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DBSystemResource(),
			},
		},
	}
}

func readDBSystems(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DBSystemDatasourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}
