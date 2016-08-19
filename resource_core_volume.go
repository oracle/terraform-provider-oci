package main

import "github.com/hashicorp/terraform/helper/schema"

func ResourceCoreVolume() *schema.Resource {
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
	client := m.(BareMetalClient)
	sync := &VolumeSync{D: d, Client: client}
	return createResource(d, sync)
}

func readVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &VolumeSync{D: d, Client: client}
	return readResource(sync)
}

func updateVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &VolumeSync{D: d, Client: client}
	return updateResource(d, sync)
}

func deleteVolume(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &VolumeSync{D: d, Client: client}
	return sync.Delete()
}
