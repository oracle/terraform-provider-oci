package main

import "github.com/hashicorp/terraform/helper/schema"

func ResourceCoreVolumeAttachment() *schema.Resource {
	return &schema.Resource{
		Create: createVolumeAttachment,
		Read:   readVolumeAttachment,
		Delete: deleteVolumeAttachment,
		Schema: map[string]*schema.Schema{
			"attachment_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": &schema.Schema{
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
			"volume_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createVolumeAttachment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &VolumeAttachmentSync{D: d, Client: client}
	return createResource(d, sync)
}

func readVolumeAttachment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &VolumeAttachmentSync{D: d, Client: client}
	return readResource(sync)
}

func deleteVolumeAttachment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &VolumeAttachmentSync{D: d, Client: client}
	return sync.Delete()
}
