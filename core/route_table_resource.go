package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func RouteTableResource() *schema.Resource {
	return &schema.Resource{
		Create: createRouteTable,
		Read:   readRouteTable,
		Update: updateRouteTable,
		Delete: deleteRouteTable,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"route_rules": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr_block": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"display_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"network_entity_id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"network_entity_type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"time_created": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"time_modified": &schema.Schema{
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
			"vcn_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createRouteTable(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &RouteTableResourceCrud{D: d, Client: client}
	return crud.CreateResource(d, crd)
}

func readRouteTable(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &RouteTableResourceCrud{D: d, Client: client}
	return crud.ReadResource(crd)
}

func updateRouteTable(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &RouteTableResourceCrud{D: d, Client: client}
	return crud.UpdateResource(d, crd)
}

func deleteRouteTable(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &RouteTableResourceCrud{D: d, Client: client}
	return crud.DeleteResource(crd)
}
