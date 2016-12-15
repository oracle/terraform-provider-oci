package database

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func DatabaseDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabase,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"db_home_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_unique_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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

func readDatabase(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DatabaseDatasourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}
