package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func VolumeResource() *schema.Resource {
	return &schema.Resource{
		Create: createVolume,
		Read:   readVolume,
		Update: updateVolume,
		Delete: deleteVolume,
		Schema: map[string]*schema.Schema{
			"availability_domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"size_in_mbs": &schema.Schema{
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

func createVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeResourceCrud{D: d, Client: client}
	return crud.CreateResource(d, sync)
}

func readVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeResourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}

func updateVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeResourceCrud{D: d, Client: client}
	return crud.UpdateResource(d, sync)
}

func deleteVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeResourceCrud{D: d, Client: client}
	return sync.Delete()
}
