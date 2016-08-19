package main

import "github.com/hashicorp/terraform/helper/schema"

func ResourceCoreVolumeAttachments() *schema.Resource {
	return &schema.Resource{
		Read: readVolumeAttachments,
		Schema: map[string]*schema.Schema{
			"availability_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_attachments": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceCoreVolumeAttachment(),
			},
		},
	}
}

func readVolumeAttachments(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &VolumeAttachmentsSync{D: d, Client: client}
	return readResource(sync)
}
