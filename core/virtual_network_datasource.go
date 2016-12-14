package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func VirtualNetworkDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readVirtualNetworks,
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
			"virtual_networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     VirtualNetworkResource(),
			},
		},
	}
}

func readVirtualNetworks(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VirtualNetworkDatasourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}
