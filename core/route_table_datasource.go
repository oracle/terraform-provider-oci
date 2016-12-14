package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func RouteTableDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readRouteTables,
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
			"route_tables": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     RouteTableResource(),
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func readRouteTables(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &RouteTableDatasourceCrud{D: d, Client: client}

	return crud.ReadResource(reader)
}
