package main

import "github.com/hashicorp/terraform/helper/schema"

func ResourceCoreCpe() *schema.Resource {
	return &schema.Resource{
		Create: createCpe,
		Read:   readCpe,
		Delete: deleteCpe,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}

}

func createCpe(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &CpeSync{D: d, Client: client}
	return createResource(d, sync)
}

func readCpe(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &CpeSync{D: d, Client: client}
	return readResource(sync)
}

func deleteCpe(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &CpeSync{D: d, Client: client}
	return sync.Delete()
}
