package database

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func DatabasesDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabases,
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
			"db_nodes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DatabaseDatasource(),
			},
		},
	}
}

func readDatabases(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DatabasesDatasourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}
