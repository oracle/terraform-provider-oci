package database

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func DBHomesDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDBHomes,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_id": &schema.Schema{
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
			"db_homes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DBHomeDatasource(),
			},
		},
	}
}

func readDBHomes(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DBHomesDatasourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}
