package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCoreSubnet() *schema.Resource {
	return &schema.Resource{
		Create: createSubnet,
		Read:   readSubnet,
		Delete: deleteSubnet,
		Schema: map[string]*schema.Schema{
			"availability_domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cidr_block": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"route_table_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcn_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"security_list_ids": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"virtual_router_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_router_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createSubnet(d *schema.ResourceData, m interface{}) (e error) {
	sync := &SubnetSync{D: d, Client: m.(client.BareMetalClient)}
	return crud.CreateResource(d, sync)
}

func readSubnet(d *schema.ResourceData, m interface{}) (e error) {
	sync := &SubnetSync{D: d, Client: m.(client.BareMetalClient)}
	return crud.ReadResource(sync)
}

func deleteSubnet(d *schema.ResourceData, m interface{}) (e error) {
	sync := &SubnetSync{D: d, Client: m.(client.BareMetalClient)}
	return crud.DeleteResource(sync)
}
