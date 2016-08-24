package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCoreIPSec() *schema.Resource {
	return &schema.Resource{
		Create: createIPSec,
		Read:   readIPSec,
		Delete: deleteIPSec,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cpe_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"drg_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"static_routes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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

func createIPSec(d *schema.ResourceData, m interface{}) (e error) {
	sync := &IPSecSync{D: d, Client: m.(client.BareMetalClient)}
	return crud.CreateResource(d, sync)
}

func readIPSec(d *schema.ResourceData, m interface{}) (e error) {
	sync := &IPSecSync{D: d, Client: m.(client.BareMetalClient)}
	return crud.ReadResource(sync)
}

func deleteIPSec(d *schema.ResourceData, m interface{}) (e error) {
	sync := &IPSecSync{D: d, Client: m.(client.BareMetalClient)}
	return crud.DeleteResource(sync)
}
