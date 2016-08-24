package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func DatasourceCoreSubnets() *schema.Resource {
	return &schema.Resource{
		Read: readSubnets,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vcn_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"subnets": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     resourceCoreSubnets(),
			},
		},
	}
}

func resourceCoreSubnets() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cidr_block": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"route_table_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcn_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_list_ids": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": &schema.Schema{
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

func readSubnets(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &SubnetsSync{
		D:      d,
		Client: client,
	}

	return crud.ReadResource(reader)
}
