package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func VolumeAttachmentResource() *schema.Resource {
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
				Computed: true,
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
				Required: true,
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
			"volume_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createVolumeAttachment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeAttachmentResourceCrud{D: d, Client: client}
	return crud.CreateResource(d, sync)
}

func readVolumeAttachment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeAttachmentResourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}

func deleteVolumeAttachment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &VolumeAttachmentResourceCrud{D: d, Client: client}
	return crud.DeleteResource(sync)
}
